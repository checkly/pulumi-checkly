// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class HeartbeatCheckAlertSettingsRunBasedEscalationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// After how many failed consecutive check runs an alert notification should be sent. Possible values are between 1 and 5. (Default `1`).
        /// </summary>
        [Input("failedRunThreshold")]
        public Input<int>? FailedRunThreshold { get; set; }

        public HeartbeatCheckAlertSettingsRunBasedEscalationGetArgs()
        {
        }
        public static new HeartbeatCheckAlertSettingsRunBasedEscalationGetArgs Empty => new HeartbeatCheckAlertSettingsRunBasedEscalationGetArgs();
    }
}
