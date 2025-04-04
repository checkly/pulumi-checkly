# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['TriggerCheckGroupArgs', 'TriggerCheckGroup']

@pulumi.input_type
class TriggerCheckGroupArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[int],
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TriggerCheckGroup resource.
        :param pulumi.Input[int] group_id: The id of the group that you want to attach the trigger to.
        :param pulumi.Input[str] token: The token value created to trigger the group
        :param pulumi.Input[str] url: The request URL to trigger the group run.
        """
        pulumi.set(__self__, "group_id", group_id)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[int]:
        """
        The id of the group that you want to attach the trigger to.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token value created to trigger the group
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The request URL to trigger the group run.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _TriggerCheckGroupState:
    def __init__(__self__, *,
                 group_id: Optional[pulumi.Input[int]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TriggerCheckGroup resources.
        :param pulumi.Input[int] group_id: The id of the group that you want to attach the trigger to.
        :param pulumi.Input[str] token: The token value created to trigger the group
        :param pulumi.Input[str] url: The request URL to trigger the group run.
        """
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[int]]:
        """
        The id of the group that you want to attach the trigger to.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token value created to trigger the group
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The request URL to trigger the group run.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class TriggerCheckGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_checkly as checkly

        test_trigger_group = checkly.TriggerCheckGroup("test_trigger_group", group_id=215)
        pulumi.export("testTriggerGroupUrl", test_trigger_group.url)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] group_id: The id of the group that you want to attach the trigger to.
        :param pulumi.Input[str] token: The token value created to trigger the group
        :param pulumi.Input[str] url: The request URL to trigger the group run.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerCheckGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_checkly as checkly

        test_trigger_group = checkly.TriggerCheckGroup("test_trigger_group", group_id=215)
        pulumi.export("testTriggerGroupUrl", test_trigger_group.url)
        ```

        :param str resource_name: The name of the resource.
        :param TriggerCheckGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerCheckGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[int]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TriggerCheckGroupArgs.__new__(TriggerCheckGroupArgs)

            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["token"] = token
            __props__.__dict__["url"] = url
        super(TriggerCheckGroup, __self__).__init__(
            'checkly:index/triggerCheckGroup:TriggerCheckGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_id: Optional[pulumi.Input[int]] = None,
            token: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'TriggerCheckGroup':
        """
        Get an existing TriggerCheckGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] group_id: The id of the group that you want to attach the trigger to.
        :param pulumi.Input[str] token: The token value created to trigger the group
        :param pulumi.Input[str] url: The request URL to trigger the group run.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TriggerCheckGroupState.__new__(_TriggerCheckGroupState)

        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["token"] = token
        __props__.__dict__["url"] = url
        return TriggerCheckGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[int]:
        """
        The id of the group that you want to attach the trigger to.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The token value created to trigger the group
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The request URL to trigger the group run.
        """
        return pulumi.get(self, "url")

