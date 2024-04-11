// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/checkly/pulumi-checkly/sdk/go/checkly/internal"
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
	case "checkly:index/alertChannel:AlertChannel":
		r = &AlertChannel{}
	case "checkly:index/check:Check":
		r = &Check{}
	case "checkly:index/checkGroup:CheckGroup":
		r = &CheckGroup{}
	case "checkly:index/dashboard:Dashboard":
		r = &Dashboard{}
	case "checkly:index/environmentVariable:EnvironmentVariable":
		r = &EnvironmentVariable{}
	case "checkly:index/heartbeat:Heartbeat":
		r = &Heartbeat{}
	case "checkly:index/maintenanceWindow:MaintenanceWindow":
		r = &MaintenanceWindow{}
	case "checkly:index/privateLocation:PrivateLocation":
		r = &PrivateLocation{}
	case "checkly:index/snippet:Snippet":
		r = &Snippet{}
	case "checkly:index/triggerCheck:TriggerCheck":
		r = &TriggerCheck{}
	case "checkly:index/triggerCheckGroup:TriggerCheckGroup":
		r = &TriggerCheckGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:checkly" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"checkly",
		"index/alertChannel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/check",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/checkGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/dashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/environmentVariable",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/heartbeat",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/maintenanceWindow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/privateLocation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/snippet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/triggerCheck",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"checkly",
		"index/triggerCheckGroup",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"checkly",
		&pkg{version},
	)
}
