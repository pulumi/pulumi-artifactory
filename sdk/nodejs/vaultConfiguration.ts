// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource enables you to configure an external vault connector to use as a centralized secret management tool for the keys used to sign packages. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/vault).
 * This feature is supported with Enterprise X and Enterprise+ licenses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 * import * as std from "@pulumi/std";
 *
 * const my_vault_config_app_role = new artifactory.VaultConfiguration("my-vault-config-app-role", {
 *     name: "my-vault-config-app-role",
 *     config: {
 *         url: "http://127.0.0.1:8200",
 *         auth: {
 *             type: "AppRole",
 *             roleId: "1b62ff05...",
 *             secretId: "acbd6657...",
 *         },
 *         mounts: [{
 *             path: "secret",
 *             type: "KV2",
 *         }],
 *     },
 * });
 * const my_vault_config_cert = new artifactory.VaultConfiguration("my-vault-config-cert", {
 *     name: "my-vault-config-cert",
 *     config: {
 *         url: "http://127.0.0.1:8200",
 *         auth: {
 *             type: "Certificate",
 *             certificate: std.file({
 *                 input: "samples/public.pem",
 *             }).then(invoke => invoke.result),
 *             certificateKey: std.file({
 *                 input: "samples/private.pem",
 *             }).then(invoke => invoke.result),
 *         },
 *         mounts: [{
 *             path: "secret",
 *             type: "KV2",
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import artifactory:index/vaultConfiguration:VaultConfiguration my-vault-config my-vault-config
 * ```
 */
export class VaultConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing VaultConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VaultConfigurationState, opts?: pulumi.CustomResourceOptions): VaultConfiguration {
        return new VaultConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/vaultConfiguration:VaultConfiguration';

    /**
     * Returns true if the given object is an instance of VaultConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VaultConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VaultConfiguration.__pulumiType;
    }

    public readonly config!: pulumi.Output<outputs.VaultConfigurationConfig>;
    /**
     * Name of the Vault configuration
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a VaultConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VaultConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VaultConfigurationArgs | VaultConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VaultConfigurationState | undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as VaultConfigurationArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VaultConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VaultConfiguration resources.
 */
export interface VaultConfigurationState {
    config?: pulumi.Input<inputs.VaultConfigurationConfig>;
    /**
     * Name of the Vault configuration
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VaultConfiguration resource.
 */
export interface VaultConfigurationArgs {
    config: pulumi.Input<inputs.VaultConfigurationConfig>;
    /**
     * Name of the Vault configuration
     */
    name?: pulumi.Input<string>;
}