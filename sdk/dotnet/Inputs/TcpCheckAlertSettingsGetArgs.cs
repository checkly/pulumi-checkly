// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class TcpCheckAlertSettingsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines what type of escalation to use. Possible values are `RUN_BASED` or `TIME_BASED`.
        /// </summary>
        [Input("escalationType")]
        public Input<string>? EscalationType { get; set; }

        [Input("parallelRunFailureThresholds")]
        private InputList<Inputs.TcpCheckAlertSettingsParallelRunFailureThresholdGetArgs>? _parallelRunFailureThresholds;
        public InputList<Inputs.TcpCheckAlertSettingsParallelRunFailureThresholdGetArgs> ParallelRunFailureThresholds
        {
            get => _parallelRunFailureThresholds ?? (_parallelRunFailureThresholds = new InputList<Inputs.TcpCheckAlertSettingsParallelRunFailureThresholdGetArgs>());
            set => _parallelRunFailureThresholds = value;
        }

        [Input("reminders")]
        private InputList<Inputs.TcpCheckAlertSettingsReminderGetArgs>? _reminders;
        public InputList<Inputs.TcpCheckAlertSettingsReminderGetArgs> Reminders
        {
            get => _reminders ?? (_reminders = new InputList<Inputs.TcpCheckAlertSettingsReminderGetArgs>());
            set => _reminders = value;
        }

        [Input("runBasedEscalations")]
        private InputList<Inputs.TcpCheckAlertSettingsRunBasedEscalationGetArgs>? _runBasedEscalations;
        public InputList<Inputs.TcpCheckAlertSettingsRunBasedEscalationGetArgs> RunBasedEscalations
        {
            get => _runBasedEscalations ?? (_runBasedEscalations = new InputList<Inputs.TcpCheckAlertSettingsRunBasedEscalationGetArgs>());
            set => _runBasedEscalations = value;
        }

        [Input("timeBasedEscalations")]
        private InputList<Inputs.TcpCheckAlertSettingsTimeBasedEscalationGetArgs>? _timeBasedEscalations;
        public InputList<Inputs.TcpCheckAlertSettingsTimeBasedEscalationGetArgs> TimeBasedEscalations
        {
            get => _timeBasedEscalations ?? (_timeBasedEscalations = new InputList<Inputs.TcpCheckAlertSettingsTimeBasedEscalationGetArgs>());
            set => _timeBasedEscalations = value;
        }

        public TcpCheckAlertSettingsGetArgs()
        {
        }
        public static new TcpCheckAlertSettingsGetArgs Empty => new TcpCheckAlertSettingsGetArgs();
    }
}
