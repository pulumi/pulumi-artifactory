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
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PermissionTargetArgs', 'PermissionTarget']

@pulumi.input_type
class PermissionTargetArgs:
    def __init__(__self__, *,
                 build: Optional[pulumi.Input['PermissionTargetBuildArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 release_bundle: Optional[pulumi.Input['PermissionTargetReleaseBundleArgs']] = None,
                 repo: Optional[pulumi.Input['PermissionTargetRepoArgs']] = None):
        """
        The set of arguments for constructing a PermissionTarget resource.
        :param pulumi.Input['PermissionTargetBuildArgs'] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[builtins.str] name: Name of permission.
        :param pulumi.Input['PermissionTargetReleaseBundleArgs'] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input['PermissionTargetRepoArgs'] repo: Repository permission configuration.
        """
        if build is not None:
            pulumi.set(__self__, "build", build)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if release_bundle is not None:
            pulumi.set(__self__, "release_bundle", release_bundle)
        if repo is not None:
            pulumi.set(__self__, "repo", repo)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['PermissionTargetBuildArgs']]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['PermissionTargetBuildArgs']]):
        pulumi.set(self, "build", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="releaseBundle")
    def release_bundle(self) -> Optional[pulumi.Input['PermissionTargetReleaseBundleArgs']]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundle")

    @release_bundle.setter
    def release_bundle(self, value: Optional[pulumi.Input['PermissionTargetReleaseBundleArgs']]):
        pulumi.set(self, "release_bundle", value)

    @property
    @pulumi.getter
    def repo(self) -> Optional[pulumi.Input['PermissionTargetRepoArgs']]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: Optional[pulumi.Input['PermissionTargetRepoArgs']]):
        pulumi.set(self, "repo", value)


@pulumi.input_type
class _PermissionTargetState:
    def __init__(__self__, *,
                 build: Optional[pulumi.Input['PermissionTargetBuildArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 release_bundle: Optional[pulumi.Input['PermissionTargetReleaseBundleArgs']] = None,
                 repo: Optional[pulumi.Input['PermissionTargetRepoArgs']] = None):
        """
        Input properties used for looking up and filtering PermissionTarget resources.
        :param pulumi.Input['PermissionTargetBuildArgs'] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[builtins.str] name: Name of permission.
        :param pulumi.Input['PermissionTargetReleaseBundleArgs'] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input['PermissionTargetRepoArgs'] repo: Repository permission configuration.
        """
        if build is not None:
            pulumi.set(__self__, "build", build)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if release_bundle is not None:
            pulumi.set(__self__, "release_bundle", release_bundle)
        if repo is not None:
            pulumi.set(__self__, "repo", repo)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['PermissionTargetBuildArgs']]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['PermissionTargetBuildArgs']]):
        pulumi.set(self, "build", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="releaseBundle")
    def release_bundle(self) -> Optional[pulumi.Input['PermissionTargetReleaseBundleArgs']]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundle")

    @release_bundle.setter
    def release_bundle(self, value: Optional[pulumi.Input['PermissionTargetReleaseBundleArgs']]):
        pulumi.set(self, "release_bundle", value)

    @property
    @pulumi.getter
    def repo(self) -> Optional[pulumi.Input['PermissionTargetRepoArgs']]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repo")

    @repo.setter
    def repo(self, value: Optional[pulumi.Input['PermissionTargetRepoArgs']]):
        pulumi.set(self, "repo", value)


@pulumi.type_token("artifactory:index/permissionTarget:PermissionTarget")
class PermissionTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[Union['PermissionTargetBuildArgs', 'PermissionTargetBuildArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 release_bundle: Optional[pulumi.Input[Union['PermissionTargetReleaseBundleArgs', 'PermissionTargetReleaseBundleArgsDict']]] = None,
                 repo: Optional[pulumi.Input[Union['PermissionTargetRepoArgs', 'PermissionTargetRepoArgsDict']]] = None,
                 __props__=None):
        """
        Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.

        !>This resource has been deprecated in favor of platform_permission resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory permission target called testpermission
        test_perm = artifactory.PermissionTarget("test-perm",
            name="test-perm",
            repo={
                "includes_patterns": ["foo/**"],
                "excludes_patterns": ["bar/**"],
                "repositories": ["example-repo-local"],
                "actions": {
                    "users": [
                        {
                            "name": "anonymous",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                        {
                            "name": "user1",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                    ],
                    "groups": [
                        {
                            "name": "readers",
                            "permissions": ["read"],
                        },
                        {
                            "name": "dev",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                    ],
                },
            },
            build={
                "includes_patterns": ["**"],
                "repositories": ["artifactory-build-info"],
                "actions": {
                    "users": [
                        {
                            "name": "anonymous",
                            "permissions": ["read"],
                        },
                        {
                            "name": "user1",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                    ],
                },
            },
            release_bundle={
                "includes_patterns": ["**"],
                "repositories": ["release-bundles"],
                "actions": {
                    "users": [{
                        "name": "anonymous",
                        "permissions": ["read"],
                    }],
                },
            })
        ```

        ## Permissions

        The provider supports the following `permission` enums:

        * `read`
        * `write`
        * `annotate`
        * `delete`
        * `manage`
        * `managedXrayMeta`
        * `distribute`

        The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):

        * `read` - matches `Read` permissions.
        * `write` - matches `  Deploy / Cache / Create ` permissions.
        * `annotate` - matches `Annotate` permissions.
        * `delete` - matches `Delete / Overwrite` permissions.
        * `manage` - matches `Manage` permissions.
        * `managedXrayMeta` - matches `Manage Xray Metadata` permissions.
        * `distribute` - matches `Distribute` permissions.

        ## Import

        Permission targets can be imported using their name, e.g.

        ```sh
        $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['PermissionTargetBuildArgs', 'PermissionTargetBuildArgsDict']] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[builtins.str] name: Name of permission.
        :param pulumi.Input[Union['PermissionTargetReleaseBundleArgs', 'PermissionTargetReleaseBundleArgsDict']] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input[Union['PermissionTargetRepoArgs', 'PermissionTargetRepoArgsDict']] repo: Repository permission configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PermissionTargetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.

        !>This resource has been deprecated in favor of platform_permission resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Create a new Artifactory permission target called testpermission
        test_perm = artifactory.PermissionTarget("test-perm",
            name="test-perm",
            repo={
                "includes_patterns": ["foo/**"],
                "excludes_patterns": ["bar/**"],
                "repositories": ["example-repo-local"],
                "actions": {
                    "users": [
                        {
                            "name": "anonymous",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                        {
                            "name": "user1",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                    ],
                    "groups": [
                        {
                            "name": "readers",
                            "permissions": ["read"],
                        },
                        {
                            "name": "dev",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                    ],
                },
            },
            build={
                "includes_patterns": ["**"],
                "repositories": ["artifactory-build-info"],
                "actions": {
                    "users": [
                        {
                            "name": "anonymous",
                            "permissions": ["read"],
                        },
                        {
                            "name": "user1",
                            "permissions": [
                                "read",
                                "write",
                            ],
                        },
                    ],
                },
            },
            release_bundle={
                "includes_patterns": ["**"],
                "repositories": ["release-bundles"],
                "actions": {
                    "users": [{
                        "name": "anonymous",
                        "permissions": ["read"],
                    }],
                },
            })
        ```

        ## Permissions

        The provider supports the following `permission` enums:

        * `read`
        * `write`
        * `annotate`
        * `delete`
        * `manage`
        * `managedXrayMeta`
        * `distribute`

        The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):

        * `read` - matches `Read` permissions.
        * `write` - matches `  Deploy / Cache / Create ` permissions.
        * `annotate` - matches `Annotate` permissions.
        * `delete` - matches `Delete / Overwrite` permissions.
        * `manage` - matches `Manage` permissions.
        * `managedXrayMeta` - matches `Manage Xray Metadata` permissions.
        * `distribute` - matches `Distribute` permissions.

        ## Import

        Permission targets can be imported using their name, e.g.

        ```sh
        $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
        ```

        :param str resource_name: The name of the resource.
        :param PermissionTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PermissionTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[Union['PermissionTargetBuildArgs', 'PermissionTargetBuildArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 release_bundle: Optional[pulumi.Input[Union['PermissionTargetReleaseBundleArgs', 'PermissionTargetReleaseBundleArgsDict']]] = None,
                 repo: Optional[pulumi.Input[Union['PermissionTargetRepoArgs', 'PermissionTargetRepoArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PermissionTargetArgs.__new__(PermissionTargetArgs)

            __props__.__dict__["build"] = build
            __props__.__dict__["name"] = name
            __props__.__dict__["release_bundle"] = release_bundle
            __props__.__dict__["repo"] = repo
        super(PermissionTarget, __self__).__init__(
            'artifactory:index/permissionTarget:PermissionTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            build: Optional[pulumi.Input[Union['PermissionTargetBuildArgs', 'PermissionTargetBuildArgsDict']]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            release_bundle: Optional[pulumi.Input[Union['PermissionTargetReleaseBundleArgs', 'PermissionTargetReleaseBundleArgsDict']]] = None,
            repo: Optional[pulumi.Input[Union['PermissionTargetRepoArgs', 'PermissionTargetRepoArgsDict']]] = None) -> 'PermissionTarget':
        """
        Get an existing PermissionTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['PermissionTargetBuildArgs', 'PermissionTargetBuildArgsDict']] build: As for repo but for artifactory-build-info permissions.
        :param pulumi.Input[builtins.str] name: Name of permission.
        :param pulumi.Input[Union['PermissionTargetReleaseBundleArgs', 'PermissionTargetReleaseBundleArgsDict']] release_bundle: As for repo for for release-bundles permissions.
        :param pulumi.Input[Union['PermissionTargetRepoArgs', 'PermissionTargetRepoArgsDict']] repo: Repository permission configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PermissionTargetState.__new__(_PermissionTargetState)

        __props__.__dict__["build"] = build
        __props__.__dict__["name"] = name
        __props__.__dict__["release_bundle"] = release_bundle
        __props__.__dict__["repo"] = repo
        return PermissionTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def build(self) -> pulumi.Output[Optional['outputs.PermissionTargetBuild']]:
        """
        As for repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of permission.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="releaseBundle")
    def release_bundle(self) -> pulumi.Output[Optional['outputs.PermissionTargetReleaseBundle']]:
        """
        As for repo for for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundle")

    @property
    @pulumi.getter
    def repo(self) -> pulumi.Output[Optional['outputs.PermissionTargetRepo']]:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repo")

