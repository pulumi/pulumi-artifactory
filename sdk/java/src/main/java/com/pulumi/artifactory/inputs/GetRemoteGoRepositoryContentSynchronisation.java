// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRemoteGoRepositoryContentSynchronisation extends com.pulumi.resources.InvokeArgs {

    public static final GetRemoteGoRepositoryContentSynchronisation Empty = new GetRemoteGoRepositoryContentSynchronisation();

    /**
     * If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
     * 
     */
    @Import(name="enabled")
    private @Nullable Boolean enabled;

    /**
     * @return If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
     * 
     */
    @Import(name="propertiesEnabled")
    private @Nullable Boolean propertiesEnabled;

    /**
     * @return If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Boolean> propertiesEnabled() {
        return Optional.ofNullable(this.propertiesEnabled);
    }

    /**
     * If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
     * 
     */
    @Import(name="sourceOriginAbsenceDetection")
    private @Nullable Boolean sourceOriginAbsenceDetection;

    /**
     * @return If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
     * 
     */
    public Optional<Boolean> sourceOriginAbsenceDetection() {
        return Optional.ofNullable(this.sourceOriginAbsenceDetection);
    }

    /**
     * If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
     * 
     */
    @Import(name="statisticsEnabled")
    private @Nullable Boolean statisticsEnabled;

    /**
     * @return If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Boolean> statisticsEnabled() {
        return Optional.ofNullable(this.statisticsEnabled);
    }

    private GetRemoteGoRepositoryContentSynchronisation() {}

    private GetRemoteGoRepositoryContentSynchronisation(GetRemoteGoRepositoryContentSynchronisation $) {
        this.enabled = $.enabled;
        this.propertiesEnabled = $.propertiesEnabled;
        this.sourceOriginAbsenceDetection = $.sourceOriginAbsenceDetection;
        this.statisticsEnabled = $.statisticsEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRemoteGoRepositoryContentSynchronisation defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRemoteGoRepositoryContentSynchronisation $;

        public Builder() {
            $ = new GetRemoteGoRepositoryContentSynchronisation();
        }

        public Builder(GetRemoteGoRepositoryContentSynchronisation defaults) {
            $ = new GetRemoteGoRepositoryContentSynchronisation(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled If set, Remote repository proxies a local or remote repository from another instance of Artifactory. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Boolean enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param propertiesEnabled If set, properties for artifacts that have been cached in this repository will be updated if they are modified in the artifact hosted at the remote Artifactory instance. The trigger to synchronize the properties is download of the artifact from the remote repository cache of the local Artifactory instance. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder propertiesEnabled(@Nullable Boolean propertiesEnabled) {
            $.propertiesEnabled = propertiesEnabled;
            return this;
        }

        /**
         * @param sourceOriginAbsenceDetection If set, Artifactory displays an indication on cached items if they have been deleted from the corresponding repository in the remote Artifactory instance. Default value is &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder sourceOriginAbsenceDetection(@Nullable Boolean sourceOriginAbsenceDetection) {
            $.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            return this;
        }

        /**
         * @param statisticsEnabled If set, Artifactory will notify the remote instance whenever an artifact in the Smart Remote Repository is downloaded locally so that it can update its download counter. Note that if this option is not set, there may be a discrepancy between the number of artifacts reported to have been downloaded in the different Artifactory instances of the proxy chain. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder statisticsEnabled(@Nullable Boolean statisticsEnabled) {
            $.statisticsEnabled = statisticsEnabled;
            return this;
        }

        public GetRemoteGoRepositoryContentSynchronisation build() {
            return $;
        }
    }

}
