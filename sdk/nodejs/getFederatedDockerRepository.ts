// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getFederatedDockerRepository(args: GetFederatedDockerRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetFederatedDockerRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getFederatedDockerRepository:getFederatedDockerRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "blockPushingSchema1": args.blockPushingSchema1,
        "cdnRedirect": args.cdnRedirect,
        "cleanupOnDelete": args.cleanupOnDelete,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueTags": args.maxUniqueTags,
        "members": args.members,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "repoLayoutRef": args.repoLayoutRef,
        "tagRetention": args.tagRetention,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getFederatedDockerRepository.
 */
export interface GetFederatedDockerRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    blockPushingSchema1?: boolean;
    cdnRedirect?: boolean;
    cleanupOnDelete?: boolean;
    description?: string;
    disableProxy?: boolean;
    downloadDirect?: boolean;
    excludesPattern?: string;
    includesPattern?: string;
    key: string;
    maxUniqueTags?: number;
    members?: inputs.GetFederatedDockerRepositoryMember[];
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    proxy?: string;
    repoLayoutRef?: string;
    tagRetention?: number;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getFederatedDockerRepository.
 */
export interface GetFederatedDockerRepositoryResult {
    readonly apiVersion: string;
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly blockPushingSchema1: boolean;
    readonly cdnRedirect?: boolean;
    readonly cleanupOnDelete?: boolean;
    readonly description?: string;
    readonly disableProxy?: boolean;
    readonly downloadDirect?: boolean;
    readonly excludesPattern?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern?: string;
    readonly key: string;
    readonly maxUniqueTags?: number;
    readonly members?: outputs.GetFederatedDockerRepositoryMember[];
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly proxy?: string;
    readonly repoLayoutRef?: string;
    readonly tagRetention?: number;
    readonly xrayIndex?: boolean;
}
export function getFederatedDockerRepositoryOutput(args: GetFederatedDockerRepositoryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetFederatedDockerRepositoryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getFederatedDockerRepository:getFederatedDockerRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "blockPushingSchema1": args.blockPushingSchema1,
        "cdnRedirect": args.cdnRedirect,
        "cleanupOnDelete": args.cleanupOnDelete,
        "description": args.description,
        "disableProxy": args.disableProxy,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueTags": args.maxUniqueTags,
        "members": args.members,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "proxy": args.proxy,
        "repoLayoutRef": args.repoLayoutRef,
        "tagRetention": args.tagRetention,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getFederatedDockerRepository.
 */
export interface GetFederatedDockerRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    blockPushingSchema1?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    cleanupOnDelete?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    disableProxy?: pulumi.Input<boolean>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    key: pulumi.Input<string>;
    maxUniqueTags?: pulumi.Input<number>;
    members?: pulumi.Input<pulumi.Input<inputs.GetFederatedDockerRepositoryMemberArgs>[]>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    repoLayoutRef?: pulumi.Input<string>;
    tagRetention?: pulumi.Input<number>;
    xrayIndex?: pulumi.Input<boolean>;
}
