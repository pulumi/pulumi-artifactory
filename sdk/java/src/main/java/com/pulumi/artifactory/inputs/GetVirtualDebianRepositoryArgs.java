// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVirtualDebianRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVirtualDebianRepositoryArgs Empty = new GetVirtualDebianRepositoryArgs();

    @Import(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    private @Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    public Optional<Output<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }

    /**
     * (Optional) Specifying  architectures will speed up Artifactory&#39;s initial metadata indexing process. The default architecture values are amd64 and i386.
     * 
     */
    @Import(name="debianDefaultArchitectures")
    private @Nullable Output<String> debianDefaultArchitectures;

    /**
     * @return (Optional) Specifying  architectures will speed up Artifactory&#39;s initial metadata indexing process. The default architecture values are amd64 and i386.
     * 
     */
    public Optional<Output<String>> debianDefaultArchitectures() {
        return Optional.ofNullable(this.debianDefaultArchitectures);
    }

    @Import(name="defaultDeploymentRepo")
    private @Nullable Output<String> defaultDeploymentRepo;

    public Optional<Output<String>> defaultDeploymentRepo() {
        return Optional.ofNullable(this.defaultDeploymentRepo);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    @Import(name="notes")
    private @Nullable Output<String> notes;

    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    /**
     * (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
     * 
     */
    @Import(name="optionalIndexCompressionFormats")
    private @Nullable Output<List<String>> optionalIndexCompressionFormats;

    /**
     * @return (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
     * 
     */
    public Optional<Output<List<String>>> optionalIndexCompressionFormats() {
        return Optional.ofNullable(this.optionalIndexCompressionFormats);
    }

    /**
     * (Optional) Primary keypair used to sign artifacts. Default is empty.
     * 
     */
    @Import(name="primaryKeypairRef")
    private @Nullable Output<String> primaryKeypairRef;

    /**
     * @return (Optional) Primary keypair used to sign artifacts. Default is empty.
     * 
     */
    public Optional<Output<String>> primaryKeypairRef() {
        return Optional.ofNullable(this.primaryKeypairRef);
    }

    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="repositories")
    private @Nullable Output<List<String>> repositories;

    public Optional<Output<List<String>>> repositories() {
        return Optional.ofNullable(this.repositories);
    }

    /**
     * (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     * 
     */
    @Import(name="retrievalCachePeriodSeconds")
    private @Nullable Output<Integer> retrievalCachePeriodSeconds;

    /**
     * @return (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
     * 
     */
    public Optional<Output<Integer>> retrievalCachePeriodSeconds() {
        return Optional.ofNullable(this.retrievalCachePeriodSeconds);
    }

    /**
     * (Optional) Secondary keypair used to sign artifacts. Default is empty.
     * 
     */
    @Import(name="secondaryKeypairRef")
    private @Nullable Output<String> secondaryKeypairRef;

    /**
     * @return (Optional) Secondary keypair used to sign artifacts. Default is empty.
     * 
     */
    public Optional<Output<String>> secondaryKeypairRef() {
        return Optional.ofNullable(this.secondaryKeypairRef);
    }

    private GetVirtualDebianRepositoryArgs() {}

    private GetVirtualDebianRepositoryArgs(GetVirtualDebianRepositoryArgs $) {
        this.artifactoryRequestsCanRetrieveRemoteArtifacts = $.artifactoryRequestsCanRetrieveRemoteArtifacts;
        this.debianDefaultArchitectures = $.debianDefaultArchitectures;
        this.defaultDeploymentRepo = $.defaultDeploymentRepo;
        this.description = $.description;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.notes = $.notes;
        this.optionalIndexCompressionFormats = $.optionalIndexCompressionFormats;
        this.primaryKeypairRef = $.primaryKeypairRef;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.repoLayoutRef = $.repoLayoutRef;
        this.repositories = $.repositories;
        this.retrievalCachePeriodSeconds = $.retrievalCachePeriodSeconds;
        this.secondaryKeypairRef = $.secondaryKeypairRef;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVirtualDebianRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVirtualDebianRepositoryArgs $;

        public Builder() {
            $ = new GetVirtualDebianRepositoryArgs();
        }

        public Builder(GetVirtualDebianRepositoryArgs defaults) {
            $ = new GetVirtualDebianRepositoryArgs(Objects.requireNonNull(defaults));
        }

        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts) {
            $.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }

        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            return artifactoryRequestsCanRetrieveRemoteArtifacts(Output.of(artifactoryRequestsCanRetrieveRemoteArtifacts));
        }

        /**
         * @param debianDefaultArchitectures (Optional) Specifying  architectures will speed up Artifactory&#39;s initial metadata indexing process. The default architecture values are amd64 and i386.
         * 
         * @return builder
         * 
         */
        public Builder debianDefaultArchitectures(@Nullable Output<String> debianDefaultArchitectures) {
            $.debianDefaultArchitectures = debianDefaultArchitectures;
            return this;
        }

        /**
         * @param debianDefaultArchitectures (Optional) Specifying  architectures will speed up Artifactory&#39;s initial metadata indexing process. The default architecture values are amd64 and i386.
         * 
         * @return builder
         * 
         */
        public Builder debianDefaultArchitectures(String debianDefaultArchitectures) {
            return debianDefaultArchitectures(Output.of(debianDefaultArchitectures));
        }

        public Builder defaultDeploymentRepo(@Nullable Output<String> defaultDeploymentRepo) {
            $.defaultDeploymentRepo = defaultDeploymentRepo;
            return this;
        }

        public Builder defaultDeploymentRepo(String defaultDeploymentRepo) {
            return defaultDeploymentRepo(Output.of(defaultDeploymentRepo));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        /**
         * @param optionalIndexCompressionFormats (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
         * 
         * @return builder
         * 
         */
        public Builder optionalIndexCompressionFormats(@Nullable Output<List<String>> optionalIndexCompressionFormats) {
            $.optionalIndexCompressionFormats = optionalIndexCompressionFormats;
            return this;
        }

        /**
         * @param optionalIndexCompressionFormats (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
         * 
         * @return builder
         * 
         */
        public Builder optionalIndexCompressionFormats(List<String> optionalIndexCompressionFormats) {
            return optionalIndexCompressionFormats(Output.of(optionalIndexCompressionFormats));
        }

        /**
         * @param optionalIndexCompressionFormats (Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are `bz2`,`lzma` and `xz`. Default value is `bz2`.
         * 
         * @return builder
         * 
         */
        public Builder optionalIndexCompressionFormats(String... optionalIndexCompressionFormats) {
            return optionalIndexCompressionFormats(List.of(optionalIndexCompressionFormats));
        }

        /**
         * @param primaryKeypairRef (Optional) Primary keypair used to sign artifacts. Default is empty.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeypairRef(@Nullable Output<String> primaryKeypairRef) {
            $.primaryKeypairRef = primaryKeypairRef;
            return this;
        }

        /**
         * @param primaryKeypairRef (Optional) Primary keypair used to sign artifacts. Default is empty.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeypairRef(String primaryKeypairRef) {
            return primaryKeypairRef(Output.of(primaryKeypairRef));
        }

        public Builder projectEnvironments(@Nullable Output<List<String>> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        public Builder projectEnvironments(List<String> projectEnvironments) {
            return projectEnvironments(Output.of(projectEnvironments));
        }

        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        public Builder repositories(@Nullable Output<List<String>> repositories) {
            $.repositories = repositories;
            return this;
        }

        public Builder repositories(List<String> repositories) {
            return repositories(Output.of(repositories));
        }

        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        /**
         * @param retrievalCachePeriodSeconds (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
         * 
         * @return builder
         * 
         */
        public Builder retrievalCachePeriodSeconds(@Nullable Output<Integer> retrievalCachePeriodSeconds) {
            $.retrievalCachePeriodSeconds = retrievalCachePeriodSeconds;
            return this;
        }

        /**
         * @param retrievalCachePeriodSeconds (Optional, Default: `7200`) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
         * 
         * @return builder
         * 
         */
        public Builder retrievalCachePeriodSeconds(Integer retrievalCachePeriodSeconds) {
            return retrievalCachePeriodSeconds(Output.of(retrievalCachePeriodSeconds));
        }

        /**
         * @param secondaryKeypairRef (Optional) Secondary keypair used to sign artifacts. Default is empty.
         * 
         * @return builder
         * 
         */
        public Builder secondaryKeypairRef(@Nullable Output<String> secondaryKeypairRef) {
            $.secondaryKeypairRef = secondaryKeypairRef;
            return this;
        }

        /**
         * @param secondaryKeypairRef (Optional) Secondary keypair used to sign artifacts. Default is empty.
         * 
         * @return builder
         * 
         */
        public Builder secondaryKeypairRef(String secondaryKeypairRef) {
            return secondaryKeypairRef(Output.of(secondaryKeypairRef));
        }

        public GetVirtualDebianRepositoryArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetVirtualDebianRepositoryArgs", "key");
            }
            return $;
        }
    }

}
