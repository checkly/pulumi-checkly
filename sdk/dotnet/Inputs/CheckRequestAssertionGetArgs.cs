// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class CheckRequestAssertionGetArgs : Pulumi.ResourceArgs
    {
        [Input("comparison", required: true)]
        public Input<string> Comparison { get; set; } = null!;

        [Input("property")]
        public Input<string>? Property { get; set; }

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        [Input("target")]
        public Input<string>? Target { get; set; }

        public CheckRequestAssertionGetArgs()
        {
        }
    }
}
