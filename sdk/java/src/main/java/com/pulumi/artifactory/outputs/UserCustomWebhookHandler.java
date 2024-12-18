// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UserCustomWebhookHandler {
    /**
     * @return HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
     * 
     */
    private @Nullable Map<String,String> httpHeaders;
    /**
     * @return This attribute is used to build the request body. Used in custom webhooks
     * 
     */
    private @Nullable String payload;
    /**
     * @return Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    private @Nullable String proxy;
    /**
     * @return Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won&#39;t work. Example:
     * 
     */
    private @Nullable Map<String,String> secrets;
    /**
     * @return Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    private String url;

    private UserCustomWebhookHandler() {}
    /**
     * @return HTTP headers you wish to use to invoke the Webhook, comprise key/value pair.
     * 
     */
    public Map<String,String> httpHeaders() {
        return this.httpHeaders == null ? Map.of() : this.httpHeaders;
    }
    /**
     * @return This attribute is used to build the request body. Used in custom webhooks
     * 
     */
    public Optional<String> payload() {
        return Optional.ofNullable(this.payload);
    }
    /**
     * @return Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
     * 
     */
    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }
    /**
     * @return Defines a set of sensitive values (such as, tokens and passwords) that can be injected in the headers and/or payload.Secrets’ values are encrypted. In the header/payload, the value can be invoked using the `{{.secrets.token}}` format, where token is the name provided for the secret value. Comprise key/value pair. **Note:** if multiple handlers are used, same secret name and different secret value for the same url won&#39;t work. Example:
     * 
     */
    public Map<String,String> secrets() {
        return this.secrets == null ? Map.of() : this.secrets;
    }
    /**
     * @return Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UserCustomWebhookHandler defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> httpHeaders;
        private @Nullable String payload;
        private @Nullable String proxy;
        private @Nullable Map<String,String> secrets;
        private String url;
        public Builder() {}
        public Builder(UserCustomWebhookHandler defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.httpHeaders = defaults.httpHeaders;
    	      this.payload = defaults.payload;
    	      this.proxy = defaults.proxy;
    	      this.secrets = defaults.secrets;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder httpHeaders(@Nullable Map<String,String> httpHeaders) {

            this.httpHeaders = httpHeaders;
            return this;
        }
        @CustomType.Setter
        public Builder payload(@Nullable String payload) {

            this.payload = payload;
            return this;
        }
        @CustomType.Setter
        public Builder proxy(@Nullable String proxy) {

            this.proxy = proxy;
            return this;
        }
        @CustomType.Setter
        public Builder secrets(@Nullable Map<String,String> secrets) {

            this.secrets = secrets;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("UserCustomWebhookHandler", "url");
            }
            this.url = url;
            return this;
        }
        public UserCustomWebhookHandler build() {
            final var _resultValue = new UserCustomWebhookHandler();
            _resultValue.httpHeaders = httpHeaders;
            _resultValue.payload = payload;
            _resultValue.proxy = proxy;
            _resultValue.secrets = secrets;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}