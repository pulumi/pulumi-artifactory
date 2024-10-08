// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.ReleaseBundleV2PromotionWebhookCriteriaArgs;
import com.pulumi.artifactory.inputs.ReleaseBundleV2PromotionWebhookHandlerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReleaseBundleV2PromotionWebhookState extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseBundleV2PromotionWebhookState Empty = new ReleaseBundleV2PromotionWebhookState();

    /**
     * Specifies where the webhook will be applied on which enviroments.
     * 
     */
    @Import(name="criteria")
    private @Nullable Output<ReleaseBundleV2PromotionWebhookCriteriaArgs> criteria;

    /**
     * @return Specifies where the webhook will be applied on which enviroments.
     * 
     */
    public Optional<Output<ReleaseBundleV2PromotionWebhookCriteriaArgs>> criteria() {
        return Optional.ofNullable(this.criteria);
    }

    /**
     * Webhook description. Max length 1000 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Webhook description. Max length 1000 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Status of webhook. Default to `true`
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
     * 
     */
    @Import(name="eventTypes")
    private @Nullable Output<List<String>> eventTypes;

    /**
     * @return List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
     * 
     */
    public Optional<Output<List<String>>> eventTypes() {
        return Optional.ofNullable(this.eventTypes);
    }

    /**
     * At least one is required.
     * 
     */
    @Import(name="handlers")
    private @Nullable Output<List<ReleaseBundleV2PromotionWebhookHandlerArgs>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Optional<Output<List<ReleaseBundleV2PromotionWebhookHandlerArgs>>> handlers() {
        return Optional.ofNullable(this.handlers);
    }

    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    private ReleaseBundleV2PromotionWebhookState() {}

    private ReleaseBundleV2PromotionWebhookState(ReleaseBundleV2PromotionWebhookState $) {
        this.criteria = $.criteria;
        this.description = $.description;
        this.enabled = $.enabled;
        this.eventTypes = $.eventTypes;
        this.handlers = $.handlers;
        this.key = $.key;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReleaseBundleV2PromotionWebhookState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseBundleV2PromotionWebhookState $;

        public Builder() {
            $ = new ReleaseBundleV2PromotionWebhookState();
        }

        public Builder(ReleaseBundleV2PromotionWebhookState defaults) {
            $ = new ReleaseBundleV2PromotionWebhookState(Objects.requireNonNull(defaults));
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which enviroments.
         * 
         * @return builder
         * 
         */
        public Builder criteria(@Nullable Output<ReleaseBundleV2PromotionWebhookCriteriaArgs> criteria) {
            $.criteria = criteria;
            return this;
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which enviroments.
         * 
         * @return builder
         * 
         */
        public Builder criteria(ReleaseBundleV2PromotionWebhookCriteriaArgs criteria) {
            return criteria(Output.of(criteria));
        }

        /**
         * @param description Webhook description. Max length 1000 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Webhook description. Max length 1000 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param enabled Status of webhook. Default to `true`
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Status of webhook. Default to `true`
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param eventTypes List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(@Nullable Output<List<String>> eventTypes) {
            $.eventTypes = eventTypes;
            return this;
        }

        /**
         * @param eventTypes List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(List<String> eventTypes) {
            return eventTypes(Output.of(eventTypes));
        }

        /**
         * @param eventTypes List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(String... eventTypes) {
            return eventTypes(List.of(eventTypes));
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(@Nullable Output<List<ReleaseBundleV2PromotionWebhookHandlerArgs>> handlers) {
            $.handlers = handlers;
            return this;
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(List<ReleaseBundleV2PromotionWebhookHandlerArgs> handlers) {
            return handlers(Output.of(handlers));
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(ReleaseBundleV2PromotionWebhookHandlerArgs... handlers) {
            return handlers(List.of(handlers));
        }

        /**
         * @param key The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public ReleaseBundleV2PromotionWebhookState build() {
            return $;
        }
    }

}
