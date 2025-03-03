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
    public sealed class AlertChannelEmail
    {
        /// <summary>
        /// The email address of this email alert channel.
        /// </summary>
        public readonly string Address;

        [OutputConstructor]
        private AlertChannelEmail(string address)
        {
            Address = address;
        }
    }
}
