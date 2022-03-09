// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class AlertChannelWebhookGetArgs : Pulumi.ResourceArgs
    {
        [Input("headers")]
        private InputMap<object>? _headers;
        public InputMap<object> Headers
        {
            get => _headers ?? (_headers = new InputMap<object>());
            set => _headers = value;
        }

        [Input("method")]
        public Input<string>? Method { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("queryParameters")]
        private InputMap<object>? _queryParameters;
        public InputMap<object> QueryParameters
        {
            get => _queryParameters ?? (_queryParameters = new InputMap<object>());
            set => _queryParameters = value;
        }

        [Input("template")]
        public Input<string>? Template { get; set; }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("webhookSecret")]
        public Input<string>? WebhookSecret { get; set; }

        public AlertChannelWebhookGetArgs()
        {
        }
    }
}