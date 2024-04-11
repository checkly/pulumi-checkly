# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PrivateLocationArgs', 'PrivateLocation']

@pulumi.input_type
class PrivateLocationArgs:
    def __init__(__self__, *,
                 slug_name: pulumi.Input[str],
                 icon: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateLocation resource.
        :param pulumi.Input[str] slug_name: Valid slug name.
        :param pulumi.Input[str] icon: Icon assigned to the private location.
        :param pulumi.Input[str] name: The private location name.
        """
        pulumi.set(__self__, "slug_name", slug_name)
        if icon is not None:
            pulumi.set(__self__, "icon", icon)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="slugName")
    def slug_name(self) -> pulumi.Input[str]:
        """
        Valid slug name.
        """
        return pulumi.get(self, "slug_name")

    @slug_name.setter
    def slug_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "slug_name", value)

    @property
    @pulumi.getter
    def icon(self) -> Optional[pulumi.Input[str]]:
        """
        Icon assigned to the private location.
        """
        return pulumi.get(self, "icon")

    @icon.setter
    def icon(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icon", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The private location name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PrivateLocationState:
    def __init__(__self__, *,
                 icon: Optional[pulumi.Input[str]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 slug_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivateLocation resources.
        :param pulumi.Input[str] icon: Icon assigned to the private location.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] keys: Private location API keys.
        :param pulumi.Input[str] name: The private location name.
        :param pulumi.Input[str] slug_name: Valid slug name.
        """
        if icon is not None:
            pulumi.set(__self__, "icon", icon)
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if slug_name is not None:
            pulumi.set(__self__, "slug_name", slug_name)

    @property
    @pulumi.getter
    def icon(self) -> Optional[pulumi.Input[str]]:
        """
        Icon assigned to the private location.
        """
        return pulumi.get(self, "icon")

    @icon.setter
    def icon(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icon", value)

    @property
    @pulumi.getter
    def keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Private location API keys.
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The private location name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="slugName")
    def slug_name(self) -> Optional[pulumi.Input[str]]:
        """
        Valid slug name.
        """
        return pulumi.get(self, "slug_name")

    @slug_name.setter
    def slug_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug_name", value)


class PrivateLocation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 icon: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 slug_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_checkly as checkly

        location = checkly.PrivateLocation("location", slug_name="new-private-location")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] icon: Icon assigned to the private location.
        :param pulumi.Input[str] name: The private location name.
        :param pulumi.Input[str] slug_name: Valid slug name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateLocationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_checkly as checkly

        location = checkly.PrivateLocation("location", slug_name="new-private-location")
        ```

        :param str resource_name: The name of the resource.
        :param PrivateLocationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateLocationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 icon: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 slug_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateLocationArgs.__new__(PrivateLocationArgs)

            __props__.__dict__["icon"] = icon
            __props__.__dict__["name"] = name
            if slug_name is None and not opts.urn:
                raise TypeError("Missing required property 'slug_name'")
            __props__.__dict__["slug_name"] = slug_name
            __props__.__dict__["keys"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["keys"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(PrivateLocation, __self__).__init__(
            'checkly:index/privateLocation:PrivateLocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            icon: Optional[pulumi.Input[str]] = None,
            keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            slug_name: Optional[pulumi.Input[str]] = None) -> 'PrivateLocation':
        """
        Get an existing PrivateLocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] icon: Icon assigned to the private location.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] keys: Private location API keys.
        :param pulumi.Input[str] name: The private location name.
        :param pulumi.Input[str] slug_name: Valid slug name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateLocationState.__new__(_PrivateLocationState)

        __props__.__dict__["icon"] = icon
        __props__.__dict__["keys"] = keys
        __props__.__dict__["name"] = name
        __props__.__dict__["slug_name"] = slug_name
        return PrivateLocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def icon(self) -> pulumi.Output[Optional[str]]:
        """
        Icon assigned to the private location.
        """
        return pulumi.get(self, "icon")

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Output[Sequence[str]]:
        """
        Private location API keys.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The private location name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="slugName")
    def slug_name(self) -> pulumi.Output[str]:
        """
        Valid slug name.
        """
        return pulumi.get(self, "slug_name")

