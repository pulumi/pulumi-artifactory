// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a local repository replication resource, also referred to as Artifactory push replication. This can be used to create and manage Artifactory local repository replications using [Push Replication API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-SetRepositoryReplicationConfiguration).
 *
 * Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near end invoking a synchronization of artifacts to the far end.
 * See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).
 *
 * This resource can create the replication of local repository to single repository on the remote server.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const config = new pulumi.Config();
 * // The base URL of the Artifactory deployment
 * const artifactoryUrl = config.require("artifactoryUrl");
 * // The username for the Artifactory
 * const artifactoryUsername = config.require("artifactoryUsername");
 * // The password for the Artifactory
 * const artifactoryPassword = config.require("artifactoryPassword");
 * // Create a replication between two artifactory local repositories
 * const providerTestSource = new artifactory.LocalMavenRepository("provider_test_source", {key: "provider_test_source"});
 * const providerTestDest = new artifactory.LocalMavenRepository("provider_test_dest", {key: "provider_test_dest"});
 * const foo_rep = new artifactory.LocalRepositorySingleReplication("foo-rep", {
 *     repoKey: providerTestSource.key,
 *     cronExp: "0 0 * * * ?",
 *     enableEventReplication: true,
 *     url: pulumi.interpolate`${artifactoryUrl}/artifactory/${providerTestDest.key}`,
 *     username: "$var.artifactory_username",
 *     password: "$var.artifactory_password",
 *     enabled: true,
 *     socketTimeoutMillis: 16000,
 *     syncDeletes: false,
 *     syncProperties: true,
 *     syncStatistics: true,
 *     includePathPrefixPattern: "/some-repo/",
 *     excludePathPrefixPattern: "/some-other-repo/",
 *     checkBinaryExistenceInFilestore: true,
 * });
 * ```
 *
 * ## Import
 *
 * Push replication configs can be imported using their repo key, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication foo-rep provider_test_source
 * ```
 */
export class LocalRepositorySingleReplication extends pulumi.CustomResource {
    /**
     * Get an existing LocalRepositorySingleReplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocalRepositorySingleReplicationState, opts?: pulumi.CustomResourceOptions): LocalRepositorySingleReplication {
        return new LocalRepositorySingleReplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/localRepositorySingleReplication:LocalRepositorySingleReplication';

    /**
     * Returns true if the given object is an instance of LocalRepositorySingleReplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocalRepositorySingleReplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocalRepositorySingleReplication.__pulumiType;
    }

    /**
     * Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     */
    public readonly checkBinaryExistenceInFilestore!: pulumi.Output<boolean>;
    /**
     * A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     */
    public readonly cronExp!: pulumi.Output<string>;
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
     */
    public readonly enableEventReplication!: pulumi.Output<boolean>;
    /**
     * When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`. By default, no artifacts are excluded.
     */
    public readonly excludePathPrefixPattern!: pulumi.Output<string>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**&#47;*)`.
     */
    public readonly includePathPrefixPattern!: pulumi.Output<string>;
    /**
     * Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     */
    public readonly proxy!: pulumi.Output<string>;
    /**
     * Replication ID, the value is unknown until the resource is created. Can't be set or updated.
     */
    public /*out*/ readonly replicationKey!: pulumi.Output<string>;
    /**
     * Repository name.
     */
    public readonly repoKey!: pulumi.Output<string>;
    /**
     * The network timeout in milliseconds to use for remote operations. Default value is `15000`.
     */
    public readonly socketTimeoutMillis!: pulumi.Output<number>;
    /**
     * When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
     */
    public readonly syncDeletes!: pulumi.Output<boolean>;
    /**
     * When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
     */
    public readonly syncProperties!: pulumi.Output<boolean>;
    /**
     * When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
     */
    public readonly syncStatistics!: pulumi.Output<boolean>;
    /**
     * The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Username on the remote Artifactory instance.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a LocalRepositorySingleReplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocalRepositorySingleReplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocalRepositorySingleReplicationArgs | LocalRepositorySingleReplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocalRepositorySingleReplicationState | undefined;
            resourceInputs["checkBinaryExistenceInFilestore"] = state ? state.checkBinaryExistenceInFilestore : undefined;
            resourceInputs["cronExp"] = state ? state.cronExp : undefined;
            resourceInputs["enableEventReplication"] = state ? state.enableEventReplication : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["excludePathPrefixPattern"] = state ? state.excludePathPrefixPattern : undefined;
            resourceInputs["includePathPrefixPattern"] = state ? state.includePathPrefixPattern : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["replicationKey"] = state ? state.replicationKey : undefined;
            resourceInputs["repoKey"] = state ? state.repoKey : undefined;
            resourceInputs["socketTimeoutMillis"] = state ? state.socketTimeoutMillis : undefined;
            resourceInputs["syncDeletes"] = state ? state.syncDeletes : undefined;
            resourceInputs["syncProperties"] = state ? state.syncProperties : undefined;
            resourceInputs["syncStatistics"] = state ? state.syncStatistics : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as LocalRepositorySingleReplicationArgs | undefined;
            if ((!args || args.cronExp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cronExp'");
            }
            if ((!args || args.repoKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repoKey'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["checkBinaryExistenceInFilestore"] = args ? args.checkBinaryExistenceInFilestore : undefined;
            resourceInputs["cronExp"] = args ? args.cronExp : undefined;
            resourceInputs["enableEventReplication"] = args ? args.enableEventReplication : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["excludePathPrefixPattern"] = args ? args.excludePathPrefixPattern : undefined;
            resourceInputs["includePathPrefixPattern"] = args ? args.includePathPrefixPattern : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["repoKey"] = args ? args.repoKey : undefined;
            resourceInputs["socketTimeoutMillis"] = args ? args.socketTimeoutMillis : undefined;
            resourceInputs["syncDeletes"] = args ? args.syncDeletes : undefined;
            resourceInputs["syncProperties"] = args ? args.syncProperties : undefined;
            resourceInputs["syncStatistics"] = args ? args.syncStatistics : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["replicationKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(LocalRepositorySingleReplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocalRepositorySingleReplication resources.
 */
export interface LocalRepositorySingleReplicationState {
    /**
     * Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     */
    checkBinaryExistenceInFilestore?: pulumi.Input<boolean>;
    /**
     * A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     */
    cronExp?: pulumi.Input<string>;
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
     */
    enableEventReplication?: pulumi.Input<boolean>;
    /**
     * When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`. By default, no artifacts are excluded.
     */
    excludePathPrefixPattern?: pulumi.Input<string>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**&#47;*)`.
     */
    includePathPrefixPattern?: pulumi.Input<string>;
    /**
     * Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     */
    password?: pulumi.Input<string>;
    /**
     * Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     */
    proxy?: pulumi.Input<string>;
    /**
     * Replication ID, the value is unknown until the resource is created. Can't be set or updated.
     */
    replicationKey?: pulumi.Input<string>;
    /**
     * Repository name.
     */
    repoKey?: pulumi.Input<string>;
    /**
     * The network timeout in milliseconds to use for remote operations. Default value is `15000`.
     */
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
     */
    syncDeletes?: pulumi.Input<boolean>;
    /**
     * When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
     */
    syncProperties?: pulumi.Input<boolean>;
    /**
     * When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
     */
    syncStatistics?: pulumi.Input<boolean>;
    /**
     * The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
     */
    url?: pulumi.Input<string>;
    /**
     * Username on the remote Artifactory instance.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LocalRepositorySingleReplication resource.
 */
export interface LocalRepositorySingleReplicationArgs {
    /**
     * Enabling the `checkBinaryExistenceInFilestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     */
    checkBinaryExistenceInFilestore?: pulumi.Input<boolean>;
    /**
     * A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
     */
    cronExp: pulumi.Input<string>;
    /**
     * When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. add, deleted or property change. Default value is `false`.
     */
    enableEventReplication?: pulumi.Input<boolean>;
    /**
     * When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`. By default, no artifacts are excluded.
     */
    excludePathPrefixPattern?: pulumi.Input<string>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**&#47;*)`.
     */
    includePathPrefixPattern?: pulumi.Input<string>;
    /**
     * Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     */
    password?: pulumi.Input<string>;
    /**
     * Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     */
    proxy?: pulumi.Input<string>;
    /**
     * Repository name.
     */
    repoKey: pulumi.Input<string>;
    /**
     * The network timeout in milliseconds to use for remote operations. Default value is `15000`.
     */
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
     */
    syncDeletes?: pulumi.Input<boolean>;
    /**
     * When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
     */
    syncProperties?: pulumi.Input<boolean>;
    /**
     * When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
     */
    syncStatistics?: pulumi.Input<boolean>;
    /**
     * The URL of the target local repository on a remote Artifactory server. Use the format `https://<artifactory_url>/artifactory/<repository_name>`.
     */
    url: pulumi.Input<string>;
    /**
     * Username on the remote Artifactory instance.
     */
    username: pulumi.Input<string>;
}
