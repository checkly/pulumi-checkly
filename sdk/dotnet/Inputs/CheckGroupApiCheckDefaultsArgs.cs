// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class CheckGroupApiCheckDefaultsArgs : global::Pulumi.ResourceArgs
    {
        [Input("assertions")]
        private InputList<Inputs.CheckGroupApiCheckDefaultsAssertionArgs>? _assertions;
        public InputList<Inputs.CheckGroupApiCheckDefaultsAssertionArgs> Assertions
        {
            get => _assertions ?? (_assertions = new InputList<Inputs.CheckGroupApiCheckDefaultsAssertionArgs>());
            set => _assertions = value;
        }

        [Input("basicAuth")]
        public Input<Inputs.CheckGroupApiCheckDefaultsBasicAuthArgs>? BasicAuth { get; set; }

        [Input("headers")]
        private InputMap<object>? _headers;
        public InputMap<object> Headers
        {
            get => _headers ?? (_headers = new InputMap<object>());
            set => _headers = value;
        }

        [Input("queryParameters")]
        private InputMap<object>? _queryParameters;
        public InputMap<object> QueryParameters
        {
            get => _queryParameters ?? (_queryParameters = new InputMap<object>());
            set => _queryParameters = value;
        }

        /// <summary>
        /// The base url for this group which you can reference with the `GROUP_BASE_URL` variable in all group checks.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public CheckGroupApiCheckDefaultsArgs()
        {
            Url = "";
        }
        public static new CheckGroupApiCheckDefaultsArgs Empty => new CheckGroupApiCheckDefaultsArgs();
    }
}
