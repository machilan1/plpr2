// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gcp:serviceaccount/account:Account":
		r = &Account{}
	case "gcp:serviceaccount/iAMBinding:IAMBinding":
		r = &IAMBinding{}
	case "gcp:serviceaccount/iAMMember:IAMMember":
		r = &IAMMember{}
	case "gcp:serviceaccount/iAMPolicy:IAMPolicy":
		r = &IAMPolicy{}
	case "gcp:serviceaccount/key:Key":
		r = &Key{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"serviceaccount/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"serviceaccount/iAMBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"serviceaccount/iAMMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"serviceaccount/iAMPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"serviceaccount/key",
		&module{version},
	)
}
