package blobstore

import "context"

type Permission string

const (
	// Read Permission allows the user to read the object.
	Read Permission = "READ"
	// Write Permission allows the user to write the object.
	Write Permission = "WRITE"
)

// BlobStore provides access to a blob storage service.
// For example, AWS S3, Google Cloud Storage, Azure Blob Storage or just a local file system.
// The implementation should be able to sign a URL for a blob object.
//
// Objects are files in a bucket, where the object name is the full path to the file.
// For example, "folder/file.txt".
// A bucket is a container for objects, similar to a folder.
type BlobStore interface {
	// SignedURL returns a signed URL that allows a user to download an object.
	SignedURL(ctx context.Context, objectName string, perm Permission) (string, error)
	// Read reads the content of the object.
	Read(ctx context.Context, objectName string) ([]byte, error)
	// Upload uploads the content to the object.
	Upload(ctx context.Context, objectName string, data []byte) error
}

// =====================================================================================================================

type NoopBlobStore struct{}

var _ BlobStore = (*NoopBlobStore)(nil)

func NewNoopBlobStore() *NoopBlobStore {
	return &NoopBlobStore{}
}

func (*NoopBlobStore) SignedURL(ctx context.Context, objectName string, perm Permission) (string, error) {
	return objectName, nil
}

func (*NoopBlobStore) Read(ctx context.Context, objectName string) ([]byte, error) {
	return []byte(objectName), nil
}

func (*NoopBlobStore) Upload(ctx context.Context, objectName string, data []byte) error {
	return nil
}
