// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory Archive Policy resource. This resource enable system administrators to define and customize policies based on specific criteria for removing unused binaries from across their JFrog platform. See [Retention Policies](https://jfrog.com/help/r/jfrog-platform-administration-documentation/retention-policies) for more details.
 *
 * ~>Currently in beta and not yet globally available. A full rollout is scheduled for Q1 2025.
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import artifactory:index/archivePolicy:ArchivePolicy my-archive-policy my-policy
 * ```
 *
 * ```sh
 * $ pulumi import artifactory:index/archivePolicy:ArchivePolicy my-archive-policy my-policy:myproj
 * ```
 */
export class ArchivePolicy extends pulumi.CustomResource {
    /**
     * Get an existing ArchivePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ArchivePolicyState, opts?: pulumi.CustomResourceOptions): ArchivePolicy {
        return new ArchivePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/archivePolicy:ArchivePolicy';

    /**
     * Returns true if the given object is an instance of ArchivePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ArchivePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ArchivePolicy.__pulumiType;
    }

    /**
     * The cron expression determines when the policy is run. This parameter is not mandatory, however if left empty the policy will not run automatically and can only be triggered manually.
     */
    public readonly cronExpression!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The maximum duration (in minutes) for policy execution, after which the policy will stop running even if not completed. While setting a maximum run duration for a policy is useful for adhering to a strict archive V2 schedule, it can cause the policy to stop before completion.
     */
    public readonly durationInMinutes!: pulumi.Output<number | undefined>;
    /**
     * Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * An ID that is used to identify the archive policy. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * This attribute is used only for project-level archive V2 policies, it is not used for global-level policies.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    public readonly searchCriteria!: pulumi.Output<outputs.ArchivePolicySearchCriteria>;
    /**
     * A `true` value means that when this policy is executed, packages will be permanently deleted. `false` means that when the policy is executed packages will be deleted to the Trash Can. Defaults to `false`.
     */
    public readonly skipTrashcan!: pulumi.Output<boolean>;

    /**
     * Create a ArchivePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArchivePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ArchivePolicyArgs | ArchivePolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ArchivePolicyState | undefined;
            resourceInputs["cronExpression"] = state ? state.cronExpression : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["durationInMinutes"] = state ? state.durationInMinutes : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["searchCriteria"] = state ? state.searchCriteria : undefined;
            resourceInputs["skipTrashcan"] = state ? state.skipTrashcan : undefined;
        } else {
            const args = argsOrState as ArchivePolicyArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.searchCriteria === undefined) && !opts.urn) {
                throw new Error("Missing required property 'searchCriteria'");
            }
            resourceInputs["cronExpression"] = args ? args.cronExpression : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["durationInMinutes"] = args ? args.durationInMinutes : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["searchCriteria"] = args ? args.searchCriteria : undefined;
            resourceInputs["skipTrashcan"] = args ? args.skipTrashcan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ArchivePolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ArchivePolicy resources.
 */
export interface ArchivePolicyState {
    /**
     * The cron expression determines when the policy is run. This parameter is not mandatory, however if left empty the policy will not run automatically and can only be triggered manually.
     */
    cronExpression?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * The maximum duration (in minutes) for policy execution, after which the policy will stop running even if not completed. While setting a maximum run duration for a policy is useful for adhering to a strict archive V2 schedule, it can cause the policy to stop before completion.
     */
    durationInMinutes?: pulumi.Input<number>;
    /**
     * Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * An ID that is used to identify the archive policy. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
     */
    key?: pulumi.Input<string>;
    /**
     * This attribute is used only for project-level archive V2 policies, it is not used for global-level policies.
     */
    projectKey?: pulumi.Input<string>;
    searchCriteria?: pulumi.Input<inputs.ArchivePolicySearchCriteria>;
    /**
     * A `true` value means that when this policy is executed, packages will be permanently deleted. `false` means that when the policy is executed packages will be deleted to the Trash Can. Defaults to `false`.
     */
    skipTrashcan?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ArchivePolicy resource.
 */
export interface ArchivePolicyArgs {
    /**
     * The cron expression determines when the policy is run. This parameter is not mandatory, however if left empty the policy will not run automatically and can only be triggered manually.
     */
    cronExpression?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * The maximum duration (in minutes) for policy execution, after which the policy will stop running even if not completed. While setting a maximum run duration for a policy is useful for adhering to a strict archive V2 schedule, it can cause the policy to stop before completion.
     */
    durationInMinutes?: pulumi.Input<number>;
    /**
     * Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * An ID that is used to identify the archive policy. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
     */
    key: pulumi.Input<string>;
    /**
     * This attribute is used only for project-level archive V2 policies, it is not used for global-level policies.
     */
    projectKey?: pulumi.Input<string>;
    searchCriteria: pulumi.Input<inputs.ArchivePolicySearchCriteria>;
    /**
     * A `true` value means that when this policy is executed, packages will be permanently deleted. `false` means that when the policy is executed packages will be deleted to the Trash Can. Defaults to `false`.
     */
    skipTrashcan?: pulumi.Input<boolean>;
}