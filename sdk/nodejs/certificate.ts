// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory certificate resource. This can be used to create and manage Artifactory certificates which can be used as client authentication against remote repositories.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 * import * as std from "@pulumi/std";
 *
 * // Create a new Artifactory certificate called my-cert
 * const my_cert = new artifactory.Certificate("my-cert", {
 *     alias: "my-cert",
 *     content: std.file({
 *         input: "/path/to/bundle.pem",
 *     }).then(invoke => invoke.result),
 * });
 * // This can then be used by a remote repository
 * const my_remote = new artifactory.RemoteMavenRepository("my-remote", {clientTlsCertificate: my_cert.alias});
 * ```
 *
 * ## Import
 *
 * Certificates can be imported using their alias, e.g.
 *
 * ```sh
 * $ pulumi import artifactory:index/certificate:Certificate my-cert my-cert
 * ```
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/certificate:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Name of certificate.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * Path to the PEM file. Cannot be set with `content` attribute simultaneously.
     */
    public readonly file!: pulumi.Output<string | undefined>;
    /**
     * SHA256 fingerprint of the certificate.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * Name of the certificate authority that issued the certificate.
     */
    public /*out*/ readonly issuedBy!: pulumi.Output<string>;
    /**
     * The time & date when the certificate is valid from.
     */
    public /*out*/ readonly issuedOn!: pulumi.Output<string>;
    /**
     * Name of whom the certificate has been issued to.
     */
    public /*out*/ readonly issuedTo!: pulumi.Output<string>;
    /**
     * The time & date when the certificate expires.
     */
    public /*out*/ readonly validUntil!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["file"] = state ? state.file : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["issuedBy"] = state ? state.issuedBy : undefined;
            resourceInputs["issuedOn"] = state ? state.issuedOn : undefined;
            resourceInputs["issuedTo"] = state ? state.issuedTo : undefined;
            resourceInputs["validUntil"] = state ? state.validUntil : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["content"] = args?.content ? pulumi.secret(args.content) : undefined;
            resourceInputs["file"] = args?.file ? pulumi.secret(args.file) : undefined;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["issuedBy"] = undefined /*out*/;
            resourceInputs["issuedOn"] = undefined /*out*/;
            resourceInputs["issuedTo"] = undefined /*out*/;
            resourceInputs["validUntil"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["content", "file"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * Name of certificate.
     */
    alias?: pulumi.Input<string>;
    /**
     * PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
     */
    content?: pulumi.Input<string>;
    /**
     * Path to the PEM file. Cannot be set with `content` attribute simultaneously.
     */
    file?: pulumi.Input<string>;
    /**
     * SHA256 fingerprint of the certificate.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * Name of the certificate authority that issued the certificate.
     */
    issuedBy?: pulumi.Input<string>;
    /**
     * The time & date when the certificate is valid from.
     */
    issuedOn?: pulumi.Input<string>;
    /**
     * Name of whom the certificate has been issued to.
     */
    issuedTo?: pulumi.Input<string>;
    /**
     * The time & date when the certificate expires.
     */
    validUntil?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Name of certificate.
     */
    alias: pulumi.Input<string>;
    /**
     * PEM-encoded client certificate and private key. Cannot be set with `file` attribute simultaneously.
     */
    content?: pulumi.Input<string>;
    /**
     * Path to the PEM file. Cannot be set with `content` attribute simultaneously.
     */
    file?: pulumi.Input<string>;
}
