// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetPermissionTargetRepoActionsGroupArgs;
import com.pulumi.artifactory.inputs.GetPermissionTargetRepoActionsUserArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPermissionTargetRepoActionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetPermissionTargetRepoActionsArgs Empty = new GetPermissionTargetRepoActionsArgs();

    /**
     * Groups this permission applies for.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<GetPermissionTargetRepoActionsGroupArgs>> groups;

    /**
     * @return Groups this permission applies for.
     * 
     */
    public Optional<Output<List<GetPermissionTargetRepoActionsGroupArgs>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * Users this permission target applies for.
     * 
     */
    @Import(name="users")
    private @Nullable Output<List<GetPermissionTargetRepoActionsUserArgs>> users;

    /**
     * @return Users this permission target applies for.
     * 
     */
    public Optional<Output<List<GetPermissionTargetRepoActionsUserArgs>>> users() {
        return Optional.ofNullable(this.users);
    }

    private GetPermissionTargetRepoActionsArgs() {}

    private GetPermissionTargetRepoActionsArgs(GetPermissionTargetRepoActionsArgs $) {
        this.groups = $.groups;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetRepoActionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetRepoActionsArgs $;

        public Builder() {
            $ = new GetPermissionTargetRepoActionsArgs();
        }

        public Builder(GetPermissionTargetRepoActionsArgs defaults) {
            $ = new GetPermissionTargetRepoActionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<GetPermissionTargetRepoActionsGroupArgs>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(List<GetPermissionTargetRepoActionsGroupArgs> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(GetPermissionTargetRepoActionsGroupArgs... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(@Nullable Output<List<GetPermissionTargetRepoActionsUserArgs>> users) {
            $.users = users;
            return this;
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(List<GetPermissionTargetRepoActionsUserArgs> users) {
            return users(Output.of(users));
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(GetPermissionTargetRepoActionsUserArgs... users) {
            return users(List.of(users));
        }

        public GetPermissionTargetRepoActionsArgs build() {
            return $;
        }
    }

}