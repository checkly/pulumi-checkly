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
    public sealed class AlertChannelSms
    {
        /// <summary>
        /// Webhook's channel name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Mobile number to receive alerts.
        /// </summary>
        public readonly string Number;

        [OutputConstructor]
        private AlertChannelSms(
            string name,

            string number)
        {
            Name = name;
            Number = number;
        }
    }
}
