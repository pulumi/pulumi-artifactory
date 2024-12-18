// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRemoteNpmRepositoryContentSynchronisationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetRemoteNpmRepositoryContentSynchronisationArgs Empty = new GetRemoteNpmRepositoryContentSynchronisationArgs();

    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
     * 
     */
    @Import(name="propertiesEnabled")
    private @Nullable Output<Boolean> propertiesEnabled;

    /**
     * @return If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Output<Boolean>> propertiesEnabled() {
        return Optional.ofNullable(this.propertiesEnabled);
    }

    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
     * 
     */
    @Import(name="sourceOriginAbsenceDetection")
    private @Nullable Output<Boolean> sourceOriginAbsenceDetection;

    /**
     * @return If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> sourceOriginAbsenceDetection() {
        return Optional.ofNullable(this.sourceOriginAbsenceDetection);
    }

    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
     * 
     */
    @Import(name="statisticsEnabled")
    private @Nullable Output<Boolean> statisticsEnabled;

    /**
     * @return If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Output<Boolean>> statisticsEnabled() {
        return Optional.ofNullable(this.statisticsEnabled);
    }

    private GetRemoteNpmRepositoryContentSynchronisationArgs() {}

    private GetRemoteNpmRepositoryContentSynchronisationArgs(GetRemoteNpmRepositoryContentSynchronisationArgs $) {
        this.enabled = $.enabled;
        this.propertiesEnabled = $.propertiesEnabled;
        this.sourceOriginAbsenceDetection = $.sourceOriginAbsenceDetection;
        this.statisticsEnabled = $.statisticsEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRemoteNpmRepositoryContentSynchronisationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRemoteNpmRepositoryContentSynchronisationArgs $;

        public Builder() {
            $ = new GetRemoteNpmRepositoryContentSynchronisationArgs();
        }

        public Builder(GetRemoteNpmRepositoryContentSynchronisationArgs defaults) {
            $ = new GetRemoteNpmRepositoryContentSynchronisationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param propertiesEnabled If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder propertiesEnabled(@Nullable Output<Boolean> propertiesEnabled) {
            $.propertiesEnabled = propertiesEnabled;
            return this;
        }

        /**
         * @param propertiesEnabled If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder propertiesEnabled(Boolean propertiesEnabled) {
            return propertiesEnabled(Output.of(propertiesEnabled));
        }

        /**
         * @param sourceOriginAbsenceDetection If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder sourceOriginAbsenceDetection(@Nullable Output<Boolean> sourceOriginAbsenceDetection) {
            $.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            return this;
        }

        /**
         * @param sourceOriginAbsenceDetection If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder sourceOriginAbsenceDetection(Boolean sourceOriginAbsenceDetection) {
            return sourceOriginAbsenceDetection(Output.of(sourceOriginAbsenceDetection));
        }

        /**
         * @param statisticsEnabled If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder statisticsEnabled(@Nullable Output<Boolean> statisticsEnabled) {
            $.statisticsEnabled = statisticsEnabled;
            return this;
        }

        /**
         * @param statisticsEnabled If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder statisticsEnabled(Boolean statisticsEnabled) {
            return statisticsEnabled(Output.of(statisticsEnabled));
        }

        public GetRemoteNpmRepositoryContentSynchronisationArgs build() {
            return $;
        }
    }

}