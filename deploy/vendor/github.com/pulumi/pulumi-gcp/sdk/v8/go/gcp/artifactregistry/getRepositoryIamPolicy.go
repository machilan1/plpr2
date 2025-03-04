// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the current IAM policy data for repository
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/artifactregistry"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactregistry.LookupRepositoryIamPolicy(ctx, &artifactregistry.LookupRepositoryIamPolicyArgs{
//				Project:    pulumi.StringRef(my_repo.Project),
//				Location:   pulumi.StringRef(my_repo.Location),
//				Repository: my_repo.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRepositoryIamPolicy(ctx *pulumi.Context, args *LookupRepositoryIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryIamPolicyResult
	err := ctx.Invoke("gcp:artifactregistry/getRepositoryIamPolicy:getRepositoryIamPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepositoryIamPolicy.
type LookupRepositoryIamPolicyArgs struct {
	// The name of the repository's location. In addition to specific regions,
	// special values for multi-region locations are `asia`, `europe`, and `us`.
	// See [here](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations),
	// or use the
	// artifactregistry.getLocations
	// data source for possible values. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository string `pulumi:"repository"`
}

// A collection of values returned by getRepositoryIamPolicy.
type LookupRepositoryIamPolicyResult struct {
	// (Computed) The etag of the IAM policy.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
	// (Required only by `artifactregistry.RepositoryIamPolicy`) The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	Repository string `pulumi:"repository"`
}

func LookupRepositoryIamPolicyOutput(ctx *pulumi.Context, args LookupRepositoryIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRepositoryIamPolicyResultOutput, error) {
			args := v.(LookupRepositoryIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gcp:artifactregistry/getRepositoryIamPolicy:getRepositoryIamPolicy", args, LookupRepositoryIamPolicyResultOutput{}, options).(LookupRepositoryIamPolicyResultOutput), nil
		}).(LookupRepositoryIamPolicyResultOutput)
}

// A collection of arguments for invoking getRepositoryIamPolicy.
type LookupRepositoryIamPolicyOutputArgs struct {
	// The name of the repository's location. In addition to specific regions,
	// special values for multi-region locations are `asia`, `europe`, and `us`.
	// See [here](https://cloud.google.com/artifact-registry/docs/repositories/repo-locations),
	// or use the
	// artifactregistry.getLocations
	// data source for possible values. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
	// location is specified, it is taken from the provider configuration.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	Repository pulumi.StringInput `pulumi:"repository"`
}

func (LookupRepositoryIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getRepositoryIamPolicy.
type LookupRepositoryIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryIamPolicyResult)(nil)).Elem()
}

func (o LookupRepositoryIamPolicyResultOutput) ToLookupRepositoryIamPolicyResultOutput() LookupRepositoryIamPolicyResultOutput {
	return o
}

func (o LookupRepositoryIamPolicyResultOutput) ToLookupRepositoryIamPolicyResultOutputWithContext(ctx context.Context) LookupRepositoryIamPolicyResultOutput {
	return o
}

// (Computed) The etag of the IAM policy.
func (o LookupRepositoryIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRepositoryIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRepositoryIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// (Required only by `artifactregistry.RepositoryIamPolicy`) The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o LookupRepositoryIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupRepositoryIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupRepositoryIamPolicyResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryIamPolicyResult) string { return v.Repository }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryIamPolicyResultOutput{})
}
