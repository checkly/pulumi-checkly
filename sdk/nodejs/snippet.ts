// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as checkly from "@checkly/pulumi";
 *
 * const example1 = new checkly.Snippet("example_1", {
 *     name: "Example 1",
 *     script: "console.log('test');",
 * });
 * // An alternative way to use multi-line script.
 * const example2 = new checkly.Snippet("example_2", {
 *     name: "Example 2",
 *     script: `    console.log('test1');
 *     console.log('test2');
 * `,
 * });
 * ```
 */
export class Snippet extends pulumi.CustomResource {
    /**
     * Get an existing Snippet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnippetState, opts?: pulumi.CustomResourceOptions): Snippet {
        return new Snippet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'checkly:index/snippet:Snippet';

    /**
     * Returns true if the given object is an instance of Snippet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snippet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snippet.__pulumiType;
    }

    /**
     * The name of the snippet
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
     */
    public readonly script!: pulumi.Output<string>;

    /**
     * Create a Snippet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnippetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnippetArgs | SnippetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnippetState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["script"] = state ? state.script : undefined;
        } else {
            const args = argsOrState as SnippetArgs | undefined;
            if ((!args || args.script === undefined) && !opts.urn) {
                throw new Error("Missing required property 'script'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["script"] = args ? args.script : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Snippet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snippet resources.
 */
export interface SnippetState {
    /**
     * The name of the snippet
     */
    name?: pulumi.Input<string>;
    /**
     * Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
     */
    script?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snippet resource.
 */
export interface SnippetArgs {
    /**
     * The name of the snippet
     */
    name?: pulumi.Input<string>;
    /**
     * Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
     */
    script: pulumi.Input<string>;
}
