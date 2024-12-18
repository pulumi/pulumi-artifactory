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

__all__ = ['ArtifactArgs', 'Artifact']

@pulumi.input_type
class ArtifactArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 content_base64: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Artifact resource.
        :param pulumi.Input[str] path: The relative path in the target repository. Must begin with a '/'. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
        :param pulumi.Input[str] repository: Name of the respository.
        :param pulumi.Input[str] content_base64: Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
        :param pulumi.Input[str] file_path: Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
        """
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "repository", repository)
        if content_base64 is not None:
            pulumi.set(__self__, "content_base64", content_base64)
        if file_path is not None:
            pulumi.set(__self__, "file_path", file_path)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The relative path in the target repository. Must begin with a '/'. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        Name of the respository.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> Optional[pulumi.Input[str]]:
        """
        Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_path", value)


@pulumi.input_type
class _ArtifactState:
    def __init__(__self__, *,
                 checksum_md5: Optional[pulumi.Input[str]] = None,
                 checksum_sha1: Optional[pulumi.Input[str]] = None,
                 checksum_sha256: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 created: Optional[pulumi.Input[str]] = None,
                 created_by: Optional[pulumi.Input[str]] = None,
                 download_uri: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 mime_type: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 uri: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Artifact resources.
        :param pulumi.Input[str] checksum_md5: MD5 checksum of the artifact.
        :param pulumi.Input[str] checksum_sha1: SHA1 checksum of the artifact.
        :param pulumi.Input[str] checksum_sha256: SHA256 checksum of the artifact.
        :param pulumi.Input[str] content_base64: Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
        :param pulumi.Input[str] created: Timestamp when artifact is created.
        :param pulumi.Input[str] created_by: User who deploys the artifact.
        :param pulumi.Input[str] download_uri: Download URI of the artifact.
        :param pulumi.Input[str] file_path: Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
        :param pulumi.Input[str] mime_type: MIME type of the artifact.
        :param pulumi.Input[str] path: The relative path in the target repository. Must begin with a '/'. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
        :param pulumi.Input[str] repository: Name of the respository.
        :param pulumi.Input[int] size: Size of the artifact, in bytes.
        :param pulumi.Input[str] uri: URI of the artifact.
        """
        if checksum_md5 is not None:
            pulumi.set(__self__, "checksum_md5", checksum_md5)
        if checksum_sha1 is not None:
            pulumi.set(__self__, "checksum_sha1", checksum_sha1)
        if checksum_sha256 is not None:
            pulumi.set(__self__, "checksum_sha256", checksum_sha256)
        if content_base64 is not None:
            pulumi.set(__self__, "content_base64", content_base64)
        if created is not None:
            pulumi.set(__self__, "created", created)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if download_uri is not None:
            pulumi.set(__self__, "download_uri", download_uri)
        if file_path is not None:
            pulumi.set(__self__, "file_path", file_path)
        if mime_type is not None:
            pulumi.set(__self__, "mime_type", mime_type)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="checksumMd5")
    def checksum_md5(self) -> Optional[pulumi.Input[str]]:
        """
        MD5 checksum of the artifact.
        """
        return pulumi.get(self, "checksum_md5")

    @checksum_md5.setter
    def checksum_md5(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "checksum_md5", value)

    @property
    @pulumi.getter(name="checksumSha1")
    def checksum_sha1(self) -> Optional[pulumi.Input[str]]:
        """
        SHA1 checksum of the artifact.
        """
        return pulumi.get(self, "checksum_sha1")

    @checksum_sha1.setter
    def checksum_sha1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "checksum_sha1", value)

    @property
    @pulumi.getter(name="checksumSha256")
    def checksum_sha256(self) -> Optional[pulumi.Input[str]]:
        """
        SHA256 checksum of the artifact.
        """
        return pulumi.get(self, "checksum_sha256")

    @checksum_sha256.setter
    def checksum_sha256(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "checksum_sha256", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> Optional[pulumi.Input[str]]:
        """
        Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp when artifact is created.
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[pulumi.Input[str]]:
        """
        User who deploys the artifact.
        """
        return pulumi.get(self, "created_by")

    @created_by.setter
    def created_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_by", value)

    @property
    @pulumi.getter(name="downloadUri")
    def download_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Download URI of the artifact.
        """
        return pulumi.get(self, "download_uri")

    @download_uri.setter
    def download_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_uri", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter(name="mimeType")
    def mime_type(self) -> Optional[pulumi.Input[str]]:
        """
        MIME type of the artifact.
        """
        return pulumi.get(self, "mime_type")

    @mime_type.setter
    def mime_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mime_type", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The relative path in the target repository. Must begin with a '/'. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the respository.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Size of the artifact, in bytes.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def uri(self) -> Optional[pulumi.Input[str]]:
        """
        URI of the artifact.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uri", value)


class Artifact(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource for deploying artifact to Artifactory repository. Support deploying a single artifact only. Changes to `repository` or `path` attributes will trigger a recreation of the resource (i.e. delete then create). See [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/deploy-a-single-artifact) for more details.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_artifact = artifactory.Artifact("my-artifact",
            repository="my-generic-local",
            path="/my-path/my-file.zip",
            file_path="/path/to/my-file.zip")
        my_base64_artifact = artifactory.Artifact("my-base64-artifact",
            repository="my-generic-local",
            path="/my-path/my-file.zip",
            content_base64="UEsDBAoAAAAAALl8alQAAAAAAAAAAAAAAAAJAAAATUVUQS1JTkYvUEsDBAoAAAAIALh8alTmUEsubQAAAIMAAAAUAAAATUVUQS1JTkYvTUFOSUZFU1QuTUbzTczLTEstLtENSy0qzszPs1Iw1DPg5XIsSs7ILEstQggH5KRWlBYrwCR4uZyLUhNLUlN0nSqtFBwLEpMzUhV8E8tS8xSM9cz0jHm5nEozc0rAsilAO1JzcjMhYim6XinZQGuA9uiZ83LxcgEAUEsDBAoAAAAAALh8alQAAAAAAAAAAAAAAAAMAAAAYXJ0aWZhY3RvcnkvUEsDBAoAAAAAALh8alQAAAAAAAAAAAAAAAARAAAAYXJ0aWZhY3RvcnkvdGVzdC9QSwMECgAAAAgAuHxqVMgzPcxdAQAALAIAAB0AAABhcnRpZmFjdG9yeS90ZXN0L011bHRpMS5jbGFzc3VRy07CQBQ9w6OlpQqCgPgEV+jCxsTEBcaNiXFRHwkGF66GOuKQPkyZmvBZutDEhR/gRxlvCwkxwVnck3vunXPPnfn++fwCcIRdExpWDdRQL6BhYg1NHes6Nhi0ExlIdcqQ7ez1GXJn4YNgKDkyEFexPxDRLR94xFSc0OVen0cyyWdkTj3JMUPT4ZGSj9xVYTSxlRgr+zL2lDzsUovPZcBQ79w7I/7CbY8HQ7unIhkMu+lAHg1JorqgzGD2wjhyxblMhhWnmgdJnwUdBR2bFrawzWBdCM8LW3dh5D20dexYaKHN0PjHFkN5Pux6MBKu+kP1JmMlfHqSMKZCbepMhvYN2VJkTnCfzFUX0Az6c5J5tHKts2hjtJGnz0hOBixZg6JBmU3ICPP7H2CvadmkqKWkhiJFa9pAuERoYBml2eXjVIy4N2Qq2Xfk5gImIWhKgRrnIgbKWCGkj007q79QSwECFAMKAAAAAAC5fGpUAAAAAAAAAAAAAAAACQAAAAAAAAAAABAA7UEAAAAATUVUQS1JTkYvUEsBAhQDCgAAAAgAuHxqVOZQSy5tAAAAgwAAABQAAAAAAAAAAAAAAKSBJwAAAE1FVEEtSU5GL01BTklGRVNULk1GUEsBAhQDCgAAAAAAuHxqVAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAQAO1BxgAAAGFydGlmYWN0b3J5L1BLAQIUAwoAAAAAALh8alQAAAAAAAAAAAAAAAARAAAAAAAAAAAAEADtQfAAAABhcnRpZmFjdG9yeS90ZXN0L1BLAQIUAwoAAAAIALh8alTIMz3MXQEAACwCAAAdAAAAAAAAAAAAAACkgR8BAABhcnRpZmFjdG9yeS90ZXN0L011bHRpMS5jbGFzc1BLBQYAAAAABQAFAD0BAAC3AgAAAAA=")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_base64: Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
        :param pulumi.Input[str] file_path: Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
        :param pulumi.Input[str] path: The relative path in the target repository. Must begin with a '/'. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
        :param pulumi.Input[str] repository: Name of the respository.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ArtifactArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource for deploying artifact to Artifactory repository. Support deploying a single artifact only. Changes to `repository` or `path` attributes will trigger a recreation of the resource (i.e. delete then create). See [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/deploy-a-single-artifact) for more details.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        my_artifact = artifactory.Artifact("my-artifact",
            repository="my-generic-local",
            path="/my-path/my-file.zip",
            file_path="/path/to/my-file.zip")
        my_base64_artifact = artifactory.Artifact("my-base64-artifact",
            repository="my-generic-local",
            path="/my-path/my-file.zip",
            content_base64="UEsDBAoAAAAAALl8alQAAAAAAAAAAAAAAAAJAAAATUVUQS1JTkYvUEsDBAoAAAAIALh8alTmUEsubQAAAIMAAAAUAAAATUVUQS1JTkYvTUFOSUZFU1QuTUbzTczLTEstLtENSy0qzszPs1Iw1DPg5XIsSs7ILEstQggH5KRWlBYrwCR4uZyLUhNLUlN0nSqtFBwLEpMzUhV8E8tS8xSM9cz0jHm5nEozc0rAsilAO1JzcjMhYim6XinZQGuA9uiZ83LxcgEAUEsDBAoAAAAAALh8alQAAAAAAAAAAAAAAAAMAAAAYXJ0aWZhY3RvcnkvUEsDBAoAAAAAALh8alQAAAAAAAAAAAAAAAARAAAAYXJ0aWZhY3RvcnkvdGVzdC9QSwMECgAAAAgAuHxqVMgzPcxdAQAALAIAAB0AAABhcnRpZmFjdG9yeS90ZXN0L011bHRpMS5jbGFzc3VRy07CQBQ9w6OlpQqCgPgEV+jCxsTEBcaNiXFRHwkGF66GOuKQPkyZmvBZutDEhR/gRxlvCwkxwVnck3vunXPPnfn++fwCcIRdExpWDdRQL6BhYg1NHes6Nhi0ExlIdcqQ7ez1GXJn4YNgKDkyEFexPxDRLR94xFSc0OVen0cyyWdkTj3JMUPT4ZGSj9xVYTSxlRgr+zL2lDzsUovPZcBQ79w7I/7CbY8HQ7unIhkMu+lAHg1JorqgzGD2wjhyxblMhhWnmgdJnwUdBR2bFrawzWBdCM8LW3dh5D20dexYaKHN0PjHFkN5Pux6MBKu+kP1JmMlfHqSMKZCbepMhvYN2VJkTnCfzFUX0Az6c5J5tHKts2hjtJGnz0hOBixZg6JBmU3ICPP7H2CvadmkqKWkhiJFa9pAuERoYBml2eXjVIy4N2Qq2Xfk5gImIWhKgRrnIgbKWCGkj007q79QSwECFAMKAAAAAAC5fGpUAAAAAAAAAAAAAAAACQAAAAAAAAAAABAA7UEAAAAATUVUQS1JTkYvUEsBAhQDCgAAAAgAuHxqVOZQSy5tAAAAgwAAABQAAAAAAAAAAAAAAKSBJwAAAE1FVEEtSU5GL01BTklGRVNULk1GUEsBAhQDCgAAAAAAuHxqVAAAAAAAAAAAAAAAAAwAAAAAAAAAAAAQAO1BxgAAAGFydGlmYWN0b3J5L1BLAQIUAwoAAAAAALh8alQAAAAAAAAAAAAAAAARAAAAAAAAAAAAEADtQfAAAABhcnRpZmFjdG9yeS90ZXN0L1BLAQIUAwoAAAAIALh8alTIMz3MXQEAACwCAAAdAAAAAAAAAAAAAACkgR8BAABhcnRpZmFjdG9yeS90ZXN0L011bHRpMS5jbGFzc1BLBQYAAAAABQAFAD0BAAC3AgAAAAA=")
        ```

        :param str resource_name: The name of the resource.
        :param ArtifactArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ArtifactArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ArtifactArgs.__new__(ArtifactArgs)

            __props__.__dict__["content_base64"] = content_base64
            __props__.__dict__["file_path"] = file_path
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["checksum_md5"] = None
            __props__.__dict__["checksum_sha1"] = None
            __props__.__dict__["checksum_sha256"] = None
            __props__.__dict__["created"] = None
            __props__.__dict__["created_by"] = None
            __props__.__dict__["download_uri"] = None
            __props__.__dict__["mime_type"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["uri"] = None
        super(Artifact, __self__).__init__(
            'artifactory:index/artifact:Artifact',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            checksum_md5: Optional[pulumi.Input[str]] = None,
            checksum_sha1: Optional[pulumi.Input[str]] = None,
            checksum_sha256: Optional[pulumi.Input[str]] = None,
            content_base64: Optional[pulumi.Input[str]] = None,
            created: Optional[pulumi.Input[str]] = None,
            created_by: Optional[pulumi.Input[str]] = None,
            download_uri: Optional[pulumi.Input[str]] = None,
            file_path: Optional[pulumi.Input[str]] = None,
            mime_type: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'Artifact':
        """
        Get an existing Artifact resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] checksum_md5: MD5 checksum of the artifact.
        :param pulumi.Input[str] checksum_sha1: SHA1 checksum of the artifact.
        :param pulumi.Input[str] checksum_sha256: SHA256 checksum of the artifact.
        :param pulumi.Input[str] content_base64: Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
        :param pulumi.Input[str] created: Timestamp when artifact is created.
        :param pulumi.Input[str] created_by: User who deploys the artifact.
        :param pulumi.Input[str] download_uri: Download URI of the artifact.
        :param pulumi.Input[str] file_path: Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
        :param pulumi.Input[str] mime_type: MIME type of the artifact.
        :param pulumi.Input[str] path: The relative path in the target repository. Must begin with a '/'. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
        :param pulumi.Input[str] repository: Name of the respository.
        :param pulumi.Input[int] size: Size of the artifact, in bytes.
        :param pulumi.Input[str] uri: URI of the artifact.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ArtifactState.__new__(_ArtifactState)

        __props__.__dict__["checksum_md5"] = checksum_md5
        __props__.__dict__["checksum_sha1"] = checksum_sha1
        __props__.__dict__["checksum_sha256"] = checksum_sha256
        __props__.__dict__["content_base64"] = content_base64
        __props__.__dict__["created"] = created
        __props__.__dict__["created_by"] = created_by
        __props__.__dict__["download_uri"] = download_uri
        __props__.__dict__["file_path"] = file_path
        __props__.__dict__["mime_type"] = mime_type
        __props__.__dict__["path"] = path
        __props__.__dict__["repository"] = repository
        __props__.__dict__["size"] = size
        __props__.__dict__["uri"] = uri
        return Artifact(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="checksumMd5")
    def checksum_md5(self) -> pulumi.Output[str]:
        """
        MD5 checksum of the artifact.
        """
        return pulumi.get(self, "checksum_md5")

    @property
    @pulumi.getter(name="checksumSha1")
    def checksum_sha1(self) -> pulumi.Output[str]:
        """
        SHA1 checksum of the artifact.
        """
        return pulumi.get(self, "checksum_sha1")

    @property
    @pulumi.getter(name="checksumSha256")
    def checksum_sha256(self) -> pulumi.Output[str]:
        """
        SHA256 checksum of the artifact.
        """
        return pulumi.get(self, "checksum_sha256")

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> pulumi.Output[Optional[str]]:
        """
        Base64 content of the source file. Conflicts with `file_path`. Either one of these attribute must be set.
        """
        return pulumi.get(self, "content_base64")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        Timestamp when artifact is created.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> pulumi.Output[str]:
        """
        User who deploys the artifact.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="downloadUri")
    def download_uri(self) -> pulumi.Output[str]:
        """
        Download URI of the artifact.
        """
        return pulumi.get(self, "download_uri")

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> pulumi.Output[Optional[str]]:
        """
        Path to the source file. Conflicts with `content_base64`. Either one of these attribute must be set.
        """
        return pulumi.get(self, "file_path")

    @property
    @pulumi.getter(name="mimeType")
    def mime_type(self) -> pulumi.Output[str]:
        """
        MIME type of the artifact.
        """
        return pulumi.get(self, "mime_type")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The relative path in the target repository. Must begin with a '/'. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        Name of the respository.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Size of the artifact, in bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        URI of the artifact.
        """
        return pulumi.get(self, "uri")
