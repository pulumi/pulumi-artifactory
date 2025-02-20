// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ReleaseBundleWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ReleaseBundleWebhookState;
import com.pulumi.artifactory.outputs.ReleaseBundleWebhookCriteria;
import com.pulumi.artifactory.outputs.ReleaseBundleWebhookHandler;
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
 * !&gt;This resource is being deprecated and replaced by `artifactory.DestinationWebhook` resource.
 * 
 * ## Example Usage
 * 
 * .
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.ReleaseBundleWebhook;
 * import com.pulumi.artifactory.ReleaseBundleWebhookArgs;
 * import com.pulumi.artifactory.inputs.ReleaseBundleWebhookCriteriaArgs;
 * import com.pulumi.artifactory.inputs.ReleaseBundleWebhookHandlerArgs;
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
 *         var release_bundle_webhook = new ReleaseBundleWebhook("release-bundle-webhook", ReleaseBundleWebhookArgs.builder()
 *             .key("release-bundle-webhook")
 *             .eventTypes(            
 *                 "created",
 *                 "signed",
 *                 "deleted")
 *             .criteria(ReleaseBundleWebhookCriteriaArgs.builder()
 *                 .anyReleaseBundle(false)
 *                 .registeredReleaseBundleNames("bundle-name")
 *                 .includePatterns("foo/**")
 *                 .excludePatterns("bar/**")
 *                 .build())
 *             .handlers(ReleaseBundleWebhookHandlerArgs.builder()
 *                 .url("http://tempurl.org/webhook")
 *                 .secret("some-secret")
 *                 .proxy("proxy-key")
 *                 .customHttpHeaders(Map.ofEntries(
 *                     Map.entry("header-1", "value-1"),
 *                     Map.entry("header-2", "value-2")
 *                 ))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="artifactory:index/releaseBundleWebhook:ReleaseBundleWebhook")
public class ReleaseBundleWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", refs={ReleaseBundleWebhookCriteria.class}, tree="[0]")
    private Output</* @Nullable */ ReleaseBundleWebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<Optional<ReleaseBundleWebhookCriteria>> criteria() {
        return Codegen.optional(this.criteria);
    }
    /**
     * Webhook description. Max length 1000 characters.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Webhook description. Max length 1000 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Status of webhook. Default to `true`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
     * 
     */
    @Export(name="eventTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", refs={List.class,ReleaseBundleWebhookHandler.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ReleaseBundleWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<Optional<List<ReleaseBundleWebhookHandler>>> handlers() {
        return Codegen.optional(this.handlers);
    }
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
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
    public ReleaseBundleWebhook(java.lang.String name) {
        this(name, ReleaseBundleWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReleaseBundleWebhook(java.lang.String name, ReleaseBundleWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReleaseBundleWebhook(java.lang.String name, ReleaseBundleWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleWebhook:ReleaseBundleWebhook", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ReleaseBundleWebhook(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/releaseBundleWebhook:ReleaseBundleWebhook", name, state, makeResourceOptions(options, id), false);
    }

    private static ReleaseBundleWebhookArgs makeArgs(ReleaseBundleWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReleaseBundleWebhookArgs.Empty : args;
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
    public static ReleaseBundleWebhook get(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseBundleWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReleaseBundleWebhook(name, id, state, options);
    }
}
