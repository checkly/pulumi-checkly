# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TriggerCheckArgs', 'TriggerCheck']

@pulumi.input_type
class TriggerCheckArgs:
    def __init__(__self__, *,
                 check_id: pulumi.Input[str],
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TriggerCheck resource.
        """
        pulumi.set(__self__, "check_id", check_id)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="checkId")
    def check_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "check_id")

    @check_id.setter
    def check_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "check_id", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _TriggerCheckState:
    def __init__(__self__, *,
                 check_id: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TriggerCheck resources.
        """
        if check_id is not None:
            pulumi.set(__self__, "check_id", check_id)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="checkId")
    def check_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "check_id")

    @check_id.setter
    def check_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "check_id", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class TriggerCheck(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 check_id: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a TriggerCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerCheckArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a TriggerCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TriggerCheckArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerCheckArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 check_id: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TriggerCheckArgs.__new__(TriggerCheckArgs)

            if check_id is None and not opts.urn:
                raise TypeError("Missing required property 'check_id'")
            __props__.__dict__["check_id"] = check_id
            __props__.__dict__["token"] = token
            __props__.__dict__["url"] = url
        super(TriggerCheck, __self__).__init__(
            'checkly:index/triggerCheck:TriggerCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            check_id: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'TriggerCheck':
        """
        Get an existing TriggerCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TriggerCheckState.__new__(_TriggerCheckState)

        __props__.__dict__["check_id"] = check_id
        __props__.__dict__["token"] = token
        __props__.__dict__["url"] = url
        return TriggerCheck(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="checkId")
    def check_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "check_id")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "url")
