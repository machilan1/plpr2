// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

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
	case "gcp:cloudrun/domainMapping:DomainMapping":
		r = &DomainMapping{}
	case "gcp:cloudrun/iamBinding:IamBinding":
		r = &IamBinding{}
	case "gcp:cloudrun/iamMember:IamMember":
		r = &IamMember{}
	case "gcp:cloudrun/iamPolicy:IamPolicy":
		r = &IamPolicy{}
	case "gcp:cloudrun/service:Service":
		r = &Service{}
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
		"cloudrun/domainMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"cloudrun/iamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"cloudrun/iamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"cloudrun/iamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"cloudrun/service",
		&module{version},
	)
}
