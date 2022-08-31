// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ArtifactoryReleaseBundleWebhookCriteria {
    /**
     * @return Trigger on any release bundle
     * 
     */
    private final Boolean anyReleaseBundle;
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;
     * 
     */
    private final @Nullable List<String> excludePatterns;
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;
     * 
     */
    private final @Nullable List<String> includePatterns;
    /**
     * @return Trigger on this list of release bundle names
     * 
     */
    private final List<String> registeredReleaseBundleNames;

    @CustomType.Constructor
    private ArtifactoryReleaseBundleWebhookCriteria(
        @CustomType.Parameter("anyReleaseBundle") Boolean anyReleaseBundle,
        @CustomType.Parameter("excludePatterns") @Nullable List<String> excludePatterns,
        @CustomType.Parameter("includePatterns") @Nullable List<String> includePatterns,
        @CustomType.Parameter("registeredReleaseBundleNames") List<String> registeredReleaseBundleNames) {
        this.anyReleaseBundle = anyReleaseBundle;
        this.excludePatterns = excludePatterns;
        this.includePatterns = includePatterns;
        this.registeredReleaseBundleNames = registeredReleaseBundleNames;
    }

    /**
     * @return Trigger on any release bundle
     * 
     */
    public Boolean anyReleaseBundle() {
        return this.anyReleaseBundle;
    }
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;
     * 
     */
    public List<String> excludePatterns() {
        return this.excludePatterns == null ? List.of() : this.excludePatterns;
    }
    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: &#34;org/apache/**&#34;
     * 
     */
    public List<String> includePatterns() {
        return this.includePatterns == null ? List.of() : this.includePatterns;
    }
    /**
     * @return Trigger on this list of release bundle names
     * 
     */
    public List<String> registeredReleaseBundleNames() {
        return this.registeredReleaseBundleNames;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ArtifactoryReleaseBundleWebhookCriteria defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean anyReleaseBundle;
        private @Nullable List<String> excludePatterns;
        private @Nullable List<String> includePatterns;
        private List<String> registeredReleaseBundleNames;

        public Builder() {
    	      // Empty
        }

        public Builder(ArtifactoryReleaseBundleWebhookCriteria defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.anyReleaseBundle = defaults.anyReleaseBundle;
    	      this.excludePatterns = defaults.excludePatterns;
    	      this.includePatterns = defaults.includePatterns;
    	      this.registeredReleaseBundleNames = defaults.registeredReleaseBundleNames;
        }

        public Builder anyReleaseBundle(Boolean anyReleaseBundle) {
            this.anyReleaseBundle = Objects.requireNonNull(anyReleaseBundle);
            return this;
        }
        public Builder excludePatterns(@Nullable List<String> excludePatterns) {
            this.excludePatterns = excludePatterns;
            return this;
        }
        public Builder excludePatterns(String... excludePatterns) {
            return excludePatterns(List.of(excludePatterns));
        }
        public Builder includePatterns(@Nullable List<String> includePatterns) {
            this.includePatterns = includePatterns;
            return this;
        }
        public Builder includePatterns(String... includePatterns) {
            return includePatterns(List.of(includePatterns));
        }
        public Builder registeredReleaseBundleNames(List<String> registeredReleaseBundleNames) {
            this.registeredReleaseBundleNames = Objects.requireNonNull(registeredReleaseBundleNames);
            return this;
        }
        public Builder registeredReleaseBundleNames(String... registeredReleaseBundleNames) {
            return registeredReleaseBundleNames(List.of(registeredReleaseBundleNames));
        }        public ArtifactoryReleaseBundleWebhookCriteria build() {
            return new ArtifactoryReleaseBundleWebhookCriteria(anyReleaseBundle, excludePatterns, includePatterns, registeredReleaseBundleNames);
        }
    }
}