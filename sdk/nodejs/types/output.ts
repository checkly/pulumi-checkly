// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

import * as utilities from "../utilities";

export interface AlertChannelCall {
    /**
     * The name of this alert channel
     */
    name: string;
    /**
     * The mobile number to receive the alerts
     */
    number: string;
}

export interface AlertChannelEmail {
    /**
     * The email address of this email alert channel.
     */
    address: string;
}

export interface AlertChannelOpsgenie {
    apiKey: string;
    name: string;
    priority: string;
    region: string;
}

export interface AlertChannelPagerduty {
    account?: string;
    serviceKey: string;
    serviceName?: string;
}

export interface AlertChannelSlack {
    /**
     * The name of the alert's Slack channel
     */
    channel: string;
    /**
     * The Slack webhook URL
     */
    url: string;
}

export interface AlertChannelSms {
    /**
     * The name of this alert channel
     */
    name: string;
    /**
     * The mobile number to receive the alerts
     */
    number: string;
}

export interface AlertChannelWebhook {
    headers: {[key: string]: any};
    /**
     * (Default `POST`)
     */
    method?: string;
    name: string;
    queryParameters: {[key: string]: any};
    template?: string;
    url: string;
    webhookSecret?: string;
    /**
     * Type of the webhook. Possible values are 'WEBHOOK*DISCORD', 'WEBHOOK*FIREHYDRANT', 'WEBHOOK*GITLAB*ALERT', 'WEBHOOK*SPIKESH', 'WEBHOOK*SPLUNK', 'WEBHOOK*MSTEAMS' and 'WEBHOOK*TELEGRAM'.
     */
    webhookType?: string;
}

export interface CheckAlertChannelSubscription {
    activated: boolean;
    channelId: number;
}

export interface CheckAlertSettings {
    /**
     * Determines what type of escalation to use. Possible values are `RUN_BASED` or `TIME_BASED`.
     */
    escalationType?: string;
    parallelRunFailureThresholds?: outputs.CheckAlertSettingsParallelRunFailureThreshold[];
    reminders?: outputs.CheckAlertSettingsReminder[];
    runBasedEscalations?: outputs.CheckAlertSettingsRunBasedEscalation[];
    /**
     * @deprecated This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.
     */
    sslCertificates?: outputs.CheckAlertSettingsSslCertificate[];
    timeBasedEscalations?: outputs.CheckAlertSettingsTimeBasedEscalation[];
}

export interface CheckAlertSettingsParallelRunFailureThreshold {
    /**
     * Applicable only for checks scheduled in parallel in multiple locations.
     */
    enabled?: boolean;
    /**
     * Possible values are `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `100`, and `100`. (Default `10`).
     */
    percentage?: number;
}

export interface CheckAlertSettingsReminder {
    /**
     * How many reminders to send out after the initial alert notification. Possible values are `0`, `1`, `2`, `3`, `4`, `5`, and `100000`
     */
    amount?: number;
    /**
     * Possible values are `5`, `10`, `15`, and `30`. (Default `5`).
     */
    interval?: number;
}

export interface CheckAlertSettingsRunBasedEscalation {
    /**
     * After how many failed consecutive check runs an alert notification should be sent. Possible values are between 1 and 5. (Default `1`).
     */
    failedRunThreshold?: number;
}

export interface CheckAlertSettingsSslCertificate {
    /**
     * How long before SSL certificate expiry to send alerts. Possible values `3`, `7`, `14`, `30`. (Default `3`).
     */
    alertThreshold?: number;
    /**
     * Applicable only for checks scheduled in parallel in multiple locations.
     */
    enabled?: boolean;
}

export interface CheckAlertSettingsTimeBasedEscalation {
    /**
     * After how many minutes after a check starts failing an alert should be sent. Possible values are `5`, `10`, `15`, and `30`. (Default `5`).
     */
    minutesFailingThreshold?: number;
}

export interface CheckEnvironmentVariable {
    key: string;
    locked?: boolean;
    value: string;
}

export interface CheckGroupAlertChannelSubscription {
    activated: boolean;
    channelId: number;
}

export interface CheckGroupAlertSettings {
    /**
     * Determines what type of escalation to use. Possible values are `RUN_BASED` or `TIME_BASED`.
     */
    escalationType?: string;
    parallelRunFailureThresholds?: outputs.CheckGroupAlertSettingsParallelRunFailureThreshold[];
    reminders?: outputs.CheckGroupAlertSettingsReminder[];
    runBasedEscalations?: outputs.CheckGroupAlertSettingsRunBasedEscalation[];
    /**
     * @deprecated This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.
     */
    sslCertificates?: outputs.CheckGroupAlertSettingsSslCertificate[];
    timeBasedEscalations?: outputs.CheckGroupAlertSettingsTimeBasedEscalation[];
}

export interface CheckGroupAlertSettingsParallelRunFailureThreshold {
    /**
     * Applicable only for checks scheduled in parallel in multiple locations.
     */
    enabled?: boolean;
    /**
     * Possible values are `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `100`, and `100`. (Default `10`).
     */
    percentage?: number;
}

export interface CheckGroupAlertSettingsReminder {
    /**
     * How many reminders to send out after the initial alert notification. Possible values are `0`, `1`, `2`, `3`, `4`, `5`, and `100000`
     */
    amount?: number;
    /**
     * Possible values are `5`, `10`, `15`, and `30`. (Default `5`).
     */
    interval?: number;
}

export interface CheckGroupAlertSettingsRunBasedEscalation {
    /**
     * After how many failed consecutive check runs an alert notification should be sent. Possible values are between 1 and 5. (Default `1`).
     */
    failedRunThreshold?: number;
}

export interface CheckGroupAlertSettingsSslCertificate {
    /**
     * At what moment in time to start alerting on SSL certificates. Possible values `3`, `7`, `14`, `30`. (Default `3`).
     */
    alertThreshold?: number;
    /**
     * Applicable only for checks scheduled in parallel in multiple locations.
     */
    enabled?: boolean;
}

export interface CheckGroupAlertSettingsTimeBasedEscalation {
    /**
     * After how many minutes after a check starts failing an alert should be sent. Possible values are `5`, `10`, `15`, and `30`. (Default `5`).
     */
    minutesFailingThreshold?: number;
}

export interface CheckGroupApiCheckDefaults {
    assertions?: outputs.CheckGroupApiCheckDefaultsAssertion[];
    basicAuth: outputs.CheckGroupApiCheckDefaultsBasicAuth;
    headers: {[key: string]: any};
    queryParameters: {[key: string]: any};
    /**
     * The base url for this group which you can reference with the `GROUP_BASE_URL` variable in all group checks.
     */
    url: string;
}
/**
 * checkGroupApiCheckDefaultsProvideDefaults sets the appropriate defaults for CheckGroupApiCheckDefaults
 */
export function checkGroupApiCheckDefaultsProvideDefaults(val: CheckGroupApiCheckDefaults): CheckGroupApiCheckDefaults {
    return {
        ...val,
        url: (val.url) ?? "",
    };
}

export interface CheckGroupApiCheckDefaultsAssertion {
    /**
     * The type of comparison to be executed between expected and actual value of the assertion. Possible values `EQUALS`, `NOT_EQUALS`, `HAS_KEY`, `NOT_HAS_KEY`, `HAS_VALUE`, `NOT_HAS_VALUE`, `IS_EMPTY`, `NOT_EMPTY`, `GREATER_THAN`, `LESS_THAN`, `CONTAINS`, `NOT_CONTAINS`, `IS_NULL`, and `NOT_NULL`.
     */
    comparison: string;
    property?: string;
    /**
     * The source of the asserted value. Possible values `STATUS_CODE`, `JSON_BODY`, `HEADERS`, `TEXT_BODY`, and `RESPONSE_TIME`.
     */
    source: string;
    target: string;
}

export interface CheckGroupApiCheckDefaultsBasicAuth {
    password: string;
    username: string;
}

export interface CheckGroupEnvironmentVariable {
    key: string;
    locked?: boolean;
    value: string;
}

export interface CheckGroupRetryStrategy {
    /**
     * The number of seconds to wait before the first retry attempt.
     */
    baseBackoffSeconds?: number;
    /**
     * The total amount of time to continue retrying the check (maximum 600 seconds).
     */
    maxDurationSeconds?: number;
    /**
     * The maximum number of times to retry the check. Value must be between 1 and 10.
     */
    maxRetries?: number;
    /**
     * Whether retries should be run in the same region as the initial check run.
     */
    sameRegion?: boolean;
    /**
     * Determines which type of retry strategy to use. Possible values are `FIXED`, `LINEAR`, or `EXPONENTIAL`.
     */
    type: string;
}

export interface CheckRequest {
    /**
     * A request can have multiple assertions.
     */
    assertions?: outputs.CheckRequestAssertion[];
    /**
     * Set up HTTP basic authentication (username & password).
     */
    basicAuth: outputs.CheckRequestBasicAuth;
    /**
     * The body of the request.
     */
    body?: string;
    /**
     * The `Content-Type` header of the request. Possible values `NONE`, `JSON`, `FORM`, `RAW`, and `GRAPHQL`.
     */
    bodyType?: string;
    followRedirects?: boolean;
    headers: {[key: string]: any};
    /**
     * The HTTP method to use for this API check. Possible values are `GET`, `POST`, `PUT`, `HEAD`, `DELETE`, `PATCH`. (Default `GET`).
     */
    method?: string;
    queryParameters: {[key: string]: any};
    skipSsl?: boolean;
    url: string;
}

export interface CheckRequestAssertion {
    /**
     * The type of comparison to be executed between expected and actual value of the assertion. Possible values `EQUALS`, `NOT_EQUALS`, `HAS_KEY`, `NOT_HAS_KEY`, `HAS_VALUE`, `NOT_HAS_VALUE`, `IS_EMPTY`, `NOT_EMPTY`, `GREATER_THAN`, `LESS_THAN`, `CONTAINS`, `NOT_CONTAINS`, `IS_NULL`, and `NOT_NULL`.
     */
    comparison: string;
    property?: string;
    /**
     * The source of the asserted value. Possible values `STATUS_CODE`, `JSON_BODY`, `HEADERS`, `TEXT_BODY`, and `RESPONSE_TIME`.
     */
    source: string;
    target?: string;
}

export interface CheckRequestBasicAuth {
    password: string;
    username: string;
}

export interface CheckRetryStrategy {
    /**
     * The number of seconds to wait before the first retry attempt.
     */
    baseBackoffSeconds?: number;
    /**
     * The total amount of time to continue retrying the check (maximum 600 seconds).
     */
    maxDurationSeconds?: number;
    /**
     * The maximum number of times to retry the check. Value must be between 1 and 10.
     */
    maxRetries?: number;
    /**
     * Whether retries should be run in the same region as the initial check run.
     */
    sameRegion?: boolean;
    /**
     * Determines which type of retry strategy to use. Possible values are `FIXED`, `LINEAR`, or `EXPONENTIAL`.
     */
    type: string;
}

export interface HeartbeatAlertChannelSubscription {
    activated: boolean;
    channelId: number;
}

export interface HeartbeatAlertSettings {
    /**
     * Determines what type of escalation to use. Possible values are `RUN_BASED` or `TIME_BASED`.
     */
    escalationType?: string;
    parallelRunFailureThresholds?: outputs.HeartbeatAlertSettingsParallelRunFailureThreshold[];
    reminders?: outputs.HeartbeatAlertSettingsReminder[];
    runBasedEscalations?: outputs.HeartbeatAlertSettingsRunBasedEscalation[];
    /**
     * @deprecated This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.
     */
    sslCertificates?: outputs.HeartbeatAlertSettingsSslCertificate[];
    timeBasedEscalations?: outputs.HeartbeatAlertSettingsTimeBasedEscalation[];
}

export interface HeartbeatAlertSettingsParallelRunFailureThreshold {
    /**
     * Applicable only for checks scheduled in parallel in multiple locations.
     */
    enabled?: boolean;
    /**
     * Possible values are `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `100`, and `100`. (Default `10`).
     */
    percentage?: number;
}

export interface HeartbeatAlertSettingsReminder {
    /**
     * How many reminders to send out after the initial alert notification. Possible values are `0`, `1`, `2`, `3`, `4`, `5`, and `100000`
     */
    amount?: number;
    /**
     * Possible values are `5`, `10`, `15`, and `30`. (Default `5`).
     */
    interval?: number;
}

export interface HeartbeatAlertSettingsRunBasedEscalation {
    /**
     * After how many failed consecutive check runs an alert notification should be sent. Possible values are between 1 and 5. (Default `1`).
     */
    failedRunThreshold?: number;
}

export interface HeartbeatAlertSettingsSslCertificate {
    /**
     * How long before SSL certificate expiry to send alerts. Possible values `3`, `7`, `14`, `30`. (Default `3`).
     */
    alertThreshold?: number;
    /**
     * Applicable only for checks scheduled in parallel in multiple locations.
     */
    enabled?: boolean;
}

export interface HeartbeatAlertSettingsTimeBasedEscalation {
    /**
     * After how many minutes after a check starts failing an alert should be sent. Possible values are `5`, `10`, `15`, and `30`. (Default `5`).
     */
    minutesFailingThreshold?: number;
}

export interface HeartbeatHeartbeat {
    /**
     * How long Checkly should wait before triggering any alerts when a ping does not arrive within the set period.
     */
    grace: number;
    /**
     * Possible values `seconds`, `minutes`, `hours` and `days`.
     */
    graceUnit: string;
    /**
     * How often you expect a ping to the ping URL.
     */
    period: number;
    /**
     * Possible values `seconds`, `minutes`, `hours` and `days`.
     */
    periodUnit: string;
    /**
     * Custom token to generate your ping URL. Checkly will expect a ping to `https://ping.checklyhq.com/[PING_TOKEN]`.
     */
    pingToken: string;
}

