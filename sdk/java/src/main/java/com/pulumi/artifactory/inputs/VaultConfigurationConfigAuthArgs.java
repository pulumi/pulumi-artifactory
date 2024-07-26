// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VaultConfigurationConfigAuthArgs extends com.pulumi.resources.ResourceArgs {

    public static final VaultConfigurationConfigAuthArgs Empty = new VaultConfigurationConfigAuthArgs();

    /**
     * Client certificate (in PEM format) for `Certificate` type.
     * 
     */
    @Import(name="certificate")
    private @Nullable Output<String> certificate;

    /**
     * @return Client certificate (in PEM format) for `Certificate` type.
     * 
     */
    public Optional<Output<String>> certificate() {
        return Optional.ofNullable(this.certificate);
    }

    /**
     * Private key (in PEM format) for `Certificate` type.
     * 
     */
    @Import(name="certificateKey")
    private @Nullable Output<String> certificateKey;

    /**
     * @return Private key (in PEM format) for `Certificate` type.
     * 
     */
    public Optional<Output<String>> certificateKey() {
        return Optional.ofNullable(this.certificateKey);
    }

    /**
     * Role ID for `AppRole` type
     * 
     */
    @Import(name="roleId")
    private @Nullable Output<String> roleId;

    /**
     * @return Role ID for `AppRole` type
     * 
     */
    public Optional<Output<String>> roleId() {
        return Optional.ofNullable(this.roleId);
    }

    /**
     * Secret ID for `AppRole` type
     * 
     */
    @Import(name="secretId")
    private @Nullable Output<String> secretId;

    /**
     * @return Secret ID for `AppRole` type
     * 
     */
    public Optional<Output<String>> secretId() {
        return Optional.ofNullable(this.secretId);
    }

    @Import(name="type", required=true)
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }

    private VaultConfigurationConfigAuthArgs() {}

    private VaultConfigurationConfigAuthArgs(VaultConfigurationConfigAuthArgs $) {
        this.certificate = $.certificate;
        this.certificateKey = $.certificateKey;
        this.roleId = $.roleId;
        this.secretId = $.secretId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VaultConfigurationConfigAuthArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VaultConfigurationConfigAuthArgs $;

        public Builder() {
            $ = new VaultConfigurationConfigAuthArgs();
        }

        public Builder(VaultConfigurationConfigAuthArgs defaults) {
            $ = new VaultConfigurationConfigAuthArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param certificate Client certificate (in PEM format) for `Certificate` type.
         * 
         * @return builder
         * 
         */
        public Builder certificate(@Nullable Output<String> certificate) {
            $.certificate = certificate;
            return this;
        }

        /**
         * @param certificate Client certificate (in PEM format) for `Certificate` type.
         * 
         * @return builder
         * 
         */
        public Builder certificate(String certificate) {
            return certificate(Output.of(certificate));
        }

        /**
         * @param certificateKey Private key (in PEM format) for `Certificate` type.
         * 
         * @return builder
         * 
         */
        public Builder certificateKey(@Nullable Output<String> certificateKey) {
            $.certificateKey = certificateKey;
            return this;
        }

        /**
         * @param certificateKey Private key (in PEM format) for `Certificate` type.
         * 
         * @return builder
         * 
         */
        public Builder certificateKey(String certificateKey) {
            return certificateKey(Output.of(certificateKey));
        }

        /**
         * @param roleId Role ID for `AppRole` type
         * 
         * @return builder
         * 
         */
        public Builder roleId(@Nullable Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId Role ID for `AppRole` type
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        /**
         * @param secretId Secret ID for `AppRole` type
         * 
         * @return builder
         * 
         */
        public Builder secretId(@Nullable Output<String> secretId) {
            $.secretId = secretId;
            return this;
        }

        /**
         * @param secretId Secret ID for `AppRole` type
         * 
         * @return builder
         * 
         */
        public Builder secretId(String secretId) {
            return secretId(Output.of(secretId));
        }

        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public VaultConfigurationConfigAuthArgs build() {
            if ($.type == null) {
                throw new MissingRequiredPropertyException("VaultConfigurationConfigAuthArgs", "type");
            }
            return $;
        }
    }

}