// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualBowerRepositoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualBowerRepositoryArgs Empty = new VirtualBowerRepositoryArgs();

    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    @Import(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    private @Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    /**
     * @return Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    public Optional<Output<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }

    /**
     * Default repository to deploy artifacts.
     * 
     */
    @Import(name="defaultDeploymentRepo")
    private @Nullable Output<String> defaultDeploymentRepo;

    /**
     * @return Default repository to deploy artifacts.
     * 
     */
    public Optional<Output<String>> defaultDeploymentRepo() {
        return Optional.ofNullable(this.defaultDeploymentRepo);
    }

    /**
     * Public description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Public description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
     * artifacts are excluded.
     * 
     */
    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
     * artifacts are excluded.
     * 
     */
    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    /**
     * When set, external dependencies are rewritten. Default value is false.
     * 
     */
    @Import(name="externalDependenciesEnabled")
    private @Nullable Output<Boolean> externalDependenciesEnabled;

    /**
     * @return When set, external dependencies are rewritten. Default value is false.
     * 
     */
    public Optional<Output<Boolean>> externalDependenciesEnabled() {
        return Optional.ofNullable(this.externalDependenciesEnabled);
    }

    /**
     * An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    @Import(name="externalDependenciesPatterns")
    private @Nullable Output<List<String>> externalDependenciesPatterns;

    /**
     * @return An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    public Optional<Output<List<String>>> externalDependenciesPatterns() {
        return Optional.ofNullable(this.externalDependenciesPatterns);
    }

    /**
     * The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    @Import(name="externalDependenciesRemoteRepo")
    private @Nullable Output<String> externalDependenciesRemoteRepo;

    /**
     * @return The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    public Optional<Output<String>> externalDependenciesRemoteRepo() {
        return Optional.ofNullable(this.externalDependenciesRemoteRepo);
    }

    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
     * 
     */
    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
     * 
     */
    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * Internal description.
     * 
     */
    @Import(name="notes")
    private @Nullable Output<String> notes;

    /**
     * @return Internal description.
     * 
     */
    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    /**
     * Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    /**
     * Repository layout key for the virtual repository
     * 
     */
    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    /**
     * @return Repository layout key for the virtual repository
     * 
     */
    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    /**
     * The effective list of actual repositories included in this virtual repository.
     * 
     */
    @Import(name="repositories")
    private @Nullable Output<List<String>> repositories;

    /**
     * @return The effective list of actual repositories included in this virtual repository.
     * 
     */
    public Optional<Output<List<String>>> repositories() {
        return Optional.ofNullable(this.repositories);
    }

    private VirtualBowerRepositoryArgs() {}

    private VirtualBowerRepositoryArgs(VirtualBowerRepositoryArgs $) {
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
    public static Builder builder(VirtualBowerRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualBowerRepositoryArgs $;

        public Builder() {
            $ = new VirtualBowerRepositoryArgs();
        }

        public Builder(VirtualBowerRepositoryArgs defaults) {
            $ = new VirtualBowerRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param artifactoryRequestsCanRetrieveRemoteArtifacts Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
         * another Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts) {
            $.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }

        /**
         * @param artifactoryRequestsCanRetrieveRemoteArtifacts Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
         * another Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            return artifactoryRequestsCanRetrieveRemoteArtifacts(Output.of(artifactoryRequestsCanRetrieveRemoteArtifacts));
        }

        /**
         * @param defaultDeploymentRepo Default repository to deploy artifacts.
         * 
         * @return builder
         * 
         */
        public Builder defaultDeploymentRepo(@Nullable Output<String> defaultDeploymentRepo) {
            $.defaultDeploymentRepo = defaultDeploymentRepo;
            return this;
        }

        /**
         * @param defaultDeploymentRepo Default repository to deploy artifacts.
         * 
         * @return builder
         * 
         */
        public Builder defaultDeploymentRepo(String defaultDeploymentRepo) {
            return defaultDeploymentRepo(Output.of(defaultDeploymentRepo));
        }

        /**
         * @param description Public description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Public description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        /**
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**&#47;z/*`.By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        /**
         * @param externalDependenciesEnabled When set, external dependencies are rewritten. Default value is false.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesEnabled(@Nullable Output<Boolean> externalDependenciesEnabled) {
            $.externalDependenciesEnabled = externalDependenciesEnabled;
            return this;
        }

        /**
         * @param externalDependenciesEnabled When set, external dependencies are rewritten. Default value is false.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesEnabled(Boolean externalDependenciesEnabled) {
            return externalDependenciesEnabled(Output.of(externalDependenciesEnabled));
        }

        /**
         * @param externalDependenciesPatterns An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(@Nullable Output<List<String>> externalDependenciesPatterns) {
            $.externalDependenciesPatterns = externalDependenciesPatterns;
            return this;
        }

        /**
         * @param externalDependenciesPatterns An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(List<String> externalDependenciesPatterns) {
            return externalDependenciesPatterns(Output.of(externalDependenciesPatterns));
        }

        /**
         * @param externalDependenciesPatterns An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(String... externalDependenciesPatterns) {
            return externalDependenciesPatterns(List.of(externalDependenciesPatterns));
        }

        /**
         * @param externalDependenciesRemoteRepo The remote repository aggregated by this virtual repository in which the external dependency will be cached.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesRemoteRepo(@Nullable Output<String> externalDependenciesRemoteRepo) {
            $.externalDependenciesRemoteRepo = externalDependenciesRemoteRepo;
            return this;
        }

        /**
         * @param externalDependenciesRemoteRepo The remote repository aggregated by this virtual repository in which the external dependency will be cached.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesRemoteRepo(String externalDependenciesRemoteRepo) {
            return externalDependenciesRemoteRepo(Output.of(externalDependenciesRemoteRepo));
        }

        /**
         * @param includesPattern List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
         * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        /**
         * @param includesPattern List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**&#47;z/*`. When
         * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**&#47;*`).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        /**
         * @param key A mandatory identifier for the repository that must be unique. It cannot begin with a number or
         * contain spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key A mandatory identifier for the repository that must be unique. It cannot begin with a number or
         * contain spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param notes Internal description.
         * 
         * @return builder
         * 
         */
        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        /**
         * @param notes Internal description.
         * 
         * @return builder
         * 
         */
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

        /**
         * @param projectKey Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
         * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        /**
         * @param projectKey Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
         * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        /**
         * @param repoLayoutRef Repository layout key for the virtual repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        /**
         * @param repoLayoutRef Repository layout key for the virtual repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(@Nullable Output<List<String>> repositories) {
            $.repositories = repositories;
            return this;
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(List<String> repositories) {
            return repositories(Output.of(repositories));
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        public VirtualBowerRepositoryArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("VirtualBowerRepositoryArgs", "key");
            }
            return $;
        }
    }

}
