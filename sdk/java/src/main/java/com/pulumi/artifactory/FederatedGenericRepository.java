// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.FederatedGenericRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.FederatedGenericRepositoryState;
import com.pulumi.artifactory.outputs.FederatedGenericRepositoryMember;
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
 * Creates a federated Generic repository.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.FederatedGenericRepository;
 * import com.pulumi.artifactory.FederatedGenericRepositoryArgs;
 * import com.pulumi.artifactory.inputs.FederatedGenericRepositoryMemberArgs;
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
 *         var terraform_federated_test_generic_repo = new FederatedGenericRepository(&#34;terraform-federated-test-generic-repo&#34;, FederatedGenericRepositoryArgs.builder()        
 *             .key(&#34;terraform-federated-test-generic-repo&#34;)
 *             .members(            
 *                 FederatedGenericRepositoryMemberArgs.builder()
 *                     .enabled(true)
 *                     .url(&#34;http://tempurl.org/artifactory/terraform-federated-test-generic-repo&#34;)
 *                     .build(),
 *                 FederatedGenericRepositoryMemberArgs.builder()
 *                     .enabled(true)
 *                     .url(&#34;http://tempurl2.org/artifactory/terraform-federated-test-generic-repo-2&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Federated repositories can be imported using their name, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/federatedGenericRepository:FederatedGenericRepository terraform-federated-test-generic-repo terraform-federated-test-generic-repo
 * ```
 * 
 */
@ResourceType(type="artifactory:index/federatedGenericRepository:FederatedGenericRepository")
public class FederatedGenericRepository extends com.pulumi.resources.CustomResource {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    @Export(name="archiveBrowsingEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> archiveBrowsingEnabled;

    /**
     * @return When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    public Output<Optional<Boolean>> archiveBrowsingEnabled() {
        return Codegen.optional(this.archiveBrowsingEnabled);
    }
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    @Export(name="blackedOut", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> blackedOut;

    /**
     * @return When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    public Output<Optional<Boolean>> blackedOut() {
        return Codegen.optional(this.blackedOut);
    }
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    @Export(name="downloadDirect", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    public Output<Optional<Boolean>> downloadDirect() {
        return Codegen.optional(this.downloadDirect);
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", type=String.class, parameters={})
    private Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    public Output<String> excludesPattern() {
        return this.excludesPattern;
    }
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", type=String.class, parameters={})
    private Output<String> includesPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<String> includesPattern() {
        return this.includesPattern;
    }
    /**
     * the identity key of the repo.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    @Export(name="members", type=List.class, parameters={FederatedGenericRepositoryMember.class})
    private Output<List<FederatedGenericRepositoryMember>> members;

    /**
     * @return The list of Federated members and must contain this repository URL (configured base URL
     * `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
     * Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
     * to set up Federated repositories correctly.
     * 
     */
    public Output<List<FederatedGenericRepositoryMember>> members() {
        return this.members;
    }
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    @Export(name="packageType", type=String.class, parameters={})
    private Output<String> packageType;

    public Output<String> packageType() {
        return this.packageType;
    }
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Export(name="priorityResolution", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Output<Optional<Boolean>> priorityResolution() {
        return Codegen.optional(this.priorityResolution);
    }
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    @Export(name="projectEnvironments", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    public Output<Optional<List<String>>> projectEnvironments() {
        return Codegen.optional(this.projectEnvironments);
    }
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * List of property set name
     * 
     */
    @Export(name="propertySets", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> propertySets;

    /**
     * @return List of property set name
     * 
     */
    public Output<Optional<List<String>>> propertySets() {
        return Codegen.optional(this.propertySets);
    }
    /**
     * Repository layout key for the local repository
     * 
     */
    @Export(name="repoLayoutRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Export(name="xrayIndex", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Output<Optional<Boolean>> xrayIndex() {
        return Codegen.optional(this.xrayIndex);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FederatedGenericRepository(String name) {
        this(name, FederatedGenericRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FederatedGenericRepository(String name, FederatedGenericRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FederatedGenericRepository(String name, FederatedGenericRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/federatedGenericRepository:FederatedGenericRepository", name, args == null ? FederatedGenericRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FederatedGenericRepository(String name, Output<String> id, @Nullable FederatedGenericRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/federatedGenericRepository:FederatedGenericRepository", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static FederatedGenericRepository get(String name, Output<String> id, @Nullable FederatedGenericRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FederatedGenericRepository(name, id, state, options);
    }
}