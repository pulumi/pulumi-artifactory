// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ReleaseBundleV2CleanupPolicyArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ReleaseBundleV2CleanupPolicyState;
import com.pulumi.artifactory.outputs.ReleaseBundleV2CleanupPolicySearchCriteria;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory Archive Policy resource. This resource enable system administrators to configure and maintain JFrog cleanup policies for Release Bundles V2. See [Cleanup Policies](https://jfrog.com/help/r/jfrog-rest-apis/cleanup-policies-release-bundles-v2-apis) for more details.
 * 
 * ~&gt;Release Bundles V2 Cleanup Policies APIs are supported on Artifactory version 7.104.2 and later.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import artifactory:index/releaseBundleV2CleanupPolicy:ReleaseBundleV2CleanupPolicy my-cleanup-policy my-policy
 * ```
 * 
 */
@ResourceType(type="artifactory:index/releaseBundleV2CleanupPolicy:ReleaseBundleV2CleanupPolicy")
public class ReleaseBundleV2CleanupPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The cron expression determines when the policy is run. This parameter is not mandatory, however if left empty the policy will not run automatically and can only be triggered manually.
     * 
     */
    @Export(name="cronExpression", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cronExpression;

    /**
     * @return The cron expression determines when the policy is run. This parameter is not mandatory, however if left empty the policy will not run automatically and can only be triggered manually.
     * 
     */
    public Output<Optional<String>> cronExpression() {
        return Codegen.optional(this.cronExpression);
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The maximum duration (in minutes) for policy execution, after which the policy will stop running even if not completed. While setting a maximum run duration for a policy is useful for adhering to a strict archive V2 schedule, it can cause the policy to stop before completion.
     * 
     */
    @Export(name="durationInMinutes", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> durationInMinutes;

    /**
     * @return The maximum duration (in minutes) for policy execution, after which the policy will stop running even if not completed. While setting a maximum run duration for a policy is useful for adhering to a strict archive V2 schedule, it can cause the policy to stop before completion.
     * 
     */
    public Output<Optional<Integer>> durationInMinutes() {
        return Codegen.optional(this.durationInMinutes);
    }
    /**
     * Enables or disabled the release bundles v2 cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Enables or disabled the release bundles v2 cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Needs to be set to releaseBundle.
     * 
     */
    @Export(name="itemType", refs={String.class}, tree="[0]")
    private Output<String> itemType;

    /**
     * @return Needs to be set to releaseBundle.
     * 
     */
    public Output<String> itemType() {
        return this.itemType;
    }
    /**
     * An ID that is used to identify the cleanup policy. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return An ID that is used to identify the cleanup policy. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    @Export(name="searchCriteria", refs={ReleaseBundleV2CleanupPolicySearchCriteria.class}, tree="[0]")
    private Output<ReleaseBundleV2CleanupPolicySearchCriteria> searchCriteria;

    public Output<ReleaseBundleV2CleanupPolicySearchCriteria> searchCriteria() {
        return this.searchCriteria;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReleaseBundleV2CleanupPolicy(java.lang.String name) {
        this(name, ReleaseBundleV2CleanupPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReleaseBundleV2CleanupPolicy(java.lang.String name, ReleaseBundleV2CleanupPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReleaseBundleV2CleanupPolicy(java.lang.String name, ReleaseBundleV2CleanupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleV2CleanupPolicy:ReleaseBundleV2CleanupPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReleaseBundleV2CleanupPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleV2CleanupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleV2CleanupPolicy:ReleaseBundleV2CleanupPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static ReleaseBundleV2CleanupPolicyArgs makeArgs(ReleaseBundleV2CleanupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReleaseBundleV2CleanupPolicyArgs.Empty : args;
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
    public static ReleaseBundleV2CleanupPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleV2CleanupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReleaseBundleV2CleanupPolicy(name, id, state, options);
    }
}
