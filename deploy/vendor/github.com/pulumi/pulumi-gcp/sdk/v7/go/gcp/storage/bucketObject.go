// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new object inside an existing bucket in Google cloud storage service (GCS).
// [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied using the `storage.ObjectACL` resource.
//
//	For more information see
//
// [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/objects).
//
// ## Example Usage
//
// Example creating a public object in an existing `image-store` bucket.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := storage.NewBucketObject(ctx, "picture", &storage.BucketObjectArgs{
//				Name:   pulumi.String("butterfly01"),
//				Source: pulumi.NewFileAsset("/images/nature/garden-tiger-moth.jpg"),
//				Bucket: pulumi.String("image-store"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Example creating an empty folder in an existing `image-store` bucket.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := storage.NewBucketObject(ctx, "empty_folder", &storage.BucketObjectArgs{
//				Name:    pulumi.String("empty_folder/"),
//				Content: pulumi.String(" "),
//				Bucket:  pulumi.String("image-store"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// This resource does not support import.
type BucketObject struct {
	pulumi.CustomResourceState

	// The name of the containing bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrOutput `pulumi:"cacheControl"`
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content pulumi.StringOutput `pulumi:"content"`
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition pulumi.StringPtrOutput `pulumi:"contentDisposition"`
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding pulumi.StringPtrOutput `pulumi:"contentEncoding"`
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage pulumi.StringPtrOutput `pulumi:"contentLanguage"`
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c pulumi.StringOutput `pulumi:"crc32c"`
	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). Google [documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	CustomerEncryption BucketObjectCustomerEncryptionPtrOutput `pulumi:"customerEncryption"`
	DetectMd5hash      pulumi.StringPtrOutput                  `pulumi:"detectMd5hash"`
	// Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold pulumi.BoolPtrOutput `pulumi:"eventBasedHold"`
	// (Computed) The content generation of this object. Used for object [versioning](https://cloud.google.com/storage/docs/object-versioning) and [soft delete](https://cloud.google.com/storage/docs/soft-delete).
	Generation pulumi.IntOutput `pulumi:"generation"`
	// The resource name of the Cloud KMS key that will be used to [encrypt](https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys) the object.
	KmsKeyName pulumi.StringOutput `pulumi:"kmsKeyName"`
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash pulumi.StringOutput `pulumi:"md5hash"`
	// (Computed) A url reference to download this object.
	MediaLink pulumi.StringOutput `pulumi:"mediaLink"`
	// User-provided metadata, in key/value pairs.
	//
	// One of the following is required:
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Computed) The name of the object. Use this field in interpolations with `storage.ObjectACL` to recreate
	// `storage.ObjectACL` resources when your `storage.BucketObject` is recreated.
	OutputName pulumi.StringOutput `pulumi:"outputName"`
	// The [object retention](http://cloud.google.com/storage/docs/object-lock) settings for the object. The retention settings allow an object to be retained until a provided date. Structure is documented below.
	Retention BucketObjectRetentionPtrOutput `pulumi:"retention"`
	// (Computed) A url reference to this object.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	//
	// ***
	Source pulumi.AssetOrArchiveOutput `pulumi:"source"`
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
	// Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold pulumi.BoolPtrOutput `pulumi:"temporaryHold"`
}

// NewBucketObject registers a new resource with the given unique name, arguments, and options.
func NewBucketObject(ctx *pulumi.Context,
	name string, args *BucketObjectArgs, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Content != nil {
		args.Content = pulumi.ToSecret(args.Content).(pulumi.StringPtrInput)
	}
	if args.CustomerEncryption != nil {
		args.CustomerEncryption = pulumi.ToSecret(args.CustomerEncryption).(BucketObjectCustomerEncryptionPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"content",
		"customerEncryption",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketObject
	err := ctx.RegisterResource("gcp:storage/bucketObject:BucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObject gets an existing BucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectState, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	var resource BucketObject
	err := ctx.ReadResource("gcp:storage/bucketObject:BucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObject resources.
type bucketObjectState struct {
	// The name of the containing bucket.
	Bucket *string `pulumi:"bucket"`
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl *string `pulumi:"cacheControl"`
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content *string `pulumi:"content"`
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType *string `pulumi:"contentType"`
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c *string `pulumi:"crc32c"`
	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). Google [documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	CustomerEncryption *BucketObjectCustomerEncryption `pulumi:"customerEncryption"`
	DetectMd5hash      *string                         `pulumi:"detectMd5hash"`
	// Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold *bool `pulumi:"eventBasedHold"`
	// (Computed) The content generation of this object. Used for object [versioning](https://cloud.google.com/storage/docs/object-versioning) and [soft delete](https://cloud.google.com/storage/docs/soft-delete).
	Generation *int `pulumi:"generation"`
	// The resource name of the Cloud KMS key that will be used to [encrypt](https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys) the object.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash *string `pulumi:"md5hash"`
	// (Computed) A url reference to download this object.
	MediaLink *string `pulumi:"mediaLink"`
	// User-provided metadata, in key/value pairs.
	//
	// One of the following is required:
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name *string `pulumi:"name"`
	// (Computed) The name of the object. Use this field in interpolations with `storage.ObjectACL` to recreate
	// `storage.ObjectACL` resources when your `storage.BucketObject` is recreated.
	OutputName *string `pulumi:"outputName"`
	// The [object retention](http://cloud.google.com/storage/docs/object-lock) settings for the object. The retention settings allow an object to be retained until a provided date. Structure is documented below.
	Retention *BucketObjectRetention `pulumi:"retention"`
	// (Computed) A url reference to this object.
	SelfLink *string `pulumi:"selfLink"`
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	//
	// ***
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass *string `pulumi:"storageClass"`
	// Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold *bool `pulumi:"temporaryHold"`
}

type BucketObjectState struct {
	// The name of the containing bucket.
	Bucket pulumi.StringPtrInput
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrInput
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content pulumi.StringPtrInput
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition pulumi.StringPtrInput
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding pulumi.StringPtrInput
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage pulumi.StringPtrInput
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringPtrInput
	// (Computed) Base 64 CRC32 hash of the uploaded data.
	Crc32c pulumi.StringPtrInput
	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). Google [documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	CustomerEncryption BucketObjectCustomerEncryptionPtrInput
	DetectMd5hash      pulumi.StringPtrInput
	// Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold pulumi.BoolPtrInput
	// (Computed) The content generation of this object. Used for object [versioning](https://cloud.google.com/storage/docs/object-versioning) and [soft delete](https://cloud.google.com/storage/docs/soft-delete).
	Generation pulumi.IntPtrInput
	// The resource name of the Cloud KMS key that will be used to [encrypt](https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys) the object.
	KmsKeyName pulumi.StringPtrInput
	// (Computed) Base 64 MD5 hash of the uploaded data.
	Md5hash pulumi.StringPtrInput
	// (Computed) A url reference to download this object.
	MediaLink pulumi.StringPtrInput
	// User-provided metadata, in key/value pairs.
	//
	// One of the following is required:
	Metadata pulumi.StringMapInput
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name pulumi.StringPtrInput
	// (Computed) The name of the object. Use this field in interpolations with `storage.ObjectACL` to recreate
	// `storage.ObjectACL` resources when your `storage.BucketObject` is recreated.
	OutputName pulumi.StringPtrInput
	// The [object retention](http://cloud.google.com/storage/docs/object-lock) settings for the object. The retention settings allow an object to be retained until a provided date. Structure is documented below.
	Retention BucketObjectRetentionPtrInput
	// (Computed) A url reference to this object.
	SelfLink pulumi.StringPtrInput
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	//
	// ***
	Source pulumi.AssetOrArchiveInput
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass pulumi.StringPtrInput
	// Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold pulumi.BoolPtrInput
}

func (BucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectState)(nil)).Elem()
}

type bucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket string `pulumi:"bucket"`
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl *string `pulumi:"cacheControl"`
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content *string `pulumi:"content"`
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType *string `pulumi:"contentType"`
	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). Google [documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	CustomerEncryption *BucketObjectCustomerEncryption `pulumi:"customerEncryption"`
	DetectMd5hash      *string                         `pulumi:"detectMd5hash"`
	// Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold *bool `pulumi:"eventBasedHold"`
	// The resource name of the Cloud KMS key that will be used to [encrypt](https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys) the object.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// User-provided metadata, in key/value pairs.
	//
	// One of the following is required:
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name *string `pulumi:"name"`
	// The [object retention](http://cloud.google.com/storage/docs/object-lock) settings for the object. The retention settings allow an object to be retained until a provided date. Structure is documented below.
	Retention *BucketObjectRetention `pulumi:"retention"`
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	//
	// ***
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass *string `pulumi:"storageClass"`
	// Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold *bool `pulumi:"temporaryHold"`
}

// The set of arguments for constructing a BucketObject resource.
type BucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket pulumi.StringInput
	// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
	// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrInput
	// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
	Content pulumi.StringPtrInput
	// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
	ContentDisposition pulumi.StringPtrInput
	// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
	ContentEncoding pulumi.StringPtrInput
	// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
	ContentLanguage pulumi.StringPtrInput
	// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringPtrInput
	// Enables object encryption with Customer-Supplied Encryption Key (CSEK). Google [documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
	// Structure is documented below.
	CustomerEncryption BucketObjectCustomerEncryptionPtrInput
	DetectMd5hash      pulumi.StringPtrInput
	// Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
	EventBasedHold pulumi.BoolPtrInput
	// The resource name of the Cloud KMS key that will be used to [encrypt](https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys) the object.
	KmsKeyName pulumi.StringPtrInput
	// User-provided metadata, in key/value pairs.
	//
	// One of the following is required:
	Metadata pulumi.StringMapInput
	// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
	Name pulumi.StringPtrInput
	// The [object retention](http://cloud.google.com/storage/docs/object-lock) settings for the object. The retention settings allow an object to be retained until a provided date. Structure is documented below.
	Retention BucketObjectRetentionPtrInput
	// A path to the data you want to upload. Must be defined
	// if `content` is not.
	//
	// ***
	Source pulumi.AssetOrArchiveInput
	// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
	// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
	// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
	StorageClass pulumi.StringPtrInput
	// Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
	TemporaryHold pulumi.BoolPtrInput
}

func (BucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectArgs)(nil)).Elem()
}

type BucketObjectInput interface {
	pulumi.Input

	ToBucketObjectOutput() BucketObjectOutput
	ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput
}

func (*BucketObject) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (i *BucketObject) ToBucketObjectOutput() BucketObjectOutput {
	return i.ToBucketObjectOutputWithContext(context.Background())
}

func (i *BucketObject) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectOutput)
}

// BucketObjectArrayInput is an input type that accepts BucketObjectArray and BucketObjectArrayOutput values.
// You can construct a concrete instance of `BucketObjectArrayInput` via:
//
//	BucketObjectArray{ BucketObjectArgs{...} }
type BucketObjectArrayInput interface {
	pulumi.Input

	ToBucketObjectArrayOutput() BucketObjectArrayOutput
	ToBucketObjectArrayOutputWithContext(context.Context) BucketObjectArrayOutput
}

type BucketObjectArray []BucketObjectInput

func (BucketObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (i BucketObjectArray) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return i.ToBucketObjectArrayOutputWithContext(context.Background())
}

func (i BucketObjectArray) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectArrayOutput)
}

// BucketObjectMapInput is an input type that accepts BucketObjectMap and BucketObjectMapOutput values.
// You can construct a concrete instance of `BucketObjectMapInput` via:
//
//	BucketObjectMap{ "key": BucketObjectArgs{...} }
type BucketObjectMapInput interface {
	pulumi.Input

	ToBucketObjectMapOutput() BucketObjectMapOutput
	ToBucketObjectMapOutputWithContext(context.Context) BucketObjectMapOutput
}

type BucketObjectMap map[string]BucketObjectInput

func (BucketObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (i BucketObjectMap) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return i.ToBucketObjectMapOutputWithContext(context.Background())
}

func (i BucketObjectMap) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectMapOutput)
}

type BucketObjectOutput struct{ *pulumi.OutputState }

func (BucketObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (o BucketObjectOutput) ToBucketObjectOutput() BucketObjectOutput {
	return o
}

func (o BucketObjectOutput) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return o
}

// The name of the containing bucket.
func (o BucketObjectOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
// directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
func (o BucketObjectOutput) CacheControl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.CacheControl }).(pulumi.StringPtrOutput)
}

// Data as `string` to be uploaded. Must be defined if `source` is not. **Note**: The `content` field is marked as sensitive.
func (o BucketObjectOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
func (o BucketObjectOutput) ContentDisposition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.ContentDisposition }).(pulumi.StringPtrOutput)
}

// [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
func (o BucketObjectOutput) ContentEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.ContentEncoding }).(pulumi.StringPtrOutput)
}

// [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
func (o BucketObjectOutput) ContentLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.ContentLanguage }).(pulumi.StringPtrOutput)
}

// [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
func (o BucketObjectOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// (Computed) Base 64 CRC32 hash of the uploaded data.
func (o BucketObjectOutput) Crc32c() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Crc32c }).(pulumi.StringOutput)
}

// Enables object encryption with Customer-Supplied Encryption Key (CSEK). Google [documentation about CSEK.](https://cloud.google.com/storage/docs/encryption/customer-supplied-keys)
// Structure is documented below.
func (o BucketObjectOutput) CustomerEncryption() BucketObjectCustomerEncryptionPtrOutput {
	return o.ApplyT(func(v *BucketObject) BucketObjectCustomerEncryptionPtrOutput { return v.CustomerEncryption }).(BucketObjectCustomerEncryptionPtrOutput)
}

func (o BucketObjectOutput) DetectMd5hash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.DetectMd5hash }).(pulumi.StringPtrOutput)
}

// Whether an object is under [event-based hold](https://cloud.google.com/storage/docs/object-holds#hold-types). Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any).
func (o BucketObjectOutput) EventBasedHold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.BoolPtrOutput { return v.EventBasedHold }).(pulumi.BoolPtrOutput)
}

// (Computed) The content generation of this object. Used for object [versioning](https://cloud.google.com/storage/docs/object-versioning) and [soft delete](https://cloud.google.com/storage/docs/soft-delete).
func (o BucketObjectOutput) Generation() pulumi.IntOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.IntOutput { return v.Generation }).(pulumi.IntOutput)
}

// The resource name of the Cloud KMS key that will be used to [encrypt](https://cloud.google.com/storage/docs/encryption/using-customer-managed-keys) the object.
func (o BucketObjectOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.KmsKeyName }).(pulumi.StringOutput)
}

// (Computed) Base 64 MD5 hash of the uploaded data.
func (o BucketObjectOutput) Md5hash() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Md5hash }).(pulumi.StringOutput)
}

// (Computed) A url reference to download this object.
func (o BucketObjectOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.MediaLink }).(pulumi.StringOutput)
}

// User-provided metadata, in key/value pairs.
//
// One of the following is required:
func (o BucketObjectOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the object. If you're interpolating the name of this object, see `outputName` instead.
func (o BucketObjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Computed) The name of the object. Use this field in interpolations with `storage.ObjectACL` to recreate
// `storage.ObjectACL` resources when your `storage.BucketObject` is recreated.
func (o BucketObjectOutput) OutputName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.OutputName }).(pulumi.StringOutput)
}

// The [object retention](http://cloud.google.com/storage/docs/object-lock) settings for the object. The retention settings allow an object to be retained until a provided date. Structure is documented below.
func (o BucketObjectOutput) Retention() BucketObjectRetentionPtrOutput {
	return o.ApplyT(func(v *BucketObject) BucketObjectRetentionPtrOutput { return v.Retention }).(BucketObjectRetentionPtrOutput)
}

// (Computed) A url reference to this object.
func (o BucketObjectOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// A path to the data you want to upload. Must be defined
// if `content` is not.
//
// ***
func (o BucketObjectOutput) Source() pulumi.AssetOrArchiveOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.AssetOrArchiveOutput { return v.Source }).(pulumi.AssetOrArchiveOutput)
}

// The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
// Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`, `ARCHIVE`. If not provided, this defaults to the bucket's default
// storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
func (o BucketObjectOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.StorageClass }).(pulumi.StringOutput)
}

// Whether an object is under [temporary hold](https://cloud.google.com/storage/docs/object-holds#hold-types). While this flag is set to true, the object is protected against deletion and overwrites.
func (o BucketObjectOutput) TemporaryHold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.BoolPtrOutput { return v.TemporaryHold }).(pulumi.BoolPtrOutput)
}

type BucketObjectArrayOutput struct{ *pulumi.OutputState }

func (BucketObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) Index(i pulumi.IntInput) BucketObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].([]*BucketObject)[vs[1].(int)]
	}).(BucketObjectOutput)
}

type BucketObjectMapOutput struct{ *pulumi.OutputState }

func (BucketObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) MapIndex(k pulumi.StringInput) BucketObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].(map[string]*BucketObject)[vs[1].(string)]
	}).(BucketObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectInput)(nil)).Elem(), &BucketObject{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectArrayInput)(nil)).Elem(), BucketObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectMapInput)(nil)).Elem(), BucketObjectMap{})
	pulumi.RegisterOutputType(BucketObjectOutput{})
	pulumi.RegisterOutputType(BucketObjectArrayOutput{})
	pulumi.RegisterOutputType(BucketObjectMapOutput{})
}
