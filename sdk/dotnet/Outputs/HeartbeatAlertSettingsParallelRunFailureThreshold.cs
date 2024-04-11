// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Outputs
{

    [OutputType]
    public sealed class HeartbeatAlertSettingsParallelRunFailureThreshold
    {
        /// <summary>
        /// Applicable only for checks scheduled in parallel in multiple locations.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Possible values are `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `100`, and `100`. (Default `10`).
        /// </summary>
        public readonly int? Percentage;

        [OutputConstructor]
        private HeartbeatAlertSettingsParallelRunFailureThreshold(
            bool? enabled,

            int? percentage)
        {
            Enabled = enabled;
            Percentage = percentage;
        }
    }
}
