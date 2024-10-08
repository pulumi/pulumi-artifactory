// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ReleaseBundleV2PromotionArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ReleaseBundleV2PromotionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource enables you to promote Release Bundle V2 version. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/promote-a-release-bundle-v2-to-a-target-environment).
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
 * import com.pulumi.artifactory.ReleaseBundleV2Promotion;
 * import com.pulumi.artifactory.ReleaseBundleV2PromotionArgs;
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
 *         var my_release_bundle_v2_promotion = new ReleaseBundleV2Promotion("my-release-bundle-v2-promotion", ReleaseBundleV2PromotionArgs.builder()
 *             .name("my-release-bundle-v2-artifacts")
 *             .version("1.0.0")
 *             .keypairName("my-keypair-name")
 *             .environment("DEV")
 *             .includedRepositoryKeys("commons-qa-maven-local")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion")
public class ReleaseBundleV2Promotion extends com.pulumi.resources.CustomResource {
    /**
     * Timestamp when the new version was created (ISO 8601 standard).
     * 
     */
    @Export(name="created", refs={String.class}, tree="[0]")
    private Output<String> created;

    /**
     * @return Timestamp when the new version was created (ISO 8601 standard).
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * Timestamp when the new version was created (in milliseconds).
     * 
     */
    @Export(name="createdMillis", refs={Integer.class}, tree="[0]")
    private Output<Integer> createdMillis;

    /**
     * @return Timestamp when the new version was created (in milliseconds).
     * 
     */
    public Output<Integer> createdMillis() {
        return this.createdMillis;
    }
    /**
     * Target environment
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output<String> environment;

    /**
     * @return Target environment
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }
    /**
     * Defines specific repositories to exclude from the promotion.
     * 
     */
    @Export(name="excludedRepositoryKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> excludedRepositoryKeys;

    /**
     * @return Defines specific repositories to exclude from the promotion.
     * 
     */
    public Output<Optional<List<String>>> excludedRepositoryKeys() {
        return Codegen.optional(this.excludedRepositoryKeys);
    }
    /**
     * Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
     * 
     */
    @Export(name="includedRepositoryKeys", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> includedRepositoryKeys;

    /**
     * @return Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excluded_repository_keys`).
     * 
     */
    public Output<Optional<List<String>>> includedRepositoryKeys() {
        return Codegen.optional(this.includedRepositoryKeys);
    }
    /**
     * Key-pair name to use for signature creation
     * 
     */
    @Export(name="keypairName", refs={String.class}, tree="[0]")
    private Output<String> keypairName;

    /**
     * @return Key-pair name to use for signature creation
     * 
     */
    public Output<String> keypairName() {
        return this.keypairName;
    }
    /**
     * Name of Release Bundle
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of Release Bundle
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Project key the Release Bundle belongs to
     * 
     */
    @Export(name="projectKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key the Release Bundle belongs to
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * Version to promote
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return Version to promote
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReleaseBundleV2Promotion(java.lang.String name) {
        this(name, ReleaseBundleV2PromotionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReleaseBundleV2Promotion(java.lang.String name, ReleaseBundleV2PromotionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReleaseBundleV2Promotion(java.lang.String name, ReleaseBundleV2PromotionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReleaseBundleV2Promotion(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleV2PromotionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion", name, state, makeResourceOptions(options, id), false);
    }

    private static ReleaseBundleV2PromotionArgs makeArgs(ReleaseBundleV2PromotionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReleaseBundleV2PromotionArgs.Empty : args;
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
    public static ReleaseBundleV2Promotion get(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleV2PromotionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReleaseBundleV2Promotion(name, id, state, options);
    }
}
