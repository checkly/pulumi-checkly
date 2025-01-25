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
    public sealed class CheckRequest
    {
        /// <summary>
        /// A request can have multiple assertions.
        /// </summary>
        public readonly ImmutableArray<Outputs.CheckRequestAssertion> Assertions;
        /// <summary>
        /// Set up HTTP basic authentication (username &amp; password).
        /// </summary>
        public readonly Outputs.CheckRequestBasicAuth? BasicAuth;
        /// <summary>
        /// The body of the request.
        /// </summary>
        public readonly string? Body;
        /// <summary>
        /// The `Content-Type` header of the request. Possible values `NONE`, `JSON`, `FORM`, `RAW`, and `GRAPHQL`.
        /// </summary>
        public readonly string? BodyType;
        public readonly bool? FollowRedirects;
        public readonly ImmutableDictionary<string, string>? Headers;
        /// <summary>
        /// IP Family to be used when executing the api check. The value can be either IPv4 or IPv6.
        /// </summary>
        public readonly string? IpFamily;
        /// <summary>
        /// The HTTP method to use for this API check. Possible values are `GET`, `POST`, `PUT`, `HEAD`, `DELETE`, `PATCH`. (Default `GET`).
        /// </summary>
        public readonly string? Method;
        public readonly ImmutableDictionary<string, string>? QueryParameters;
        public readonly bool? SkipSsl;
        public readonly string Url;

        [OutputConstructor]
        private CheckRequest(
            ImmutableArray<Outputs.CheckRequestAssertion> assertions,

            Outputs.CheckRequestBasicAuth? basicAuth,

            string? body,

            string? bodyType,

            bool? followRedirects,

            ImmutableDictionary<string, string>? headers,

            string? ipFamily,

            string? method,

            ImmutableDictionary<string, string>? queryParameters,

            bool? skipSsl,

            string url)
        {
            Assertions = assertions;
            BasicAuth = basicAuth;
            Body = body;
            BodyType = bodyType;
            FollowRedirects = followRedirects;
            Headers = headers;
            IpFamily = ipFamily;
            Method = method;
            QueryParameters = queryParameters;
            SkipSsl = skipSsl;
            Url = url;
        }
    }
}
