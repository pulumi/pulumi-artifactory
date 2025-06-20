// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 *
 * ## Example Usage
 *
 * .
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my_generic_local = new artifactory.LocalGenericRepository("my-generic-local", {key: "my-generic-local"});
 * const artifact_webhook = new artifactory.ArtifactWebhook("artifact-webhook", {
 *     key: "artifact-webhook",
 *     eventTypes: [
 *         "deployed",
 *         "deleted",
 *         "moved",
 *         "copied",
 *     ],
 *     criteria: {
 *         anyLocal: true,
 *         anyRemote: false,
 *         anyFederated: false,
 *         repoKeys: [my_generic_local.key],
 *         includePatterns: ["foo/**"],
 *         excludePatterns: ["bar/**"],
 *     },
 *     handlers: [{
 *         url: "http://tempurl.org/webhook",
 *         secret: "some-secret",
 *         proxy: "proxy-key",
 *         customHttpHeaders: {
 *             "header-1": "value-1",
 *             "header-2": "value-2",
 *         },
 *     }],
 * }, {
 *     dependsOn: [my_generic_local],
 * });
 * ```
 */
export class ArtifactWebhook extends pulumi.CustomResource {
    /**
     * Get an existing ArtifactWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ArtifactWebhookState, opts?: pulumi.CustomResourceOptions): ArtifactWebhook {
        return new ArtifactWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/artifactWebhook:ArtifactWebhook';

    /**
     * Returns true if the given object is an instance of ArtifactWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ArtifactWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ArtifactWebhook.__pulumiType;
    }

    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    public readonly criteria!: pulumi.Output<outputs.ArtifactWebhookCriteria | undefined>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Status of webhook. Default to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
     */
    public readonly eventTypes!: pulumi.Output<string[]>;
    /**
     * At least one is required.
     */
    public readonly handlers!: pulumi.Output<outputs.ArtifactWebhookHandler[] | undefined>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    public readonly key!: pulumi.Output<string>;

    /**
     * Create a ArtifactWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArtifactWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ArtifactWebhookArgs | ArtifactWebhookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ArtifactWebhookState | undefined;
            resourceInputs["criteria"] = state ? state.criteria : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["eventTypes"] = state ? state.eventTypes : undefined;
            resourceInputs["handlers"] = state ? state.handlers : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
        } else {
            const args = argsOrState as ArtifactWebhookArgs | undefined;
            if ((!args || args.eventTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventTypes'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["criteria"] = args ? args.criteria : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventTypes"] = args ? args.eventTypes : undefined;
            resourceInputs["handlers"] = args ? args.handlers : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ArtifactWebhook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ArtifactWebhook resources.
 */
export interface ArtifactWebhookState {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria?: pulumi.Input<inputs.ArtifactWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
     */
    eventTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.ArtifactWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ArtifactWebhook resource.
 */
export interface ArtifactWebhookArgs {
    /**
     * Specifies where the webhook will be applied on which repositories.
     */
    criteria?: pulumi.Input<inputs.ArtifactWebhookCriteria>;
    /**
     * Webhook description. Max length 1000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Status of webhook. Default to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `deployed`, `deleted`, `moved`, `copied`, `cached`.
     */
    eventTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * At least one is required.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.ArtifactWebhookHandler>[]>;
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     */
    key: pulumi.Input<string>;
}
