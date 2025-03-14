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

__all__ = [
    'GetStaticIpsResult',
    'AwaitableGetStaticIpsResult',
    'get_static_ips',
    'get_static_ips_output',
]

@pulumi.output_type
class GetStaticIpsResult:
    """
    A collection of values returned by getStaticIps.
    """
    def __init__(__self__, addresses=None, id=None, ip_family=None, locations=None):
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_family and not isinstance(ip_family, str):
            raise TypeError("Expected argument 'ip_family' to be a str")
        pulumi.set(__self__, "ip_family", ip_family)
        if locations and not isinstance(locations, list):
            raise TypeError("Expected argument 'locations' to be a list")
        pulumi.set(__self__, "locations", locations)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence[str]:
        """
        Static IP addresses for Checkly's runner infrastructure.
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the static IPs data source.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipFamily")
    def ip_family(self) -> Optional[str]:
        """
        Specify the IP families you want to get static IPs for. Only `IPv6` or `IPv4` are valid options.
        """
        return pulumi.get(self, "ip_family")

    @property
    @pulumi.getter
    def locations(self) -> Optional[Sequence[str]]:
        """
        Specify the locations you want to get static IPs for.
        """
        return pulumi.get(self, "locations")


class AwaitableGetStaticIpsResult(GetStaticIpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStaticIpsResult(
            addresses=self.addresses,
            id=self.id,
            ip_family=self.ip_family,
            locations=self.locations)


def get_static_ips(ip_family: Optional[str] = None,
                   locations: Optional[Sequence[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStaticIpsResult:
    """
    Use this data source to access information about an existing resource.

    :param str ip_family: Specify the IP families you want to get static IPs for. Only `IPv6` or `IPv4` are valid options.
    :param Sequence[str] locations: Specify the locations you want to get static IPs for.
    """
    __args__ = dict()
    __args__['ipFamily'] = ip_family
    __args__['locations'] = locations
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('checkly:index/getStaticIps:getStaticIps', __args__, opts=opts, typ=GetStaticIpsResult).value

    return AwaitableGetStaticIpsResult(
        addresses=pulumi.get(__ret__, 'addresses'),
        id=pulumi.get(__ret__, 'id'),
        ip_family=pulumi.get(__ret__, 'ip_family'),
        locations=pulumi.get(__ret__, 'locations'))
def get_static_ips_output(ip_family: Optional[pulumi.Input[Optional[str]]] = None,
                          locations: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStaticIpsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str ip_family: Specify the IP families you want to get static IPs for. Only `IPv6` or `IPv4` are valid options.
    :param Sequence[str] locations: Specify the locations you want to get static IPs for.
    """
    __args__ = dict()
    __args__['ipFamily'] = ip_family
    __args__['locations'] = locations
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('checkly:index/getStaticIps:getStaticIps', __args__, opts=opts, typ=GetStaticIpsResult)
    return __ret__.apply(lambda __response__: GetStaticIpsResult(
        addresses=pulumi.get(__response__, 'addresses'),
        id=pulumi.get(__response__, 'id'),
        ip_family=pulumi.get(__response__, 'ip_family'),
        locations=pulumi.get(__response__, 'locations')))
