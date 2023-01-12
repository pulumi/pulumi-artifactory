// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.PropertySetPropertyArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PropertySetState extends com.pulumi.resources.ResourceArgs {

    public static final PropertySetState Empty = new PropertySetState();

    /**
     * Predefined property name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Predefined property name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A list of properties that will be part of the property set.
     * 
     */
    @Import(name="properties")
    private @Nullable Output<List<PropertySetPropertyArgs>> properties;

    /**
     * @return A list of properties that will be part of the property set.
     * 
     */
    public Optional<Output<List<PropertySetPropertyArgs>>> properties() {
        return Optional.ofNullable(this.properties);
    }

    /**
     * Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
     * 
     */
    @Import(name="visible")
    private @Nullable Output<Boolean> visible;

    /**
     * @return Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> visible() {
        return Optional.ofNullable(this.visible);
    }

    private PropertySetState() {}

    private PropertySetState(PropertySetState $) {
        this.name = $.name;
        this.properties = $.properties;
        this.visible = $.visible;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PropertySetState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PropertySetState $;

        public Builder() {
            $ = new PropertySetState();
        }

        public Builder(PropertySetState defaults) {
            $ = new PropertySetState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Predefined property name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Predefined property name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param properties A list of properties that will be part of the property set.
         * 
         * @return builder
         * 
         */
        public Builder properties(@Nullable Output<List<PropertySetPropertyArgs>> properties) {
            $.properties = properties;
            return this;
        }

        /**
         * @param properties A list of properties that will be part of the property set.
         * 
         * @return builder
         * 
         */
        public Builder properties(List<PropertySetPropertyArgs> properties) {
            return properties(Output.of(properties));
        }

        /**
         * @param properties A list of properties that will be part of the property set.
         * 
         * @return builder
         * 
         */
        public Builder properties(PropertySetPropertyArgs... properties) {
            return properties(List.of(properties));
        }

        /**
         * @param visible Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder visible(@Nullable Output<Boolean> visible) {
            $.visible = visible;
            return this;
        }

        /**
         * @param visible Defines if the list visible and assignable to the repository or artifact. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder visible(Boolean visible) {
            return visible(Output.of(visible));
        }

        public PropertySetState build() {
            return $;
        }
    }

}