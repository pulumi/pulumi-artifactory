// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SingleReplicationConfigState extends com.pulumi.resources.ResourceArgs {

    public static final SingleReplicationConfigState Empty = new SingleReplicationConfigState();

    @Import(name="cronExp")
    private @Nullable Output<String> cronExp;

    public Optional<Output<String>> cronExp() {
        return Optional.ofNullable(this.cronExp);
    }

    @Import(name="enableEventReplication")
    private @Nullable Output<Boolean> enableEventReplication;

    public Optional<Output<Boolean>> enableEventReplication() {
        return Optional.ofNullable(this.enableEventReplication);
    }

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Requires password encryption to be turned off `POST /api/system/decrypt`.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="pathPrefix")
    private @Nullable Output<String> pathPrefix;

    public Optional<Output<String>> pathPrefix() {
        return Optional.ofNullable(this.pathPrefix);
    }

    /**
     * Proxy key from Artifactory Proxies setting.
     * 
     */
    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies setting.
     * 
     */
    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    @Import(name="repoKey")
    private @Nullable Output<String> repoKey;

    public Optional<Output<String>> repoKey() {
        return Optional.ofNullable(this.repoKey);
    }

    @Import(name="socketTimeoutMillis")
    private @Nullable Output<Integer> socketTimeoutMillis;

    public Optional<Output<Integer>> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }

    @Import(name="syncDeletes")
    private @Nullable Output<Boolean> syncDeletes;

    public Optional<Output<Boolean>> syncDeletes() {
        return Optional.ofNullable(this.syncDeletes);
    }

    @Import(name="syncProperties")
    private @Nullable Output<Boolean> syncProperties;

    public Optional<Output<Boolean>> syncProperties() {
        return Optional.ofNullable(this.syncProperties);
    }

    @Import(name="syncStatistics")
    private @Nullable Output<Boolean> syncStatistics;

    public Optional<Output<Boolean>> syncStatistics() {
        return Optional.ofNullable(this.syncStatistics);
    }

    @Import(name="url")
    private @Nullable Output<String> url;

    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    @Import(name="username")
    private @Nullable Output<String> username;

    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private SingleReplicationConfigState() {}

    private SingleReplicationConfigState(SingleReplicationConfigState $) {
        this.cronExp = $.cronExp;
        this.enableEventReplication = $.enableEventReplication;
        this.enabled = $.enabled;
        this.password = $.password;
        this.pathPrefix = $.pathPrefix;
        this.proxy = $.proxy;
        this.repoKey = $.repoKey;
        this.socketTimeoutMillis = $.socketTimeoutMillis;
        this.syncDeletes = $.syncDeletes;
        this.syncProperties = $.syncProperties;
        this.syncStatistics = $.syncStatistics;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SingleReplicationConfigState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SingleReplicationConfigState $;

        public Builder() {
            $ = new SingleReplicationConfigState();
        }

        public Builder(SingleReplicationConfigState defaults) {
            $ = new SingleReplicationConfigState(Objects.requireNonNull(defaults));
        }

        public Builder cronExp(@Nullable Output<String> cronExp) {
            $.cronExp = cronExp;
            return this;
        }

        public Builder cronExp(String cronExp) {
            return cronExp(Output.of(cronExp));
        }

        public Builder enableEventReplication(@Nullable Output<Boolean> enableEventReplication) {
            $.enableEventReplication = enableEventReplication;
            return this;
        }

        public Builder enableEventReplication(Boolean enableEventReplication) {
            return enableEventReplication(Output.of(enableEventReplication));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param password Requires password encryption to be turned off `POST /api/system/decrypt`.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Requires password encryption to be turned off `POST /api/system/decrypt`.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder pathPrefix(@Nullable Output<String> pathPrefix) {
            $.pathPrefix = pathPrefix;
            return this;
        }

        public Builder pathPrefix(String pathPrefix) {
            return pathPrefix(Output.of(pathPrefix));
        }

        /**
         * @param proxy Proxy key from Artifactory Proxies setting.
         * 
         * @return builder
         * 
         */
        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        /**
         * @param proxy Proxy key from Artifactory Proxies setting.
         * 
         * @return builder
         * 
         */
        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        public Builder repoKey(@Nullable Output<String> repoKey) {
            $.repoKey = repoKey;
            return this;
        }

        public Builder repoKey(String repoKey) {
            return repoKey(Output.of(repoKey));
        }

        public Builder socketTimeoutMillis(@Nullable Output<Integer> socketTimeoutMillis) {
            $.socketTimeoutMillis = socketTimeoutMillis;
            return this;
        }

        public Builder socketTimeoutMillis(Integer socketTimeoutMillis) {
            return socketTimeoutMillis(Output.of(socketTimeoutMillis));
        }

        public Builder syncDeletes(@Nullable Output<Boolean> syncDeletes) {
            $.syncDeletes = syncDeletes;
            return this;
        }

        public Builder syncDeletes(Boolean syncDeletes) {
            return syncDeletes(Output.of(syncDeletes));
        }

        public Builder syncProperties(@Nullable Output<Boolean> syncProperties) {
            $.syncProperties = syncProperties;
            return this;
        }

        public Builder syncProperties(Boolean syncProperties) {
            return syncProperties(Output.of(syncProperties));
        }

        public Builder syncStatistics(@Nullable Output<Boolean> syncStatistics) {
            $.syncStatistics = syncStatistics;
            return this;
        }

        public Builder syncStatistics(Boolean syncStatistics) {
            return syncStatistics(Output.of(syncStatistics));
        }

        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        public Builder url(String url) {
            return url(Output.of(url));
        }

        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public SingleReplicationConfigState build() {
            return $;
        }
    }

}