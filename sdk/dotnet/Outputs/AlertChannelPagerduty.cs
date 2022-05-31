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
    public sealed class AlertChannelPagerduty
    {
        public readonly string? Account;
        public readonly string ServiceKey;
        public readonly string? ServiceName;

        [OutputConstructor]
        private AlertChannelPagerduty(
            string? account,

            string serviceKey,

            string? serviceName)
        {
            Account = account;
            ServiceKey = serviceKey;
            ServiceName = serviceName;
        }
    }
}
