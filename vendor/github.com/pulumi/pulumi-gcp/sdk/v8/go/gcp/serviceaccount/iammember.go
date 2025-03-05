// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource, such as allowing the members to run operations as or modify the service account. To configure permissions for a service account on other GCP resources, use the googleProjectIam set of resources.
//
// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
//
// * `serviceaccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
// * `serviceaccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
// * `serviceaccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
//
// > **Note:** `serviceaccount.IAMPolicy` **cannot** be used in conjunction with `serviceaccount.IAMBinding` and `serviceaccount.IAMMember` or they will fight over what your policy should be.
//
// > **Note:** `serviceaccount.IAMBinding` resources **can be** used in conjunction with `serviceaccount.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## Example Usage
//
// ### Service Account IAM Policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/iam.serviceAccountUser",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that only Jane can interact with"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMPolicy(ctx, "admin-account-iam", &serviceaccount.IAMPolicyArgs{
//				ServiceAccountId: sa.Name,
//				PolicyData:       pulumi.String(admin.PolicyData),
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
// ### Service Account IAM Binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that only Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMBinding(ctx, "admin-account-iam", &serviceaccount.IAMBindingArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
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
// ### Service Account IAM Binding With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that only Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMBinding(ctx, "admin-account-iam", &serviceaccount.IAMBindingArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &serviceaccount.IAMBindingConditionArgs{
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
// ### Service Account IAM Member
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := compute.GetDefaultServiceAccount(ctx, &compute.GetDefaultServiceAccountArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMMember(ctx, "admin-account-iam", &serviceaccount.IAMMemberArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Member:           pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			// Allow SA service account use the default GCE account
//			_, err = serviceaccount.NewIAMMember(ctx, "gce-default-account-iam", &serviceaccount.IAMMemberArgs{
//				ServiceAccountId: pulumi.String(_default.Name),
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Member: sa.Email.ApplyT(func(email string) (string, error) {
//					return fmt.Sprintf("serviceAccount:%v", email), nil
//				}).(pulumi.StringOutput),
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
// ### Service Account IAM Member With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMMember(ctx, "admin-account-iam", &serviceaccount.IAMMemberArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Member:           pulumi.String("user:jane@example.com"),
//				Condition: &serviceaccount.IAMMemberConditionArgs{
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
// ### Additional Examples
//
// ### Service Account IAM Policy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/iam.serviceAccountUser",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that only Jane can interact with"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMPolicy(ctx, "admin-account-iam", &serviceaccount.IAMPolicyArgs{
//				ServiceAccountId: sa.Name,
//				PolicyData:       pulumi.String(admin.PolicyData),
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
// ### Service Account IAM Binding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that only Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMBinding(ctx, "admin-account-iam", &serviceaccount.IAMBindingArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
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
// ### Service Account IAM Binding With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that only Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMBinding(ctx, "admin-account-iam", &serviceaccount.IAMBindingArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &serviceaccount.IAMBindingConditionArgs{
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
// ### Service Account IAM Member
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := compute.GetDefaultServiceAccount(ctx, &compute.GetDefaultServiceAccountArgs{}, nil)
//			if err != nil {
//				return err
//			}
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMMember(ctx, "admin-account-iam", &serviceaccount.IAMMemberArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Member:           pulumi.String("user:jane@example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			// Allow SA service account use the default GCE account
//			_, err = serviceaccount.NewIAMMember(ctx, "gce-default-account-iam", &serviceaccount.IAMMemberArgs{
//				ServiceAccountId: pulumi.String(_default.Name),
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Member: sa.Email.ApplyT(func(email string) (string, error) {
//					return fmt.Sprintf("serviceAccount:%v", email), nil
//				}).(pulumi.StringOutput),
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
// ### Service Account IAM Member With IAM Conditions:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			sa, err := serviceaccount.NewAccount(ctx, "sa", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("my-service-account"),
//				DisplayName: pulumi.String("A service account that Jane can use"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = serviceaccount.NewIAMMember(ctx, "admin-account-iam", &serviceaccount.IAMMemberArgs{
//				ServiceAccountId: sa.Name,
//				Role:             pulumi.String("roles/iam.serviceAccountUser"),
//				Member:           pulumi.String("user:jane@example.com"),
//				Condition: &serviceaccount.IAMMemberConditionArgs{
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
// ### Importing with conditions:
//
// Here are examples of importing IAM memberships and bindings that include conditions:
//
// ```sh
// $ pulumi import gcp:serviceaccount/iAMMember:IAMMember admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser expires_after_2019_12_31"
// ```
//
// ```sh
// $ pulumi import gcp:serviceaccount/iAMMember:IAMMember admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
// ```
type IAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the service account IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewIAMMember registers a new resource with the given unique name, arguments, and options.
func NewIAMMember(ctx *pulumi.Context,
	name string, args *IAMMemberArgs, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("gcp:serviceAccount/iAMMember:IAMMember"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IAMMember
	err := ctx.RegisterResource("gcp:serviceaccount/iAMMember:IAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMMember gets an existing IAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMMemberState, opts ...pulumi.ResourceOption) (*IAMMember, error) {
	var resource IAMMember
	err := ctx.ReadResource("gcp:serviceaccount/iAMMember:IAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMMember resources.
type iammemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the service account IAM policy.
	Etag *string `pulumi:"etag"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type IAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	// (Computed) The etag of the service account IAM policy.
	Etag pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringPtrInput
}

func (IAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberState)(nil)).Elem()
}

type iammemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IAMMemberCondition `pulumi:"condition"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member string `pulumi:"member"`
	// The role that should be applied. Only one
	// `serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a IAMMember resource.
type IAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IAMMemberConditionPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringInput
	// The role that should be applied. Only one
	// `serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringInput
}

func (IAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iammemberArgs)(nil)).Elem()
}

type IAMMemberInput interface {
	pulumi.Input

	ToIAMMemberOutput() IAMMemberOutput
	ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput
}

func (*IAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (i *IAMMember) ToIAMMemberOutput() IAMMemberOutput {
	return i.ToIAMMemberOutputWithContext(context.Background())
}

func (i *IAMMember) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberOutput)
}

// IAMMemberArrayInput is an input type that accepts IAMMemberArray and IAMMemberArrayOutput values.
// You can construct a concrete instance of `IAMMemberArrayInput` via:
//
//	IAMMemberArray{ IAMMemberArgs{...} }
type IAMMemberArrayInput interface {
	pulumi.Input

	ToIAMMemberArrayOutput() IAMMemberArrayOutput
	ToIAMMemberArrayOutputWithContext(context.Context) IAMMemberArrayOutput
}

type IAMMemberArray []IAMMemberInput

func (IAMMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (i IAMMemberArray) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return i.ToIAMMemberArrayOutputWithContext(context.Background())
}

func (i IAMMemberArray) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberArrayOutput)
}

// IAMMemberMapInput is an input type that accepts IAMMemberMap and IAMMemberMapOutput values.
// You can construct a concrete instance of `IAMMemberMapInput` via:
//
//	IAMMemberMap{ "key": IAMMemberArgs{...} }
type IAMMemberMapInput interface {
	pulumi.Input

	ToIAMMemberMapOutput() IAMMemberMapOutput
	ToIAMMemberMapOutputWithContext(context.Context) IAMMemberMapOutput
}

type IAMMemberMap map[string]IAMMemberInput

func (IAMMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (i IAMMemberMap) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return i.ToIAMMemberMapOutputWithContext(context.Background())
}

func (i IAMMemberMap) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMMemberMapOutput)
}

type IAMMemberOutput struct{ *pulumi.OutputState }

func (IAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMMember)(nil)).Elem()
}

func (o IAMMemberOutput) ToIAMMemberOutput() IAMMemberOutput {
	return o
}

func (o IAMMemberOutput) ToIAMMemberOutputWithContext(ctx context.Context) IAMMemberOutput {
	return o
}

// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
// Structure is documented below.
func (o IAMMemberOutput) Condition() IAMMemberConditionPtrOutput {
	return o.ApplyT(func(v *IAMMember) IAMMemberConditionPtrOutput { return v.Condition }).(IAMMemberConditionPtrOutput)
}

// (Computed) The etag of the service account IAM policy.
func (o IAMMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in `role`.
// Each entry can have one of the following values:
// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o IAMMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The role that should be applied. Only one
// `serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (o IAMMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The fully-qualified name of the service account to apply policy to.
func (o IAMMemberOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMMember) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

type IAMMemberArrayOutput struct{ *pulumi.OutputState }

func (IAMMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMMember)(nil)).Elem()
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutput() IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) ToIAMMemberArrayOutputWithContext(ctx context.Context) IAMMemberArrayOutput {
	return o
}

func (o IAMMemberArrayOutput) Index(i pulumi.IntInput) IAMMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].([]*IAMMember)[vs[1].(int)]
	}).(IAMMemberOutput)
}

type IAMMemberMapOutput struct{ *pulumi.OutputState }

func (IAMMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMMember)(nil)).Elem()
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutput() IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) ToIAMMemberMapOutputWithContext(ctx context.Context) IAMMemberMapOutput {
	return o
}

func (o IAMMemberMapOutput) MapIndex(k pulumi.StringInput) IAMMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMMember {
		return vs[0].(map[string]*IAMMember)[vs[1].(string)]
	}).(IAMMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberInput)(nil)).Elem(), &IAMMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberArrayInput)(nil)).Elem(), IAMMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMMemberMapInput)(nil)).Elem(), IAMMemberMap{})
	pulumi.RegisterOutputType(IAMMemberOutput{})
	pulumi.RegisterOutputType(IAMMemberArrayOutput{})
	pulumi.RegisterOutputType(IAMMemberMapOutput{})
}
