// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves a virtual Debian repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const virtual_debian = artifactory.getVirtualDebianRepository({
 *     key: "virtual-debian",
 * });
 * ```
 */
export function getVirtualDebianRepository(args: GetVirtualDebianRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualDebianRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getVirtualDebianRepository:getVirtualDebianRepository", {
        "artifactoryRequestsCanRetrieveRemoteArtifacts": args.artifactoryRequestsCanRetrieveRemoteArtifacts,
        "debianDefaultArchitectures": args.debianDefaultArchitectures,
        "defaultDeploymentRepo": args.defaultDeploymentRepo,
        "description": args.description,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "notes": args.notes,
        "optionalIndexCompressionFormats": args.optionalIndexCompressionFormats,
        "primaryKeypairRef": args.primaryKeypairRef,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "repoLayoutRef": args.repoLayoutRef,
        "repositories": args.repositories,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "secondaryKeypairRef": args.secondaryKeypairRef,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualDebianRepository.
 */
export interface GetVirtualDebianRepositoryArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: boolean;
    /**
     * (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
     */
    debianDefaultArchitectures?: string;
    defaultDeploymentRepo?: string;
    description?: string;
    excludesPattern?: string;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    notes?: string;
    /**
     * (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
     */
    optionalIndexCompressionFormats?: string[];
    /**
     * (Optional) Primary keypair used to sign artifacts. Default is empty.
     */
    primaryKeypairRef?: string;
    projectEnvironments?: string[];
    projectKey?: string;
    repoLayoutRef?: string;
    repositories?: string[];
    /**
     * (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: number;
    /**
     * (Optional) Secondary keypair used to sign artifacts. Default is empty.
     */
    secondaryKeypairRef?: string;
}

/**
 * A collection of values returned by getVirtualDebianRepository.
 */
export interface GetVirtualDebianRepositoryResult {
    readonly artifactoryRequestsCanRetrieveRemoteArtifacts?: boolean;
    /**
     * (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
     */
    readonly debianDefaultArchitectures?: string;
    readonly defaultDeploymentRepo?: string;
    readonly description?: string;
    readonly excludesPattern?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    readonly notes?: string;
    /**
     * (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
     */
    readonly optionalIndexCompressionFormats: string[];
    readonly packageType: string;
    /**
     * (Optional) Primary keypair used to sign artifacts. Default is empty.
     */
    readonly primaryKeypairRef?: string;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly repoLayoutRef?: string;
    readonly repositories?: string[];
    /**
     * (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     */
    readonly retrievalCachePeriodSeconds?: number;
    /**
     * (Optional) Secondary keypair used to sign artifacts. Default is empty.
     */
    readonly secondaryKeypairRef?: string;
}
/**
 * Retrieves a virtual Debian repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const virtual_debian = artifactory.getVirtualDebianRepository({
 *     key: "virtual-debian",
 * });
 * ```
 */
export function getVirtualDebianRepositoryOutput(args: GetVirtualDebianRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVirtualDebianRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getVirtualDebianRepository:getVirtualDebianRepository", {
        "artifactoryRequestsCanRetrieveRemoteArtifacts": args.artifactoryRequestsCanRetrieveRemoteArtifacts,
        "debianDefaultArchitectures": args.debianDefaultArchitectures,
        "defaultDeploymentRepo": args.defaultDeploymentRepo,
        "description": args.description,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "notes": args.notes,
        "optionalIndexCompressionFormats": args.optionalIndexCompressionFormats,
        "primaryKeypairRef": args.primaryKeypairRef,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "repoLayoutRef": args.repoLayoutRef,
        "repositories": args.repositories,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "secondaryKeypairRef": args.secondaryKeypairRef,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualDebianRepository.
 */
export interface GetVirtualDebianRepositoryOutputArgs {
    artifactoryRequestsCanRetrieveRemoteArtifacts?: pulumi.Input<boolean>;
    /**
     * (Optional) Specifying  architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.
     */
    debianDefaultArchitectures?: pulumi.Input<string>;
    defaultDeploymentRepo?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    /**
     * (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
     */
    optionalIndexCompressionFormats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional) Primary keypair used to sign artifacts. Default is empty.
     */
    primaryKeypairRef?: pulumi.Input<string>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    repositories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    /**
     * (Optional) Secondary keypair used to sign artifacts. Default is empty.
     */
    secondaryKeypairRef?: pulumi.Input<string>;
}
