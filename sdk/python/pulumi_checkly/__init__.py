# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alert_channel import *
from .check import *
from .check_group import *
from .client_certificate import *
from .dashboard import *
from .environment_variable import *
from .get_static_ips import *
from .heartbeat_check import *
from .maintenance_window import *
from .private_location import *
from .provider import *
from .snippet import *
from .status_page import *
from .status_page_service import *
from .tcp_check import *
from .trigger_check import *
from .trigger_check_group import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_checkly.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_checkly.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "checkly",
  "mod": "index/alertChannel",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/alertChannel:AlertChannel": "AlertChannel"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/check",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/check:Check": "Check"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/checkGroup",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/checkGroup:CheckGroup": "CheckGroup"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/clientCertificate",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/clientCertificate:ClientCertificate": "ClientCertificate"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/dashboard",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/dashboard:Dashboard": "Dashboard"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/environmentVariable",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/environmentVariable:EnvironmentVariable": "EnvironmentVariable"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/heartbeatCheck",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/heartbeatCheck:HeartbeatCheck": "HeartbeatCheck"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/maintenanceWindow",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/maintenanceWindow:MaintenanceWindow": "MaintenanceWindow"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/privateLocation",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/privateLocation:PrivateLocation": "PrivateLocation"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/snippet",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/snippet:Snippet": "Snippet"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/statusPage",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/statusPage:StatusPage": "StatusPage"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/statusPageService",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/statusPageService:StatusPageService": "StatusPageService"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/tcpCheck",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/tcpCheck:TcpCheck": "TcpCheck"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/triggerCheck",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/triggerCheck:TriggerCheck": "TriggerCheck"
  }
 },
 {
  "pkg": "checkly",
  "mod": "index/triggerCheckGroup",
  "fqn": "pulumi_checkly",
  "classes": {
   "checkly:index/triggerCheckGroup:TriggerCheckGroup": "TriggerCheckGroup"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "checkly",
  "token": "pulumi:providers:checkly",
  "fqn": "pulumi_checkly",
  "class": "Provider"
 }
]
"""
)
