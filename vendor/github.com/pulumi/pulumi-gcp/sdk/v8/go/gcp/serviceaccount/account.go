// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows management of a Google Cloud service account.
//
// * [API documentation](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/access/service-accounts)
//
// > **Warning:**  If you delete and recreate a service account, you must reapply any IAM roles that it had before.
//
// > Creation of service accounts is eventually consistent, and that can lead to
// errors when you try to apply ACLs to service accounts immediately after
// creation.
//
// ## Example Usage
//
// This snippet creates a service account in a project.
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
//			_, err := serviceaccount.NewAccount(ctx, "service_account", &serviceaccount.AccountArgs{
//				AccountId:   pulumi.String("service-account-id"),
//				DisplayName: pulumi.String("Service Account"),
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
// Service accounts can be imported using their URI, e.g.
//
// * `projects/{{project_id}}/serviceAccounts/{{email}}`
//
// When using the `pulumi import` command, service accounts can be imported using one of the formats above. For example:
//
// ```sh
// $ pulumi import gcp:serviceaccount/account:Account default projects/{{project_id}}/serviceAccounts/{{email}}
// ```
type Account struct {
	pulumi.CustomResourceState

	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists pulumi.BoolPtrOutput `pulumi:"createIgnoreAlreadyExists"`
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email pulumi.StringOutput `pulumi:"email"`
	// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
	Member pulumi.StringOutput `pulumi:"member"`
	// The fully-qualified name of the service account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringOutput `pulumi:"project"`
	// The unique id of the service account.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		args = &AccountArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("gcp:serviceAccount/account:Account"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Account
	err := ctx.RegisterResource("gcp:serviceaccount/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("gcp:serviceaccount/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId *string `pulumi:"accountId"`
	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists *bool `pulumi:"createIgnoreAlreadyExists"`
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description *string `pulumi:"description"`
	// Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	Disabled *bool `pulumi:"disabled"`
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName *string `pulumi:"displayName"`
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email *string `pulumi:"email"`
	// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
	Member *string `pulumi:"member"`
	// The fully-qualified name of the service account.
	Name *string `pulumi:"name"`
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
	// The unique id of the service account.
	UniqueId *string `pulumi:"uniqueId"`
}

type AccountState struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId pulumi.StringPtrInput
	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists pulumi.BoolPtrInput
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description pulumi.StringPtrInput
	// Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	Disabled pulumi.BoolPtrInput
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName pulumi.StringPtrInput
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email pulumi.StringPtrInput
	// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
	Member pulumi.StringPtrInput
	// The fully-qualified name of the service account.
	Name pulumi.StringPtrInput
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringPtrInput
	// The unique id of the service account.
	UniqueId pulumi.StringPtrInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId *string `pulumi:"accountId"`
	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists *bool `pulumi:"createIgnoreAlreadyExists"`
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description *string `pulumi:"description"`
	// Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	Disabled *bool `pulumi:"disabled"`
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId pulumi.StringPtrInput
	// If set to true, skip service account creation if a service account with the same email already exists.
	CreateIgnoreAlreadyExists pulumi.BoolPtrInput
	// A text description of the service account.
	// Must be less than or equal to 256 UTF-8 bytes.
	Description pulumi.StringPtrInput
	// Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
	// Must be set after creation to disable a service account.
	Disabled pulumi.BoolPtrInput
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName pulumi.StringPtrInput
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project pulumi.StringPtrInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

// AccountArrayInput is an input type that accepts AccountArray and AccountArrayOutput values.
// You can construct a concrete instance of `AccountArrayInput` via:
//
//	AccountArray{ AccountArgs{...} }
type AccountArrayInput interface {
	pulumi.Input

	ToAccountArrayOutput() AccountArrayOutput
	ToAccountArrayOutputWithContext(context.Context) AccountArrayOutput
}

type AccountArray []AccountInput

func (AccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (i AccountArray) ToAccountArrayOutput() AccountArrayOutput {
	return i.ToAccountArrayOutputWithContext(context.Background())
}

func (i AccountArray) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountArrayOutput)
}

// AccountMapInput is an input type that accepts AccountMap and AccountMapOutput values.
// You can construct a concrete instance of `AccountMapInput` via:
//
//	AccountMap{ "key": AccountArgs{...} }
type AccountMapInput interface {
	pulumi.Input

	ToAccountMapOutput() AccountMapOutput
	ToAccountMapOutputWithContext(context.Context) AccountMapOutput
}

type AccountMap map[string]AccountInput

func (AccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (i AccountMap) ToAccountMapOutput() AccountMapOutput {
	return i.ToAccountMapOutputWithContext(context.Background())
}

func (i AccountMap) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountMapOutput)
}

type AccountOutput struct{ *pulumi.OutputState }

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Account)(nil)).Elem()
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

// The account id that is used to generate the service
// account email address and a stable unique id. It is unique within a project,
// must be 6-30 characters long, and match the regular expression `a-z`
// to comply with RFC1035. Changing this forces a new service account to be created.
func (o AccountOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// If set to true, skip service account creation if a service account with the same email already exists.
func (o AccountOutput) CreateIgnoreAlreadyExists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.BoolPtrOutput { return v.CreateIgnoreAlreadyExists }).(pulumi.BoolPtrOutput)
}

// A text description of the service account.
// Must be less than or equal to 256 UTF-8 bytes.
func (o AccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether a service account is disabled or not. Defaults to `false`. This field has no effect during creation.
// Must be set after creation to disable a service account.
func (o AccountOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The display name for the service account.
// Can be updated without creating a new resource.
func (o AccountOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Account) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The e-mail address of the service account. This value
// should be referenced from any `organizations.getIAMPolicy` data sources
// that would grant the service account privileges.
func (o AccountOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
func (o AccountOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The fully-qualified name of the service account.
func (o AccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project that the service account will be created in.
// Defaults to the provider project configuration.
func (o AccountOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The unique id of the service account.
func (o AccountOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Account) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

type AccountArrayOutput struct{ *pulumi.OutputState }

func (AccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Account)(nil)).Elem()
}

func (o AccountArrayOutput) ToAccountArrayOutput() AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) ToAccountArrayOutputWithContext(ctx context.Context) AccountArrayOutput {
	return o
}

func (o AccountArrayOutput) Index(i pulumi.IntInput) AccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Account {
		return vs[0].([]*Account)[vs[1].(int)]
	}).(AccountOutput)
}

type AccountMapOutput struct{ *pulumi.OutputState }

func (AccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Account)(nil)).Elem()
}

func (o AccountMapOutput) ToAccountMapOutput() AccountMapOutput {
	return o
}

func (o AccountMapOutput) ToAccountMapOutputWithContext(ctx context.Context) AccountMapOutput {
	return o
}

func (o AccountMapOutput) MapIndex(k pulumi.StringInput) AccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Account {
		return vs[0].(map[string]*Account)[vs[1].(string)]
	}).(AccountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountInput)(nil)).Elem(), &Account{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountArrayInput)(nil)).Elem(), AccountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountMapInput)(nil)).Elem(), AccountMap{})
	pulumi.RegisterOutputType(AccountOutput{})
	pulumi.RegisterOutputType(AccountArrayOutput{})
	pulumi.RegisterOutputType(AccountMapOutput{})
}
