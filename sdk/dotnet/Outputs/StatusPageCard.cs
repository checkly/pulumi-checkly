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
    public sealed class StatusPageCard
    {
        /// <summary>
        /// The name of the card.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of services to attach to the card.
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusPageCardServiceAttachment> ServiceAttachments;

        [OutputConstructor]
        private StatusPageCard(
            string name,

            ImmutableArray<Outputs.StatusPageCardServiceAttachment> serviceAttachments)
        {
            Name = name;
            ServiceAttachments = serviceAttachments;
        }
    }
}
