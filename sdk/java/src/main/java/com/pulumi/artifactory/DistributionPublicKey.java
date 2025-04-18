// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.DistributionPublicKeyArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.DistributionPublicKeyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory Distribution Public Key resource. This can be used to create and manage Artifactory Distribution Public Keys.
 * 
 * See [API description](https://jfrog.com/help/r/jfrog-rest-apis/set-distributionpublic-gpg-key) in the Artifactory documentation for more details. Also the [UI documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/managing-webstart-and-jar-signing) has further details on where to find these keys in Artifactory.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.DistributionPublicKey;
 * import com.pulumi.artifactory.DistributionPublicKeyArgs;
 * import com.pulumi.std.StdFunctions;
 * import com.pulumi.std.inputs.FileArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var my_key = new DistributionPublicKey("my-key", DistributionPublicKeyArgs.builder()
 *             .alias("my-key")
 *             .publicKey(StdFunctions.file(FileArgs.builder()
 *                 .input("samples/rsa.pub")
 *                 .build()).result())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Distribution Public Key can be imported using the key id, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/distributionPublicKey:DistributionPublicKey my-key keyid
 * ```
 * 
 */
@ResourceType(type="artifactory:index/distributionPublicKey:DistributionPublicKey")
public class DistributionPublicKey extends com.pulumi.resources.CustomResource {
    /**
     * Will be used as an identifier when uploading/retrieving the public key via REST API.
     * 
     */
    @Export(name="alias", refs={String.class}, tree="[0]")
    private Output<String> alias;

    /**
     * @return Will be used as an identifier when uploading/retrieving the public key via REST API.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }
    /**
     * Returns the computed key fingerprint
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return Returns the computed key fingerprint
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * Returns the name and eMail address of issuer
     * 
     */
    @Export(name="issuedBy", refs={String.class}, tree="[0]")
    private Output<String> issuedBy;

    /**
     * @return Returns the name and eMail address of issuer
     * 
     */
    public Output<String> issuedBy() {
        return this.issuedBy;
    }
    /**
     * Returns the date/time when this GPG key was created
     * 
     */
    @Export(name="issuedOn", refs={String.class}, tree="[0]")
    private Output<String> issuedOn;

    /**
     * @return Returns the date/time when this GPG key was created
     * 
     */
    public Output<String> issuedOn() {
        return this.issuedOn;
    }
    /**
     * Returns the key id by which this key is referenced in Artifactory
     * 
     */
    @Export(name="keyId", refs={String.class}, tree="[0]")
    private Output<String> keyId;

    /**
     * @return Returns the key id by which this key is referenced in Artifactory
     * 
     */
    public Output<String> keyId() {
        return this.keyId;
    }
    /**
     * The Public key to add as a trusted distribution GPG key.
     * 
     * The following additional attributes are exported:
     * 
     */
    @Export(name="publicKey", refs={String.class}, tree="[0]")
    private Output<String> publicKey;

    /**
     * @return The Public key to add as a trusted distribution GPG key.
     * 
     * The following additional attributes are exported:
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }
    /**
     * Returns the date/time when this GPG key expires.
     * 
     */
    @Export(name="validUntil", refs={String.class}, tree="[0]")
    private Output<String> validUntil;

    /**
     * @return Returns the date/time when this GPG key expires.
     * 
     */
    public Output<String> validUntil() {
        return this.validUntil;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DistributionPublicKey(java.lang.String name) {
        this(name, DistributionPublicKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DistributionPublicKey(java.lang.String name, DistributionPublicKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DistributionPublicKey(java.lang.String name, DistributionPublicKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/distributionPublicKey:DistributionPublicKey", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DistributionPublicKey(java.lang.String name, Output<java.lang.String> id, @Nullable DistributionPublicKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/distributionPublicKey:DistributionPublicKey", name, state, makeResourceOptions(options, id), false);
    }

    private static DistributionPublicKeyArgs makeArgs(DistributionPublicKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DistributionPublicKeyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DistributionPublicKey get(java.lang.String name, Output<java.lang.String> id, @Nullable DistributionPublicKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DistributionPublicKey(name, id, state, options);
    }
}
