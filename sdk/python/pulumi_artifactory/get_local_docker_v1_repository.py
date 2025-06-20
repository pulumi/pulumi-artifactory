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

__all__ = [
    'GetLocalDockerV1RepositoryResult',
    'AwaitableGetLocalDockerV1RepositoryResult',
    'get_local_docker_v1_repository',
    'get_local_docker_v1_repository_output',
]

@pulumi.output_type
class GetLocalDockerV1RepositoryResult:
    """
    A collection of values returned by getLocalDockerV1Repository.
    """
    def __init__(__self__, api_version=None, archive_browsing_enabled=None, blacked_out=None, block_pushing_schema1=None, cdn_redirect=None, description=None, download_direct=None, excludes_pattern=None, id=None, includes_pattern=None, key=None, max_unique_tags=None, notes=None, package_type=None, priority_resolution=None, project_environments=None, project_key=None, property_sets=None, repo_layout_ref=None, tag_retention=None, xray_index=None):
        if api_version and not isinstance(api_version, str):
            raise TypeError("Expected argument 'api_version' to be a str")
        pulumi.set(__self__, "api_version", api_version)
        if archive_browsing_enabled and not isinstance(archive_browsing_enabled, bool):
            raise TypeError("Expected argument 'archive_browsing_enabled' to be a bool")
        pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out and not isinstance(blacked_out, bool):
            raise TypeError("Expected argument 'blacked_out' to be a bool")
        pulumi.set(__self__, "blacked_out", blacked_out)
        if block_pushing_schema1 and not isinstance(block_pushing_schema1, bool):
            raise TypeError("Expected argument 'block_pushing_schema1' to be a bool")
        pulumi.set(__self__, "block_pushing_schema1", block_pushing_schema1)
        if cdn_redirect and not isinstance(cdn_redirect, bool):
            raise TypeError("Expected argument 'cdn_redirect' to be a bool")
        pulumi.set(__self__, "cdn_redirect", cdn_redirect)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if download_direct and not isinstance(download_direct, bool):
            raise TypeError("Expected argument 'download_direct' to be a bool")
        pulumi.set(__self__, "download_direct", download_direct)
        if excludes_pattern and not isinstance(excludes_pattern, str):
            raise TypeError("Expected argument 'excludes_pattern' to be a str")
        pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if includes_pattern and not isinstance(includes_pattern, str):
            raise TypeError("Expected argument 'includes_pattern' to be a str")
        pulumi.set(__self__, "includes_pattern", includes_pattern)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if max_unique_tags and not isinstance(max_unique_tags, int):
            raise TypeError("Expected argument 'max_unique_tags' to be a int")
        pulumi.set(__self__, "max_unique_tags", max_unique_tags)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if package_type and not isinstance(package_type, str):
            raise TypeError("Expected argument 'package_type' to be a str")
        pulumi.set(__self__, "package_type", package_type)
        if priority_resolution and not isinstance(priority_resolution, bool):
            raise TypeError("Expected argument 'priority_resolution' to be a bool")
        pulumi.set(__self__, "priority_resolution", priority_resolution)
        if project_environments and not isinstance(project_environments, list):
            raise TypeError("Expected argument 'project_environments' to be a list")
        pulumi.set(__self__, "project_environments", project_environments)
        if project_key and not isinstance(project_key, str):
            raise TypeError("Expected argument 'project_key' to be a str")
        pulumi.set(__self__, "project_key", project_key)
        if property_sets and not isinstance(property_sets, list):
            raise TypeError("Expected argument 'property_sets' to be a list")
        pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref and not isinstance(repo_layout_ref, str):
            raise TypeError("Expected argument 'repo_layout_ref' to be a str")
        pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if tag_retention and not isinstance(tag_retention, int):
            raise TypeError("Expected argument 'tag_retention' to be a int")
        pulumi.set(__self__, "tag_retention", tag_retention)
        if xray_index and not isinstance(xray_index, bool):
            raise TypeError("Expected argument 'xray_index' to be a bool")
        pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> builtins.str:
        """
        The Docker API version in use.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "archive_browsing_enabled")

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "blacked_out")

    @property
    @pulumi.getter(name="blockPushingSchema1")
    def block_pushing_schema1(self) -> builtins.bool:
        """
        When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to
        this repository.
        """
        return pulumi.get(self, "block_pushing_schema1")

    @property
    @pulumi.getter(name="cdnRedirect")
    def cdn_redirect(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "cdn_redirect")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "download_direct")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[builtins.str]:
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[builtins.str]:
        return pulumi.get(self, "includes_pattern")

    @property
    @pulumi.getter
    def key(self) -> builtins.str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="maxUniqueTags")
    def max_unique_tags(self) -> builtins.int:
        """
        The maximum number of unique tags of a single Docker image to store in this repository. Once the
        number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
        limit. This only applies to manifest v2.
        """
        return pulumi.get(self, "max_unique_tags")

    @property
    @pulumi.getter
    def notes(self) -> Optional[builtins.str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> builtins.str:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="priorityResolution")
    def priority_resolution(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "priority_resolution")

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "project_environments")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[builtins.str]:
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> Optional[Sequence[builtins.str]]:
        return pulumi.get(self, "property_sets")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[builtins.str]:
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter(name="tagRetention")
    def tag_retention(self) -> builtins.int:
        """
        If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This
        only applies to manifest V2.
        """
        return pulumi.get(self, "tag_retention")

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "xray_index")


class AwaitableGetLocalDockerV1RepositoryResult(GetLocalDockerV1RepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocalDockerV1RepositoryResult(
            api_version=self.api_version,
            archive_browsing_enabled=self.archive_browsing_enabled,
            blacked_out=self.blacked_out,
            block_pushing_schema1=self.block_pushing_schema1,
            cdn_redirect=self.cdn_redirect,
            description=self.description,
            download_direct=self.download_direct,
            excludes_pattern=self.excludes_pattern,
            id=self.id,
            includes_pattern=self.includes_pattern,
            key=self.key,
            max_unique_tags=self.max_unique_tags,
            notes=self.notes,
            package_type=self.package_type,
            priority_resolution=self.priority_resolution,
            project_environments=self.project_environments,
            project_key=self.project_key,
            property_sets=self.property_sets,
            repo_layout_ref=self.repo_layout_ref,
            tag_retention=self.tag_retention,
            xray_index=self.xray_index)


def get_local_docker_v1_repository(archive_browsing_enabled: Optional[builtins.bool] = None,
                                   blacked_out: Optional[builtins.bool] = None,
                                   cdn_redirect: Optional[builtins.bool] = None,
                                   description: Optional[builtins.str] = None,
                                   download_direct: Optional[builtins.bool] = None,
                                   excludes_pattern: Optional[builtins.str] = None,
                                   includes_pattern: Optional[builtins.str] = None,
                                   key: Optional[builtins.str] = None,
                                   max_unique_tags: Optional[builtins.int] = None,
                                   notes: Optional[builtins.str] = None,
                                   priority_resolution: Optional[builtins.bool] = None,
                                   project_environments: Optional[Sequence[builtins.str]] = None,
                                   project_key: Optional[builtins.str] = None,
                                   property_sets: Optional[Sequence[builtins.str]] = None,
                                   repo_layout_ref: Optional[builtins.str] = None,
                                   xray_index: Optional[builtins.bool] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocalDockerV1RepositoryResult:
    """
    Retrieves a local Docker (v1) repository resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    artifactory_local_test_docker_v1_repository = artifactory.get_local_docker_v1_repository(key="artifactory_local_test_docker_v1_repository")
    ```


    :param builtins.str key: the identity key of the repo.
    :param builtins.int max_unique_tags: The maximum number of unique tags of a single Docker image to store in this repository. Once the
           number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
           limit. This only applies to manifest v2.
    """
    __args__ = dict()
    __args__['archiveBrowsingEnabled'] = archive_browsing_enabled
    __args__['blackedOut'] = blacked_out
    __args__['cdnRedirect'] = cdn_redirect
    __args__['description'] = description
    __args__['downloadDirect'] = download_direct
    __args__['excludesPattern'] = excludes_pattern
    __args__['includesPattern'] = includes_pattern
    __args__['key'] = key
    __args__['maxUniqueTags'] = max_unique_tags
    __args__['notes'] = notes
    __args__['priorityResolution'] = priority_resolution
    __args__['projectEnvironments'] = project_environments
    __args__['projectKey'] = project_key
    __args__['propertySets'] = property_sets
    __args__['repoLayoutRef'] = repo_layout_ref
    __args__['xrayIndex'] = xray_index
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getLocalDockerV1Repository:getLocalDockerV1Repository', __args__, opts=opts, typ=GetLocalDockerV1RepositoryResult).value

    return AwaitableGetLocalDockerV1RepositoryResult(
        api_version=pulumi.get(__ret__, 'api_version'),
        archive_browsing_enabled=pulumi.get(__ret__, 'archive_browsing_enabled'),
        blacked_out=pulumi.get(__ret__, 'blacked_out'),
        block_pushing_schema1=pulumi.get(__ret__, 'block_pushing_schema1'),
        cdn_redirect=pulumi.get(__ret__, 'cdn_redirect'),
        description=pulumi.get(__ret__, 'description'),
        download_direct=pulumi.get(__ret__, 'download_direct'),
        excludes_pattern=pulumi.get(__ret__, 'excludes_pattern'),
        id=pulumi.get(__ret__, 'id'),
        includes_pattern=pulumi.get(__ret__, 'includes_pattern'),
        key=pulumi.get(__ret__, 'key'),
        max_unique_tags=pulumi.get(__ret__, 'max_unique_tags'),
        notes=pulumi.get(__ret__, 'notes'),
        package_type=pulumi.get(__ret__, 'package_type'),
        priority_resolution=pulumi.get(__ret__, 'priority_resolution'),
        project_environments=pulumi.get(__ret__, 'project_environments'),
        project_key=pulumi.get(__ret__, 'project_key'),
        property_sets=pulumi.get(__ret__, 'property_sets'),
        repo_layout_ref=pulumi.get(__ret__, 'repo_layout_ref'),
        tag_retention=pulumi.get(__ret__, 'tag_retention'),
        xray_index=pulumi.get(__ret__, 'xray_index'))
def get_local_docker_v1_repository_output(archive_browsing_enabled: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                          blacked_out: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                          cdn_redirect: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                          description: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          download_direct: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                          excludes_pattern: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          includes_pattern: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          key: Optional[pulumi.Input[builtins.str]] = None,
                                          max_unique_tags: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                          notes: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          priority_resolution: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                          project_environments: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                                          project_key: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          property_sets: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                                          repo_layout_ref: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          xray_index: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLocalDockerV1RepositoryResult]:
    """
    Retrieves a local Docker (v1) repository resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    artifactory_local_test_docker_v1_repository = artifactory.get_local_docker_v1_repository(key="artifactory_local_test_docker_v1_repository")
    ```


    :param builtins.str key: the identity key of the repo.
    :param builtins.int max_unique_tags: The maximum number of unique tags of a single Docker image to store in this repository. Once the
           number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no
           limit. This only applies to manifest v2.
    """
    __args__ = dict()
    __args__['archiveBrowsingEnabled'] = archive_browsing_enabled
    __args__['blackedOut'] = blacked_out
    __args__['cdnRedirect'] = cdn_redirect
    __args__['description'] = description
    __args__['downloadDirect'] = download_direct
    __args__['excludesPattern'] = excludes_pattern
    __args__['includesPattern'] = includes_pattern
    __args__['key'] = key
    __args__['maxUniqueTags'] = max_unique_tags
    __args__['notes'] = notes
    __args__['priorityResolution'] = priority_resolution
    __args__['projectEnvironments'] = project_environments
    __args__['projectKey'] = project_key
    __args__['propertySets'] = property_sets
    __args__['repoLayoutRef'] = repo_layout_ref
    __args__['xrayIndex'] = xray_index
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('artifactory:index/getLocalDockerV1Repository:getLocalDockerV1Repository', __args__, opts=opts, typ=GetLocalDockerV1RepositoryResult)
    return __ret__.apply(lambda __response__: GetLocalDockerV1RepositoryResult(
        api_version=pulumi.get(__response__, 'api_version'),
        archive_browsing_enabled=pulumi.get(__response__, 'archive_browsing_enabled'),
        blacked_out=pulumi.get(__response__, 'blacked_out'),
        block_pushing_schema1=pulumi.get(__response__, 'block_pushing_schema1'),
        cdn_redirect=pulumi.get(__response__, 'cdn_redirect'),
        description=pulumi.get(__response__, 'description'),
        download_direct=pulumi.get(__response__, 'download_direct'),
        excludes_pattern=pulumi.get(__response__, 'excludes_pattern'),
        id=pulumi.get(__response__, 'id'),
        includes_pattern=pulumi.get(__response__, 'includes_pattern'),
        key=pulumi.get(__response__, 'key'),
        max_unique_tags=pulumi.get(__response__, 'max_unique_tags'),
        notes=pulumi.get(__response__, 'notes'),
        package_type=pulumi.get(__response__, 'package_type'),
        priority_resolution=pulumi.get(__response__, 'priority_resolution'),
        project_environments=pulumi.get(__response__, 'project_environments'),
        project_key=pulumi.get(__response__, 'project_key'),
        property_sets=pulumi.get(__response__, 'property_sets'),
        repo_layout_ref=pulumi.get(__response__, 'repo_layout_ref'),
        tag_retention=pulumi.get(__response__, 'tag_retention'),
        xray_index=pulumi.get(__response__, 'xray_index')))
