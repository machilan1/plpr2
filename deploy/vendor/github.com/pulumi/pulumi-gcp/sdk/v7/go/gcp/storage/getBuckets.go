// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a list of existing GCS buckets.
// See [the official documentation](https://cloud.google.com/storage/docs/introduction)
// and [API](https://cloud.google.com/storage/docs/json_api/v1/buckets/list).
//
// ## Example Usage
//
// Example GCS buckets.
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
//			_, err := storage.GetBuckets(ctx, &storage.GetBucketsArgs{
//				Project: pulumi.StringRef("example-project"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetBuckets(ctx *pulumi.Context, args *GetBucketsArgs, opts ...pulumi.InvokeOption) (*GetBucketsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBucketsResult
	err := ctx.Invoke("gcp:storage/getBuckets:getBuckets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBuckets.
type GetBucketsArgs struct {
	// Filter results to buckets whose names begin with this prefix.
	Prefix *string `pulumi:"prefix"`
	// The ID of the project. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getBuckets.
type GetBucketsResult struct {
	// A list of all retrieved GCS buckets. Structure is defined below.
	Buckets []GetBucketsBucket `pulumi:"buckets"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Prefix  *string `pulumi:"prefix"`
	Project *string `pulumi:"project"`
}

func GetBucketsOutput(ctx *pulumi.Context, args GetBucketsOutputArgs, opts ...pulumi.InvokeOption) GetBucketsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBucketsResult, error) {
			args := v.(GetBucketsArgs)
			r, err := GetBuckets(ctx, &args, opts...)
			var s GetBucketsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBucketsResultOutput)
}

// A collection of arguments for invoking getBuckets.
type GetBucketsOutputArgs struct {
	// Filter results to buckets whose names begin with this prefix.
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	// The ID of the project. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetBucketsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketsArgs)(nil)).Elem()
}

// A collection of values returned by getBuckets.
type GetBucketsResultOutput struct{ *pulumi.OutputState }

func (GetBucketsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBucketsResult)(nil)).Elem()
}

func (o GetBucketsResultOutput) ToGetBucketsResultOutput() GetBucketsResultOutput {
	return o
}

func (o GetBucketsResultOutput) ToGetBucketsResultOutputWithContext(ctx context.Context) GetBucketsResultOutput {
	return o
}

// A list of all retrieved GCS buckets. Structure is defined below.
func (o GetBucketsResultOutput) Buckets() GetBucketsBucketArrayOutput {
	return o.ApplyT(func(v GetBucketsResult) []GetBucketsBucket { return v.Buckets }).(GetBucketsBucketArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBucketsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBucketsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBucketsResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketsResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o GetBucketsResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBucketsResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBucketsResultOutput{})
}
