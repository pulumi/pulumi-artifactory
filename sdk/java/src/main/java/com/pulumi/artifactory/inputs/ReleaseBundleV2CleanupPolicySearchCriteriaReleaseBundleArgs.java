// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs Empty = new ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs();

    /**
     * The name of the release bundle. Set `**` to include all bundles. Example: `name = &#34;**&#34;`
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the release bundle. Set `**` to include all bundles. Example: `name = &#34;**&#34;`
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The project identifier associated with the release bundle. This key is obtained from the Project Settings screen. Leave the field blank to apply at a global level.
     * 
     */
    @Import(name="projectKey", required=true)
    private Output<String> projectKey;

    /**
     * @return The project identifier associated with the release bundle. This key is obtained from the Project Settings screen. Leave the field blank to apply at a global level.
     * 
     */
    public Output<String> projectKey() {
        return this.projectKey;
    }

    private ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs() {}

    private ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs(ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs $) {
        this.name = $.name;
        this.projectKey = $.projectKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs $;

        public Builder() {
            $ = new ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs();
        }

        public Builder(ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs defaults) {
            $ = new ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the release bundle. Set `**` to include all bundles. Example: `name = &#34;**&#34;`
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the release bundle. Set `**` to include all bundles. Example: `name = &#34;**&#34;`
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectKey The project identifier associated with the release bundle. This key is obtained from the Project Settings screen. Leave the field blank to apply at a global level.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        /**
         * @param projectKey The project identifier associated with the release bundle. This key is obtained from the Project Settings screen. Leave the field blank to apply at a global level.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        public ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs", "name");
            }
            if ($.projectKey == null) {
                throw new MissingRequiredPropertyException("ReleaseBundleV2CleanupPolicySearchCriteriaReleaseBundleArgs", "projectKey");
            }
            return $;
        }
    }

}
