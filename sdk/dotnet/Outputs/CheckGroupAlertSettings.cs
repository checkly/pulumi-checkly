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
    public sealed class CheckGroupAlertSettings
    {
        /// <summary>
        /// Determines what type of escalation to use. Possible values are `RUN_BASED` or `TIME_BASED`.
        /// </summary>
        public readonly string? EscalationType;
        public readonly ImmutableArray<Outputs.CheckGroupAlertSettingsParallelRunFailureThreshold> ParallelRunFailureThresholds;
        public readonly ImmutableArray<Outputs.CheckGroupAlertSettingsReminder> Reminders;
        public readonly ImmutableArray<Outputs.CheckGroupAlertSettingsRunBasedEscalation> RunBasedEscalations;
        public readonly ImmutableArray<Outputs.CheckGroupAlertSettingsSslCertificate> SslCertificates;
        public readonly ImmutableArray<Outputs.CheckGroupAlertSettingsTimeBasedEscalation> TimeBasedEscalations;

        [OutputConstructor]
        private CheckGroupAlertSettings(
            string? escalationType,

            ImmutableArray<Outputs.CheckGroupAlertSettingsParallelRunFailureThreshold> parallelRunFailureThresholds,

            ImmutableArray<Outputs.CheckGroupAlertSettingsReminder> reminders,

            ImmutableArray<Outputs.CheckGroupAlertSettingsRunBasedEscalation> runBasedEscalations,

            ImmutableArray<Outputs.CheckGroupAlertSettingsSslCertificate> sslCertificates,

            ImmutableArray<Outputs.CheckGroupAlertSettingsTimeBasedEscalation> timeBasedEscalations)
        {
            EscalationType = escalationType;
            ParallelRunFailureThresholds = parallelRunFailureThresholds;
            Reminders = reminders;
            RunBasedEscalations = runBasedEscalations;
            SslCertificates = sslCertificates;
            TimeBasedEscalations = timeBasedEscalations;
        }
    }
}
