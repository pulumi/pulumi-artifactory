// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # Artifactory Group Data Source
 *
 * Provides an Artifactory group datasource. This can be used to read the configuration of groups in artifactory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * //
 * const myGroup = artifactory.getGroup({
 *     name: "my_group",
 *     includeUsers: "true",
 * });
 * ```
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getGroup:getGroup", {
        "adminPrivileges": args.adminPrivileges,
        "autoJoin": args.autoJoin,
        "description": args.description,
        "externalId": args.externalId,
        "includeUsers": args.includeUsers,
        "name": args.name,
        "policyManager": args.policyManager,
        "realm": args.realm,
        "realmAttributes": args.realmAttributes,
        "reportsManager": args.reportsManager,
        "usersNames": args.usersNames,
        "watchManager": args.watchManager,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     */
    adminPrivileges?: boolean;
    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     */
    autoJoin?: boolean;
    /**
     * A description for the group
     */
    description?: string;
    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     */
    externalId?: string;
    /**
     * Determines if the group's associated user list will return as an attribute. Default is `false`.
     */
    includeUsers?: string;
    /**
     * Name of the group.
     */
    name: string;
    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     */
    policyManager?: boolean;
    /**
     * The realm for the group.
     */
    realm?: string;
    /**
     * The realm attributes for the group.
     */
    realmAttributes?: string;
    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     */
    reportsManager?: boolean;
    /**
     * List of users assigned to the group. Set includeUsers to `true` to retrieve this list.
     */
    usersNames?: string[];
    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     */
    watchManager?: boolean;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     */
    readonly adminPrivileges: boolean;
    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     */
    readonly autoJoin: boolean;
    /**
     * A description for the group
     */
    readonly description?: string;
    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     */
    readonly externalId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeUsers?: string;
    readonly name: string;
    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     */
    readonly policyManager?: boolean;
    /**
     * The realm for the group.
     */
    readonly realm: string;
    /**
     * The realm attributes for the group.
     */
    readonly realmAttributes?: string;
    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     */
    readonly reportsManager?: boolean;
    /**
     * List of users assigned to the group. Set includeUsers to `true` to retrieve this list.
     */
    readonly usersNames?: string[];
    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     */
    readonly watchManager?: boolean;
}
/**
 * ## # Artifactory Group Data Source
 *
 * Provides an Artifactory group datasource. This can be used to read the configuration of groups in artifactory.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * //
 * const myGroup = artifactory.getGroup({
 *     name: "my_group",
 *     includeUsers: "true",
 * });
 * ```
 */
export function getGroupOutput(args: GetGroupOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("artifactory:index/getGroup:getGroup", {
        "adminPrivileges": args.adminPrivileges,
        "autoJoin": args.autoJoin,
        "description": args.description,
        "externalId": args.externalId,
        "includeUsers": args.includeUsers,
        "name": args.name,
        "policyManager": args.policyManager,
        "realm": args.realm,
        "realmAttributes": args.realmAttributes,
        "reportsManager": args.reportsManager,
        "usersNames": args.usersNames,
        "watchManager": args.watchManager,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupOutputArgs {
    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     */
    adminPrivileges?: pulumi.Input<boolean>;
    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     */
    autoJoin?: pulumi.Input<boolean>;
    /**
     * A description for the group
     */
    description?: pulumi.Input<string>;
    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     */
    externalId?: pulumi.Input<string>;
    /**
     * Determines if the group's associated user list will return as an attribute. Default is `false`.
     */
    includeUsers?: pulumi.Input<string>;
    /**
     * Name of the group.
     */
    name: pulumi.Input<string>;
    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     */
    policyManager?: pulumi.Input<boolean>;
    /**
     * The realm for the group.
     */
    realm?: pulumi.Input<string>;
    /**
     * The realm attributes for the group.
     */
    realmAttributes?: pulumi.Input<string>;
    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     */
    reportsManager?: pulumi.Input<boolean>;
    /**
     * List of users assigned to the group. Set includeUsers to `true` to retrieve this list.
     */
    usersNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     */
    watchManager?: pulumi.Input<boolean>;
}
