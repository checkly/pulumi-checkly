// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class CheckGroupRetryStrategyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds to wait before the first retry attempt.
        /// </summary>
        [Input("baseBackoffSeconds")]
        public Input<int>? BaseBackoffSeconds { get; set; }

        /// <summary>
        /// The total amount of time to continue retrying the check (maximum 600 seconds).
        /// </summary>
        [Input("maxDurationSeconds")]
        public Input<int>? MaxDurationSeconds { get; set; }

        /// <summary>
        /// The maximum number of times to retry the check. Value must be between 1 and 10.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// Whether retries should be run in the same region as the initial check run.
        /// </summary>
        [Input("sameRegion")]
        public Input<bool>? SameRegion { get; set; }

        /// <summary>
        /// Determines which type of retry strategy to use. Possible values are `FIXED`, `LINEAR`, or `EXPONENTIAL`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CheckGroupRetryStrategyArgs()
        {
        }
        public static new CheckGroupRetryStrategyArgs Empty => new CheckGroupRetryStrategyArgs();
    }
}
