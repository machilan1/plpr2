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

// Three different resources help you manage your IAM policy for Cloud Storage ManagedFolder. Each of these resources serves a different use case:
//
// * `storage.ManagedFolderIamPolicy`: Authoritative. Sets the IAM policy for the managedfolder and replaces any existing policy already attached.
// * `storage.ManagedFolderIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedfolder are preserved.
// * `storage.ManagedFolderIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedfolder are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `storage.ManagedFolderIamPolicy`: Retrieves the IAM policy for the managedfolder
//
// > **Note:** `storage.ManagedFolderIamPolicy` **cannot** be used in conjunction with `storage.ManagedFolderIamBinding` and `storage.ManagedFolderIamMember` or they will fight over what your policy should be.
//
// > **Note:** `storage.ManagedFolderIamBinding` resources **can be** used in conjunction with `storage.ManagedFolderIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## storage.ManagedFolderIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/storage.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewManagedFolderIamPolicy(ctx, "policy", &storage.ManagedFolderIamPolicyArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				PolicyData:    pulumi.String(admin.PolicyData),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/storage.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewManagedFolderIamPolicy(ctx, "policy", &storage.ManagedFolderIamPolicyArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				PolicyData:    pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## storage.ManagedFolderIamBinding
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
//			_, err := storage.NewManagedFolderIamBinding(ctx, "binding", &storage.ManagedFolderIamBindingArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
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
//			_, err := storage.NewManagedFolderIamBinding(ctx, "binding", &storage.ManagedFolderIamBindingArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &storage.ManagedFolderIamBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## storage.ManagedFolderIamMember
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
//			_, err := storage.NewManagedFolderIamMember(ctx, "member", &storage.ManagedFolderIamMemberArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Member:        pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
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
//			_, err := storage.NewManagedFolderIamMember(ctx, "member", &storage.ManagedFolderIamMemberArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Member:        pulumi.String("user:jane@example.com"),
//				Condition: &storage.ManagedFolderIamMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## > **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
// ---
//
// # IAM policy for Cloud Storage ManagedFolder
// Three different resources help you manage your IAM policy for Cloud Storage ManagedFolder. Each of these resources serves a different use case:
//
// * `storage.ManagedFolderIamPolicy`: Authoritative. Sets the IAM policy for the managedfolder and replaces any existing policy already attached.
// * `storage.ManagedFolderIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedfolder are preserved.
// * `storage.ManagedFolderIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedfolder are preserved.
//
// # A data source can be used to retrieve policy data in advent you do not need creation
//
// * `storage.ManagedFolderIamPolicy`: Retrieves the IAM policy for the managedfolder
//
// > **Note:** `storage.ManagedFolderIamPolicy` **cannot** be used in conjunction with `storage.ManagedFolderIamBinding` and `storage.ManagedFolderIamMember` or they will fight over what your policy should be.
//
// > **Note:** `storage.ManagedFolderIamBinding` resources **can be** used in conjunction with `storage.ManagedFolderIamMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
//
// ## storage.ManagedFolderIamPolicy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/storage.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewManagedFolderIamPolicy(ctx, "policy", &storage.ManagedFolderIamPolicyArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				PolicyData:    pulumi.String(admin.PolicyData),
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
// With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/storage.admin",
//						Members: []string{
//							"user:jane@example.com",
//						},
//						Condition: {
//							Title:       "expires_after_2019_12_31",
//							Description: pulumi.StringRef("Expiring at midnight of 2019-12-31"),
//							Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewManagedFolderIamPolicy(ctx, "policy", &storage.ManagedFolderIamPolicyArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				PolicyData:    pulumi.String(admin.PolicyData),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## storage.ManagedFolderIamBinding
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
//			_, err := storage.NewManagedFolderIamBinding(ctx, "binding", &storage.ManagedFolderIamBindingArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
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
//			_, err := storage.NewManagedFolderIamBinding(ctx, "binding", &storage.ManagedFolderIamBindingArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &storage.ManagedFolderIamBindingConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// ## storage.ManagedFolderIamMember
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
//			_, err := storage.NewManagedFolderIamMember(ctx, "member", &storage.ManagedFolderIamMemberArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Member:        pulumi.String("user:jane@example.com"),
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
// With IAM Conditions:
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
//			_, err := storage.NewManagedFolderIamMember(ctx, "member", &storage.ManagedFolderIamMemberArgs{
//				Bucket:        pulumi.Any(folder.Bucket),
//				ManagedFolder: pulumi.Any(folder.Name),
//				Role:          pulumi.String("roles/storage.admin"),
//				Member:        pulumi.String("user:jane@example.com"),
//				Condition: &storage.ManagedFolderIamMemberConditionArgs{
//					Title:       pulumi.String("expires_after_2019_12_31"),
//					Description: pulumi.String("Expiring at midnight of 2019-12-31"),
//					Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms:
//
// * b/{{bucket}}/managedFolders/{{managed_folder}}
//
// * {{bucket}}/{{managed_folder}}
//
// Any variables not passed in the import command will be taken from the provider configuration.
//
// Cloud Storage managedfolder IAM resources can be imported using the resource identifiers, role, and member.
//
// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
//
// ```sh
// $ pulumi import gcp:storage/managedFolderIamPolicy:ManagedFolderIamPolicy editor "b/{{bucket}}/managedFolders/{{managed_folder}} roles/storage.objectViewer user:jane@example.com"
// ```
//
// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
//
// ```sh
// $ pulumi import gcp:storage/managedFolderIamPolicy:ManagedFolderIamPolicy editor "b/{{bucket}}/managedFolders/{{managed_folder}} roles/storage.objectViewer"
// ```
//
// IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
// $ pulumi import gcp:storage/managedFolderIamPolicy:ManagedFolderIamPolicy editor b/{{bucket}}/managedFolders/{{managed_folder}}
// ```
//
// -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
//
//	full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ManagedFolderIamPolicy struct {
	pulumi.CustomResourceState

	// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedFolder pulumi.StringOutput `pulumi:"managedFolder"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewManagedFolderIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagedFolderIamPolicy(ctx *pulumi.Context,
	name string, args *ManagedFolderIamPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedFolderIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.ManagedFolder == nil {
		return nil, errors.New("invalid value for required argument 'ManagedFolder'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ManagedFolderIamPolicy
	err := ctx.RegisterResource("gcp:storage/managedFolderIamPolicy:ManagedFolderIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedFolderIamPolicy gets an existing ManagedFolderIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedFolderIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedFolderIamPolicyState, opts ...pulumi.ResourceOption) (*ManagedFolderIamPolicy, error) {
	var resource ManagedFolderIamPolicy
	err := ctx.ReadResource("gcp:storage/managedFolderIamPolicy:ManagedFolderIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedFolderIamPolicy resources.
type managedFolderIamPolicyState struct {
	// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
	Bucket *string `pulumi:"bucket"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedFolder *string `pulumi:"managedFolder"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type ManagedFolderIamPolicyState struct {
	// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	ManagedFolder pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (ManagedFolderIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedFolderIamPolicyState)(nil)).Elem()
}

type managedFolderIamPolicyArgs struct {
	// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
	Bucket string `pulumi:"bucket"`
	// Used to find the parent resource to bind the IAM policy to
	ManagedFolder string `pulumi:"managedFolder"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a ManagedFolderIamPolicy resource.
type ManagedFolderIamPolicyArgs struct {
	// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	ManagedFolder pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (ManagedFolderIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedFolderIamPolicyArgs)(nil)).Elem()
}

type ManagedFolderIamPolicyInput interface {
	pulumi.Input

	ToManagedFolderIamPolicyOutput() ManagedFolderIamPolicyOutput
	ToManagedFolderIamPolicyOutputWithContext(ctx context.Context) ManagedFolderIamPolicyOutput
}

func (*ManagedFolderIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedFolderIamPolicy)(nil)).Elem()
}

func (i *ManagedFolderIamPolicy) ToManagedFolderIamPolicyOutput() ManagedFolderIamPolicyOutput {
	return i.ToManagedFolderIamPolicyOutputWithContext(context.Background())
}

func (i *ManagedFolderIamPolicy) ToManagedFolderIamPolicyOutputWithContext(ctx context.Context) ManagedFolderIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedFolderIamPolicyOutput)
}

// ManagedFolderIamPolicyArrayInput is an input type that accepts ManagedFolderIamPolicyArray and ManagedFolderIamPolicyArrayOutput values.
// You can construct a concrete instance of `ManagedFolderIamPolicyArrayInput` via:
//
//	ManagedFolderIamPolicyArray{ ManagedFolderIamPolicyArgs{...} }
type ManagedFolderIamPolicyArrayInput interface {
	pulumi.Input

	ToManagedFolderIamPolicyArrayOutput() ManagedFolderIamPolicyArrayOutput
	ToManagedFolderIamPolicyArrayOutputWithContext(context.Context) ManagedFolderIamPolicyArrayOutput
}

type ManagedFolderIamPolicyArray []ManagedFolderIamPolicyInput

func (ManagedFolderIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedFolderIamPolicy)(nil)).Elem()
}

func (i ManagedFolderIamPolicyArray) ToManagedFolderIamPolicyArrayOutput() ManagedFolderIamPolicyArrayOutput {
	return i.ToManagedFolderIamPolicyArrayOutputWithContext(context.Background())
}

func (i ManagedFolderIamPolicyArray) ToManagedFolderIamPolicyArrayOutputWithContext(ctx context.Context) ManagedFolderIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedFolderIamPolicyArrayOutput)
}

// ManagedFolderIamPolicyMapInput is an input type that accepts ManagedFolderIamPolicyMap and ManagedFolderIamPolicyMapOutput values.
// You can construct a concrete instance of `ManagedFolderIamPolicyMapInput` via:
//
//	ManagedFolderIamPolicyMap{ "key": ManagedFolderIamPolicyArgs{...} }
type ManagedFolderIamPolicyMapInput interface {
	pulumi.Input

	ToManagedFolderIamPolicyMapOutput() ManagedFolderIamPolicyMapOutput
	ToManagedFolderIamPolicyMapOutputWithContext(context.Context) ManagedFolderIamPolicyMapOutput
}

type ManagedFolderIamPolicyMap map[string]ManagedFolderIamPolicyInput

func (ManagedFolderIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedFolderIamPolicy)(nil)).Elem()
}

func (i ManagedFolderIamPolicyMap) ToManagedFolderIamPolicyMapOutput() ManagedFolderIamPolicyMapOutput {
	return i.ToManagedFolderIamPolicyMapOutputWithContext(context.Background())
}

func (i ManagedFolderIamPolicyMap) ToManagedFolderIamPolicyMapOutputWithContext(ctx context.Context) ManagedFolderIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedFolderIamPolicyMapOutput)
}

type ManagedFolderIamPolicyOutput struct{ *pulumi.OutputState }

func (ManagedFolderIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedFolderIamPolicy)(nil)).Elem()
}

func (o ManagedFolderIamPolicyOutput) ToManagedFolderIamPolicyOutput() ManagedFolderIamPolicyOutput {
	return o
}

func (o ManagedFolderIamPolicyOutput) ToManagedFolderIamPolicyOutputWithContext(ctx context.Context) ManagedFolderIamPolicyOutput {
	return o
}

// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
func (o ManagedFolderIamPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedFolderIamPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// (Computed) The etag of the IAM policy.
func (o ManagedFolderIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedFolderIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Used to find the parent resource to bind the IAM policy to
func (o ManagedFolderIamPolicyOutput) ManagedFolder() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedFolderIamPolicy) pulumi.StringOutput { return v.ManagedFolder }).(pulumi.StringOutput)
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (o ManagedFolderIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedFolderIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

type ManagedFolderIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (ManagedFolderIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedFolderIamPolicy)(nil)).Elem()
}

func (o ManagedFolderIamPolicyArrayOutput) ToManagedFolderIamPolicyArrayOutput() ManagedFolderIamPolicyArrayOutput {
	return o
}

func (o ManagedFolderIamPolicyArrayOutput) ToManagedFolderIamPolicyArrayOutputWithContext(ctx context.Context) ManagedFolderIamPolicyArrayOutput {
	return o
}

func (o ManagedFolderIamPolicyArrayOutput) Index(i pulumi.IntInput) ManagedFolderIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedFolderIamPolicy {
		return vs[0].([]*ManagedFolderIamPolicy)[vs[1].(int)]
	}).(ManagedFolderIamPolicyOutput)
}

type ManagedFolderIamPolicyMapOutput struct{ *pulumi.OutputState }

func (ManagedFolderIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedFolderIamPolicy)(nil)).Elem()
}

func (o ManagedFolderIamPolicyMapOutput) ToManagedFolderIamPolicyMapOutput() ManagedFolderIamPolicyMapOutput {
	return o
}

func (o ManagedFolderIamPolicyMapOutput) ToManagedFolderIamPolicyMapOutputWithContext(ctx context.Context) ManagedFolderIamPolicyMapOutput {
	return o
}

func (o ManagedFolderIamPolicyMapOutput) MapIndex(k pulumi.StringInput) ManagedFolderIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedFolderIamPolicy {
		return vs[0].(map[string]*ManagedFolderIamPolicy)[vs[1].(string)]
	}).(ManagedFolderIamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedFolderIamPolicyInput)(nil)).Elem(), &ManagedFolderIamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedFolderIamPolicyArrayInput)(nil)).Elem(), ManagedFolderIamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedFolderIamPolicyMapInput)(nil)).Elem(), ManagedFolderIamPolicyMap{})
	pulumi.RegisterOutputType(ManagedFolderIamPolicyOutput{})
	pulumi.RegisterOutputType(ManagedFolderIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(ManagedFolderIamPolicyMapOutput{})
}
