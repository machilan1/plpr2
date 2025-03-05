// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Latest Recovery Time for a given instance. For more information see the
// [official documentation](https://cloud.google.com/sql/)
// and
// [API](https://cloud.google.com/sql/docs/postgres/backup-recovery/pitr#get-the-latest-recovery-time).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/sql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := sql.GetDatabaseInstanceLatestRecoveryTime(ctx, &sql.GetDatabaseInstanceLatestRecoveryTimeArgs{
//				Instance: "sample-instance",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("latestRecoveryTime", _default)
//			return nil
//		})
//	}
//
// ```
func GetDatabaseInstanceLatestRecoveryTime(ctx *pulumi.Context, args *GetDatabaseInstanceLatestRecoveryTimeArgs, opts ...pulumi.InvokeOption) (*GetDatabaseInstanceLatestRecoveryTimeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabaseInstanceLatestRecoveryTimeResult
	err := ctx.Invoke("gcp:sql/getDatabaseInstanceLatestRecoveryTime:getDatabaseInstanceLatestRecoveryTime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseInstanceLatestRecoveryTime.
type GetDatabaseInstanceLatestRecoveryTimeArgs struct {
	// The name of the instance.
	Instance string `pulumi:"instance"`
	// The ID of the project in which the resource belongs.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDatabaseInstanceLatestRecoveryTime.
type GetDatabaseInstanceLatestRecoveryTimeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the instance.
	Instance string `pulumi:"instance"`
	// Timestamp, identifies the latest recovery time of the source instance.
	LatestRecoveryTime string `pulumi:"latestRecoveryTime"`
	// The ID of the project in which the resource belongs.
	Project string `pulumi:"project"`
}

func GetDatabaseInstanceLatestRecoveryTimeOutput(ctx *pulumi.Context, args GetDatabaseInstanceLatestRecoveryTimeOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseInstanceLatestRecoveryTimeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDatabaseInstanceLatestRecoveryTimeResultOutput, error) {
			args := v.(GetDatabaseInstanceLatestRecoveryTimeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("gcp:sql/getDatabaseInstanceLatestRecoveryTime:getDatabaseInstanceLatestRecoveryTime", args, GetDatabaseInstanceLatestRecoveryTimeResultOutput{}, options).(GetDatabaseInstanceLatestRecoveryTimeResultOutput), nil
		}).(GetDatabaseInstanceLatestRecoveryTimeResultOutput)
}

// A collection of arguments for invoking getDatabaseInstanceLatestRecoveryTime.
type GetDatabaseInstanceLatestRecoveryTimeOutputArgs struct {
	// The name of the instance.
	Instance pulumi.StringInput `pulumi:"instance"`
	// The ID of the project in which the resource belongs.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetDatabaseInstanceLatestRecoveryTimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseInstanceLatestRecoveryTimeArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseInstanceLatestRecoveryTime.
type GetDatabaseInstanceLatestRecoveryTimeResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseInstanceLatestRecoveryTimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseInstanceLatestRecoveryTimeResult)(nil)).Elem()
}

func (o GetDatabaseInstanceLatestRecoveryTimeResultOutput) ToGetDatabaseInstanceLatestRecoveryTimeResultOutput() GetDatabaseInstanceLatestRecoveryTimeResultOutput {
	return o
}

func (o GetDatabaseInstanceLatestRecoveryTimeResultOutput) ToGetDatabaseInstanceLatestRecoveryTimeResultOutputWithContext(ctx context.Context) GetDatabaseInstanceLatestRecoveryTimeResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabaseInstanceLatestRecoveryTimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstanceLatestRecoveryTimeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the instance.
func (o GetDatabaseInstanceLatestRecoveryTimeResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstanceLatestRecoveryTimeResult) string { return v.Instance }).(pulumi.StringOutput)
}

// Timestamp, identifies the latest recovery time of the source instance.
func (o GetDatabaseInstanceLatestRecoveryTimeResultOutput) LatestRecoveryTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstanceLatestRecoveryTimeResult) string { return v.LatestRecoveryTime }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
func (o GetDatabaseInstanceLatestRecoveryTimeResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseInstanceLatestRecoveryTimeResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseInstanceLatestRecoveryTimeResultOutput{})
}
