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
        /// The name of this alert channel
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The mobile number to receive the alerts
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
