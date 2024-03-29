# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'AlertChannelEmailArgs',
    'AlertChannelOpsgenieArgs',
    'AlertChannelPagerdutyArgs',
    'AlertChannelSlackArgs',
    'AlertChannelSmsArgs',
    'AlertChannelWebhookArgs',
    'CheckAlertChannelSubscriptionArgs',
    'CheckAlertSettingsArgs',
    'CheckAlertSettingsReminderArgs',
    'CheckAlertSettingsRunBasedEscalationArgs',
    'CheckAlertSettingsSslCertificateArgs',
    'CheckAlertSettingsTimeBasedEscalationArgs',
    'CheckGroupAlertChannelSubscriptionArgs',
    'CheckGroupAlertSettingsArgs',
    'CheckGroupAlertSettingsReminderArgs',
    'CheckGroupAlertSettingsRunBasedEscalationArgs',
    'CheckGroupAlertSettingsSslCertificateArgs',
    'CheckGroupAlertSettingsTimeBasedEscalationArgs',
    'CheckGroupApiCheckDefaultsArgs',
    'CheckGroupApiCheckDefaultsAssertionArgs',
    'CheckGroupApiCheckDefaultsBasicAuthArgs',
    'CheckRequestArgs',
    'CheckRequestAssertionArgs',
    'CheckRequestBasicAuthArgs',
]

@pulumi.input_type
class AlertChannelEmailArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str]):
        pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)


@pulumi.input_type
class AlertChannelOpsgenieArgs:
    def __init__(__self__, *,
                 api_key: pulumi.Input[str],
                 name: pulumi.Input[str],
                 priority: pulumi.Input[str],
                 region: pulumi.Input[str]):
        pulumi.set(__self__, "api_key", api_key)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[str]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[str]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class AlertChannelPagerdutyArgs:
    def __init__(__self__, *,
                 service_key: pulumi.Input[str],
                 account: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "service_key", service_key)
        if account is not None:
            pulumi.set(__self__, "account", account)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="serviceKey")
    def service_key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_key")

    @service_key.setter
    def service_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_key", value)

    @property
    @pulumi.getter
    def account(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "account")

    @account.setter
    def account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class AlertChannelSlackArgs:
    def __init__(__self__, *,
                 channel: pulumi.Input[str],
                 url: pulumi.Input[str]):
        pulumi.set(__self__, "channel", channel)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def channel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "channel")

    @channel.setter
    def channel(self, value: pulumi.Input[str]):
        pulumi.set(self, "channel", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class AlertChannelSmsArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 number: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "number", number)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def number(self) -> pulumi.Input[str]:
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: pulumi.Input[str]):
        pulumi.set(self, "number", value)


@pulumi.input_type
class AlertChannelWebhookArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 url: pulumi.Input[str],
                 headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 query_parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 webhook_secret: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "url", url)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if query_parameters is not None:
            pulumi.set(__self__, "query_parameters", query_parameters)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if webhook_secret is not None:
            pulumi.set(__self__, "webhook_secret", webhook_secret)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter(name="queryParameters")
    def query_parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "query_parameters")

    @query_parameters.setter
    def query_parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "query_parameters", value)

    @property
    @pulumi.getter
    def template(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter(name="webhookSecret")
    def webhook_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "webhook_secret")

    @webhook_secret.setter
    def webhook_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "webhook_secret", value)


@pulumi.input_type
class CheckAlertChannelSubscriptionArgs:
    def __init__(__self__, *,
                 activated: pulumi.Input[bool],
                 channel_id: pulumi.Input[int]):
        pulumi.set(__self__, "activated", activated)
        pulumi.set(__self__, "channel_id", channel_id)

    @property
    @pulumi.getter
    def activated(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "activated")

    @activated.setter
    def activated(self, value: pulumi.Input[bool]):
        pulumi.set(self, "activated", value)

    @property
    @pulumi.getter(name="channelId")
    def channel_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "channel_id")

    @channel_id.setter
    def channel_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "channel_id", value)


@pulumi.input_type
class CheckAlertSettingsArgs:
    def __init__(__self__, *,
                 escalation_type: Optional[pulumi.Input[str]] = None,
                 reminders: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsReminderArgs']]]] = None,
                 run_based_escalations: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsRunBasedEscalationArgs']]]] = None,
                 ssl_certificates: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsSslCertificateArgs']]]] = None,
                 time_based_escalations: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsTimeBasedEscalationArgs']]]] = None):
        if escalation_type is not None:
            pulumi.set(__self__, "escalation_type", escalation_type)
        if reminders is not None:
            pulumi.set(__self__, "reminders", reminders)
        if run_based_escalations is not None:
            pulumi.set(__self__, "run_based_escalations", run_based_escalations)
        if ssl_certificates is not None:
            warnings.warn("""This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.""", DeprecationWarning)
            pulumi.log.warn("""ssl_certificates is deprecated: This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.""")
        if ssl_certificates is not None:
            pulumi.set(__self__, "ssl_certificates", ssl_certificates)
        if time_based_escalations is not None:
            pulumi.set(__self__, "time_based_escalations", time_based_escalations)

    @property
    @pulumi.getter(name="escalationType")
    def escalation_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "escalation_type")

    @escalation_type.setter
    def escalation_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "escalation_type", value)

    @property
    @pulumi.getter
    def reminders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsReminderArgs']]]]:
        return pulumi.get(self, "reminders")

    @reminders.setter
    def reminders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsReminderArgs']]]]):
        pulumi.set(self, "reminders", value)

    @property
    @pulumi.getter(name="runBasedEscalations")
    def run_based_escalations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsRunBasedEscalationArgs']]]]:
        return pulumi.get(self, "run_based_escalations")

    @run_based_escalations.setter
    def run_based_escalations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsRunBasedEscalationArgs']]]]):
        pulumi.set(self, "run_based_escalations", value)

    @property
    @pulumi.getter(name="sslCertificates")
    def ssl_certificates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsSslCertificateArgs']]]]:
        return pulumi.get(self, "ssl_certificates")

    @ssl_certificates.setter
    def ssl_certificates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsSslCertificateArgs']]]]):
        pulumi.set(self, "ssl_certificates", value)

    @property
    @pulumi.getter(name="timeBasedEscalations")
    def time_based_escalations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsTimeBasedEscalationArgs']]]]:
        return pulumi.get(self, "time_based_escalations")

    @time_based_escalations.setter
    def time_based_escalations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckAlertSettingsTimeBasedEscalationArgs']]]]):
        pulumi.set(self, "time_based_escalations", value)


@pulumi.input_type
class CheckAlertSettingsReminderArgs:
    def __init__(__self__, *,
                 amount: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None):
        if amount is not None:
            pulumi.set(__self__, "amount", amount)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)

    @property
    @pulumi.getter
    def amount(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)


@pulumi.input_type
class CheckAlertSettingsRunBasedEscalationArgs:
    def __init__(__self__, *,
                 failed_run_threshold: Optional[pulumi.Input[int]] = None):
        if failed_run_threshold is not None:
            pulumi.set(__self__, "failed_run_threshold", failed_run_threshold)

    @property
    @pulumi.getter(name="failedRunThreshold")
    def failed_run_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "failed_run_threshold")

    @failed_run_threshold.setter
    def failed_run_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "failed_run_threshold", value)


@pulumi.input_type
class CheckAlertSettingsSslCertificateArgs:
    def __init__(__self__, *,
                 alert_threshold: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        if alert_threshold is not None:
            pulumi.set(__self__, "alert_threshold", alert_threshold)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="alertThreshold")
    def alert_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "alert_threshold")

    @alert_threshold.setter
    def alert_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alert_threshold", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class CheckAlertSettingsTimeBasedEscalationArgs:
    def __init__(__self__, *,
                 minutes_failing_threshold: Optional[pulumi.Input[int]] = None):
        if minutes_failing_threshold is not None:
            pulumi.set(__self__, "minutes_failing_threshold", minutes_failing_threshold)

    @property
    @pulumi.getter(name="minutesFailingThreshold")
    def minutes_failing_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "minutes_failing_threshold")

    @minutes_failing_threshold.setter
    def minutes_failing_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minutes_failing_threshold", value)


@pulumi.input_type
class CheckGroupAlertChannelSubscriptionArgs:
    def __init__(__self__, *,
                 activated: pulumi.Input[bool],
                 channel_id: pulumi.Input[int]):
        pulumi.set(__self__, "activated", activated)
        pulumi.set(__self__, "channel_id", channel_id)

    @property
    @pulumi.getter
    def activated(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "activated")

    @activated.setter
    def activated(self, value: pulumi.Input[bool]):
        pulumi.set(self, "activated", value)

    @property
    @pulumi.getter(name="channelId")
    def channel_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "channel_id")

    @channel_id.setter
    def channel_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "channel_id", value)


@pulumi.input_type
class CheckGroupAlertSettingsArgs:
    def __init__(__self__, *,
                 escalation_type: Optional[pulumi.Input[str]] = None,
                 reminders: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsReminderArgs']]]] = None,
                 run_based_escalations: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsRunBasedEscalationArgs']]]] = None,
                 ssl_certificates: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsSslCertificateArgs']]]] = None,
                 time_based_escalations: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsTimeBasedEscalationArgs']]]] = None):
        if escalation_type is not None:
            pulumi.set(__self__, "escalation_type", escalation_type)
        if reminders is not None:
            pulumi.set(__self__, "reminders", reminders)
        if run_based_escalations is not None:
            pulumi.set(__self__, "run_based_escalations", run_based_escalations)
        if ssl_certificates is not None:
            warnings.warn("""This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.""", DeprecationWarning)
            pulumi.log.warn("""ssl_certificates is deprecated: This property is deprecated and it's ignored by the Checkly Public API. It will be removed in a future version.""")
        if ssl_certificates is not None:
            pulumi.set(__self__, "ssl_certificates", ssl_certificates)
        if time_based_escalations is not None:
            pulumi.set(__self__, "time_based_escalations", time_based_escalations)

    @property
    @pulumi.getter(name="escalationType")
    def escalation_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "escalation_type")

    @escalation_type.setter
    def escalation_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "escalation_type", value)

    @property
    @pulumi.getter
    def reminders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsReminderArgs']]]]:
        return pulumi.get(self, "reminders")

    @reminders.setter
    def reminders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsReminderArgs']]]]):
        pulumi.set(self, "reminders", value)

    @property
    @pulumi.getter(name="runBasedEscalations")
    def run_based_escalations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsRunBasedEscalationArgs']]]]:
        return pulumi.get(self, "run_based_escalations")

    @run_based_escalations.setter
    def run_based_escalations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsRunBasedEscalationArgs']]]]):
        pulumi.set(self, "run_based_escalations", value)

    @property
    @pulumi.getter(name="sslCertificates")
    def ssl_certificates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsSslCertificateArgs']]]]:
        return pulumi.get(self, "ssl_certificates")

    @ssl_certificates.setter
    def ssl_certificates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsSslCertificateArgs']]]]):
        pulumi.set(self, "ssl_certificates", value)

    @property
    @pulumi.getter(name="timeBasedEscalations")
    def time_based_escalations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsTimeBasedEscalationArgs']]]]:
        return pulumi.get(self, "time_based_escalations")

    @time_based_escalations.setter
    def time_based_escalations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupAlertSettingsTimeBasedEscalationArgs']]]]):
        pulumi.set(self, "time_based_escalations", value)


@pulumi.input_type
class CheckGroupAlertSettingsReminderArgs:
    def __init__(__self__, *,
                 amount: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None):
        if amount is not None:
            pulumi.set(__self__, "amount", amount)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)

    @property
    @pulumi.getter
    def amount(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)


@pulumi.input_type
class CheckGroupAlertSettingsRunBasedEscalationArgs:
    def __init__(__self__, *,
                 failed_run_threshold: Optional[pulumi.Input[int]] = None):
        if failed_run_threshold is not None:
            pulumi.set(__self__, "failed_run_threshold", failed_run_threshold)

    @property
    @pulumi.getter(name="failedRunThreshold")
    def failed_run_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "failed_run_threshold")

    @failed_run_threshold.setter
    def failed_run_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "failed_run_threshold", value)


@pulumi.input_type
class CheckGroupAlertSettingsSslCertificateArgs:
    def __init__(__self__, *,
                 alert_threshold: Optional[pulumi.Input[int]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        if alert_threshold is not None:
            pulumi.set(__self__, "alert_threshold", alert_threshold)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="alertThreshold")
    def alert_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "alert_threshold")

    @alert_threshold.setter
    def alert_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alert_threshold", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class CheckGroupAlertSettingsTimeBasedEscalationArgs:
    def __init__(__self__, *,
                 minutes_failing_threshold: Optional[pulumi.Input[int]] = None):
        if minutes_failing_threshold is not None:
            pulumi.set(__self__, "minutes_failing_threshold", minutes_failing_threshold)

    @property
    @pulumi.getter(name="minutesFailingThreshold")
    def minutes_failing_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "minutes_failing_threshold")

    @minutes_failing_threshold.setter
    def minutes_failing_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minutes_failing_threshold", value)


@pulumi.input_type
class CheckGroupApiCheckDefaultsArgs:
    def __init__(__self__, *,
                 assertions: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupApiCheckDefaultsAssertionArgs']]]] = None,
                 basic_auth: Optional[pulumi.Input['CheckGroupApiCheckDefaultsBasicAuthArgs']] = None,
                 headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 query_parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        if assertions is not None:
            pulumi.set(__self__, "assertions", assertions)
        if basic_auth is not None:
            pulumi.set(__self__, "basic_auth", basic_auth)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if query_parameters is not None:
            pulumi.set(__self__, "query_parameters", query_parameters)
        if url is None:
            url = ''
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def assertions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupApiCheckDefaultsAssertionArgs']]]]:
        return pulumi.get(self, "assertions")

    @assertions.setter
    def assertions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckGroupApiCheckDefaultsAssertionArgs']]]]):
        pulumi.set(self, "assertions", value)

    @property
    @pulumi.getter(name="basicAuth")
    def basic_auth(self) -> Optional[pulumi.Input['CheckGroupApiCheckDefaultsBasicAuthArgs']]:
        return pulumi.get(self, "basic_auth")

    @basic_auth.setter
    def basic_auth(self, value: Optional[pulumi.Input['CheckGroupApiCheckDefaultsBasicAuthArgs']]):
        pulumi.set(self, "basic_auth", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="queryParameters")
    def query_parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "query_parameters")

    @query_parameters.setter
    def query_parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "query_parameters", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class CheckGroupApiCheckDefaultsAssertionArgs:
    def __init__(__self__, *,
                 comparison: pulumi.Input[str],
                 source: pulumi.Input[str],
                 target: pulumi.Input[str],
                 property: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "comparison", comparison)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "target", target)
        if property is not None:
            pulumi.set(__self__, "property", property)

    @property
    @pulumi.getter
    def comparison(self) -> pulumi.Input[str]:
        return pulumi.get(self, "comparison")

    @comparison.setter
    def comparison(self, value: pulumi.Input[str]):
        pulumi.set(self, "comparison", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def property(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "property")

    @property.setter
    def property(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property", value)


@pulumi.input_type
class CheckGroupApiCheckDefaultsBasicAuthArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class CheckRequestArgs:
    def __init__(__self__, *,
                 url: pulumi.Input[str],
                 assertions: Optional[pulumi.Input[Sequence[pulumi.Input['CheckRequestAssertionArgs']]]] = None,
                 basic_auth: Optional[pulumi.Input['CheckRequestBasicAuthArgs']] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 body_type: Optional[pulumi.Input[str]] = None,
                 follow_redirects: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 query_parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 skip_ssl: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "url", url)
        if assertions is not None:
            pulumi.set(__self__, "assertions", assertions)
        if basic_auth is not None:
            pulumi.set(__self__, "basic_auth", basic_auth)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if body_type is not None:
            pulumi.set(__self__, "body_type", body_type)
        if follow_redirects is not None:
            pulumi.set(__self__, "follow_redirects", follow_redirects)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if query_parameters is not None:
            pulumi.set(__self__, "query_parameters", query_parameters)
        if skip_ssl is not None:
            pulumi.set(__self__, "skip_ssl", skip_ssl)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def assertions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CheckRequestAssertionArgs']]]]:
        return pulumi.get(self, "assertions")

    @assertions.setter
    def assertions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CheckRequestAssertionArgs']]]]):
        pulumi.set(self, "assertions", value)

    @property
    @pulumi.getter(name="basicAuth")
    def basic_auth(self) -> Optional[pulumi.Input['CheckRequestBasicAuthArgs']]:
        return pulumi.get(self, "basic_auth")

    @basic_auth.setter
    def basic_auth(self, value: Optional[pulumi.Input['CheckRequestBasicAuthArgs']]):
        pulumi.set(self, "basic_auth", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter(name="bodyType")
    def body_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "body_type")

    @body_type.setter
    def body_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body_type", value)

    @property
    @pulumi.getter(name="followRedirects")
    def follow_redirects(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "follow_redirects")

    @follow_redirects.setter
    def follow_redirects(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "follow_redirects", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter(name="queryParameters")
    def query_parameters(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "query_parameters")

    @query_parameters.setter
    def query_parameters(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "query_parameters", value)

    @property
    @pulumi.getter(name="skipSsl")
    def skip_ssl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "skip_ssl")

    @skip_ssl.setter
    def skip_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_ssl", value)


@pulumi.input_type
class CheckRequestAssertionArgs:
    def __init__(__self__, *,
                 comparison: pulumi.Input[str],
                 source: pulumi.Input[str],
                 property: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "comparison", comparison)
        pulumi.set(__self__, "source", source)
        if property is not None:
            pulumi.set(__self__, "property", property)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def comparison(self) -> pulumi.Input[str]:
        return pulumi.get(self, "comparison")

    @comparison.setter
    def comparison(self, value: pulumi.Input[str]):
        pulumi.set(self, "comparison", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def property(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "property")

    @property.setter
    def property(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "property", value)


@pulumi.input_type
class CheckRequestBasicAuthArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


