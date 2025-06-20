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

__all__ = ['UserLockPolicyArgs', 'UserLockPolicy']

@pulumi.input_type
class UserLockPolicyArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[builtins.bool],
                 login_attempts: pulumi.Input[builtins.int],
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a UserLockPolicy resource.
        :param pulumi.Input[builtins.bool] enabled: Enable User Lock Policy. Lock user after exceeding max failed login attempts.
        :param pulumi.Input[builtins.int] login_attempts: Max failed login attempts.
        :param pulumi.Input[builtins.str] name: Name of the resource. Only used for importing.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "login_attempts", login_attempts)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[builtins.bool]:
        """
        Enable User Lock Policy. Lock user after exceeding max failed login attempts.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="loginAttempts")
    def login_attempts(self) -> pulumi.Input[builtins.int]:
        """
        Max failed login attempts.
        """
        return pulumi.get(self, "login_attempts")

    @login_attempts.setter
    def login_attempts(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "login_attempts", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the resource. Only used for importing.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _UserLockPolicyState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 login_attempts: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering UserLockPolicy resources.
        :param pulumi.Input[builtins.bool] enabled: Enable User Lock Policy. Lock user after exceeding max failed login attempts.
        :param pulumi.Input[builtins.int] login_attempts: Max failed login attempts.
        :param pulumi.Input[builtins.str] name: Name of the resource. Only used for importing.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if login_attempts is not None:
            pulumi.set(__self__, "login_attempts", login_attempts)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Enable User Lock Policy. Lock user after exceeding max failed login attempts.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="loginAttempts")
    def login_attempts(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Max failed login attempts.
        """
        return pulumi.get(self, "login_attempts")

    @login_attempts.setter
    def login_attempts(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "login_attempts", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the resource. Only used for importing.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.type_token("artifactory:index/userLockPolicy:UserLockPolicy")
class UserLockPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 login_attempts: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Provides an Artifactory User Lock Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_user_lock_policy = artifactory.UserLockPolicy("my-user-lock-policy",
            name="my-user-lock-policy",
            enabled=True,
            login_attempts=10)
        ```

        ## Import

        ```sh
        $ pulumi import artifactory:index/userLockPolicy:UserLockPolicy my-user-lock-policy my-user-lock-policy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] enabled: Enable User Lock Policy. Lock user after exceeding max failed login attempts.
        :param pulumi.Input[builtins.int] login_attempts: Max failed login attempts.
        :param pulumi.Input[builtins.str] name: Name of the resource. Only used for importing.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserLockPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Artifactory User Lock Policy resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_user_lock_policy = artifactory.UserLockPolicy("my-user-lock-policy",
            name="my-user-lock-policy",
            enabled=True,
            login_attempts=10)
        ```

        ## Import

        ```sh
        $ pulumi import artifactory:index/userLockPolicy:UserLockPolicy my-user-lock-policy my-user-lock-policy
        ```

        :param str resource_name: The name of the resource.
        :param UserLockPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserLockPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 login_attempts: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserLockPolicyArgs.__new__(UserLockPolicyArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if login_attempts is None and not opts.urn:
                raise TypeError("Missing required property 'login_attempts'")
            __props__.__dict__["login_attempts"] = login_attempts
            __props__.__dict__["name"] = name
        super(UserLockPolicy, __self__).__init__(
            'artifactory:index/userLockPolicy:UserLockPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            login_attempts: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None) -> 'UserLockPolicy':
        """
        Get an existing UserLockPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] enabled: Enable User Lock Policy. Lock user after exceeding max failed login attempts.
        :param pulumi.Input[builtins.int] login_attempts: Max failed login attempts.
        :param pulumi.Input[builtins.str] name: Name of the resource. Only used for importing.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserLockPolicyState.__new__(_UserLockPolicyState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["login_attempts"] = login_attempts
        __props__.__dict__["name"] = name
        return UserLockPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[builtins.bool]:
        """
        Enable User Lock Policy. Lock user after exceeding max failed login attempts.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="loginAttempts")
    def login_attempts(self) -> pulumi.Output[builtins.int]:
        """
        Max failed login attempts.
        """
        return pulumi.get(self, "login_attempts")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the resource. Only used for importing.
        """
        return pulumi.get(self, "name")

