# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from .. import _utilities

import types

__config__ = pulumi.Config('artifactory')


class _ExportableConfig(types.ModuleType):
    @property
    def access_token(self) -> Optional[str]:
        """
        This is a access token that can be given to you by your admin under `User Management > Access Tokens`. If not set, the
        'api_key' attribute value will be used.
        """
        return __config__.get('accessToken')

    @property
    def api_key(self) -> Optional[str]:
        """
        API key. If `access_token` attribute, `JFROG_ACCESS_TOKEN` or `ARTIFACTORY_ACCESS_TOKEN` environment variable is set,
        the provider will ignore this attribute.
        """
        return __config__.get('apiKey')

    @property
    def oidc_provider_name(self) -> Optional[str]:
        """
        OIDC provider name. See [Configure an OIDC
        Integration](https://jfrog.com/help/r/jfrog-platform-administration-documentation/configure-an-oidc-integration) for
        more details.
        """
        return __config__.get('oidcProviderName')

    @property
    def tfc_credential_tag_name(self) -> Optional[str]:
        return __config__.get('tfcCredentialTagName')

    @property
    def url(self) -> Optional[str]:
        """
        Artifactory URL.
        """
        return __config__.get('url')

