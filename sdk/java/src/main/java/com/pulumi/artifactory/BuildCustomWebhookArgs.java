// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.inputs.BuildCustomWebhookCriteriaArgs;
import com.pulumi.artifactory.inputs.BuildCustomWebhookHandlerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BuildCustomWebhookArgs extends com.pulumi.resources.ResourceArgs {

    public static final BuildCustomWebhookArgs Empty = new BuildCustomWebhookArgs();

    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Import(name="criteria")
    private @Nullable Output<BuildCustomWebhookCriteriaArgs> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Optional<Output<BuildCustomWebhookCriteriaArgs>> criteria() {
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
     * Status of webhook. Default to `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
     * 
     */
    @Import(name="eventTypes", required=true)
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }

    /**
     * At least one is required.
     * 
     */
    @Import(name="handlers")
    private @Nullable Output<List<BuildCustomWebhookHandlerArgs>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Optional<Output<List<BuildCustomWebhookHandlerArgs>>> handlers() {
        return Optional.ofNullable(this.handlers);
    }

    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    private BuildCustomWebhookArgs() {}

    private BuildCustomWebhookArgs(BuildCustomWebhookArgs $) {
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
    public static Builder builder(BuildCustomWebhookArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BuildCustomWebhookArgs $;

        public Builder() {
            $ = new BuildCustomWebhookArgs();
        }

        public Builder(BuildCustomWebhookArgs defaults) {
            $ = new BuildCustomWebhookArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(@Nullable Output<BuildCustomWebhookCriteriaArgs> criteria) {
            $.criteria = criteria;
            return this;
        }

        /**
         * @param criteria Specifies where the webhook will be applied on which repositories.
         * 
         * @return builder
         * 
         */
        public Builder criteria(BuildCustomWebhookCriteriaArgs criteria) {
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
         * @param enabled Status of webhook. Default to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Status of webhook. Default to `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(Output<List<String>> eventTypes) {
            $.eventTypes = eventTypes;
            return this;
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(List<String> eventTypes) {
            return eventTypes(Output.of(eventTypes));
        }

        /**
         * @param eventTypes List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `uploaded`, `deleted`, `promoted`.
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
        public Builder handlers(@Nullable Output<List<BuildCustomWebhookHandlerArgs>> handlers) {
            $.handlers = handlers;
            return this;
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(List<BuildCustomWebhookHandlerArgs> handlers) {
            return handlers(Output.of(handlers));
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(BuildCustomWebhookHandlerArgs... handlers) {
            return handlers(List.of(handlers));
        }

        /**
         * @param key The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
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

        public BuildCustomWebhookArgs build() {
            if ($.eventTypes == null) {
                throw new MissingRequiredPropertyException("BuildCustomWebhookArgs", "eventTypes");
            }
            if ($.key == null) {
                throw new MissingRequiredPropertyException("BuildCustomWebhookArgs", "key");
            }
            return $;
        }
    }

}
