// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get a flat (the default) or deep listing of the files and folders (not included by default) within a folder. For deep listing you can specify an optional depth to limit the results. Optionally include a map of metadata timestamp values as part of the result.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my-repo-file-list = artifactory.getFileList({
 *     folderPath: "path/to/artifact",
 *     repositoryKey: "my-generic-local",
 * });
 * ```
 */
export function getFileList(args: GetFileListArgs, opts?: pulumi.InvokeOptions): Promise<GetFileListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getFileList:getFileList", {
        "deepListing": args.deepListing,
        "depth": args.depth,
        "folderPath": args.folderPath,
        "includeRootPath": args.includeRootPath,
        "listFolders": args.listFolders,
        "metadataTimestamps": args.metadataTimestamps,
        "repositoryKey": args.repositoryKey,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileList.
 */
export interface GetFileListArgs {
    /**
     * Get deep listing
     */
    deepListing?: boolean;
    /**
     * Depth of the deep listing
     */
    depth?: number;
    /**
     * Path of the folder
     */
    folderPath: string;
    /**
     * Include root path
     */
    includeRootPath?: boolean;
    /**
     * Include folders
     */
    listFolders?: boolean;
    /**
     * File metadata
     */
    metadataTimestamps?: boolean;
    /**
     * Repository key
     */
    repositoryKey: string;
}

/**
 * A collection of values returned by getFileList.
 */
export interface GetFileListResult {
    /**
     * Creation time
     */
    readonly created: string;
    /**
     * Get deep listing
     */
    readonly deepListing?: boolean;
    /**
     * Depth of the deep listing
     */
    readonly depth?: number;
    /**
     * A list of files.
     */
    readonly files: outputs.GetFileListFile[];
    /**
     * Path of the folder
     */
    readonly folderPath: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Include root path
     */
    readonly includeRootPath?: boolean;
    /**
     * Include folders
     */
    readonly listFolders?: boolean;
    /**
     * Include metadata timestamps
     */
    readonly metadataTimestamps?: boolean;
    /**
     * Repository key
     */
    readonly repositoryKey: string;
    /**
     * URL to file/path
     */
    readonly uri: string;
}
/**
 * Get a flat (the default) or deep listing of the files and folders (not included by default) within a folder. For deep listing you can specify an optional depth to limit the results. Optionally include a map of metadata timestamp values as part of the result.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my-repo-file-list = artifactory.getFileList({
 *     folderPath: "path/to/artifact",
 *     repositoryKey: "my-generic-local",
 * });
 * ```
 */
export function getFileListOutput(args: GetFileListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFileListResult> {
    return pulumi.output(args).apply((a: any) => getFileList(a, opts))
}

/**
 * A collection of arguments for invoking getFileList.
 */
export interface GetFileListOutputArgs {
    /**
     * Get deep listing
     */
    deepListing?: pulumi.Input<boolean>;
    /**
     * Depth of the deep listing
     */
    depth?: pulumi.Input<number>;
    /**
     * Path of the folder
     */
    folderPath: pulumi.Input<string>;
    /**
     * Include root path
     */
    includeRootPath?: pulumi.Input<boolean>;
    /**
     * Include folders
     */
    listFolders?: pulumi.Input<boolean>;
    /**
     * File metadata
     */
    metadataTimestamps?: pulumi.Input<boolean>;
    /**
     * Repository key
     */
    repositoryKey: pulumi.Input<string>;
}