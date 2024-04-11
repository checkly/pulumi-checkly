// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class CheckAlertSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines what type of escalation to use. Possible values are `RUN_BASED` or `TIME_BASED`.
        /// </summary>
        [Input("escalationType")]
        public Input<string>? EscalationType { get; set; }

        [Input("parallelRunFailureThresholds")]
        private InputList<Inputs.CheckAlertSettingsParallelRunFailureThresholdGetArgs>? _parallelRunFailureThresholds;
        public InputList<Inputs.CheckAlertSettingsParallelRunFailureThresholdGetArgs> ParallelRunFailureThresholds
        {
            get => _parallelRunFailureThresholds ?? (_parallelRunFailureThresholds = new InputList<Inputs.CheckAlertSettingsParallelRunFailureThresholdGetArgs>());
            set => _parallelRunFailureThresholds = value;
        }

        [Input("reminders")]
        private InputList<Inputs.CheckAlertSettingsReminderGetArgs>? _reminders;
        public InputList<Inputs.CheckAlertSettingsReminderGetArgs> Reminders
        {
            get => _reminders ?? (_reminders = new InputList<Inputs.CheckAlertSettingsReminderGetArgs>());
            set => _reminders = value;
        }

        [Input("runBasedEscalations")]
        private InputList<Inputs.CheckAlertSettingsRunBasedEscalationGetArgs>? _runBasedEscalations;
        public InputList<Inputs.CheckAlertSettingsRunBasedEscalationGetArgs> RunBasedEscalations
        {
            get => _runBasedEscalations ?? (_runBasedEscalations = new InputList<Inputs.CheckAlertSettingsRunBasedEscalationGetArgs>());
            set => _runBasedEscalations = value;
        }

        [Input("sslCertificates")]
        private InputList<Inputs.CheckAlertSettingsSslCertificateGetArgs>? _sslCertificates;
        [Obsolete(@"This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.")]
        public InputList<Inputs.CheckAlertSettingsSslCertificateGetArgs> SslCertificates
        {
            get => _sslCertificates ?? (_sslCertificates = new InputList<Inputs.CheckAlertSettingsSslCertificateGetArgs>());
            set => _sslCertificates = value;
        }

        [Input("timeBasedEscalations")]
        private InputList<Inputs.CheckAlertSettingsTimeBasedEscalationGetArgs>? _timeBasedEscalations;
        public InputList<Inputs.CheckAlertSettingsTimeBasedEscalationGetArgs> TimeBasedEscalations
        {
            get => _timeBasedEscalations ?? (_timeBasedEscalations = new InputList<Inputs.CheckAlertSettingsTimeBasedEscalationGetArgs>());
            set => _timeBasedEscalations = value;
        }

        public CheckAlertSettingsGetArgs()
        {
        }
        public static new CheckAlertSettingsGetArgs Empty => new CheckAlertSettingsGetArgs();
    }
}
