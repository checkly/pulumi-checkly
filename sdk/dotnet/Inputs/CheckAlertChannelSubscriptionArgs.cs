// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class CheckAlertChannelSubscriptionArgs : Pulumi.ResourceArgs
    {
        [Input("activated", required: true)]
        public Input<bool> Activated { get; set; } = null!;

        [Input("channelId", required: true)]
        public Input<int> ChannelId { get; set; } = null!;

        public CheckAlertChannelSubscriptionArgs()
        {
        }
    }
}
