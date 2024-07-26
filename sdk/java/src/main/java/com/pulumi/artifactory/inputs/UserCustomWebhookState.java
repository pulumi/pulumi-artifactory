// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.UserCustomWebhookHandlerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserCustomWebhookState extends com.pulumi.resources.ResourceArgs {

    public static final UserCustomWebhookState Empty = new UserCustomWebhookState();

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
     * List of event triggers for the Webhook. Allow values: `locked`
     * 
     */
    @Import(name="eventTypes")
    private @Nullable Output<List<String>> eventTypes;

    /**
     * @return List of event triggers for the Webhook. Allow values: `locked`
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
    private @Nullable Output<List<UserCustomWebhookHandlerArgs>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Optional<Output<List<UserCustomWebhookHandlerArgs>>> handlers() {
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

    private UserCustomWebhookState() {}

    private UserCustomWebhookState(UserCustomWebhookState $) {
        this.description = $.description;
        this.enabled = $.enabled;
        this.eventTypes = $.eventTypes;
        this.handlers = $.handlers;
        this.key = $.key;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserCustomWebhookState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserCustomWebhookState $;

        public Builder() {
            $ = new UserCustomWebhookState();
        }

        public Builder(UserCustomWebhookState defaults) {
            $ = new UserCustomWebhookState(Objects.requireNonNull(defaults));
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
         * @param eventTypes List of event triggers for the Webhook. Allow values: `locked`
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(@Nullable Output<List<String>> eventTypes) {
            $.eventTypes = eventTypes;
            return this;
        }

        /**
         * @param eventTypes List of event triggers for the Webhook. Allow values: `locked`
         * 
         * @return builder
         * 
         */
        public Builder eventTypes(List<String> eventTypes) {
            return eventTypes(Output.of(eventTypes));
        }

        /**
         * @param eventTypes List of event triggers for the Webhook. Allow values: `locked`
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
        public Builder handlers(@Nullable Output<List<UserCustomWebhookHandlerArgs>> handlers) {
            $.handlers = handlers;
            return this;
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(List<UserCustomWebhookHandlerArgs> handlers) {
            return handlers(Output.of(handlers));
        }

        /**
         * @param handlers At least one is required.
         * 
         * @return builder
         * 
         */
        public Builder handlers(UserCustomWebhookHandlerArgs... handlers) {
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

        public UserCustomWebhookState build() {
            return $;
        }
    }

}