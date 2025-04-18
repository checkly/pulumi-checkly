// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class HeartbeatCheckAlertSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines what type of escalation to use. Possible values are `RUN_BASED` or `TIME_BASED`.
        /// </summary>
        [Input("escalationType")]
        public Input<string>? EscalationType { get; set; }

        [Input("parallelRunFailureThresholds")]
        private InputList<Inputs.HeartbeatCheckAlertSettingsParallelRunFailureThresholdArgs>? _parallelRunFailureThresholds;
        public InputList<Inputs.HeartbeatCheckAlertSettingsParallelRunFailureThresholdArgs> ParallelRunFailureThresholds
        {
            get => _parallelRunFailureThresholds ?? (_parallelRunFailureThresholds = new InputList<Inputs.HeartbeatCheckAlertSettingsParallelRunFailureThresholdArgs>());
            set => _parallelRunFailureThresholds = value;
        }

        [Input("reminders")]
        private InputList<Inputs.HeartbeatCheckAlertSettingsReminderArgs>? _reminders;
        public InputList<Inputs.HeartbeatCheckAlertSettingsReminderArgs> Reminders
        {
            get => _reminders ?? (_reminders = new InputList<Inputs.HeartbeatCheckAlertSettingsReminderArgs>());
            set => _reminders = value;
        }

        [Input("runBasedEscalations")]
        private InputList<Inputs.HeartbeatCheckAlertSettingsRunBasedEscalationArgs>? _runBasedEscalations;
        public InputList<Inputs.HeartbeatCheckAlertSettingsRunBasedEscalationArgs> RunBasedEscalations
        {
            get => _runBasedEscalations ?? (_runBasedEscalations = new InputList<Inputs.HeartbeatCheckAlertSettingsRunBasedEscalationArgs>());
            set => _runBasedEscalations = value;
        }

        [Input("sslCertificates")]
        private InputList<Inputs.HeartbeatCheckAlertSettingsSslCertificateArgs>? _sslCertificates;
        [Obsolete(@"This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.")]
        public InputList<Inputs.HeartbeatCheckAlertSettingsSslCertificateArgs> SslCertificates
        {
            get => _sslCertificates ?? (_sslCertificates = new InputList<Inputs.HeartbeatCheckAlertSettingsSslCertificateArgs>());
            set => _sslCertificates = value;
        }

        [Input("timeBasedEscalations")]
        private InputList<Inputs.HeartbeatCheckAlertSettingsTimeBasedEscalationArgs>? _timeBasedEscalations;
        public InputList<Inputs.HeartbeatCheckAlertSettingsTimeBasedEscalationArgs> TimeBasedEscalations
        {
            get => _timeBasedEscalations ?? (_timeBasedEscalations = new InputList<Inputs.HeartbeatCheckAlertSettingsTimeBasedEscalationArgs>());
            set => _timeBasedEscalations = value;
        }

        public HeartbeatCheckAlertSettingsArgs()
        {
        }
        public static new HeartbeatCheckAlertSettingsArgs Empty => new HeartbeatCheckAlertSettingsArgs();
    }
}
