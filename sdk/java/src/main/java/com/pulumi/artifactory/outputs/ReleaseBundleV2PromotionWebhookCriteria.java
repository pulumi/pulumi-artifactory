// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ReleaseBundleV2PromotionWebhookCriteria {
    private @Nullable List<String> excludePatterns;
    private @Nullable List<String> includePatterns;
    /**
     * @return Trigger on this list of environment names.
     * 
     */
    private List<String> selectedEnvironments;

    private ReleaseBundleV2PromotionWebhookCriteria() {}
    public List<String> excludePatterns() {
        return this.excludePatterns == null ? List.of() : this.excludePatterns;
    }
    public List<String> includePatterns() {
        return this.includePatterns == null ? List.of() : this.includePatterns;
    }
    /**
     * @return Trigger on this list of environment names.
     * 
     */
    public List<String> selectedEnvironments() {
        return this.selectedEnvironments;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReleaseBundleV2PromotionWebhookCriteria defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> excludePatterns;
        private @Nullable List<String> includePatterns;
        private List<String> selectedEnvironments;
        public Builder() {}
        public Builder(ReleaseBundleV2PromotionWebhookCriteria defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.excludePatterns = defaults.excludePatterns;
    	      this.includePatterns = defaults.includePatterns;
    	      this.selectedEnvironments = defaults.selectedEnvironments;
        }

        @CustomType.Setter
        public Builder excludePatterns(@Nullable List<String> excludePatterns) {

            this.excludePatterns = excludePatterns;
            return this;
        }
        public Builder excludePatterns(String... excludePatterns) {
            return excludePatterns(List.of(excludePatterns));
        }
        @CustomType.Setter
        public Builder includePatterns(@Nullable List<String> includePatterns) {

            this.includePatterns = includePatterns;
            return this;
        }
        public Builder includePatterns(String... includePatterns) {
            return includePatterns(List.of(includePatterns));
        }
        @CustomType.Setter
        public Builder selectedEnvironments(List<String> selectedEnvironments) {
            if (selectedEnvironments == null) {
              throw new MissingRequiredPropertyException("ReleaseBundleV2PromotionWebhookCriteria", "selectedEnvironments");
            }
            this.selectedEnvironments = selectedEnvironments;
            return this;
        }
        public Builder selectedEnvironments(String... selectedEnvironments) {
            return selectedEnvironments(List.of(selectedEnvironments));
        }
        public ReleaseBundleV2PromotionWebhookCriteria build() {
            final var _resultValue = new ReleaseBundleV2PromotionWebhookCriteria();
            _resultValue.excludePatterns = excludePatterns;
            _resultValue.includePatterns = includePatterns;
            _resultValue.selectedEnvironments = selectedEnvironments;
            return _resultValue;
        }
    }
}