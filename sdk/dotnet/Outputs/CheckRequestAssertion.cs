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
    public sealed class CheckRequestAssertion
    {
        public readonly string Comparison;
        public readonly string? Property;
        public readonly string Source;
        public readonly string? Target;

        [OutputConstructor]
        private CheckRequestAssertion(
            string comparison,

            string? property,

            string source,

            string? target)
        {
            Comparison = comparison;
            Property = property;
            Source = source;
            Target = target;
        }
    }
}
