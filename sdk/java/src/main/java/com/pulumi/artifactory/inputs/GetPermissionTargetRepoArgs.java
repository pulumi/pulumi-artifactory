// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetPermissionTargetRepoActionsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPermissionTargetRepoArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetPermissionTargetRepoArgs Empty = new GetPermissionTargetRepoArgs();

    @Import(name="actions")
    private @Nullable Output<GetPermissionTargetRepoActionsArgs> actions;

    public Optional<Output<GetPermissionTargetRepoActionsArgs>> actions() {
        return Optional.ofNullable(this.actions);
    }

    /**
     * Pattern of artifacts to exclude.
     * 
     */
    @Import(name="excludesPatterns")
    private @Nullable Output<List<String>> excludesPatterns;

    /**
     * @return Pattern of artifacts to exclude.
     * 
     */
    public Optional<Output<List<String>>> excludesPatterns() {
        return Optional.ofNullable(this.excludesPatterns);
    }

    /**
     * Pattern of artifacts to include.
     * 
     */
    @Import(name="includesPatterns")
    private @Nullable Output<List<String>> includesPatterns;

    /**
     * @return Pattern of artifacts to include.
     * 
     */
    public Optional<Output<List<String>>> includesPatterns() {
        return Optional.ofNullable(this.includesPatterns);
    }

    /**
     * List of repositories this permission target is applicable for. You can specify the
     * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
     * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
     * 
     */
    @Import(name="repositories", required=true)
    private Output<List<String>> repositories;

    /**
     * @return List of repositories this permission target is applicable for. You can specify the
     * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
     * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
     * 
     */
    public Output<List<String>> repositories() {
        return this.repositories;
    }

    private GetPermissionTargetRepoArgs() {}

    private GetPermissionTargetRepoArgs(GetPermissionTargetRepoArgs $) {
        this.actions = $.actions;
        this.excludesPatterns = $.excludesPatterns;
        this.includesPatterns = $.includesPatterns;
        this.repositories = $.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetRepoArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetRepoArgs $;

        public Builder() {
            $ = new GetPermissionTargetRepoArgs();
        }

        public Builder(GetPermissionTargetRepoArgs defaults) {
            $ = new GetPermissionTargetRepoArgs(Objects.requireNonNull(defaults));
        }

        public Builder actions(@Nullable Output<GetPermissionTargetRepoActionsArgs> actions) {
            $.actions = actions;
            return this;
        }

        public Builder actions(GetPermissionTargetRepoActionsArgs actions) {
            return actions(Output.of(actions));
        }

        /**
         * @param excludesPatterns Pattern of artifacts to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludesPatterns(@Nullable Output<List<String>> excludesPatterns) {
            $.excludesPatterns = excludesPatterns;
            return this;
        }

        /**
         * @param excludesPatterns Pattern of artifacts to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludesPatterns(List<String> excludesPatterns) {
            return excludesPatterns(Output.of(excludesPatterns));
        }

        /**
         * @param excludesPatterns Pattern of artifacts to exclude.
         * 
         * @return builder
         * 
         */
        public Builder excludesPatterns(String... excludesPatterns) {
            return excludesPatterns(List.of(excludesPatterns));
        }

        /**
         * @param includesPatterns Pattern of artifacts to include.
         * 
         * @return builder
         * 
         */
        public Builder includesPatterns(@Nullable Output<List<String>> includesPatterns) {
            $.includesPatterns = includesPatterns;
            return this;
        }

        /**
         * @param includesPatterns Pattern of artifacts to include.
         * 
         * @return builder
         * 
         */
        public Builder includesPatterns(List<String> includesPatterns) {
            return includesPatterns(Output.of(includesPatterns));
        }

        /**
         * @param includesPatterns Pattern of artifacts to include.
         * 
         * @return builder
         * 
         */
        public Builder includesPatterns(String... includesPatterns) {
            return includesPatterns(List.of(includesPatterns));
        }

        /**
         * @param repositories List of repositories this permission target is applicable for. You can specify the
         * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
         * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
         * 
         * @return builder
         * 
         */
        public Builder repositories(Output<List<String>> repositories) {
            $.repositories = repositories;
            return this;
        }

        /**
         * @param repositories List of repositories this permission target is applicable for. You can specify the
         * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
         * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
         * 
         * @return builder
         * 
         */
        public Builder repositories(List<String> repositories) {
            return repositories(Output.of(repositories));
        }

        /**
         * @param repositories List of repositories this permission target is applicable for. You can specify the
         * name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
         * repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
         * 
         * @return builder
         * 
         */
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        public GetPermissionTargetRepoArgs build() {
            $.repositories = Objects.requireNonNull($.repositories, "expected parameter 'repositories' to be non-null");
            return $;
        }
    }

}