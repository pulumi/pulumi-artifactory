// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.GroupArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.GroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * ```sh
 * $ pulumi import artifactory:index/group:Group terraform-group mygroup
 * ```
 * 
 * ~&gt; `users_names` can&#39;t be imported due to API limitations.
 * 
 */
@ResourceType(type="artifactory:index/group:Group")
public class Group extends com.pulumi.resources.CustomResource {
    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     * 
     */
    @Export(name="adminPrivileges", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> adminPrivileges;

    /**
     * @return Any users added to this group will automatically be assigned with admin privileges in the system.
     * 
     */
    public Output<Boolean> adminPrivileges() {
        return this.adminPrivileges;
    }
    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     * 
     */
    @Export(name="autoJoin", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> autoJoin;

    /**
     * @return When this parameter is set, any new users defined in the system are automatically assigned to this group.
     * 
     */
    public Output<Boolean> autoJoin() {
        return this.autoJoin;
    }
    /**
     * A description for the group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return A description for the group.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * When this is set to `true`, an empty or missing usernames array will detach all users from the group.
     * 
     */
    @Export(name="detachAllUsers", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> detachAllUsers;

    /**
     * @return When this is set to `true`, an empty or missing usernames array will detach all users from the group.
     * 
     */
    public Output<Optional<Boolean>> detachAllUsers() {
        return Codegen.optional(this.detachAllUsers);
    }
    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     * 
     */
    @Export(name="externalId", refs={String.class}, tree="[0]")
    private Output<String> externalId;

    /**
     * @return New external group ID used to configure the corresponding group in Azure AD.
     * 
     */
    public Output<String> externalId() {
        return this.externalId;
    }
    /**
     * Name of the group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     * 
     */
    @Export(name="policyManager", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> policyManager;

    /**
     * @return When this override is set, User in the group can set Xray security and compliance policies. Default value is `false`.
     * 
     */
    public Output<Boolean> policyManager() {
        return this.policyManager;
    }
    /**
     * The realm for the group.
     * 
     */
    @Export(name="realm", refs={String.class}, tree="[0]")
    private Output<String> realm;

    /**
     * @return The realm for the group.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }
    /**
     * The realm attributes for the group.
     * 
     */
    @Export(name="realmAttributes", refs={String.class}, tree="[0]")
    private Output<String> realmAttributes;

    /**
     * @return The realm attributes for the group.
     * 
     */
    public Output<String> realmAttributes() {
        return this.realmAttributes;
    }
    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     * 
     */
    @Export(name="reportsManager", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reportsManager;

    /**
     * @return When this override is set, User in the group can manage Xray Reports on any resource type. Default value is `false`.
     * 
     */
    public Output<Boolean> reportsManager() {
        return this.reportsManager;
    }
    @Export(name="usersNames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> usersNames;

    public Output<List<String>> usersNames() {
        return this.usersNames;
    }
    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     * 
     */
    @Export(name="watchManager", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> watchManager;

    /**
     * @return When this override is set, User in the group can manage Xray Watches on any resource type. Default value is `false`.
     * 
     */
    public Output<Boolean> watchManager() {
        return this.watchManager;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Group(java.lang.String name) {
        this(name, GroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Group(java.lang.String name, @Nullable GroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Group(java.lang.String name, @Nullable GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/group:Group", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Group(java.lang.String name, Output<java.lang.String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/group:Group", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupArgs makeArgs(@Nullable GroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Group get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Group(name, id, state, options);
    }
}
