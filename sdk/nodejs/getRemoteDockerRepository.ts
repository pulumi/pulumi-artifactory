// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a remote Docker repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote_docker = artifactory.getRemoteDockerRepository({
 *     key: "remote-docker",
 * });
 * ```
 */
export function getRemoteDockerRepository(args: GetRemoteDockerRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteDockerRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getRemoteDockerRepository:getRemoteDockerRepository", {
        "allowAnyHostAuth": args.allowAnyHostAuth,
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "assumedOfflinePeriodSecs": args.assumedOfflinePeriodSecs,
        "blackedOut": args.blackedOut,
        "blockMismatchingMimeTypes": args.blockMismatchingMimeTypes,
        "blockPushingSchema1": args.blockPushingSchema1,
        "bypassHeadRequests": args.bypassHeadRequests,
        "cdnRedirect": args.cdnRedirect,
        "clientTlsCertificate": args.clientTlsCertificate,
        "contentSynchronisation": args.contentSynchronisation,
        "curated": args.curated,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "disableUrlNormalization": args.disableUrlNormalization,
        "downloadDirect": args.downloadDirect,
        "enableCookieManagement": args.enableCookieManagement,
        "enableTokenAuthentication": args.enableTokenAuthentication,
        "excludesPattern": args.excludesPattern,
        "externalDependenciesEnabled": args.externalDependenciesEnabled,
        "externalDependenciesPatterns": args.externalDependenciesPatterns,
        "hardFail": args.hardFail,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "listRemoteFolderItems": args.listRemoteFolderItems,
        "localAddress": args.localAddress,
        "metadataRetrievalTimeoutSecs": args.metadataRetrievalTimeoutSecs,
        "mismatchingMimeTypesOverrideList": args.mismatchingMimeTypesOverrideList,
        "missedCachePeriodSeconds": args.missedCachePeriodSeconds,
        "notes": args.notes,
        "offline": args.offline,
        "password": args.password,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectId": args.projectId,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "queryParams": args.queryParams,
        "remoteRepoLayoutRef": args.remoteRepoLayoutRef,
        "repoLayoutRef": args.repoLayoutRef,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "shareConfiguration": args.shareConfiguration,
        "socketTimeoutMillis": args.socketTimeoutMillis,
        "storeArtifactsLocally": args.storeArtifactsLocally,
        "synchronizeProperties": args.synchronizeProperties,
        "unusedArtifactsCleanupPeriodHours": args.unusedArtifactsCleanupPeriodHours,
        "url": args.url,
        "username": args.username,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getRemoteDockerRepository.
 */
export interface GetRemoteDockerRepositoryArgs {
    allowAnyHostAuth?: boolean;
    archiveBrowsingEnabled?: boolean;
    assumedOfflinePeriodSecs?: number;
    blackedOut?: boolean;
    blockMismatchingMimeTypes?: boolean;
    /**
     * (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
     */
    blockPushingSchema1?: boolean;
    bypassHeadRequests?: boolean;
    cdnRedirect?: boolean;
    clientTlsCertificate?: string;
    contentSynchronisation?: inputs.GetRemoteDockerRepositoryContentSynchronisation;
    curated?: boolean;
    description?: string;
    disableProxy?: boolean;
    /**
     * (Optional) Whether to disable URL normalization.
     */
    disableUrlNormalization?: boolean;
    downloadDirect?: boolean;
    enableCookieManagement?: boolean;
    /**
     * (Optional) Enable token (Bearer) based authentication.
     */
    enableTokenAuthentication?: boolean;
    excludesPattern?: string;
    /**
     * (Optional) Also known as 'Foreign Layers Caching' on the UI.
     */
    externalDependenciesEnabled?: boolean;
    /**
     * (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
     */
    externalDependenciesPatterns?: string[];
    hardFail?: boolean;
    includesPattern?: string;
    /**
     * the identity key of the repo.
     */
    key: string;
    listRemoteFolderItems?: boolean;
    localAddress?: string;
    metadataRetrievalTimeoutSecs?: number;
    mismatchingMimeTypesOverrideList?: string;
    missedCachePeriodSeconds?: number;
    notes?: string;
    offline?: boolean;
    password?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectId?: string;
    projectKey?: string;
    propertySets?: string[];
    proxy?: string;
    queryParams?: string;
    remoteRepoLayoutRef?: string;
    repoLayoutRef?: string;
    retrievalCachePeriodSeconds?: number;
    shareConfiguration?: boolean;
    socketTimeoutMillis?: number;
    storeArtifactsLocally?: boolean;
    synchronizeProperties?: boolean;
    unusedArtifactsCleanupPeriodHours?: number;
    url?: string;
    username?: string;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getRemoteDockerRepository.
 */
export interface GetRemoteDockerRepositoryResult {
    readonly allowAnyHostAuth?: boolean;
    readonly archiveBrowsingEnabled?: boolean;
    readonly assumedOfflinePeriodSecs?: number;
    readonly blackedOut?: boolean;
    readonly blockMismatchingMimeTypes?: boolean;
    /**
     * (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
     */
    readonly blockPushingSchema1: boolean;
    readonly bypassHeadRequests?: boolean;
    readonly cdnRedirect?: boolean;
    readonly clientTlsCertificate: string;
    readonly contentSynchronisation: outputs.GetRemoteDockerRepositoryContentSynchronisation;
    readonly curated?: boolean;
    readonly description?: string;
    readonly disableProxy?: boolean;
    /**
     * (Optional) Whether to disable URL normalization.
     */
    readonly disableUrlNormalization?: boolean;
    readonly downloadDirect?: boolean;
    readonly enableCookieManagement?: boolean;
    /**
     * (Optional) Enable token (Bearer) based authentication.
     */
    readonly enableTokenAuthentication: boolean;
    readonly excludesPattern?: string;
    /**
     * (Optional) Also known as 'Foreign Layers Caching' on the UI.
     */
    readonly externalDependenciesEnabled?: boolean;
    /**
     * (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
     */
    readonly externalDependenciesPatterns?: string[];
    readonly hardFail?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    readonly listRemoteFolderItems?: boolean;
    readonly localAddress?: string;
    readonly metadataRetrievalTimeoutSecs?: number;
    readonly mismatchingMimeTypesOverrideList?: string;
    readonly missedCachePeriodSeconds?: number;
    readonly notes?: string;
    readonly offline?: boolean;
    readonly packageType: string;
    readonly password?: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectId?: string;
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly proxy?: string;
    readonly queryParams?: string;
    readonly remoteRepoLayoutRef?: string;
    readonly repoLayoutRef?: string;
    readonly retrievalCachePeriodSeconds?: number;
    readonly shareConfiguration: boolean;
    readonly socketTimeoutMillis?: number;
    readonly storeArtifactsLocally?: boolean;
    readonly synchronizeProperties?: boolean;
    readonly unusedArtifactsCleanupPeriodHours?: number;
    readonly url?: string;
    readonly username?: string;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a remote Docker repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote_docker = artifactory.getRemoteDockerRepository({
 *     key: "remote-docker",
 * });
 * ```
 */
export function getRemoteDockerRepositoryOutput(args: GetRemoteDockerRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRemoteDockerRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getRemoteDockerRepository:getRemoteDockerRepository", {
        "allowAnyHostAuth": args.allowAnyHostAuth,
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "assumedOfflinePeriodSecs": args.assumedOfflinePeriodSecs,
        "blackedOut": args.blackedOut,
        "blockMismatchingMimeTypes": args.blockMismatchingMimeTypes,
        "blockPushingSchema1": args.blockPushingSchema1,
        "bypassHeadRequests": args.bypassHeadRequests,
        "cdnRedirect": args.cdnRedirect,
        "clientTlsCertificate": args.clientTlsCertificate,
        "contentSynchronisation": args.contentSynchronisation,
        "curated": args.curated,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "disableUrlNormalization": args.disableUrlNormalization,
        "downloadDirect": args.downloadDirect,
        "enableCookieManagement": args.enableCookieManagement,
        "enableTokenAuthentication": args.enableTokenAuthentication,
        "excludesPattern": args.excludesPattern,
        "externalDependenciesEnabled": args.externalDependenciesEnabled,
        "externalDependenciesPatterns": args.externalDependenciesPatterns,
        "hardFail": args.hardFail,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "listRemoteFolderItems": args.listRemoteFolderItems,
        "localAddress": args.localAddress,
        "metadataRetrievalTimeoutSecs": args.metadataRetrievalTimeoutSecs,
        "mismatchingMimeTypesOverrideList": args.mismatchingMimeTypesOverrideList,
        "missedCachePeriodSeconds": args.missedCachePeriodSeconds,
        "notes": args.notes,
        "offline": args.offline,
        "password": args.password,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectId": args.projectId,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "queryParams": args.queryParams,
        "remoteRepoLayoutRef": args.remoteRepoLayoutRef,
        "repoLayoutRef": args.repoLayoutRef,
        "retrievalCachePeriodSeconds": args.retrievalCachePeriodSeconds,
        "shareConfiguration": args.shareConfiguration,
        "socketTimeoutMillis": args.socketTimeoutMillis,
        "storeArtifactsLocally": args.storeArtifactsLocally,
        "synchronizeProperties": args.synchronizeProperties,
        "unusedArtifactsCleanupPeriodHours": args.unusedArtifactsCleanupPeriodHours,
        "url": args.url,
        "username": args.username,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getRemoteDockerRepository.
 */
export interface GetRemoteDockerRepositoryOutputArgs {
    allowAnyHostAuth?: pulumi.Input<boolean>;
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    blackedOut?: pulumi.Input<boolean>;
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    /**
     * (Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.
     */
    blockPushingSchema1?: pulumi.Input<boolean>;
    bypassHeadRequests?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    contentSynchronisation?: pulumi.Input<inputs.GetRemoteDockerRepositoryContentSynchronisationArgs>;
    curated?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    disableProxy?: pulumi.Input<boolean>;
    /**
     * (Optional) Whether to disable URL normalization.
     */
    disableUrlNormalization?: pulumi.Input<boolean>;
    downloadDirect?: pulumi.Input<boolean>;
    enableCookieManagement?: pulumi.Input<boolean>;
    /**
     * (Optional) Enable token (Bearer) based authentication.
     */
    enableTokenAuthentication?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * (Optional) Also known as 'Foreign Layers Caching' on the UI.
     */
    externalDependenciesEnabled?: pulumi.Input<boolean>;
    /**
     * (Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to `[**]` in the UI, which means that remote modules may be downloaded from any external VCS source. Due to SDKv2 limitations, we can't set the default value for the list. This value `[**]` must be assigned to the attribute manually, if user don't specify any other non-default values. We don't want to make this attribute required, but it must be set to avoid the state drift on update. Note: Artifactory assigns `[**]` on update if HCL doesn't have the attribute set or the list is empty.
     */
    externalDependenciesPatterns?: pulumi.Input<pulumi.Input<string>[]>;
    hardFail?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    listRemoteFolderItems?: pulumi.Input<boolean>;
    localAddress?: pulumi.Input<string>;
    metadataRetrievalTimeoutSecs?: pulumi.Input<number>;
    mismatchingMimeTypesOverrideList?: pulumi.Input<string>;
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    offline?: pulumi.Input<boolean>;
    password?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectId?: pulumi.Input<string>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    queryParams?: pulumi.Input<string>;
    remoteRepoLayoutRef?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    socketTimeoutMillis?: pulumi.Input<number>;
    storeArtifactsLocally?: pulumi.Input<boolean>;
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    xrayIndex?: pulumi.Input<boolean>;
}
