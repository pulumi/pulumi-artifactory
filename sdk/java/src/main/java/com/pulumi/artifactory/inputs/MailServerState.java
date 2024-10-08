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


public final class MailServerState extends com.pulumi.resources.ResourceArgs {

    public static final MailServerState Empty = new MailServerState();

    /**
     * The Artifactory URL to to link to in all outgoing messages.
     * 
     */
    @Import(name="artifactoryUrl")
    private @Nullable Output<String> artifactoryUrl;

    /**
     * @return The Artifactory URL to to link to in all outgoing messages.
     * 
     */
    public Optional<Output<String>> artifactoryUrl() {
        return Optional.ofNullable(this.artifactoryUrl);
    }

    /**
     * When set, mail notifications are enabled.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return When set, mail notifications are enabled.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The &#39;from&#39; address header to use in all outgoing messages.
     * 
     */
    @Import(name="from")
    private @Nullable Output<String> from;

    /**
     * @return The &#39;from&#39; address header to use in all outgoing messages.
     * 
     */
    public Optional<Output<String>> from() {
        return Optional.ofNullable(this.from);
    }

    /**
     * The mail server IP address / DNS.
     * 
     */
    @Import(name="host")
    private @Nullable Output<String> host;

    /**
     * @return The mail server IP address / DNS.
     * 
     */
    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    /**
     * The password for authentication with the mail server.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password for authentication with the mail server.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The port number of the mail server.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port number of the mail server.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * A prefix to use for the subject of all outgoing mails.
     * 
     */
    @Import(name="subjectPrefix")
    private @Nullable Output<String> subjectPrefix;

    /**
     * @return A prefix to use for the subject of all outgoing mails.
     * 
     */
    public Optional<Output<String>> subjectPrefix() {
        return Optional.ofNullable(this.subjectPrefix);
    }

    /**
     * When set to &#39;true&#39;, uses a secure connection to the mail server.
     * 
     */
    @Import(name="useSsl")
    private @Nullable Output<Boolean> useSsl;

    /**
     * @return When set to &#39;true&#39;, uses a secure connection to the mail server.
     * 
     */
    public Optional<Output<Boolean>> useSsl() {
        return Optional.ofNullable(this.useSsl);
    }

    /**
     * When set to &#39;true&#39;, uses Transport Layer Security when connecting to the mail server.
     * 
     */
    @Import(name="useTls")
    private @Nullable Output<Boolean> useTls;

    /**
     * @return When set to &#39;true&#39;, uses Transport Layer Security when connecting to the mail server.
     * 
     */
    public Optional<Output<Boolean>> useTls() {
        return Optional.ofNullable(this.useTls);
    }

    /**
     * The username for authentication with the mail server.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username for authentication with the mail server.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private MailServerState() {}

    private MailServerState(MailServerState $) {
        this.artifactoryUrl = $.artifactoryUrl;
        this.enabled = $.enabled;
        this.from = $.from;
        this.host = $.host;
        this.password = $.password;
        this.port = $.port;
        this.subjectPrefix = $.subjectPrefix;
        this.useSsl = $.useSsl;
        this.useTls = $.useTls;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MailServerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MailServerState $;

        public Builder() {
            $ = new MailServerState();
        }

        public Builder(MailServerState defaults) {
            $ = new MailServerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param artifactoryUrl The Artifactory URL to to link to in all outgoing messages.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryUrl(@Nullable Output<String> artifactoryUrl) {
            $.artifactoryUrl = artifactoryUrl;
            return this;
        }

        /**
         * @param artifactoryUrl The Artifactory URL to to link to in all outgoing messages.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryUrl(String artifactoryUrl) {
            return artifactoryUrl(Output.of(artifactoryUrl));
        }

        /**
         * @param enabled When set, mail notifications are enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled When set, mail notifications are enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param from The &#39;from&#39; address header to use in all outgoing messages.
         * 
         * @return builder
         * 
         */
        public Builder from(@Nullable Output<String> from) {
            $.from = from;
            return this;
        }

        /**
         * @param from The &#39;from&#39; address header to use in all outgoing messages.
         * 
         * @return builder
         * 
         */
        public Builder from(String from) {
            return from(Output.of(from));
        }

        /**
         * @param host The mail server IP address / DNS.
         * 
         * @return builder
         * 
         */
        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The mail server IP address / DNS.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param password The password for authentication with the mail server.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password for authentication with the mail server.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param port The port number of the mail server.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port number of the mail server.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param subjectPrefix A prefix to use for the subject of all outgoing mails.
         * 
         * @return builder
         * 
         */
        public Builder subjectPrefix(@Nullable Output<String> subjectPrefix) {
            $.subjectPrefix = subjectPrefix;
            return this;
        }

        /**
         * @param subjectPrefix A prefix to use for the subject of all outgoing mails.
         * 
         * @return builder
         * 
         */
        public Builder subjectPrefix(String subjectPrefix) {
            return subjectPrefix(Output.of(subjectPrefix));
        }

        /**
         * @param useSsl When set to &#39;true&#39;, uses a secure connection to the mail server.
         * 
         * @return builder
         * 
         */
        public Builder useSsl(@Nullable Output<Boolean> useSsl) {
            $.useSsl = useSsl;
            return this;
        }

        /**
         * @param useSsl When set to &#39;true&#39;, uses a secure connection to the mail server.
         * 
         * @return builder
         * 
         */
        public Builder useSsl(Boolean useSsl) {
            return useSsl(Output.of(useSsl));
        }

        /**
         * @param useTls When set to &#39;true&#39;, uses Transport Layer Security when connecting to the mail server.
         * 
         * @return builder
         * 
         */
        public Builder useTls(@Nullable Output<Boolean> useTls) {
            $.useTls = useTls;
            return this;
        }

        /**
         * @param useTls When set to &#39;true&#39;, uses Transport Layer Security when connecting to the mail server.
         * 
         * @return builder
         * 
         */
        public Builder useTls(Boolean useTls) {
            return useTls(Output.of(useTls));
        }

        /**
         * @param username The username for authentication with the mail server.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username for authentication with the mail server.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public MailServerState build() {
            return $;
        }
    }

}
