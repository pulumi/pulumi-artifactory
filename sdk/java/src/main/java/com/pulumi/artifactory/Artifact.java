// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ArtifactArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.ArtifactState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource for deploying artifact to Artifactory repository. Support deploying a single artifact only. Changes to `repository` or `path` attributes will trigger a recreation of the resource (i.e. delete then create). See [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/deploy-a-single-artifact) for more details.
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
 * import com.pulumi.artifactory.Artifact;
 * import com.pulumi.artifactory.ArtifactArgs;
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
 *         var my_artifact = new Artifact("my-artifact", ArtifactArgs.builder()
 *             .repository("my-generic-local")
 *             .path("/my-path/my-file.zip")
 *             .filePath("/path/to/my-file.zip")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="artifactory:index/artifact:Artifact")
public class Artifact extends com.pulumi.resources.CustomResource {
    /**
     * MD5 checksum of the artifact.
     * 
     */
    @Export(name="checksumMd5", refs={String.class}, tree="[0]")
    private Output<String> checksumMd5;

    /**
     * @return MD5 checksum of the artifact.
     * 
     */
    public Output<String> checksumMd5() {
        return this.checksumMd5;
    }
    /**
     * SHA1 checksum of the artifact.
     * 
     */
    @Export(name="checksumSha1", refs={String.class}, tree="[0]")
    private Output<String> checksumSha1;

    /**
     * @return SHA1 checksum of the artifact.
     * 
     */
    public Output<String> checksumSha1() {
        return this.checksumSha1;
    }
    /**
     * SHA256 checksum of the artifact.
     * 
     */
    @Export(name="checksumSha256", refs={String.class}, tree="[0]")
    private Output<String> checksumSha256;

    /**
     * @return SHA256 checksum of the artifact.
     * 
     */
    public Output<String> checksumSha256() {
        return this.checksumSha256;
    }
    /**
     * Timestamp when artifact is created.
     * 
     */
    @Export(name="created", refs={String.class}, tree="[0]")
    private Output<String> created;

    /**
     * @return Timestamp when artifact is created.
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * User who deploys the artifact.
     * 
     */
    @Export(name="createdBy", refs={String.class}, tree="[0]")
    private Output<String> createdBy;

    /**
     * @return User who deploys the artifact.
     * 
     */
    public Output<String> createdBy() {
        return this.createdBy;
    }
    /**
     * Download URI of the artifact.
     * 
     */
    @Export(name="downloadUri", refs={String.class}, tree="[0]")
    private Output<String> downloadUri;

    /**
     * @return Download URI of the artifact.
     * 
     */
    public Output<String> downloadUri() {
        return this.downloadUri;
    }
    /**
     * Path to the source file.
     * 
     */
    @Export(name="filePath", refs={String.class}, tree="[0]")
    private Output<String> filePath;

    /**
     * @return Path to the source file.
     * 
     */
    public Output<String> filePath() {
        return this.filePath;
    }
    /**
     * MIME type of the artifact.
     * 
     */
    @Export(name="mimeType", refs={String.class}, tree="[0]")
    private Output<String> mimeType;

    /**
     * @return MIME type of the artifact.
     * 
     */
    public Output<String> mimeType() {
        return this.mimeType;
    }
    /**
     * The relative path in the target repository. Must begin with a &#39;/&#39;. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return The relative path in the target repository. Must begin with a &#39;/&#39;. You can add key-value matrix parameters to deploy the artifacts with properties. For more details, please refer to [Introducing Matrix Parameters](https://jfrog.com/help/r/jfrog-artifactory-documentation/using-properties-in-deployment-and-resolution).
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Name of the respository.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return Name of the respository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * Size of the artifact, in bytes.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return Size of the artifact, in bytes.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * URI of the artifact.
     * 
     */
    @Export(name="uri", refs={String.class}, tree="[0]")
    private Output<String> uri;

    /**
     * @return URI of the artifact.
     * 
     */
    public Output<String> uri() {
        return this.uri;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Artifact(java.lang.String name) {
        this(name, ArtifactArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Artifact(java.lang.String name, ArtifactArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Artifact(java.lang.String name, ArtifactArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifact:Artifact", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Artifact(java.lang.String name, Output<java.lang.String> id, @Nullable ArtifactState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/artifact:Artifact", name, state, makeResourceOptions(options, id), false);
    }

    private static ArtifactArgs makeArgs(ArtifactArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ArtifactArgs.Empty : args;
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
    public static Artifact get(java.lang.String name, Output<java.lang.String> id, @Nullable ArtifactState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Artifact(name, id, state, options);
    }
}
