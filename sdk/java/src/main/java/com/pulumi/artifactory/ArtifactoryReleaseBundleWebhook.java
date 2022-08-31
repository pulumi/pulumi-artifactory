// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ArtifactoryReleaseBundleWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ArtifactoryReleaseBundleWebhookState;
import com.pulumi.artifactory.outputs.ArtifactoryReleaseBundleWebhookCriteria;
import com.pulumi.artifactory.outputs.ArtifactoryReleaseBundleWebhookHandler;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 * 
 * ## Example Usage
 * 
 * .
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ArtifactoryReleaseBundleWebhook;
 * import com.pulumi.artifactory.ArtifactoryReleaseBundleWebhookArgs;
 * import com.pulumi.artifactory.inputs.ArtifactoryReleaseBundleWebhookCriteriaArgs;
 * import com.pulumi.artifactory.inputs.ArtifactoryReleaseBundleWebhookHandlerArgs;
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
 *         var artifactory_release_bundle_webhook = new ArtifactoryReleaseBundleWebhook(&#34;artifactory-release-bundle-webhook&#34;, ArtifactoryReleaseBundleWebhookArgs.builder()        
 *             .criteria(ArtifactoryReleaseBundleWebhookCriteriaArgs.builder()
 *                 .anyReleaseBundle(false)
 *                 .excludePatterns(&#34;bar/**&#34;)
 *                 .includePatterns(&#34;foo/**&#34;)
 *                 .registeredReleaseBundleNames(&#34;bundle-name&#34;)
 *                 .build())
 *             .eventTypes(            
 *                 &#34;received&#34;,
 *                 &#34;delete_started&#34;,
 *                 &#34;delete_completed&#34;,
 *                 &#34;delete_failed&#34;)
 *             .handlers(ArtifactoryReleaseBundleWebhookHandlerArgs.builder()
 *                 .customHttpHeaders(Map.ofEntries(
 *                     Map.entry(&#34;header-1&#34;, &#34;value-1&#34;),
 *                     Map.entry(&#34;header-2&#34;, &#34;value-2&#34;)
 *                 ))
 *                 .proxy(&#34;proxy-key&#34;)
 *                 .secret(&#34;some-secret&#34;)
 *                 .url(&#34;http://tempurl.org/webhook&#34;)
 *                 .build())
 *             .key(&#34;artifactory-release-bundle-webhook&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="artifactory:index/artifactoryReleaseBundleWebhook:ArtifactoryReleaseBundleWebhook")
public class ArtifactoryReleaseBundleWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", type=ArtifactoryReleaseBundleWebhookCriteria.class, parameters={})
    private Output<ArtifactoryReleaseBundleWebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<ArtifactoryReleaseBundleWebhookCriteria> criteria() {
        return this.criteria;
    }
    /**
     * Webhook description. Max length 1000 characters.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Webhook description. Max length 1000 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Status of webhook. Default to &#39;true&#39;
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Status of webhook. Default to &#39;true&#39;
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: &#34;received&#34;, &#34;delete_started&#34;, &#34;delete_completed&#34;, &#34;delete_failed&#34;
     * 
     */
    @Export(name="eventTypes", type=List.class, parameters={String.class})
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: &#34;received&#34;, &#34;delete_started&#34;, &#34;delete_completed&#34;, &#34;delete_failed&#34;
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", type=List.class, parameters={ArtifactoryReleaseBundleWebhookHandler.class})
    private Output<List<ArtifactoryReleaseBundleWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<ArtifactoryReleaseBundleWebhookHandler>> handlers() {
        return this.handlers;
    }
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ArtifactoryReleaseBundleWebhook(String name) {
        this(name, ArtifactoryReleaseBundleWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ArtifactoryReleaseBundleWebhook(String name, ArtifactoryReleaseBundleWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ArtifactoryReleaseBundleWebhook(String name, ArtifactoryReleaseBundleWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactoryReleaseBundleWebhook:ArtifactoryReleaseBundleWebhook", name, args == null ? ArtifactoryReleaseBundleWebhookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ArtifactoryReleaseBundleWebhook(String name, Output<String> id, @Nullable ArtifactoryReleaseBundleWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifactoryReleaseBundleWebhook:ArtifactoryReleaseBundleWebhook", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static ArtifactoryReleaseBundleWebhook get(String name, Output<String> id, @Nullable ArtifactoryReleaseBundleWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ArtifactoryReleaseBundleWebhook(name, id, state, options);
    }
}