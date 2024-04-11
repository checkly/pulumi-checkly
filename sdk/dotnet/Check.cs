// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly
{
    /// <summary>
    /// Checks allows you to monitor key webapp flows, backend API's and set up alerting, so you get a notification when things break or slow down.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Checkly = Pulumi.Checkly;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Basic API Check
    ///     var exampleCheckCheck = new Checkly.Check("exampleCheckCheck", new()
    ///     {
    ///         Type = "API",
    ///         Activated = true,
    ///         ShouldFail = false,
    ///         Frequency = 1,
    ///         UseGlobalAlertSettings = true,
    ///         Locations = new[]
    ///         {
    ///             "us-west-1",
    ///         },
    ///         Request = new Checkly.Inputs.CheckRequestArgs
    ///         {
    ///             Url = "https://api.example.com/",
    ///             FollowRedirects = true,
    ///             SkipSsl = false,
    ///             Assertions = new[]
    ///             {
    ///                 new Checkly.Inputs.CheckRequestAssertionArgs
    ///                 {
    ///                     Source = "STATUS_CODE",
    ///                     Comparison = "EQUALS",
    ///                     Target = "200",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     // A more complex example using more assertions and setting alerts
    ///     var exampleCheck2 = new Checkly.Check("exampleCheck2", new()
    ///     {
    ///         Type = "API",
    ///         Activated = true,
    ///         ShouldFail = true,
    ///         Frequency = 1,
    ///         DegradedResponseTime = 5000,
    ///         MaxResponseTime = 10000,
    ///         Locations = new[]
    ///         {
    ///             "us-west-1",
    ///             "ap-northeast-1",
    ///             "ap-south-1",
    ///         },
    ///         AlertSettings = new Checkly.Inputs.CheckAlertSettingsArgs
    ///         {
    ///             EscalationType = "RUN_BASED",
    ///             RunBasedEscalations = new[]
    ///             {
    ///                 new Checkly.Inputs.CheckAlertSettingsRunBasedEscalationArgs
    ///                 {
    ///                     FailedRunThreshold = 1,
    ///                 },
    ///             },
    ///             Reminders = new[]
    ///             {
    ///                 new Checkly.Inputs.CheckAlertSettingsReminderArgs
    ///                 {
    ///                     Amount = 1,
    ///                 },
    ///             },
    ///         },
    ///         RetryStrategy = new Checkly.Inputs.CheckRetryStrategyArgs
    ///         {
    ///             Type = "FIXED",
    ///             BaseBackoffSeconds = 60,
    ///             MaxDurationSeconds = 600,
    ///             MaxRetries = 3,
    ///             SameRegion = false,
    ///         },
    ///         Request = new Checkly.Inputs.CheckRequestArgs
    ///         {
    ///             FollowRedirects = true,
    ///             SkipSsl = false,
    ///             Url = "http://api.example.com/",
    ///             QueryParameters = 
    ///             {
    ///                 { "search", "foo" },
    ///             },
    ///             Headers = 
    ///             {
    ///                 { "X-Bogus", "bogus" },
    ///             },
    ///             Assertions = new[]
    ///             {
    ///                 new Checkly.Inputs.CheckRequestAssertionArgs
    ///                 {
    ///                     Source = "JSON_BODY",
    ///                     Property = "code",
    ///                     Comparison = "HAS_VALUE",
    ///                     Target = "authentication.failed",
    ///                 },
    ///                 new Checkly.Inputs.CheckRequestAssertionArgs
    ///                 {
    ///                     Source = "STATUS_CODE",
    ///                     Property = "",
    ///                     Comparison = "EQUALS",
    ///                     Target = "401",
    ///                 },
    ///             },
    ///             BasicAuth = new Checkly.Inputs.CheckRequestBasicAuthArgs
    ///             {
    ///                 Username = "",
    ///                 Password = "",
    ///             },
    ///         },
    ///     });
    /// 
    ///     // Basic Browser  Check
    ///     var browserCheck1 = new Checkly.Check("browserCheck1", new()
    ///     {
    ///         Type = "BROWSER",
    ///         Activated = true,
    ///         ShouldFail = false,
    ///         Frequency = 10,
    ///         UseGlobalAlertSettings = true,
    ///         Locations = new[]
    ///         {
    ///             "us-west-1",
    ///         },
    ///         RuntimeId = "2023.02",
    ///         Script = @"const { expect, test } = require('@playwright/test')
    /// 
    /// test.use({ actionTimeout: 10000 })
    /// 
    /// test('visit page and take screenshot', async ({ page }) =&gt; {
    ///     const response = await page.goto(process.env.ENVIRONMENT_URL || 'https://checklyhq.com')
    ///     await page.screenshot({ path: 'screenshot.jpg' })
    ///     expect(response.status(), 'should respond with correct status code').toBeLessThan(400)
    /// })
    /// ",
    ///     });
    /// 
    ///     // Connection checks with alert channels
    ///     var emailAc1 = new Checkly.AlertChannel("emailAc1", new()
    ///     {
    ///         Email = new Checkly.Inputs.AlertChannelEmailArgs
    ///         {
    ///             Address = "info1@example.com",
    ///         },
    ///     });
    /// 
    ///     var emailAc2 = new Checkly.AlertChannel("emailAc2", new()
    ///     {
    ///         Email = new Checkly.Inputs.AlertChannelEmailArgs
    ///         {
    ///             Address = "info2@example.com",
    ///         },
    ///     });
    /// 
    ///     var exampleCheckIndex_checkCheck = new Checkly.Check("exampleCheckIndex/checkCheck", new()
    ///     {
    ///         AlertChannelSubscriptions = new[]
    ///         {
    ///             new Checkly.Inputs.CheckAlertChannelSubscriptionArgs
    ///             {
    ///                 ChannelId = emailAc1.Id,
    ///                 Activated = true,
    ///             },
    ///             new Checkly.Inputs.CheckAlertChannelSubscriptionArgs
    ///             {
    ///                 ChannelId = emailAc2.Id,
    ///                 Activated = true,
    ///             },
    ///         },
    ///     });
    /// 
    ///     // resource "checkly_check" "browser_check_1" {
    ///     //   name                      = "Example check"
    ///     //   type                      = "BROWSER"
    ///     //   activated                 = true
    ///     //   should_fail               = false
    ///     //   frequency                 = 10
    ///     //   use_global_alert_settings = true
    ///     //   locations = [
    ///     //     "us-west-1"
    ///     //   ]
    /// });
    /// ```
    /// </summary>
    [ChecklyResourceType("checkly:index/check:Check")]
    public partial class Check : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Determines if the check is running or not. Possible values `true`, and `false`.
        /// </summary>
        [Output("activated")]
        public Output<bool> Activated { get; private set; } = null!;

        /// <summary>
        /// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
        /// </summary>
        [Output("alertChannelSubscriptions")]
        public Output<ImmutableArray<Outputs.CheckAlertChannelSubscription>> AlertChannelSubscriptions { get; private set; } = null!;

        [Output("alertSettings")]
        public Output<Outputs.CheckAlertSettings> AlertSettings { get; private set; } = null!;

        /// <summary>
        /// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 30000. (Default `15000`).
        /// </summary>
        [Output("degradedResponseTime")]
        public Output<int?> DegradedResponseTime { get; private set; } = null!;

        /// <summary>
        /// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
        /// </summary>
        [Output("doubleCheck")]
        public Output<bool?> DoubleCheck { get; private set; } = null!;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden. These are only relevant for browser checks. Use global environment variables whenever possible.
        /// </summary>
        [Output("environmentVariable")]
        public Output<ImmutableArray<Outputs.CheckEnvironmentVariable>> EnvironmentVariable { get; private set; } = null!;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
        /// </summary>
        [Output("environmentVariables")]
        public Output<ImmutableDictionary<string, object>?> EnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
        /// </summary>
        [Output("frequency")]
        public Output<int> Frequency { get; private set; } = null!;

        /// <summary>
        /// This property only valid for API high frequency checks. To create a hight frequency check, the property `frequency` must be `0` and `frequency_offset` could be `10`, `20` or `30`.
        /// </summary>
        [Output("frequencyOffset")]
        public Output<int?> FrequencyOffset { get; private set; } = null!;

        /// <summary>
        /// The id of the check group this check is part of.
        /// </summary>
        [Output("groupId")]
        public Output<int?> GroupId { get; private set; } = null!;

        /// <summary>
        /// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
        /// </summary>
        [Output("groupOrder")]
        public Output<int?> GroupOrder { get; private set; } = null!;

        /// <summary>
        /// A valid piece of Node.js code to run in the setup phase.
        /// </summary>
        [Output("localSetupScript")]
        public Output<string?> LocalSetupScript { get; private set; } = null!;

        /// <summary>
        /// A valid piece of Node.js code to run in the teardown phase.
        /// </summary>
        [Output("localTeardownScript")]
        public Output<string?> LocalTeardownScript { get; private set; } = null!;

        /// <summary>
        /// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
        /// </summary>
        [Output("locations")]
        public Output<ImmutableArray<string>> Locations { get; private set; } = null!;

        /// <summary>
        /// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 30000. (Default `30000`).
        /// </summary>
        [Output("maxResponseTime")]
        public Output<int?> MaxResponseTime { get; private set; } = null!;

        /// <summary>
        /// Determines if any notifications will be sent out when a check fails/degrades/recovers.
        /// </summary>
        [Output("muted")]
        public Output<bool?> Muted { get; private set; } = null!;

        /// <summary>
        /// The name of the check.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An array of one or more private locations slugs.
        /// </summary>
        [Output("privateLocations")]
        public Output<ImmutableArray<string>> PrivateLocations { get; private set; } = null!;

        /// <summary>
        /// An API check might have one request config.
        /// </summary>
        [Output("request")]
        public Output<Outputs.CheckRequest?> Request { get; private set; } = null!;

        /// <summary>
        /// A strategy for retrying failed check runs.
        /// </summary>
        [Output("retryStrategy")]
        public Output<Outputs.CheckRetryStrategy> RetryStrategy { get; private set; } = null!;

        /// <summary>
        /// Determines if the check should run in all selected locations in parallel or round-robin.
        /// </summary>
        [Output("runParallel")]
        public Output<bool?> RunParallel { get; private set; } = null!;

        /// <summary>
        /// The id of the runtime to use for this check.
        /// </summary>
        [Output("runtimeId")]
        public Output<string?> RuntimeId { get; private set; } = null!;

        /// <summary>
        /// A valid piece of Node.js JavaScript code describing a browser interaction with the Puppeteer/Playwright framework or a reference to an external JavaScript file.
        /// </summary>
        [Output("script")]
        public Output<string?> Script { get; private set; } = null!;

        /// <summary>
        /// An ID reference to a snippet to use in the setup phase of an API check.
        /// </summary>
        [Output("setupSnippetId")]
        public Output<int?> SetupSnippetId { get; private set; } = null!;

        /// <summary>
        /// Allows to invert the behaviour of when a check is considered to fail. Allows for validating error status like 404.
        /// </summary>
        [Output("shouldFail")]
        public Output<bool?> ShouldFail { get; private set; } = null!;

        /// <summary>
        /// Determines if the SSL certificate should be validated for expiry.
        /// </summary>
        [Output("sslCheck")]
        public Output<bool?> SslCheck { get; private set; } = null!;

        /// <summary>
        /// A valid fully qualified domain name (FQDN) to check its SSL certificate.
        /// </summary>
        [Output("sslCheckDomain")]
        public Output<string?> SslCheckDomain { get; private set; } = null!;

        /// <summary>
        /// A list of tags for organizing and filtering checks.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// An ID reference to a snippet to use in the teardown phase of an API check.
        /// </summary>
        [Output("teardownSnippetId")]
        public Output<int?> TeardownSnippetId { get; private set; } = null!;

        /// <summary>
        /// Determines which type of retry strategy to use. Possible values are `FIXED`, `LINEAR`, or `EXPONENTIAL`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// When true, the account level alert settings will be used, not the alert setting defined on this check.
        /// </summary>
        [Output("useGlobalAlertSettings")]
        public Output<bool?> UseGlobalAlertSettings { get; private set; } = null!;


        /// <summary>
        /// Create a Check resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Check(string name, CheckArgs args, CustomResourceOptions? options = null)
            : base("checkly:index/check:Check", name, args ?? new CheckArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Check(string name, Input<string> id, CheckState? state = null, CustomResourceOptions? options = null)
            : base("checkly:index/check:Check", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/checkly",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Check resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Check Get(string name, Input<string> id, CheckState? state = null, CustomResourceOptions? options = null)
        {
            return new Check(name, id, state, options);
        }
    }

    public sealed class CheckArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the check is running or not. Possible values `true`, and `false`.
        /// </summary>
        [Input("activated", required: true)]
        public Input<bool> Activated { get; set; } = null!;

        [Input("alertChannelSubscriptions")]
        private InputList<Inputs.CheckAlertChannelSubscriptionArgs>? _alertChannelSubscriptions;

        /// <summary>
        /// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
        /// </summary>
        public InputList<Inputs.CheckAlertChannelSubscriptionArgs> AlertChannelSubscriptions
        {
            get => _alertChannelSubscriptions ?? (_alertChannelSubscriptions = new InputList<Inputs.CheckAlertChannelSubscriptionArgs>());
            set => _alertChannelSubscriptions = value;
        }

        [Input("alertSettings")]
        public Input<Inputs.CheckAlertSettingsArgs>? AlertSettings { get; set; }

        /// <summary>
        /// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 30000. (Default `15000`).
        /// </summary>
        [Input("degradedResponseTime")]
        public Input<int>? DegradedResponseTime { get; set; }

        /// <summary>
        /// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
        /// </summary>
        [Input("doubleCheck")]
        public Input<bool>? DoubleCheck { get; set; }

        [Input("environmentVariable")]
        private InputList<Inputs.CheckEnvironmentVariableArgs>? _environmentVariable;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden. These are only relevant for browser checks. Use global environment variables whenever possible.
        /// </summary>
        public InputList<Inputs.CheckEnvironmentVariableArgs> EnvironmentVariable
        {
            get => _environmentVariable ?? (_environmentVariable = new InputList<Inputs.CheckEnvironmentVariableArgs>());
            set => _environmentVariable = value;
        }

        [Input("environmentVariables")]
        private InputMap<object>? _environmentVariables;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
        /// </summary>
        [Obsolete(@"The property `environment_variables` is deprecated and will be removed in a future version. Consider using the new `environment_variable` list.")]
        public InputMap<object> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<object>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
        /// </summary>
        [Input("frequency", required: true)]
        public Input<int> Frequency { get; set; } = null!;

        /// <summary>
        /// This property only valid for API high frequency checks. To create a hight frequency check, the property `frequency` must be `0` and `frequency_offset` could be `10`, `20` or `30`.
        /// </summary>
        [Input("frequencyOffset")]
        public Input<int>? FrequencyOffset { get; set; }

        /// <summary>
        /// The id of the check group this check is part of.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
        /// </summary>
        [Input("groupOrder")]
        public Input<int>? GroupOrder { get; set; }

        /// <summary>
        /// A valid piece of Node.js code to run in the setup phase.
        /// </summary>
        [Input("localSetupScript")]
        public Input<string>? LocalSetupScript { get; set; }

        /// <summary>
        /// A valid piece of Node.js code to run in the teardown phase.
        /// </summary>
        [Input("localTeardownScript")]
        public Input<string>? LocalTeardownScript { get; set; }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 30000. (Default `30000`).
        /// </summary>
        [Input("maxResponseTime")]
        public Input<int>? MaxResponseTime { get; set; }

        /// <summary>
        /// Determines if any notifications will be sent out when a check fails/degrades/recovers.
        /// </summary>
        [Input("muted")]
        public Input<bool>? Muted { get; set; }

        /// <summary>
        /// The name of the check.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateLocations")]
        private InputList<string>? _privateLocations;

        /// <summary>
        /// An array of one or more private locations slugs.
        /// </summary>
        public InputList<string> PrivateLocations
        {
            get => _privateLocations ?? (_privateLocations = new InputList<string>());
            set => _privateLocations = value;
        }

        /// <summary>
        /// An API check might have one request config.
        /// </summary>
        [Input("request")]
        public Input<Inputs.CheckRequestArgs>? Request { get; set; }

        /// <summary>
        /// A strategy for retrying failed check runs.
        /// </summary>
        [Input("retryStrategy")]
        public Input<Inputs.CheckRetryStrategyArgs>? RetryStrategy { get; set; }

        /// <summary>
        /// Determines if the check should run in all selected locations in parallel or round-robin.
        /// </summary>
        [Input("runParallel")]
        public Input<bool>? RunParallel { get; set; }

        /// <summary>
        /// The id of the runtime to use for this check.
        /// </summary>
        [Input("runtimeId")]
        public Input<string>? RuntimeId { get; set; }

        /// <summary>
        /// A valid piece of Node.js JavaScript code describing a browser interaction with the Puppeteer/Playwright framework or a reference to an external JavaScript file.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// An ID reference to a snippet to use in the setup phase of an API check.
        /// </summary>
        [Input("setupSnippetId")]
        public Input<int>? SetupSnippetId { get; set; }

        /// <summary>
        /// Allows to invert the behaviour of when a check is considered to fail. Allows for validating error status like 404.
        /// </summary>
        [Input("shouldFail")]
        public Input<bool>? ShouldFail { get; set; }

        /// <summary>
        /// Determines if the SSL certificate should be validated for expiry.
        /// </summary>
        [Input("sslCheck")]
        public Input<bool>? SslCheck { get; set; }

        /// <summary>
        /// A valid fully qualified domain name (FQDN) to check its SSL certificate.
        /// </summary>
        [Input("sslCheckDomain")]
        public Input<string>? SslCheckDomain { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags for organizing and filtering checks.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// An ID reference to a snippet to use in the teardown phase of an API check.
        /// </summary>
        [Input("teardownSnippetId")]
        public Input<int>? TeardownSnippetId { get; set; }

        /// <summary>
        /// Determines which type of retry strategy to use. Possible values are `FIXED`, `LINEAR`, or `EXPONENTIAL`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// When true, the account level alert settings will be used, not the alert setting defined on this check.
        /// </summary>
        [Input("useGlobalAlertSettings")]
        public Input<bool>? UseGlobalAlertSettings { get; set; }

        public CheckArgs()
        {
        }
        public static new CheckArgs Empty => new CheckArgs();
    }

    public sealed class CheckState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the check is running or not. Possible values `true`, and `false`.
        /// </summary>
        [Input("activated")]
        public Input<bool>? Activated { get; set; }

        [Input("alertChannelSubscriptions")]
        private InputList<Inputs.CheckAlertChannelSubscriptionGetArgs>? _alertChannelSubscriptions;

        /// <summary>
        /// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
        /// </summary>
        public InputList<Inputs.CheckAlertChannelSubscriptionGetArgs> AlertChannelSubscriptions
        {
            get => _alertChannelSubscriptions ?? (_alertChannelSubscriptions = new InputList<Inputs.CheckAlertChannelSubscriptionGetArgs>());
            set => _alertChannelSubscriptions = value;
        }

        [Input("alertSettings")]
        public Input<Inputs.CheckAlertSettingsGetArgs>? AlertSettings { get; set; }

        /// <summary>
        /// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 30000. (Default `15000`).
        /// </summary>
        [Input("degradedResponseTime")]
        public Input<int>? DegradedResponseTime { get; set; }

        /// <summary>
        /// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
        /// </summary>
        [Input("doubleCheck")]
        public Input<bool>? DoubleCheck { get; set; }

        [Input("environmentVariable")]
        private InputList<Inputs.CheckEnvironmentVariableGetArgs>? _environmentVariable;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden. These are only relevant for browser checks. Use global environment variables whenever possible.
        /// </summary>
        public InputList<Inputs.CheckEnvironmentVariableGetArgs> EnvironmentVariable
        {
            get => _environmentVariable ?? (_environmentVariable = new InputList<Inputs.CheckEnvironmentVariableGetArgs>());
            set => _environmentVariable = value;
        }

        [Input("environmentVariables")]
        private InputMap<object>? _environmentVariables;

        /// <summary>
        /// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
        /// </summary>
        [Obsolete(@"The property `environment_variables` is deprecated and will be removed in a future version. Consider using the new `environment_variable` list.")]
        public InputMap<object> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<object>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
        /// </summary>
        [Input("frequency")]
        public Input<int>? Frequency { get; set; }

        /// <summary>
        /// This property only valid for API high frequency checks. To create a hight frequency check, the property `frequency` must be `0` and `frequency_offset` could be `10`, `20` or `30`.
        /// </summary>
        [Input("frequencyOffset")]
        public Input<int>? FrequencyOffset { get; set; }

        /// <summary>
        /// The id of the check group this check is part of.
        /// </summary>
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        /// <summary>
        /// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
        /// </summary>
        [Input("groupOrder")]
        public Input<int>? GroupOrder { get; set; }

        /// <summary>
        /// A valid piece of Node.js code to run in the setup phase.
        /// </summary>
        [Input("localSetupScript")]
        public Input<string>? LocalSetupScript { get; set; }

        /// <summary>
        /// A valid piece of Node.js code to run in the teardown phase.
        /// </summary>
        [Input("localTeardownScript")]
        public Input<string>? LocalTeardownScript { get; set; }

        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        /// <summary>
        /// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 30000. (Default `30000`).
        /// </summary>
        [Input("maxResponseTime")]
        public Input<int>? MaxResponseTime { get; set; }

        /// <summary>
        /// Determines if any notifications will be sent out when a check fails/degrades/recovers.
        /// </summary>
        [Input("muted")]
        public Input<bool>? Muted { get; set; }

        /// <summary>
        /// The name of the check.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateLocations")]
        private InputList<string>? _privateLocations;

        /// <summary>
        /// An array of one or more private locations slugs.
        /// </summary>
        public InputList<string> PrivateLocations
        {
            get => _privateLocations ?? (_privateLocations = new InputList<string>());
            set => _privateLocations = value;
        }

        /// <summary>
        /// An API check might have one request config.
        /// </summary>
        [Input("request")]
        public Input<Inputs.CheckRequestGetArgs>? Request { get; set; }

        /// <summary>
        /// A strategy for retrying failed check runs.
        /// </summary>
        [Input("retryStrategy")]
        public Input<Inputs.CheckRetryStrategyGetArgs>? RetryStrategy { get; set; }

        /// <summary>
        /// Determines if the check should run in all selected locations in parallel or round-robin.
        /// </summary>
        [Input("runParallel")]
        public Input<bool>? RunParallel { get; set; }

        /// <summary>
        /// The id of the runtime to use for this check.
        /// </summary>
        [Input("runtimeId")]
        public Input<string>? RuntimeId { get; set; }

        /// <summary>
        /// A valid piece of Node.js JavaScript code describing a browser interaction with the Puppeteer/Playwright framework or a reference to an external JavaScript file.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// An ID reference to a snippet to use in the setup phase of an API check.
        /// </summary>
        [Input("setupSnippetId")]
        public Input<int>? SetupSnippetId { get; set; }

        /// <summary>
        /// Allows to invert the behaviour of when a check is considered to fail. Allows for validating error status like 404.
        /// </summary>
        [Input("shouldFail")]
        public Input<bool>? ShouldFail { get; set; }

        /// <summary>
        /// Determines if the SSL certificate should be validated for expiry.
        /// </summary>
        [Input("sslCheck")]
        public Input<bool>? SslCheck { get; set; }

        /// <summary>
        /// A valid fully qualified domain name (FQDN) to check its SSL certificate.
        /// </summary>
        [Input("sslCheckDomain")]
        public Input<string>? SslCheckDomain { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags for organizing and filtering checks.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// An ID reference to a snippet to use in the teardown phase of an API check.
        /// </summary>
        [Input("teardownSnippetId")]
        public Input<int>? TeardownSnippetId { get; set; }

        /// <summary>
        /// Determines which type of retry strategy to use. Possible values are `FIXED`, `LINEAR`, or `EXPONENTIAL`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// When true, the account level alert settings will be used, not the alert setting defined on this check.
        /// </summary>
        [Input("useGlobalAlertSettings")]
        public Input<bool>? UseGlobalAlertSettings { get; set; }

        public CheckState()
        {
        }
        public static new CheckState Empty => new CheckState();
    }
}
