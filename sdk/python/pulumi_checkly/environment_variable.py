# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EnvironmentVariableArgs', 'EnvironmentVariable']

@pulumi.input_type
class EnvironmentVariableArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str],
                 locked: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a EnvironmentVariable resource.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)


@pulumi.input_type
class _EnvironmentVariableState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnvironmentVariable resources.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if locked is not None:
            pulumi.set(__self__, "locked", locked)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def locked(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "locked")

    @locked.setter
    def locked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "locked", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class EnvironmentVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_checkly as checkly

        # Simple Enviroment Variable example
        variable1 = checkly.EnvironmentVariable("variable1",
            key="API_KEY",
            locked=True,
            value="loZd9hOGHDUrGvmW")
        variable2 = checkly.EnvironmentVariable("variable2",
            key="API_URL",
            value="http://localhost:3000")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_checkly as checkly

        # Simple Enviroment Variable example
        variable1 = checkly.EnvironmentVariable("variable1",
            key="API_KEY",
            locked=True,
            value="loZd9hOGHDUrGvmW")
        variable2 = checkly.EnvironmentVariable("variable2",
            key="API_URL",
            value="http://localhost:3000")
        ```

        :param str resource_name: The name of the resource.
        :param EnvironmentVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 locked: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentVariableArgs.__new__(EnvironmentVariableArgs)

            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["locked"] = locked
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(EnvironmentVariable, __self__).__init__(
            'checkly:index/environmentVariable:EnvironmentVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[str]] = None,
            locked: Optional[pulumi.Input[bool]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'EnvironmentVariable':
        """
        Get an existing EnvironmentVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnvironmentVariableState.__new__(_EnvironmentVariableState)

        __props__.__dict__["key"] = key
        __props__.__dict__["locked"] = locked
        __props__.__dict__["value"] = value
        return EnvironmentVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def locked(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "locked")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        return pulumi.get(self, "value")

