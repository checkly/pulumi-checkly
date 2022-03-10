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
    public sealed class CheckAlertSettingsSslCertificate
    {
        /// <summary>
        /// At what moment in time to start alerting on SSL certificates. Possible values `3`, `7`, `14`, `30`. Defaults to `3`.
        /// </summary>
        public readonly int? AlertThreshold;
        /// <summary>
        /// Determines if alert notifications should be send for expiring SSL certificates. Possible values `true`, and `false`. Defaults to `true`.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private CheckAlertSettingsSslCertificate(
            int? alertThreshold,

            bool? enabled)
        {
            AlertThreshold = alertThreshold;
            Enabled = enabled;
        }
    }
}
