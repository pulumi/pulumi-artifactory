// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a remote Cargo repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote-cargo = artifactory.getRemoteCargoRepository({
 *     key: "remote-cargo",
 * });
 * ```
 */
export function getRemoteCargoRepository(args: GetRemoteCargoRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteCargoRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getRemoteCargoRepository:getRemoteCargoRepository", {
        "allowAnyHostAuth": args.allowAnyHostAuth,
        "anonymousAccess": args.anonymousAccess,
        "assumedOfflinePeriodSecs": args.assumedOfflinePeriodSecs,
        "blackedOut": args.blackedOut,
        "blockMismatchingMimeTypes": args.blockMismatchingMimeTypes,
        "bypassHeadRequests": args.bypassHeadRequests,
        "cdnRedirect": args.cdnRedirect,
        "clientTlsCertificate": args.clientTlsCertificate,
        "contentSynchronisation": args.contentSynchronisation,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "enableCookieManagement": args.enableCookieManagement,
        "enableSparseIndex": args.enableSparseIndex,
        "excludesPattern": args.excludesPattern,
        "gitRegistryUrl": args.gitRegistryUrl,
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
 * A collection of arguments for invoking getRemoteCargoRepository.
 */
export interface GetRemoteCargoRepositoryArgs {
    allowAnyHostAuth?: boolean;
    /**
     * (Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
     */
    anonymousAccess?: boolean;
    assumedOfflinePeriodSecs?: number;
    blackedOut?: boolean;
    blockMismatchingMimeTypes?: boolean;
    bypassHeadRequests?: boolean;
    cdnRedirect?: boolean;
    clientTlsCertificate?: string;
    contentSynchronisation?: inputs.GetRemoteCargoRepositoryContentSynchronisation;
    description?: string;
    downloadDirect?: boolean;
    enableCookieManagement?: boolean;
    /**
     * (Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
     */
    enableSparseIndex?: boolean;
    excludesPattern?: string;
    /**
     * (Optional) This is the index url, expected to be a git repository. Default value is `https://github.com/rust-lang/crates.io-index`.
     */
    gitRegistryUrl?: string;
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
 * A collection of values returned by getRemoteCargoRepository.
 */
export interface GetRemoteCargoRepositoryResult {
    readonly allowAnyHostAuth?: boolean;
    /**
     * (Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
     */
    readonly anonymousAccess?: boolean;
    readonly assumedOfflinePeriodSecs?: number;
    readonly blackedOut?: boolean;
    readonly blockMismatchingMimeTypes?: boolean;
    readonly bypassHeadRequests?: boolean;
    readonly cdnRedirect?: boolean;
    readonly clientTlsCertificate: string;
    readonly contentSynchronisation: outputs.GetRemoteCargoRepositoryContentSynchronisation;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly enableCookieManagement?: boolean;
    /**
     * (Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
     */
    readonly enableSparseIndex?: boolean;
    readonly excludesPattern?: string;
    /**
     * (Optional) This is the index url, expected to be a git repository. Default value is `https://github.com/rust-lang/crates.io-index`.
     */
    readonly gitRegistryUrl?: string;
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
 * Retrieves a remote Cargo repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const remote-cargo = artifactory.getRemoteCargoRepository({
 *     key: "remote-cargo",
 * });
 * ```
 */
export function getRemoteCargoRepositoryOutput(args: GetRemoteCargoRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRemoteCargoRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getRemoteCargoRepository(a, opts))
}

/**
 * A collection of arguments for invoking getRemoteCargoRepository.
 */
export interface GetRemoteCargoRepositoryOutputArgs {
    allowAnyHostAuth?: pulumi.Input<boolean>;
    /**
     * (Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is `false`.
     */
    anonymousAccess?: pulumi.Input<boolean>;
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    blackedOut?: pulumi.Input<boolean>;
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    bypassHeadRequests?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    contentSynchronisation?: pulumi.Input<inputs.GetRemoteCargoRepositoryContentSynchronisationArgs>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    enableCookieManagement?: pulumi.Input<boolean>;
    /**
     * (Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is `false`.
     */
    enableSparseIndex?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    /**
     * (Optional) This is the index url, expected to be a git repository. Default value is `https://github.com/rust-lang/crates.io-index`.
     */
    gitRegistryUrl?: pulumi.Input<string>;
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