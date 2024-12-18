// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetPermissionTargetBuildActionsGroup;
import com.pulumi.artifactory.inputs.GetPermissionTargetBuildActionsUser;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPermissionTargetBuildActions extends com.pulumi.resources.InvokeArgs {

    public static final GetPermissionTargetBuildActions Empty = new GetPermissionTargetBuildActions();

    /**
     * Groups this permission applies for.
     * 
     */
    @Import(name="groups")
    private @Nullable List<GetPermissionTargetBuildActionsGroup> groups;

    /**
     * @return Groups this permission applies for.
     * 
     */
    public Optional<List<GetPermissionTargetBuildActionsGroup>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * Users this permission target applies for.
     * 
     */
    @Import(name="users")
    private @Nullable List<GetPermissionTargetBuildActionsUser> users;

    /**
     * @return Users this permission target applies for.
     * 
     */
    public Optional<List<GetPermissionTargetBuildActionsUser>> users() {
        return Optional.ofNullable(this.users);
    }

    private GetPermissionTargetBuildActions() {}

    private GetPermissionTargetBuildActions(GetPermissionTargetBuildActions $) {
        this.groups = $.groups;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetBuildActions defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetBuildActions $;

        public Builder() {
            $ = new GetPermissionTargetBuildActions();
        }

        public Builder(GetPermissionTargetBuildActions defaults) {
            $ = new GetPermissionTargetBuildActions(Objects.requireNonNull(defaults));
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable List<GetPermissionTargetBuildActionsGroup> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(GetPermissionTargetBuildActionsGroup... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(@Nullable List<GetPermissionTargetBuildActionsUser> users) {
            $.users = users;
            return this;
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(GetPermissionTargetBuildActionsUser... users) {
            return users(List.of(users));
        }

        public GetPermissionTargetBuildActions build() {
            return $;
        }
    }

}