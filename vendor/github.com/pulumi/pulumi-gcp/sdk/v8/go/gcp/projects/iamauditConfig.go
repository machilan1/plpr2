// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
//
// * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
// * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
// * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
// * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
//
// > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:** The underlying API method `projects.setIamPolicy` has a lot of constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
//
//	IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning 400 error code so please review these if you encounter errors with this resource.
//
// ## projects.IAMPolicy
//
// !> **Be careful!** You can accidentally lock yourself out of your project
//
//	using this resource. Deleting a `projects.IAMPolicy` removes access
//	from anyone without organization-level access to the project. Proceed with caution.
//	It's not recommended to use `projects.IAMPolicy` with your provider project
//	to avoid locking yourself out, and it should generally only be used with projects
//	fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
//	applying the change.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/editor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				Project:    pulumi.String("your-project-id"),
//				PolicyData: pulumi.String(admin.PolicyData),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/compute.admin",
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
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				Project:    pulumi.String("your-project-id"),
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## projects.IAMBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/container.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &projects.IAMBindingConditionArgs{
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
// ## projects.IAMMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
//				Member:  pulumi.String("user:jane@example.com"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/firebase.admin"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Condition: &projects.IAMMemberConditionArgs{
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
// ## projects.IAMAuditConfig
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMAuditConfig(ctx, "project", &projects.IAMAuditConfigArgs{
//				Project: pulumi.String("your-project-id"),
//				Service: pulumi.String("allServices"),
//				AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("ADMIN_READ"),
//					},
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("DATA_READ"),
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@example.com"),
//						},
//					},
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
// ## projects.IAMPolicy
//
// !> **Be careful!** You can accidentally lock yourself out of your project
//
//	using this resource. Deleting a `projects.IAMPolicy` removes access
//	from anyone without organization-level access to the project. Proceed with caution.
//	It's not recommended to use `projects.IAMPolicy` with your provider project
//	to avoid locking yourself out, and it should generally only be used with projects
//	fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
//	applying the change.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/editor",
//						Members: []string{
//							"user:jane@example.com",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				Project:    pulumi.String("your-project-id"),
//				PolicyData: pulumi.String(admin.PolicyData),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
//				Bindings: []organizations.GetIAMPolicyBinding{
//					{
//						Role: "roles/compute.admin",
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
//			_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
//				Project:    pulumi.String("your-project-id"),
//				PolicyData: pulumi.String(admin.PolicyData),
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
// ## projects.IAMBinding
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/container.admin"),
//				Members: pulumi.StringArray{
//					pulumi.String("user:jane@example.com"),
//				},
//				Condition: &projects.IAMBindingConditionArgs{
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
// ## projects.IAMMember
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/editor"),
//				Member:  pulumi.String("user:jane@example.com"),
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
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/firebase.admin"),
//				Member:  pulumi.String("user:jane@example.com"),
//				Condition: &projects.IAMMemberConditionArgs{
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
// ## projects.IAMAuditConfig
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := projects.NewIAMAuditConfig(ctx, "project", &projects.IAMAuditConfigArgs{
//				Project: pulumi.String("your-project-id"),
//				Service: pulumi.String("allServices"),
//				AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("ADMIN_READ"),
//					},
//					&projects.IAMAuditConfigAuditLogConfigArgs{
//						LogType: pulumi.String("DATA_READ"),
//						ExemptedMembers: pulumi.StringArray{
//							pulumi.String("user:joebloggs@example.com"),
//						},
//					},
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
// ### Importing Audit Configs
//
// An audit config can be imported into a `google_project_iam_audit_config` resource using the resource's `project_id` and the `service`, e.g:
//
// * `"{{project_id}} foo.googleapis.com"`
//
// An `import` block (Terraform v1.5.0 and later) can be used to import audit configs:
//
// tf
//
// import {
//
//	id = "{{project_id}} foo.googleapis.com"
//
//	to = google_project_iam_audit_config.default
//
// }
//
// The `pulumi import` command can also be used:
//
// ```sh
// $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig default "{{project_id}} foo.googleapis.com"
// ```
type IAMAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IAMAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringOutput `pulumi:"project"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are projects.IAMAuditConfig resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIAMAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewIAMAuditConfig(ctx *pulumi.Context,
	name string, args *IAMAuditConfigArgs, opts ...pulumi.ResourceOption) (*IAMAuditConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditLogConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AuditLogConfigs'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IAMAuditConfig
	err := ctx.RegisterResource("gcp:projects/iAMAuditConfig:IAMAuditConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMAuditConfig gets an existing IAMAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMAuditConfigState, opts ...pulumi.ResourceOption) (*IAMAuditConfig, error) {
	var resource IAMAuditConfig
	err := ctx.ReadResource("gcp:projects/iAMAuditConfig:IAMAuditConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMAuditConfig resources.
type iamauditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IAMAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// (Computed) The etag of the project's IAM policy.
	Etag *string `pulumi:"etag"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project *string `pulumi:"project"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are projects.IAMAuditConfig resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service *string `pulumi:"service"`
}

type IAMAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IAMAuditConfigAuditLogConfigArrayInput
	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringPtrInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringPtrInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are projects.IAMAuditConfig resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringPtrInput
}

func (IAMAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamauditConfigState)(nil)).Elem()
}

type iamauditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IAMAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The project id of the target project. This is not
	// inferred from the provider.
	Project string `pulumi:"project"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are projects.IAMAuditConfig resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IAMAuditConfig resource.
type IAMAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IAMAuditConfigAuditLogConfigArrayInput
	// The project id of the target project. This is not
	// inferred from the provider.
	Project pulumi.StringInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are projects.IAMAuditConfig resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringInput
}

func (IAMAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamauditConfigArgs)(nil)).Elem()
}

type IAMAuditConfigInput interface {
	pulumi.Input

	ToIAMAuditConfigOutput() IAMAuditConfigOutput
	ToIAMAuditConfigOutputWithContext(ctx context.Context) IAMAuditConfigOutput
}

func (*IAMAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAuditConfig)(nil)).Elem()
}

func (i *IAMAuditConfig) ToIAMAuditConfigOutput() IAMAuditConfigOutput {
	return i.ToIAMAuditConfigOutputWithContext(context.Background())
}

func (i *IAMAuditConfig) ToIAMAuditConfigOutputWithContext(ctx context.Context) IAMAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigOutput)
}

// IAMAuditConfigArrayInput is an input type that accepts IAMAuditConfigArray and IAMAuditConfigArrayOutput values.
// You can construct a concrete instance of `IAMAuditConfigArrayInput` via:
//
//	IAMAuditConfigArray{ IAMAuditConfigArgs{...} }
type IAMAuditConfigArrayInput interface {
	pulumi.Input

	ToIAMAuditConfigArrayOutput() IAMAuditConfigArrayOutput
	ToIAMAuditConfigArrayOutputWithContext(context.Context) IAMAuditConfigArrayOutput
}

type IAMAuditConfigArray []IAMAuditConfigInput

func (IAMAuditConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMAuditConfig)(nil)).Elem()
}

func (i IAMAuditConfigArray) ToIAMAuditConfigArrayOutput() IAMAuditConfigArrayOutput {
	return i.ToIAMAuditConfigArrayOutputWithContext(context.Background())
}

func (i IAMAuditConfigArray) ToIAMAuditConfigArrayOutputWithContext(ctx context.Context) IAMAuditConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigArrayOutput)
}

// IAMAuditConfigMapInput is an input type that accepts IAMAuditConfigMap and IAMAuditConfigMapOutput values.
// You can construct a concrete instance of `IAMAuditConfigMapInput` via:
//
//	IAMAuditConfigMap{ "key": IAMAuditConfigArgs{...} }
type IAMAuditConfigMapInput interface {
	pulumi.Input

	ToIAMAuditConfigMapOutput() IAMAuditConfigMapOutput
	ToIAMAuditConfigMapOutputWithContext(context.Context) IAMAuditConfigMapOutput
}

type IAMAuditConfigMap map[string]IAMAuditConfigInput

func (IAMAuditConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMAuditConfig)(nil)).Elem()
}

func (i IAMAuditConfigMap) ToIAMAuditConfigMapOutput() IAMAuditConfigMapOutput {
	return i.ToIAMAuditConfigMapOutputWithContext(context.Background())
}

func (i IAMAuditConfigMap) ToIAMAuditConfigMapOutputWithContext(ctx context.Context) IAMAuditConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigMapOutput)
}

type IAMAuditConfigOutput struct{ *pulumi.OutputState }

func (IAMAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAuditConfig)(nil)).Elem()
}

func (o IAMAuditConfigOutput) ToIAMAuditConfigOutput() IAMAuditConfigOutput {
	return o
}

func (o IAMAuditConfigOutput) ToIAMAuditConfigOutputWithContext(ctx context.Context) IAMAuditConfigOutput {
	return o
}

// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
func (o IAMAuditConfigOutput) AuditLogConfigs() IAMAuditConfigAuditLogConfigArrayOutput {
	return o.ApplyT(func(v *IAMAuditConfig) IAMAuditConfigAuditLogConfigArrayOutput { return v.AuditLogConfigs }).(IAMAuditConfigAuditLogConfigArrayOutput)
}

// (Computed) The etag of the project's IAM policy.
func (o IAMAuditConfigOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAuditConfig) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The project id of the target project. This is not
// inferred from the provider.
func (o IAMAuditConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAuditConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are projects.IAMAuditConfig resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
func (o IAMAuditConfigOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAuditConfig) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

type IAMAuditConfigArrayOutput struct{ *pulumi.OutputState }

func (IAMAuditConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMAuditConfig)(nil)).Elem()
}

func (o IAMAuditConfigArrayOutput) ToIAMAuditConfigArrayOutput() IAMAuditConfigArrayOutput {
	return o
}

func (o IAMAuditConfigArrayOutput) ToIAMAuditConfigArrayOutputWithContext(ctx context.Context) IAMAuditConfigArrayOutput {
	return o
}

func (o IAMAuditConfigArrayOutput) Index(i pulumi.IntInput) IAMAuditConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMAuditConfig {
		return vs[0].([]*IAMAuditConfig)[vs[1].(int)]
	}).(IAMAuditConfigOutput)
}

type IAMAuditConfigMapOutput struct{ *pulumi.OutputState }

func (IAMAuditConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMAuditConfig)(nil)).Elem()
}

func (o IAMAuditConfigMapOutput) ToIAMAuditConfigMapOutput() IAMAuditConfigMapOutput {
	return o
}

func (o IAMAuditConfigMapOutput) ToIAMAuditConfigMapOutputWithContext(ctx context.Context) IAMAuditConfigMapOutput {
	return o
}

func (o IAMAuditConfigMapOutput) MapIndex(k pulumi.StringInput) IAMAuditConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMAuditConfig {
		return vs[0].(map[string]*IAMAuditConfig)[vs[1].(string)]
	}).(IAMAuditConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAuditConfigInput)(nil)).Elem(), &IAMAuditConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAuditConfigArrayInput)(nil)).Elem(), IAMAuditConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAuditConfigMapInput)(nil)).Elem(), IAMAuditConfigMap{})
	pulumi.RegisterOutputType(IAMAuditConfigOutput{})
	pulumi.RegisterOutputType(IAMAuditConfigArrayOutput{})
	pulumi.RegisterOutputType(IAMAuditConfigMapOutput{})
}
