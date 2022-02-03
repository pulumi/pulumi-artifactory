CHANGELOG
=========

## Notice (2022-01-06)

*As of this notice, using CHANGELOG.md is DEPRECATED. We will be using [GitHub Releases](https://github.com/pulumi/pulumi-artifactory/releases) for this repository*

## HEAD (Unreleased)
_(none)_

---

## 0.1.2 (2022-01-06)
* Upgrade to v2.6.24 of the Jfrog Artifactory Terraform Provider
  **PLEASE NOTE:** The following breaking changes:
  * `artifactory.DockerV1Repository` no longer supports `indexCompressionFormats`
  * `artifactory.DockerV2Repository` no longer supports `indexCompressionFormats`
  * `artifactory.RemoteHelmRepository` no longer supports `failedRetrievalCachePeriodSecs` input
  * `artifactory.RemoteCargoRepository` no longer supports `failedRetrievalCachePeriodSecs` input
  * `artifactory.RemoteDockerRepository` no longer supports `failedRetrievalCachePeriodSecs` input

## 0.1.1 (2021-11-18)
* Upgrade to v2.6.18 of the Jfrog Artifactory Terraform Provider
    * (breaking) Replication password is now a [computed field](https://github.com/jfrog/terraform-provider-artifactory/pull/206) in ReplicationConfigArgs and can no longer be edited.

## 0.1.0 (2021-11-15)
* Scaffolded the provider based on v2.6.17 of the Artifactory Terraform Provider
