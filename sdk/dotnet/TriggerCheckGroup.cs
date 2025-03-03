// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Checkly = Pulumi.Checkly;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testTriggerGroup = new Checkly.TriggerCheckGroup("test_trigger_group", new()
    ///     {
    ///         GroupId = 215,
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["testTriggerGroupUrl"] = testTriggerGroup.Url,
    ///     };
    /// });
    /// ```
    /// </summary>
    [ChecklyResourceType("checkly:index/triggerCheckGroup:TriggerCheckGroup")]
    public partial class TriggerCheckGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the group that you want to attach the trigger to.
        /// </summary>
        [Output("groupId")]
        public Output<int> GroupId { get; private set; } = null!;

        /// <summary>
        /// The token value created to trigger the group
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The request URL to trigger the group run.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a TriggerCheckGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TriggerCheckGroup(string name, TriggerCheckGroupArgs args, CustomResourceOptions? options = null)
            : base("checkly:index/triggerCheckGroup:TriggerCheckGroup", name, args ?? new TriggerCheckGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TriggerCheckGroup(string name, Input<string> id, TriggerCheckGroupState? state = null, CustomResourceOptions? options = null)
            : base("checkly:index/triggerCheckGroup:TriggerCheckGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/checkly",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TriggerCheckGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TriggerCheckGroup Get(string name, Input<string> id, TriggerCheckGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new TriggerCheckGroup(name, id, state, options);
        }
    }

    public sealed class TriggerCheckGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the group that you want to attach the trigger to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<int> GroupId { get; set; } = null!;

        /// <summary>
        /// The token value created to trigger the group
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The request URL to trigger the group run.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public TriggerCheckGroupArgs()
        {
        }
        public static new TriggerCheckGroupArgs Empty => new TriggerCheckGroupArgs();
    }

    public sealed class TriggerCheckGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the group that you want to attach the trigger to.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// The token value created to trigger the group
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The request URL to trigger the group run.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public TriggerCheckGroupState()
        {
        }
        public static new TriggerCheckGroupState Empty => new TriggerCheckGroupState();
    }
}
