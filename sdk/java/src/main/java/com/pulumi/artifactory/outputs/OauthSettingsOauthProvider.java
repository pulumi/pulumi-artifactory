// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OauthSettingsOauthProvider {
    /**
     * @return OAuth user info endpoint for the IdP.
     * 
     */
    private final String apiUrl;
    /**
     * @return OAuth authorization endpoint for the IdP.
     * 
     */
    private final String authUrl;
    /**
     * @return OAuth client ID configured on the IdP.
     * 
     */
    private final String clientId;
    /**
     * @return OAuth client secret configured on the IdP.
     * 
     */
    private final String clientSecret;
    /**
     * @return Enable the Artifactory OAuth provider.  Default value is `true`.
     * 
     */
    private final @Nullable Boolean enabled;
    /**
     * @return Name of the Artifactory OAuth provider.
     * 
     */
    private final String name;
    /**
     * @return OAuth token endpoint for the IdP.
     * 
     */
    private final String tokenUrl;
    /**
     * @return Type of OAuth provider. (e.g., `github`, `google`, `cloudfoundry`, or `openId`)
     * 
     */
    private final String type;

    @CustomType.Constructor
    private OauthSettingsOauthProvider(
        @CustomType.Parameter("apiUrl") String apiUrl,
        @CustomType.Parameter("authUrl") String authUrl,
        @CustomType.Parameter("clientId") String clientId,
        @CustomType.Parameter("clientSecret") String clientSecret,
        @CustomType.Parameter("enabled") @Nullable Boolean enabled,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("tokenUrl") String tokenUrl,
        @CustomType.Parameter("type") String type) {
        this.apiUrl = apiUrl;
        this.authUrl = authUrl;
        this.clientId = clientId;
        this.clientSecret = clientSecret;
        this.enabled = enabled;
        this.name = name;
        this.tokenUrl = tokenUrl;
        this.type = type;
    }

    /**
     * @return OAuth user info endpoint for the IdP.
     * 
     */
    public String apiUrl() {
        return this.apiUrl;
    }
    /**
     * @return OAuth authorization endpoint for the IdP.
     * 
     */
    public String authUrl() {
        return this.authUrl;
    }
    /**
     * @return OAuth client ID configured on the IdP.
     * 
     */
    public String clientId() {
        return this.clientId;
    }
    /**
     * @return OAuth client secret configured on the IdP.
     * 
     */
    public String clientSecret() {
        return this.clientSecret;
    }
    /**
     * @return Enable the Artifactory OAuth provider.  Default value is `true`.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return Name of the Artifactory OAuth provider.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return OAuth token endpoint for the IdP.
     * 
     */
    public String tokenUrl() {
        return this.tokenUrl;
    }
    /**
     * @return Type of OAuth provider. (e.g., `github`, `google`, `cloudfoundry`, or `openId`)
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OauthSettingsOauthProvider defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String apiUrl;
        private String authUrl;
        private String clientId;
        private String clientSecret;
        private @Nullable Boolean enabled;
        private String name;
        private String tokenUrl;
        private String type;

        public Builder() {
    	      // Empty
        }

        public Builder(OauthSettingsOauthProvider defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiUrl = defaults.apiUrl;
    	      this.authUrl = defaults.authUrl;
    	      this.clientId = defaults.clientId;
    	      this.clientSecret = defaults.clientSecret;
    	      this.enabled = defaults.enabled;
    	      this.name = defaults.name;
    	      this.tokenUrl = defaults.tokenUrl;
    	      this.type = defaults.type;
        }

        public Builder apiUrl(String apiUrl) {
            this.apiUrl = Objects.requireNonNull(apiUrl);
            return this;
        }
        public Builder authUrl(String authUrl) {
            this.authUrl = Objects.requireNonNull(authUrl);
            return this;
        }
        public Builder clientId(String clientId) {
            this.clientId = Objects.requireNonNull(clientId);
            return this;
        }
        public Builder clientSecret(String clientSecret) {
            this.clientSecret = Objects.requireNonNull(clientSecret);
            return this;
        }
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder tokenUrl(String tokenUrl) {
            this.tokenUrl = Objects.requireNonNull(tokenUrl);
            return this;
        }
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }        public OauthSettingsOauthProvider build() {
            return new OauthSettingsOauthProvider(apiUrl, authUrl, clientId, clientSecret, enabled, name, tokenUrl, type);
        }
    }
}