// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.GetPermissionTargetReleaseBundleActionsGroup;
import com.pulumi.artifactory.inputs.GetPermissionTargetReleaseBundleActionsUser;
import com.pulumi.core.annotations.Import;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPermissionTargetReleaseBundleActions extends com.pulumi.resources.InvokeArgs {

    public static final GetPermissionTargetReleaseBundleActions Empty = new GetPermissionTargetReleaseBundleActions();

    /**
     * Groups this permission applies for.
     * 
     */
    @Import(name="groups")
    private @Nullable List<GetPermissionTargetReleaseBundleActionsGroup> groups;

    /**
     * @return Groups this permission applies for.
     * 
     */
    public Optional<List<GetPermissionTargetReleaseBundleActionsGroup>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * Users this permission target applies for.
     * 
     */
    @Import(name="users")
    private @Nullable List<GetPermissionTargetReleaseBundleActionsUser> users;

    /**
     * @return Users this permission target applies for.
     * 
     */
    public Optional<List<GetPermissionTargetReleaseBundleActionsUser>> users() {
        return Optional.ofNullable(this.users);
    }

    private GetPermissionTargetReleaseBundleActions() {}

    private GetPermissionTargetReleaseBundleActions(GetPermissionTargetReleaseBundleActions $) {
        this.groups = $.groups;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPermissionTargetReleaseBundleActions defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPermissionTargetReleaseBundleActions $;

        public Builder() {
            $ = new GetPermissionTargetReleaseBundleActions();
        }

        public Builder(GetPermissionTargetReleaseBundleActions defaults) {
            $ = new GetPermissionTargetReleaseBundleActions(Objects.requireNonNull(defaults));
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable List<GetPermissionTargetReleaseBundleActionsGroup> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups Groups this permission applies for.
         * 
         * @return builder
         * 
         */
        public Builder groups(GetPermissionTargetReleaseBundleActionsGroup... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(@Nullable List<GetPermissionTargetReleaseBundleActionsUser> users) {
            $.users = users;
            return this;
        }

        /**
         * @param users Users this permission target applies for.
         * 
         * @return builder
         * 
         */
        public Builder users(GetPermissionTargetReleaseBundleActionsUser... users) {
            return users(List.of(users));
        }

        public GetPermissionTargetReleaseBundleActions build() {
            return $;
        }
    }

}
