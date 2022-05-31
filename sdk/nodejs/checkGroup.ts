// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class CheckGroup extends pulumi.CustomResource {
    /**
     * Get an existing CheckGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CheckGroupState, opts?: pulumi.CustomResourceOptions): CheckGroup {
        return new CheckGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'checkly:index/checkGroup:CheckGroup';

    /**
     * Returns true if the given object is an instance of CheckGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CheckGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CheckGroup.__pulumiType;
    }

    /**
     * Determines if the checks in the group are running or not.
     */
    public readonly activated!: pulumi.Output<boolean>;
    public readonly alertChannelSubscriptions!: pulumi.Output<outputs.CheckGroupAlertChannelSubscription[] | undefined>;
    public readonly alertSettings!: pulumi.Output<outputs.CheckGroupAlertSettings>;
    public readonly apiCheckDefaults!: pulumi.Output<outputs.CheckGroupApiCheckDefaults>;
    /**
     * Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
     */
    public readonly concurrency!: pulumi.Output<number>;
    /**
     * Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected
     * region before marking the check as failed.
     */
    public readonly doubleCheck!: pulumi.Output<boolean | undefined>;
    /**
     * Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks.
     * Use global environment variables whenever possible.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A valid piece of Node.js code to run in the setup phase of an API check in this group.
     */
    public readonly localSetupScript!: pulumi.Output<string | undefined>;
    /**
     * A valid piece of Node.js code to run in the teardown phase of an API check in this group.
     */
    public readonly localTeardownScript!: pulumi.Output<string | undefined>;
    /**
     * An array of one or more data center locations where to run the checks.
     */
    public readonly locations!: pulumi.Output<string[]>;
    /**
     * Determines if any notifications will be sent out when a check in this group fails and/or recovers.
     */
    public readonly muted!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the check group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the runtime to use for this group.
     */
    public readonly runtimeId!: pulumi.Output<string | undefined>;
    /**
     * An ID reference to a snippet to use in the setup phase of an API check.
     */
    public readonly setupSnippetId!: pulumi.Output<number | undefined>;
    /**
     * Tags for organizing and filtering checks.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * An ID reference to a snippet to use in the teardown phase of an API check.
     */
    public readonly teardownSnippetId!: pulumi.Output<number | undefined>;
    /**
     * When true, the account level alert settings will be used, not the alert setting defined on this check group.
     */
    public readonly useGlobalAlertSettings!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CheckGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CheckGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CheckGroupArgs | CheckGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CheckGroupState | undefined;
            resourceInputs["activated"] = state ? state.activated : undefined;
            resourceInputs["alertChannelSubscriptions"] = state ? state.alertChannelSubscriptions : undefined;
            resourceInputs["alertSettings"] = state ? state.alertSettings : undefined;
            resourceInputs["apiCheckDefaults"] = state ? state.apiCheckDefaults : undefined;
            resourceInputs["concurrency"] = state ? state.concurrency : undefined;
            resourceInputs["doubleCheck"] = state ? state.doubleCheck : undefined;
            resourceInputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            resourceInputs["localSetupScript"] = state ? state.localSetupScript : undefined;
            resourceInputs["localTeardownScript"] = state ? state.localTeardownScript : undefined;
            resourceInputs["locations"] = state ? state.locations : undefined;
            resourceInputs["muted"] = state ? state.muted : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["runtimeId"] = state ? state.runtimeId : undefined;
            resourceInputs["setupSnippetId"] = state ? state.setupSnippetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["teardownSnippetId"] = state ? state.teardownSnippetId : undefined;
            resourceInputs["useGlobalAlertSettings"] = state ? state.useGlobalAlertSettings : undefined;
        } else {
            const args = argsOrState as CheckGroupArgs | undefined;
            if ((!args || args.activated === undefined) && !opts.urn) {
                throw new Error("Missing required property 'activated'");
            }
            if ((!args || args.concurrency === undefined) && !opts.urn) {
                throw new Error("Missing required property 'concurrency'");
            }
            if ((!args || args.locations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'locations'");
            }
            resourceInputs["activated"] = args ? args.activated : undefined;
            resourceInputs["alertChannelSubscriptions"] = args ? args.alertChannelSubscriptions : undefined;
            resourceInputs["alertSettings"] = args ? args.alertSettings : undefined;
            resourceInputs["apiCheckDefaults"] = args ? (args.apiCheckDefaults ? pulumi.output(args.apiCheckDefaults).apply(inputs.checkGroupApiCheckDefaultsProvideDefaults) : undefined) : undefined;
            resourceInputs["concurrency"] = args ? args.concurrency : undefined;
            resourceInputs["doubleCheck"] = args ? args.doubleCheck : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["localSetupScript"] = args ? args.localSetupScript : undefined;
            resourceInputs["localTeardownScript"] = args ? args.localTeardownScript : undefined;
            resourceInputs["locations"] = args ? args.locations : undefined;
            resourceInputs["muted"] = args ? args.muted : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["runtimeId"] = args ? args.runtimeId : undefined;
            resourceInputs["setupSnippetId"] = args ? args.setupSnippetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["teardownSnippetId"] = args ? args.teardownSnippetId : undefined;
            resourceInputs["useGlobalAlertSettings"] = args ? args.useGlobalAlertSettings : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CheckGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CheckGroup resources.
 */
export interface CheckGroupState {
    /**
     * Determines if the checks in the group are running or not.
     */
    activated?: pulumi.Input<boolean>;
    alertChannelSubscriptions?: pulumi.Input<pulumi.Input<inputs.CheckGroupAlertChannelSubscription>[]>;
    alertSettings?: pulumi.Input<inputs.CheckGroupAlertSettings>;
    apiCheckDefaults?: pulumi.Input<inputs.CheckGroupApiCheckDefaults>;
    /**
     * Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
     */
    concurrency?: pulumi.Input<number>;
    /**
     * Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected
     * region before marking the check as failed.
     */
    doubleCheck?: pulumi.Input<boolean>;
    /**
     * Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks.
     * Use global environment variables whenever possible.
     */
    environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * A valid piece of Node.js code to run in the setup phase of an API check in this group.
     */
    localSetupScript?: pulumi.Input<string>;
    /**
     * A valid piece of Node.js code to run in the teardown phase of an API check in this group.
     */
    localTeardownScript?: pulumi.Input<string>;
    /**
     * An array of one or more data center locations where to run the checks.
     */
    locations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines if any notifications will be sent out when a check in this group fails and/or recovers.
     */
    muted?: pulumi.Input<boolean>;
    /**
     * The name of the check group.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the runtime to use for this group.
     */
    runtimeId?: pulumi.Input<string>;
    /**
     * An ID reference to a snippet to use in the setup phase of an API check.
     */
    setupSnippetId?: pulumi.Input<number>;
    /**
     * Tags for organizing and filtering checks.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An ID reference to a snippet to use in the teardown phase of an API check.
     */
    teardownSnippetId?: pulumi.Input<number>;
    /**
     * When true, the account level alert settings will be used, not the alert setting defined on this check group.
     */
    useGlobalAlertSettings?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a CheckGroup resource.
 */
export interface CheckGroupArgs {
    /**
     * Determines if the checks in the group are running or not.
     */
    activated: pulumi.Input<boolean>;
    alertChannelSubscriptions?: pulumi.Input<pulumi.Input<inputs.CheckGroupAlertChannelSubscription>[]>;
    alertSettings?: pulumi.Input<inputs.CheckGroupAlertSettings>;
    apiCheckDefaults?: pulumi.Input<inputs.CheckGroupApiCheckDefaults>;
    /**
     * Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
     */
    concurrency: pulumi.Input<number>;
    /**
     * Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected
     * region before marking the check as failed.
     */
    doubleCheck?: pulumi.Input<boolean>;
    /**
     * Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks.
     * Use global environment variables whenever possible.
     */
    environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * A valid piece of Node.js code to run in the setup phase of an API check in this group.
     */
    localSetupScript?: pulumi.Input<string>;
    /**
     * A valid piece of Node.js code to run in the teardown phase of an API check in this group.
     */
    localTeardownScript?: pulumi.Input<string>;
    /**
     * An array of one or more data center locations where to run the checks.
     */
    locations: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Determines if any notifications will be sent out when a check in this group fails and/or recovers.
     */
    muted?: pulumi.Input<boolean>;
    /**
     * The name of the check group.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the runtime to use for this group.
     */
    runtimeId?: pulumi.Input<string>;
    /**
     * An ID reference to a snippet to use in the setup phase of an API check.
     */
    setupSnippetId?: pulumi.Input<number>;
    /**
     * Tags for organizing and filtering checks.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An ID reference to a snippet to use in the teardown phase of an API check.
     */
    teardownSnippetId?: pulumi.Input<number>;
    /**
     * When true, the account level alert settings will be used, not the alert setting defined on this check group.
     */
    useGlobalAlertSettings?: pulumi.Input<boolean>;
}
