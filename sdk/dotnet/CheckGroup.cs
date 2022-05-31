// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly
{
    [ChecklyResourceType("checkly:index/checkGroup:CheckGroup")]
    public partial class CheckGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Determines if the checks in the group are running or not.
        /// </summary>
        [Output("activated")]
        public Output<bool> Activated { get; private set; } = null!;

        [Output("alertChannelSubscriptions")]
        public Output<ImmutableArray<Outputs.CheckGroupAlertChannelSubscription>> AlertChannelSubscriptions { get; private set; } = null!;

        [Output("alertSettings")]
        public Output<Outputs.CheckGroupAlertSettings> AlertSettings { get; private set; } = null!;

        [Output("apiCheckDefaults")]
        public Output<Outputs.CheckGroupApiCheckDefaults> ApiCheckDefaults { get; private set; } = null!;

        /// <summary>
        /// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
        /// </summary>
        [Output("concurrency")]
        public Output<int> Concurrency { get; private set; } = null!;

        /// <summary>
        /// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected
        /// region before marking the check as failed.
        /// </summary>
        [Output("doubleCheck")]
        public Output<bool?> DoubleCheck { get; private set; } = null!;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks.
        /// Use global environment variables whenever possible.
        /// </summary>
        [Output("environmentVariables")]
        public Output<ImmutableDictionary<string, object>?> EnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// A valid piece of Node.js code to run in the setup phase of an API check in this group.
        /// </summary>
        [Output("localSetupScript")]
        public Output<string?> LocalSetupScript { get; private set; } = null!;

        /// <summary>
        /// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
        /// </summary>
        [Output("localTeardownScript")]
        public Output<string?> LocalTeardownScript { get; private set; } = null!;

        /// <summary>
        /// An array of one or more data center locations where to run the checks.
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<string>> Locations { get; private set; } = null!;

        /// <summary>
        /// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
        /// </summary>
        [Output("muted")]
        public Output<bool?> Muted { get; private set; } = null!;

        /// <summary>
        /// The name of the check group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the runtime to use for this group.
        /// </summary>
        [Output("runtimeId")]
        public Output<string?> RuntimeId { get; private set; } = null!;

        /// <summary>
        /// An ID reference to a snippet to use in the setup phase of an API check.
        /// </summary>
        [Output("setupSnippetId")]
        public Output<int?> SetupSnippetId { get; private set; } = null!;

        /// <summary>
        /// Tags for organizing and filtering checks.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// An ID reference to a snippet to use in the teardown phase of an API check.
        /// </summary>
        [Output("teardownSnippetId")]
        public Output<int?> TeardownSnippetId { get; private set; } = null!;

        /// <summary>
        /// When true, the account level alert settings will be used, not the alert setting defined on this check group.
        /// </summary>
        [Output("useGlobalAlertSettings")]
        public Output<bool?> UseGlobalAlertSettings { get; private set; } = null!;


        /// <summary>
        /// Create a CheckGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CheckGroup(string name, CheckGroupArgs args, CustomResourceOptions? options = null)
            : base("checkly:index/checkGroup:CheckGroup", name, args ?? new CheckGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CheckGroup(string name, Input<string> id, CheckGroupState? state = null, CustomResourceOptions? options = null)
            : base("checkly:index/checkGroup:CheckGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/checkly/pulumi-checkly/releases/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CheckGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CheckGroup Get(string name, Input<string> id, CheckGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new CheckGroup(name, id, state, options);
        }
    }

    public sealed class CheckGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the checks in the group are running or not.
        /// </summary>
        [Input("activated", required: true)]
        public Input<bool> Activated { get; set; } = null!;

        [Input("alertChannelSubscriptions")]
        private InputList<Inputs.CheckGroupAlertChannelSubscriptionArgs>? _alertChannelSubscriptions;
        public InputList<Inputs.CheckGroupAlertChannelSubscriptionArgs> AlertChannelSubscriptions
        {
            get => _alertChannelSubscriptions ?? (_alertChannelSubscriptions = new InputList<Inputs.CheckGroupAlertChannelSubscriptionArgs>());
            set => _alertChannelSubscriptions = value;
        }

        [Input("alertSettings")]
        public Input<Inputs.CheckGroupAlertSettingsArgs>? AlertSettings { get; set; }

        [Input("apiCheckDefaults")]
        public Input<Inputs.CheckGroupApiCheckDefaultsArgs>? ApiCheckDefaults { get; set; }

        /// <summary>
        /// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
        /// </summary>
        [Input("concurrency", required: true)]
        public Input<int> Concurrency { get; set; } = null!;

        /// <summary>
        /// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected
        /// region before marking the check as failed.
        /// </summary>
        [Input("doubleCheck")]
        public Input<bool>? DoubleCheck { get; set; }

        [Input("environmentVariables")]
        private InputMap<object>? _environmentVariables;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks.
        /// Use global environment variables whenever possible.
        /// </summary>
        public InputMap<object> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<object>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// A valid piece of Node.js code to run in the setup phase of an API check in this group.
        /// </summary>
        [Input("localSetupScript")]
        public Input<string>? LocalSetupScript { get; set; }

        /// <summary>
        /// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
        /// </summary>
        [Input("localTeardownScript")]
        public Input<string>? LocalTeardownScript { get; set; }

        [Input("locations", required: true)]
        private InputList<string>? _locations;

        /// <summary>
        /// An array of one or more data center locations where to run the checks.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
        /// </summary>
        [Input("muted")]
        public Input<bool>? Muted { get; set; }

        /// <summary>
        /// The name of the check group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the runtime to use for this group.
        /// </summary>
        [Input("runtimeId")]
        public Input<string>? RuntimeId { get; set; }

        /// <summary>
        /// An ID reference to a snippet to use in the setup phase of an API check.
        /// </summary>
        [Input("setupSnippetId")]
        public Input<int>? SetupSnippetId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags for organizing and filtering checks.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// An ID reference to a snippet to use in the teardown phase of an API check.
        /// </summary>
        [Input("teardownSnippetId")]
        public Input<int>? TeardownSnippetId { get; set; }

        /// <summary>
        /// When true, the account level alert settings will be used, not the alert setting defined on this check group.
        /// </summary>
        [Input("useGlobalAlertSettings")]
        public Input<bool>? UseGlobalAlertSettings { get; set; }

        public CheckGroupArgs()
        {
        }
    }

    public sealed class CheckGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the checks in the group are running or not.
        /// </summary>
        [Input("activated")]
        public Input<bool>? Activated { get; set; }

        [Input("alertChannelSubscriptions")]
        private InputList<Inputs.CheckGroupAlertChannelSubscriptionGetArgs>? _alertChannelSubscriptions;
        public InputList<Inputs.CheckGroupAlertChannelSubscriptionGetArgs> AlertChannelSubscriptions
        {
            get => _alertChannelSubscriptions ?? (_alertChannelSubscriptions = new InputList<Inputs.CheckGroupAlertChannelSubscriptionGetArgs>());
            set => _alertChannelSubscriptions = value;
        }

        [Input("alertSettings")]
        public Input<Inputs.CheckGroupAlertSettingsGetArgs>? AlertSettings { get; set; }

        [Input("apiCheckDefaults")]
        public Input<Inputs.CheckGroupApiCheckDefaultsGetArgs>? ApiCheckDefaults { get; set; }

        /// <summary>
        /// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
        /// </summary>
        [Input("concurrency")]
        public Input<int>? Concurrency { get; set; }

        /// <summary>
        /// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected
        /// region before marking the check as failed.
        /// </summary>
        [Input("doubleCheck")]
        public Input<bool>? DoubleCheck { get; set; }

        [Input("environmentVariables")]
        private InputMap<object>? _environmentVariables;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks.
        /// Use global environment variables whenever possible.
        /// </summary>
        public InputMap<object> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<object>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// A valid piece of Node.js code to run in the setup phase of an API check in this group.
        /// </summary>
        [Input("localSetupScript")]
        public Input<string>? LocalSetupScript { get; set; }

        /// <summary>
        /// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
        /// </summary>
        [Input("localTeardownScript")]
        public Input<string>? LocalTeardownScript { get; set; }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// An array of one or more data center locations where to run the checks.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
        /// </summary>
        [Input("muted")]
        public Input<bool>? Muted { get; set; }

        /// <summary>
        /// The name of the check group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the runtime to use for this group.
        /// </summary>
        [Input("runtimeId")]
        public Input<string>? RuntimeId { get; set; }

        /// <summary>
        /// An ID reference to a snippet to use in the setup phase of an API check.
        /// </summary>
        [Input("setupSnippetId")]
        public Input<int>? SetupSnippetId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags for organizing and filtering checks.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// An ID reference to a snippet to use in the teardown phase of an API check.
        /// </summary>
        [Input("teardownSnippetId")]
        public Input<int>? TeardownSnippetId { get; set; }

        /// <summary>
        /// When true, the account level alert settings will be used, not the alert setting defined on this check group.
        /// </summary>
        [Input("useGlobalAlertSettings")]
        public Input<bool>? UseGlobalAlertSettings { get; set; }

        public CheckGroupState()
        {
        }
    }
}
