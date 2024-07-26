// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.VaultConfigurationConfigAuthArgs;
import com.pulumi.artifactory.inputs.VaultConfigurationConfigMountArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class VaultConfigurationConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final VaultConfigurationConfigArgs Empty = new VaultConfigurationConfigArgs();

    @Import(name="auth", required=true)
    private Output<VaultConfigurationConfigAuthArgs> auth;

    public Output<VaultConfigurationConfigAuthArgs> auth() {
        return this.auth;
    }

    @Import(name="mounts", required=true)
    private Output<List<VaultConfigurationConfigMountArgs>> mounts;

    public Output<List<VaultConfigurationConfigMountArgs>> mounts() {
        return this.mounts;
    }

    /**
     * The base URL of the Vault server.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The base URL of the Vault server.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private VaultConfigurationConfigArgs() {}

    private VaultConfigurationConfigArgs(VaultConfigurationConfigArgs $) {
        this.auth = $.auth;
        this.mounts = $.mounts;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VaultConfigurationConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VaultConfigurationConfigArgs $;

        public Builder() {
            $ = new VaultConfigurationConfigArgs();
        }

        public Builder(VaultConfigurationConfigArgs defaults) {
            $ = new VaultConfigurationConfigArgs(Objects.requireNonNull(defaults));
        }

        public Builder auth(Output<VaultConfigurationConfigAuthArgs> auth) {
            $.auth = auth;
            return this;
        }

        public Builder auth(VaultConfigurationConfigAuthArgs auth) {
            return auth(Output.of(auth));
        }

        public Builder mounts(Output<List<VaultConfigurationConfigMountArgs>> mounts) {
            $.mounts = mounts;
            return this;
        }

        public Builder mounts(List<VaultConfigurationConfigMountArgs> mounts) {
            return mounts(Output.of(mounts));
        }

        public Builder mounts(VaultConfigurationConfigMountArgs... mounts) {
            return mounts(List.of(mounts));
        }

        /**
         * @param url The base URL of the Vault server.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The base URL of the Vault server.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public VaultConfigurationConfigArgs build() {
            if ($.auth == null) {
                throw new MissingRequiredPropertyException("VaultConfigurationConfigArgs", "auth");
            }
            if ($.mounts == null) {
                throw new MissingRequiredPropertyException("VaultConfigurationConfigArgs", "mounts");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("VaultConfigurationConfigArgs", "url");
            }
            return $;
        }
    }

}