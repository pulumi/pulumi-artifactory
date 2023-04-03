// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieves a federated Debian repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const federated-test-debian-repo = artifactory.getFederatedDebianRepository({
 *     key: "federated-test-debian-repo",
 * });
 * ```
 */
export function getFederatedDebianRepository(args: GetFederatedDebianRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetFederatedDebianRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getFederatedDebianRepository:getFederatedDebianRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "includesPattern": args.includesPattern,
        "indexCompressionFormats": args.indexCompressionFormats,
        "key": args.key,
        "members": args.members,
        "notes": args.notes,
        "primaryKeypairRef": args.primaryKeypairRef,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "repoLayoutRef": args.repoLayoutRef,
        "secondaryKeypairRef": args.secondaryKeypairRef,
        "trivialLayout": args.trivialLayout,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getFederatedDebianRepository.
 */
export interface GetFederatedDebianRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    description?: string;
    downloadDirect?: boolean;
    excludesPattern?: string;
    includesPattern?: string;
    indexCompressionFormats?: string[];
    /**
     * the identity key of the repo.
     */
    key: string;
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    members?: inputs.GetFederatedDebianRepositoryMember[];
    notes?: string;
    primaryKeypairRef?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    secondaryKeypairRef?: string;
    /**
     * @deprecated You shouldn't be using this
     */
    trivialLayout?: boolean;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getFederatedDebianRepository.
 */
export interface GetFederatedDebianRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly excludesPattern: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern: string;
    readonly indexCompressionFormats?: string[];
    readonly key: string;
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    readonly members?: outputs.GetFederatedDebianRepositoryMember[];
    readonly notes?: string;
    readonly packageType: string;
    readonly primaryKeypairRef?: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    readonly secondaryKeypairRef?: string;
    /**
     * @deprecated You shouldn't be using this
     */
    readonly trivialLayout?: boolean;
    readonly xrayIndex?: boolean;
}
/**
 * Retrieves a federated Debian repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const federated-test-debian-repo = artifactory.getFederatedDebianRepository({
 *     key: "federated-test-debian-repo",
 * });
 * ```
 */
export function getFederatedDebianRepositoryOutput(args: GetFederatedDebianRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFederatedDebianRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getFederatedDebianRepository(a, opts))
}

/**
 * A collection of arguments for invoking getFederatedDebianRepository.
 */
export interface GetFederatedDebianRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    includesPattern?: pulumi.Input<string>;
    indexCompressionFormats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     */
    members?: pulumi.Input<pulumi.Input<inputs.GetFederatedDebianRepositoryMemberArgs>[]>;
    notes?: pulumi.Input<string>;
    primaryKeypairRef?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    secondaryKeypairRef?: pulumi.Input<string>;
    /**
     * @deprecated You shouldn't be using this
     */
    trivialLayout?: pulumi.Input<boolean>;
    xrayIndex?: pulumi.Input<boolean>;
}