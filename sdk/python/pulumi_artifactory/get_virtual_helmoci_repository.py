# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVirtualHelmociRepositoryResult',
    'AwaitableGetVirtualHelmociRepositoryResult',
    'get_virtual_helmoci_repository',
    'get_virtual_helmoci_repository_output',
]

@pulumi.output_type
class GetVirtualHelmociRepositoryResult:
    """
    A collection of values returned by getVirtualHelmociRepository.
    """
    def __init__(__self__, artifactory_requests_can_retrieve_remote_artifacts=None, default_deployment_repo=None, description=None, excludes_pattern=None, id=None, includes_pattern=None, key=None, notes=None, package_type=None, project_environments=None, project_key=None, repo_layout_ref=None, repositories=None, resolve_oci_tags_by_timestamp=None):
        if artifactory_requests_can_retrieve_remote_artifacts and not isinstance(artifactory_requests_can_retrieve_remote_artifacts, bool):
            raise TypeError("Expected argument 'artifactory_requests_can_retrieve_remote_artifacts' to be a bool")
        pulumi.set(__self__, "artifactory_requests_can_retrieve_remote_artifacts", artifactory_requests_can_retrieve_remote_artifacts)
        if default_deployment_repo and not isinstance(default_deployment_repo, str):
            raise TypeError("Expected argument 'default_deployment_repo' to be a str")
        pulumi.set(__self__, "default_deployment_repo", default_deployment_repo)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
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
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if package_type and not isinstance(package_type, str):
            raise TypeError("Expected argument 'package_type' to be a str")
        pulumi.set(__self__, "package_type", package_type)
        if project_environments and not isinstance(project_environments, list):
            raise TypeError("Expected argument 'project_environments' to be a list")
        pulumi.set(__self__, "project_environments", project_environments)
        if project_key and not isinstance(project_key, str):
            raise TypeError("Expected argument 'project_key' to be a str")
        pulumi.set(__self__, "project_key", project_key)
        if repo_layout_ref and not isinstance(repo_layout_ref, str):
            raise TypeError("Expected argument 'repo_layout_ref' to be a str")
        pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if repositories and not isinstance(repositories, list):
            raise TypeError("Expected argument 'repositories' to be a list")
        pulumi.set(__self__, "repositories", repositories)
        if resolve_oci_tags_by_timestamp and not isinstance(resolve_oci_tags_by_timestamp, bool):
            raise TypeError("Expected argument 'resolve_oci_tags_by_timestamp' to be a bool")
        pulumi.set(__self__, "resolve_oci_tags_by_timestamp", resolve_oci_tags_by_timestamp)

    @property
    @pulumi.getter(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    def artifactory_requests_can_retrieve_remote_artifacts(self) -> Optional[bool]:
        return pulumi.get(self, "artifactory_requests_can_retrieve_remote_artifacts")

    @property
    @pulumi.getter(name="defaultDeploymentRepo")
    def default_deployment_repo(self) -> Optional[str]:
        return pulumi.get(self, "default_deployment_repo")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[str]:
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[str]:
        return pulumi.get(self, "includes_pattern")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> str:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="projectEnvironments")
    def project_environments(self) -> Sequence[str]:
        return pulumi.get(self, "project_environments")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[str]:
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[str]:
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter
    def repositories(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "repositories")

    @property
    @pulumi.getter(name="resolveOciTagsByTimestamp")
    def resolve_oci_tags_by_timestamp(self) -> Optional[bool]:
        """
        (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
        """
        return pulumi.get(self, "resolve_oci_tags_by_timestamp")


class AwaitableGetVirtualHelmociRepositoryResult(GetVirtualHelmociRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualHelmociRepositoryResult(
            artifactory_requests_can_retrieve_remote_artifacts=self.artifactory_requests_can_retrieve_remote_artifacts,
            default_deployment_repo=self.default_deployment_repo,
            description=self.description,
            excludes_pattern=self.excludes_pattern,
            id=self.id,
            includes_pattern=self.includes_pattern,
            key=self.key,
            notes=self.notes,
            package_type=self.package_type,
            project_environments=self.project_environments,
            project_key=self.project_key,
            repo_layout_ref=self.repo_layout_ref,
            repositories=self.repositories,
            resolve_oci_tags_by_timestamp=self.resolve_oci_tags_by_timestamp)


def get_virtual_helmoci_repository(artifactory_requests_can_retrieve_remote_artifacts: Optional[bool] = None,
                                   default_deployment_repo: Optional[str] = None,
                                   description: Optional[str] = None,
                                   excludes_pattern: Optional[str] = None,
                                   includes_pattern: Optional[str] = None,
                                   key: Optional[str] = None,
                                   notes: Optional[str] = None,
                                   project_environments: Optional[Sequence[str]] = None,
                                   project_key: Optional[str] = None,
                                   repo_layout_ref: Optional[str] = None,
                                   repositories: Optional[Sequence[str]] = None,
                                   resolve_oci_tags_by_timestamp: Optional[bool] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualHelmociRepositoryResult:
    """
    Retrieves a virtual Helm OCI repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_helmoci_virtual = artifactory.get_virtual_helmoci_repository(key="my-helmoci-virtual")
    ```


    :param str key: the identity key of the repo.
    :param bool resolve_oci_tags_by_timestamp: (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
    """
    __args__ = dict()
    __args__['artifactoryRequestsCanRetrieveRemoteArtifacts'] = artifactory_requests_can_retrieve_remote_artifacts
    __args__['defaultDeploymentRepo'] = default_deployment_repo
    __args__['description'] = description
    __args__['excludesPattern'] = excludes_pattern
    __args__['includesPattern'] = includes_pattern
    __args__['key'] = key
    __args__['notes'] = notes
    __args__['projectEnvironments'] = project_environments
    __args__['projectKey'] = project_key
    __args__['repoLayoutRef'] = repo_layout_ref
    __args__['repositories'] = repositories
    __args__['resolveOciTagsByTimestamp'] = resolve_oci_tags_by_timestamp
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getVirtualHelmociRepository:getVirtualHelmociRepository', __args__, opts=opts, typ=GetVirtualHelmociRepositoryResult).value

    return AwaitableGetVirtualHelmociRepositoryResult(
        artifactory_requests_can_retrieve_remote_artifacts=pulumi.get(__ret__, 'artifactory_requests_can_retrieve_remote_artifacts'),
        default_deployment_repo=pulumi.get(__ret__, 'default_deployment_repo'),
        description=pulumi.get(__ret__, 'description'),
        excludes_pattern=pulumi.get(__ret__, 'excludes_pattern'),
        id=pulumi.get(__ret__, 'id'),
        includes_pattern=pulumi.get(__ret__, 'includes_pattern'),
        key=pulumi.get(__ret__, 'key'),
        notes=pulumi.get(__ret__, 'notes'),
        package_type=pulumi.get(__ret__, 'package_type'),
        project_environments=pulumi.get(__ret__, 'project_environments'),
        project_key=pulumi.get(__ret__, 'project_key'),
        repo_layout_ref=pulumi.get(__ret__, 'repo_layout_ref'),
        repositories=pulumi.get(__ret__, 'repositories'),
        resolve_oci_tags_by_timestamp=pulumi.get(__ret__, 'resolve_oci_tags_by_timestamp'))


@_utilities.lift_output_func(get_virtual_helmoci_repository)
def get_virtual_helmoci_repository_output(artifactory_requests_can_retrieve_remote_artifacts: Optional[pulumi.Input[Optional[bool]]] = None,
                                          default_deployment_repo: Optional[pulumi.Input[Optional[str]]] = None,
                                          description: Optional[pulumi.Input[Optional[str]]] = None,
                                          excludes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                          includes_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                                          key: Optional[pulumi.Input[str]] = None,
                                          notes: Optional[pulumi.Input[Optional[str]]] = None,
                                          project_environments: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                          project_key: Optional[pulumi.Input[Optional[str]]] = None,
                                          repo_layout_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                          repositories: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                          resolve_oci_tags_by_timestamp: Optional[pulumi.Input[Optional[bool]]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVirtualHelmociRepositoryResult]:
    """
    Retrieves a virtual Helm OCI repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    my_helmoci_virtual = artifactory.get_virtual_helmoci_repository(key="my-helmoci-virtual")
    ```


    :param str key: the identity key of the repo.
    :param bool resolve_oci_tags_by_timestamp: (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
    """
    ...