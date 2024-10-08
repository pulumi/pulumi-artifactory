// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.PermissionTargetArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.PermissionTargetState;
import com.pulumi.artifactory.outputs.PermissionTargetBuild;
import com.pulumi.artifactory.outputs.PermissionTargetReleaseBundle;
import com.pulumi.artifactory.outputs.PermissionTargetRepo;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.
 * 
 * !&gt;This resource has been deprecated in favor of platform_permission resource.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.PermissionTarget;
 * import com.pulumi.artifactory.PermissionTargetArgs;
 * import com.pulumi.artifactory.inputs.PermissionTargetRepoArgs;
 * import com.pulumi.artifactory.inputs.PermissionTargetRepoActionsArgs;
 * import com.pulumi.artifactory.inputs.PermissionTargetBuildArgs;
 * import com.pulumi.artifactory.inputs.PermissionTargetBuildActionsArgs;
 * import com.pulumi.artifactory.inputs.PermissionTargetReleaseBundleArgs;
 * import com.pulumi.artifactory.inputs.PermissionTargetReleaseBundleActionsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         // Create a new Artifactory permission target called testpermission
 *         var test_perm = new PermissionTarget("test-perm", PermissionTargetArgs.builder()
 *             .name("test-perm")
 *             .repo(PermissionTargetRepoArgs.builder()
 *                 .includesPatterns("foo/**")
 *                 .excludesPatterns("bar/**")
 *                 .repositories("example-repo-local")
 *                 .actions(PermissionTargetRepoActionsArgs.builder()
 *                     .users(                    
 *                         PermissionTargetRepoActionsUserArgs.builder()
 *                             .name("anonymous")
 *                             .permissions(                            
 *                                 "read",
 *                                 "write")
 *                             .build(),
 *                         PermissionTargetRepoActionsUserArgs.builder()
 *                             .name("user1")
 *                             .permissions(                            
 *                                 "read",
 *                                 "write")
 *                             .build())
 *                     .groups(                    
 *                         PermissionTargetRepoActionsGroupArgs.builder()
 *                             .name("readers")
 *                             .permissions("read")
 *                             .build(),
 *                         PermissionTargetRepoActionsGroupArgs.builder()
 *                             .name("dev")
 *                             .permissions(                            
 *                                 "read",
 *                                 "write")
 *                             .build())
 *                     .build())
 *                 .build())
 *             .build(PermissionTargetBuildArgs.builder()
 *                 .includesPatterns("**")
 *                 .repositories("artifactory-build-info")
 *                 .actions(PermissionTargetBuildActionsArgs.builder()
 *                     .users(                    
 *                         PermissionTargetBuildActionsUserArgs.builder()
 *                             .name("anonymous")
 *                             .permissions("read")
 *                             .build(),
 *                         PermissionTargetBuildActionsUserArgs.builder()
 *                             .name("user1")
 *                             .permissions(                            
 *                                 "read",
 *                                 "write")
 *                             .build())
 *                     .build())
 *                 .build())
 *             .releaseBundle(PermissionTargetReleaseBundleArgs.builder()
 *                 .includesPatterns("**")
 *                 .repositories("release-bundles")
 *                 .actions(PermissionTargetReleaseBundleActionsArgs.builder()
 *                     .users(PermissionTargetReleaseBundleActionsUserArgs.builder()
 *                         .name("anonymous")
 *                         .permissions("read")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Permissions
 * 
 * The provider supports the following `permission` enums:
 * 
 * * `read`
 * * `write`
 * * `annotate`
 * * `delete`
 * * `manage`
 * * `managedXrayMeta`
 * * `distribute`
 * 
 * The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):
 * 
 * * `read` - matches `Read` permissions.
 * * `write` - matches `  Deploy / Cache / Create ` permissions.
 * * `annotate` - matches `Annotate` permissions.
 * * `delete` - matches `Delete / Overwrite` permissions.
 * * `manage` - matches `Manage` permissions.
 * * `managedXrayMeta` - matches `Manage Xray Metadata` permissions.
 * * `distribute` - matches `Distribute` permissions.
 * 
 * ## Import
 * 
 * Permission targets can be imported using their name, e.g.
 * 
 * ```sh
 * $ pulumi import artifactory:index/permissionTarget:PermissionTarget terraform-test-permission mypermission
 * ```
 * 
 */
@ResourceType(type="artifactory:index/permissionTarget:PermissionTarget")
public class PermissionTarget extends com.pulumi.resources.CustomResource {
    /**
     * As for repo but for artifactory-build-info permissions.
     * 
     */
    @Export(name="build", refs={PermissionTargetBuild.class}, tree="[0]")
    private Output</* @Nullable */ PermissionTargetBuild> build;

    /**
     * @return As for repo but for artifactory-build-info permissions.
     * 
     */
    public Output<Optional<PermissionTargetBuild>> build() {
        return Codegen.optional(this.build);
    }
    /**
     * Name of permission.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of permission.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * As for repo for for release-bundles permissions.
     * 
     */
    @Export(name="releaseBundle", refs={PermissionTargetReleaseBundle.class}, tree="[0]")
    private Output</* @Nullable */ PermissionTargetReleaseBundle> releaseBundle;

    /**
     * @return As for repo for for release-bundles permissions.
     * 
     */
    public Output<Optional<PermissionTargetReleaseBundle>> releaseBundle() {
        return Codegen.optional(this.releaseBundle);
    }
    /**
     * Repository permission configuration.
     * 
     */
    @Export(name="repo", refs={PermissionTargetRepo.class}, tree="[0]")
    private Output</* @Nullable */ PermissionTargetRepo> repo;

    /**
     * @return Repository permission configuration.
     * 
     */
    public Output<Optional<PermissionTargetRepo>> repo() {
        return Codegen.optional(this.repo);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PermissionTarget(java.lang.String name) {
        this(name, PermissionTargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PermissionTarget(java.lang.String name, @Nullable PermissionTargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PermissionTarget(java.lang.String name, @Nullable PermissionTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/permissionTarget:PermissionTarget", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PermissionTarget(java.lang.String name, Output<java.lang.String> id, @Nullable PermissionTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/permissionTarget:PermissionTarget", name, state, makeResourceOptions(options, id), false);
    }

    private static PermissionTargetArgs makeArgs(@Nullable PermissionTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PermissionTargetArgs.Empty : args;
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
    public static PermissionTarget get(java.lang.String name, Output<java.lang.String> id, @Nullable PermissionTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PermissionTarget(name, id, state, options);
    }
}
