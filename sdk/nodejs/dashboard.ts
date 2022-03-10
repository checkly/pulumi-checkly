// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'checkly:index/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    public readonly customDomain!: pulumi.Output<string>;
    public readonly customUrl!: pulumi.Output<string>;
    public readonly header!: pulumi.Output<string>;
    public readonly hideTags!: pulumi.Output<boolean>;
    public readonly logo!: pulumi.Output<string>;
    public readonly paginate!: pulumi.Output<boolean>;
    public readonly paginationRate!: pulumi.Output<number>;
    public readonly refreshRate!: pulumi.Output<number>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public readonly width!: pulumi.Output<string | undefined>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardState | undefined;
            resourceInputs["customDomain"] = state ? state.customDomain : undefined;
            resourceInputs["customUrl"] = state ? state.customUrl : undefined;
            resourceInputs["header"] = state ? state.header : undefined;
            resourceInputs["hideTags"] = state ? state.hideTags : undefined;
            resourceInputs["logo"] = state ? state.logo : undefined;
            resourceInputs["paginate"] = state ? state.paginate : undefined;
            resourceInputs["paginationRate"] = state ? state.paginationRate : undefined;
            resourceInputs["refreshRate"] = state ? state.refreshRate : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["width"] = state ? state.width : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
            if ((!args || args.customDomain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customDomain'");
            }
            if ((!args || args.customUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customUrl'");
            }
            if ((!args || args.header === undefined) && !opts.urn) {
                throw new Error("Missing required property 'header'");
            }
            if ((!args || args.hideTags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hideTags'");
            }
            if ((!args || args.logo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'logo'");
            }
            if ((!args || args.paginate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paginate'");
            }
            if ((!args || args.paginationRate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paginationRate'");
            }
            if ((!args || args.refreshRate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'refreshRate'");
            }
            resourceInputs["customDomain"] = args ? args.customDomain : undefined;
            resourceInputs["customUrl"] = args ? args.customUrl : undefined;
            resourceInputs["header"] = args ? args.header : undefined;
            resourceInputs["hideTags"] = args ? args.hideTags : undefined;
            resourceInputs["logo"] = args ? args.logo : undefined;
            resourceInputs["paginate"] = args ? args.paginate : undefined;
            resourceInputs["paginationRate"] = args ? args.paginationRate : undefined;
            resourceInputs["refreshRate"] = args ? args.refreshRate : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["width"] = args ? args.width : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Dashboard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    customDomain?: pulumi.Input<string>;
    customUrl?: pulumi.Input<string>;
    header?: pulumi.Input<string>;
    hideTags?: pulumi.Input<boolean>;
    logo?: pulumi.Input<string>;
    paginate?: pulumi.Input<boolean>;
    paginationRate?: pulumi.Input<number>;
    refreshRate?: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    width?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    customDomain: pulumi.Input<string>;
    customUrl: pulumi.Input<string>;
    header: pulumi.Input<string>;
    hideTags: pulumi.Input<boolean>;
    logo: pulumi.Input<string>;
    paginate: pulumi.Input<boolean>;
    paginationRate: pulumi.Input<number>;
    refreshRate: pulumi.Input<number>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    width?: pulumi.Input<string>;
}
