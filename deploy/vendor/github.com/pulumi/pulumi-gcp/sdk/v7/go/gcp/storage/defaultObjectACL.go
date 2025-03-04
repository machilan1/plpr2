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

// Authoritatively manages the default object ACLs for a Google Cloud Storage bucket
// without managing the bucket itself.
//
// > Note that for each object, its creator will have the `"OWNER"` role in addition
// to the default ACL that has been defined.
//
// For more information see
// [the official documentation](https://cloud.google.com/storage/docs/access-control/lists)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls).
//
// > Want fine-grained control over default object ACLs? Use `storage.DefaultObjectAccessControl`
// to control individual role entity pairs.
//
// ## Example Usage
//
// Example creating a default object ACL on a bucket with one owner, and one reader.
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
//			_, err := storage.NewBucket(ctx, "image-store", &storage.BucketArgs{
//				Name:     pulumi.String("image-store-bucket"),
//				Location: pulumi.String("EU"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewDefaultObjectACL(ctx, "image-store-default-acl", &storage.DefaultObjectACLArgs{
//				Bucket: image_store.Name,
//				RoleEntities: pulumi.StringArray{
//					pulumi.String("OWNER:user-my.email@gmail.com"),
//					pulumi.String("READER:group-mygroup"),
//				},
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
type DefaultObjectACL struct {
	pulumi.CustomResourceState

	// The name of the bucket it applies to.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities pulumi.StringArrayOutput `pulumi:"roleEntities"`
}

// NewDefaultObjectACL registers a new resource with the given unique name, arguments, and options.
func NewDefaultObjectACL(ctx *pulumi.Context,
	name string, args *DefaultObjectACLArgs, opts ...pulumi.ResourceOption) (*DefaultObjectACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultObjectACL
	err := ctx.RegisterResource("gcp:storage/defaultObjectACL:DefaultObjectACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultObjectACL gets an existing DefaultObjectACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultObjectACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultObjectACLState, opts ...pulumi.ResourceOption) (*DefaultObjectACL, error) {
	var resource DefaultObjectACL
	err := ctx.ReadResource("gcp:storage/defaultObjectACL:DefaultObjectACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultObjectACL resources.
type defaultObjectACLState struct {
	// The name of the bucket it applies to.
	Bucket *string `pulumi:"bucket"`
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities []string `pulumi:"roleEntities"`
}

type DefaultObjectACLState struct {
	// The name of the bucket it applies to.
	Bucket pulumi.StringPtrInput
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities pulumi.StringArrayInput
}

func (DefaultObjectACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectACLState)(nil)).Elem()
}

type defaultObjectACLArgs struct {
	// The name of the bucket it applies to.
	Bucket string `pulumi:"bucket"`
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities []string `pulumi:"roleEntities"`
}

// The set of arguments for constructing a DefaultObjectACL resource.
type DefaultObjectACLArgs struct {
	// The name of the bucket it applies to.
	Bucket pulumi.StringInput
	// List of role/entity pairs in the form `ROLE:entity`.
	// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
	// Omitting the field is the same as providing an empty list.
	RoleEntities pulumi.StringArrayInput
}

func (DefaultObjectACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultObjectACLArgs)(nil)).Elem()
}

type DefaultObjectACLInput interface {
	pulumi.Input

	ToDefaultObjectACLOutput() DefaultObjectACLOutput
	ToDefaultObjectACLOutputWithContext(ctx context.Context) DefaultObjectACLOutput
}

func (*DefaultObjectACL) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultObjectACL)(nil)).Elem()
}

func (i *DefaultObjectACL) ToDefaultObjectACLOutput() DefaultObjectACLOutput {
	return i.ToDefaultObjectACLOutputWithContext(context.Background())
}

func (i *DefaultObjectACL) ToDefaultObjectACLOutputWithContext(ctx context.Context) DefaultObjectACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectACLOutput)
}

// DefaultObjectACLArrayInput is an input type that accepts DefaultObjectACLArray and DefaultObjectACLArrayOutput values.
// You can construct a concrete instance of `DefaultObjectACLArrayInput` via:
//
//	DefaultObjectACLArray{ DefaultObjectACLArgs{...} }
type DefaultObjectACLArrayInput interface {
	pulumi.Input

	ToDefaultObjectACLArrayOutput() DefaultObjectACLArrayOutput
	ToDefaultObjectACLArrayOutputWithContext(context.Context) DefaultObjectACLArrayOutput
}

type DefaultObjectACLArray []DefaultObjectACLInput

func (DefaultObjectACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultObjectACL)(nil)).Elem()
}

func (i DefaultObjectACLArray) ToDefaultObjectACLArrayOutput() DefaultObjectACLArrayOutput {
	return i.ToDefaultObjectACLArrayOutputWithContext(context.Background())
}

func (i DefaultObjectACLArray) ToDefaultObjectACLArrayOutputWithContext(ctx context.Context) DefaultObjectACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectACLArrayOutput)
}

// DefaultObjectACLMapInput is an input type that accepts DefaultObjectACLMap and DefaultObjectACLMapOutput values.
// You can construct a concrete instance of `DefaultObjectACLMapInput` via:
//
//	DefaultObjectACLMap{ "key": DefaultObjectACLArgs{...} }
type DefaultObjectACLMapInput interface {
	pulumi.Input

	ToDefaultObjectACLMapOutput() DefaultObjectACLMapOutput
	ToDefaultObjectACLMapOutputWithContext(context.Context) DefaultObjectACLMapOutput
}

type DefaultObjectACLMap map[string]DefaultObjectACLInput

func (DefaultObjectACLMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultObjectACL)(nil)).Elem()
}

func (i DefaultObjectACLMap) ToDefaultObjectACLMapOutput() DefaultObjectACLMapOutput {
	return i.ToDefaultObjectACLMapOutputWithContext(context.Background())
}

func (i DefaultObjectACLMap) ToDefaultObjectACLMapOutputWithContext(ctx context.Context) DefaultObjectACLMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultObjectACLMapOutput)
}

type DefaultObjectACLOutput struct{ *pulumi.OutputState }

func (DefaultObjectACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultObjectACL)(nil)).Elem()
}

func (o DefaultObjectACLOutput) ToDefaultObjectACLOutput() DefaultObjectACLOutput {
	return o
}

func (o DefaultObjectACLOutput) ToDefaultObjectACLOutputWithContext(ctx context.Context) DefaultObjectACLOutput {
	return o
}

// The name of the bucket it applies to.
func (o DefaultObjectACLOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultObjectACL) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// List of role/entity pairs in the form `ROLE:entity`.
// See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
// Omitting the field is the same as providing an empty list.
func (o DefaultObjectACLOutput) RoleEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultObjectACL) pulumi.StringArrayOutput { return v.RoleEntities }).(pulumi.StringArrayOutput)
}

type DefaultObjectACLArrayOutput struct{ *pulumi.OutputState }

func (DefaultObjectACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultObjectACL)(nil)).Elem()
}

func (o DefaultObjectACLArrayOutput) ToDefaultObjectACLArrayOutput() DefaultObjectACLArrayOutput {
	return o
}

func (o DefaultObjectACLArrayOutput) ToDefaultObjectACLArrayOutputWithContext(ctx context.Context) DefaultObjectACLArrayOutput {
	return o
}

func (o DefaultObjectACLArrayOutput) Index(i pulumi.IntInput) DefaultObjectACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultObjectACL {
		return vs[0].([]*DefaultObjectACL)[vs[1].(int)]
	}).(DefaultObjectACLOutput)
}

type DefaultObjectACLMapOutput struct{ *pulumi.OutputState }

func (DefaultObjectACLMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultObjectACL)(nil)).Elem()
}

func (o DefaultObjectACLMapOutput) ToDefaultObjectACLMapOutput() DefaultObjectACLMapOutput {
	return o
}

func (o DefaultObjectACLMapOutput) ToDefaultObjectACLMapOutputWithContext(ctx context.Context) DefaultObjectACLMapOutput {
	return o
}

func (o DefaultObjectACLMapOutput) MapIndex(k pulumi.StringInput) DefaultObjectACLOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultObjectACL {
		return vs[0].(map[string]*DefaultObjectACL)[vs[1].(string)]
	}).(DefaultObjectACLOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultObjectACLInput)(nil)).Elem(), &DefaultObjectACL{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultObjectACLArrayInput)(nil)).Elem(), DefaultObjectACLArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultObjectACLMapInput)(nil)).Elem(), DefaultObjectACLMap{})
	pulumi.RegisterOutputType(DefaultObjectACLOutput{})
	pulumi.RegisterOutputType(DefaultObjectACLArrayOutput{})
	pulumi.RegisterOutputType(DefaultObjectACLMapOutput{})
}
