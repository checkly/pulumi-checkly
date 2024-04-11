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
    public sealed class CheckGroupApiCheckDefaults
    {
        public readonly ImmutableArray<Outputs.CheckGroupApiCheckDefaultsAssertion> Assertions;
        public readonly Outputs.CheckGroupApiCheckDefaultsBasicAuth? BasicAuth;
        public readonly ImmutableDictionary<string, object>? Headers;
        public readonly ImmutableDictionary<string, object>? QueryParameters;
        /// <summary>
        /// The base url for this group which you can reference with the `GROUP_BASE_URL` variable in all group checks.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private CheckGroupApiCheckDefaults(
            ImmutableArray<Outputs.CheckGroupApiCheckDefaultsAssertion> assertions,

            Outputs.CheckGroupApiCheckDefaultsBasicAuth? basicAuth,

            ImmutableDictionary<string, object>? headers,

            ImmutableDictionary<string, object>? queryParameters,

            string? url)
        {
            Assertions = assertions;
            BasicAuth = basicAuth;
            Headers = headers;
            QueryParameters = queryParameters;
            Url = url;
        }
    }
}
