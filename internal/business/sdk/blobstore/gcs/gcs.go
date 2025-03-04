package gcs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/machilan1/plpr2/internal/business/sdk/blobstore"
	"github.com/machilan1/plpr2/internal/business/sdk/mimetype"
	"google.golang.org/api/option"
)

var (
	// emulatorHost is the host of the Google Cloud Storage emulator.
	// We use this environment variable to detect if the emulator is enabled, which
	// is recommended by the official documentation.
	emulatorHost = os.Getenv("STORAGE_EMULATOR_HOST")
	// defaultSignedURLExpiry is the default duration for which the signed URL is valid.
	defaultSignedURLExpiry = 15 * time.Minute
)

type Config struct {
	Bucket string
	// SignedURLExpiry is the duration for which the signed URL is valid.
	// If zero, a default value is used.
	SignedURLExpiry time.Duration
}

type GoogleCloudStorage struct {
	client          *storage.Client
	bucketName      string
	bucket          *storage.BucketHandle
	signedURLExpiry time.Duration
	// emulator is a flag indicating whether the Google Cloud Storage emulator is enabled.
	// It's determined by the presence of the `STORAGE_EMULATOR_HOST` environment variable.
	emulator bool
}

func NewGoogleCloudStorage(ctx context.Context, cfg Config) (*GoogleCloudStorage, error) {
	var opts []option.ClientOption
	if emulatorHost != "" {
		// The `public host` config for the emulator.
		// This is the host that the application outside the container will use to access the emulator.
		const publicEmulatorHost = "http://0.0.0.0:44443"

		// The emulator host is the one that the gcs client inside the container will use to access the emulator.
		// We need to parse it to create a proxy URL, which will be used by the client to access the emulator.
		proxyURL, err := url.Parse(emulatorHost)
		if err != nil {
			return nil, fmt.Errorf("cannot parse proxy URL: %w", err)
		}
		opts = append(opts,
			// We MUST use the public host of the emulator here, otherwise all requests from gcs client will fail.
			// The `/storage/v1/` is the required path prefix for Google Cloud Storage API.
			option.WithEndpoint(publicEmulatorHost+"/storage/v1/"),
			option.WithoutAuthentication(),
			option.WithHTTPClient(&http.Client{
				Transport: &http.Transport{
					// We need to convert the public host to use the correct emulator host,
					// since this application is running inside the container.
					Proxy: http.ProxyURL(proxyURL),
				},
			}),
		)
	}

	client, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("cannot create Google Cloud Storage client: %w", err)
	}

	signedURLExpiry := defaultSignedURLExpiry
	if cfg.SignedURLExpiry != 0 {
		signedURLExpiry = cfg.SignedURLExpiry
	}

	return &GoogleCloudStorage{
		client:          client,
		bucketName:      cfg.Bucket,
		bucket:          client.Bucket(cfg.Bucket),
		signedURLExpiry: signedURLExpiry,
		emulator:        emulatorHost != "",
	}, nil
}

var _ blobstore.BlobStore = (*GoogleCloudStorage)(nil)

// SignedURL returns a signed URL for the object with the given name.
// The object name should comply with the Google Cloud Storage object name requirements,
// see https://cloud.google.com/storage/docs/objects#naming.
func (s *GoogleCloudStorage) SignedURL(ctx context.Context, objectName string, perm blobstore.Permission) (string, error) {
	// Signing a URL requires credentials authorized to sign a URL. You can pass
	// these in through SignedURLOptions with one of the following options:
	//    a. a Google service account private key, obtainable from the Google Developers Console
	//    b. a Google Access ID with iam.serviceAccounts.signBlob permissions
	//    c. a SignBytes function implementing custom signing.
	// In this example, none of these options are used, which means the SignedURL
	// function attempts to use the same authentication that was used to instantiate
	// the Storage client. This authentication must include a private key or have
	// iam.serviceAccounts.signBlob permissions.
	opts := storage.SignedURLOptions{
		Scheme:      storage.SigningSchemeV4,
		Method:      http.MethodPut,
		Expires:     time.Now().Add(s.signedURLExpiry),
		ContentType: mimetype.DetectFilePath(objectName),
		Headers:     []string{},
	}
	if perm == blobstore.Read {
		opts.Method = http.MethodGet
		opts.ContentType = "" // TODO: test if this removal is necessary
	}

	if s.emulator {
		host := emulatorHost
		splitted := strings.Split(host, "//")
		if len(splitted) > 1 {
			// Remove the protocol prefix
			host = splitted[1]
		}

		// The host.docker.internal is a special DNS name which resolves to the internal IP address used by the host,
		// which is accessible from within the containers.
		// However, we need to replace it with localhost to make it work for frontend clients which are not in the same network.
		host = strings.Replace(host, "host.docker.internal", "0.0.0.0", 1)

		opts.GoogleAccessID = "demo@example.com"
		opts.SignBytes = func(_ []byte) ([]byte, error) {
			fakeSignedBytes := []byte("emulated-signature")
			return fakeSignedBytes, nil
		}
		opts.Insecure = true
		opts.Style = storage.BucketBoundHostname(host + "/" + s.bucketName)
	}

	u, err := s.bucket.SignedURL(objectName, &opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", s.bucketName, err)
	}

	if s.emulator {
		// The `BoundHostname` style will encode the hostname like `localhost%2F80/...`,
		// so we need to decode it back to the original form `localhost:80/...`.
		// This is a workaround for the emulator.
		u = strings.Replace(u, "%2F", "/", 1)
	}

	return u, nil
}

func (s *GoogleCloudStorage) Read(ctx context.Context, objectName string) ([]byte, error) {
	r, err := s.bucket.Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Bucket(%q).Object(%q).NewReader: %w", s.bucketName, objectName, err)
	}
	defer r.Close()

	return io.ReadAll(r)
}

func (s *GoogleCloudStorage) Upload(ctx context.Context, objectName string, content []byte) error {
	w := s.bucket.Object(objectName).NewWriter(ctx)
	mime := mimetype.DetectFilePath(objectName)
	w.ContentType = mime
	if _, err := w.Write(content); err != nil {
		return fmt.Errorf("Bucket(%q).Object(%q).NewWriter: %w", s.bucketName, objectName, err)
	}

	return w.Close()
}

func (s *GoogleCloudStorage) Close() error {
	return s.client.Close()
}
