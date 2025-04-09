// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Status page services represent functional pieces of your application or website, such as landing page, API, support portal etc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as checkly from "@checkly/pulumi";
 *
 * const backend = new checkly.StatusPageService("backend", {name: "Backend"});
 * const frontend = new checkly.StatusPageService("frontend", {name: "Frontend"});
 * ```
 */
export class StatusPageService extends pulumi.CustomResource {
    /**
     * Get an existing StatusPageService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StatusPageServiceState, opts?: pulumi.CustomResourceOptions): StatusPageService {
        return new StatusPageService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'checkly:index/statusPageService:StatusPageService';

    /**
     * Returns true if the given object is an instance of StatusPageService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StatusPageService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StatusPageService.__pulumiType;
    }

    /**
     * The name of the service.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a StatusPageService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StatusPageServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StatusPageServiceArgs | StatusPageServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StatusPageServiceState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as StatusPageServiceArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StatusPageService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StatusPageService resources.
 */
export interface StatusPageServiceState {
    /**
     * The name of the service.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StatusPageService resource.
 */
export interface StatusPageServiceArgs {
    /**
     * The name of the service.
     */
    name?: pulumi.Input<string>;
}
