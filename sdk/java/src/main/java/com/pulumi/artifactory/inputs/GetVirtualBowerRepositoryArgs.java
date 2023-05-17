// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVirtualBowerRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVirtualBowerRepositoryArgs Empty = new GetVirtualBowerRepositoryArgs();

    @Import(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    private @Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    public Optional<Output<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
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

    /**
     * (Optional) When set, external dependencies are rewritten. Default value is false.
     * 
     */
    @Import(name="externalDependenciesEnabled")
    private @Nullable Output<Boolean> externalDependenciesEnabled;

    /**
     * @return (Optional) When set, external dependencies are rewritten. Default value is false.
     * 
     */
    public Optional<Output<Boolean>> externalDependenciesEnabled() {
        return Optional.ofNullable(this.externalDependenciesEnabled);
    }

    /**
     * (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    @Import(name="externalDependenciesPatterns")
    private @Nullable Output<List<String>> externalDependenciesPatterns;

    /**
     * @return (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    public Optional<Output<List<String>>> externalDependenciesPatterns() {
        return Optional.ofNullable(this.externalDependenciesPatterns);
    }

    /**
     * (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    @Import(name="externalDependenciesRemoteRepo")
    private @Nullable Output<String> externalDependenciesRemoteRepo;

    /**
     * @return (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    public Optional<Output<String>> externalDependenciesRemoteRepo() {
        return Optional.ofNullable(this.externalDependenciesRemoteRepo);
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

    private GetVirtualBowerRepositoryArgs() {}

    private GetVirtualBowerRepositoryArgs(GetVirtualBowerRepositoryArgs $) {
        this.artifactoryRequestsCanRetrieveRemoteArtifacts = $.artifactoryRequestsCanRetrieveRemoteArtifacts;
        this.defaultDeploymentRepo = $.defaultDeploymentRepo;
        this.description = $.description;
        this.excludesPattern = $.excludesPattern;
        this.externalDependenciesEnabled = $.externalDependenciesEnabled;
        this.externalDependenciesPatterns = $.externalDependenciesPatterns;
        this.externalDependenciesRemoteRepo = $.externalDependenciesRemoteRepo;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.notes = $.notes;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.repoLayoutRef = $.repoLayoutRef;
        this.repositories = $.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVirtualBowerRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVirtualBowerRepositoryArgs $;

        public Builder() {
            $ = new GetVirtualBowerRepositoryArgs();
        }

        public Builder(GetVirtualBowerRepositoryArgs defaults) {
            $ = new GetVirtualBowerRepositoryArgs(Objects.requireNonNull(defaults));
        }

        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts) {
            $.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }

        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            return artifactoryRequestsCanRetrieveRemoteArtifacts(Output.of(artifactoryRequestsCanRetrieveRemoteArtifacts));
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

        /**
         * @param externalDependenciesEnabled (Optional) When set, external dependencies are rewritten. Default value is false.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesEnabled(@Nullable Output<Boolean> externalDependenciesEnabled) {
            $.externalDependenciesEnabled = externalDependenciesEnabled;
            return this;
        }

        /**
         * @param externalDependenciesEnabled (Optional) When set, external dependencies are rewritten. Default value is false.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesEnabled(Boolean externalDependenciesEnabled) {
            return externalDependenciesEnabled(Output.of(externalDependenciesEnabled));
        }

        /**
         * @param externalDependenciesPatterns (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(@Nullable Output<List<String>> externalDependenciesPatterns) {
            $.externalDependenciesPatterns = externalDependenciesPatterns;
            return this;
        }

        /**
         * @param externalDependenciesPatterns (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(List<String> externalDependenciesPatterns) {
            return externalDependenciesPatterns(Output.of(externalDependenciesPatterns));
        }

        /**
         * @param externalDependenciesPatterns (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(String... externalDependenciesPatterns) {
            return externalDependenciesPatterns(List.of(externalDependenciesPatterns));
        }

        /**
         * @param externalDependenciesRemoteRepo (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesRemoteRepo(@Nullable Output<String> externalDependenciesRemoteRepo) {
            $.externalDependenciesRemoteRepo = externalDependenciesRemoteRepo;
            return this;
        }

        /**
         * @param externalDependenciesRemoteRepo (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesRemoteRepo(String externalDependenciesRemoteRepo) {
            return externalDependenciesRemoteRepo(Output.of(externalDependenciesRemoteRepo));
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

        public GetVirtualBowerRepositoryArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}