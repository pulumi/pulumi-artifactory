// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFileListArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFileListArgs Empty = new GetFileListArgs();

    /**
     * Get deep listing
     * 
     */
    @Import(name="deepListing")
    private @Nullable Output<Boolean> deepListing;

    /**
     * @return Get deep listing
     * 
     */
    public Optional<Output<Boolean>> deepListing() {
        return Optional.ofNullable(this.deepListing);
    }

    /**
     * Depth of the deep listing
     * 
     */
    @Import(name="depth")
    private @Nullable Output<Integer> depth;

    /**
     * @return Depth of the deep listing
     * 
     */
    public Optional<Output<Integer>> depth() {
        return Optional.ofNullable(this.depth);
    }

    /**
     * Path of the folder
     * 
     */
    @Import(name="folderPath", required=true)
    private Output<String> folderPath;

    /**
     * @return Path of the folder
     * 
     */
    public Output<String> folderPath() {
        return this.folderPath;
    }

    /**
     * Include root path
     * 
     */
    @Import(name="includeRootPath")
    private @Nullable Output<Boolean> includeRootPath;

    /**
     * @return Include root path
     * 
     */
    public Optional<Output<Boolean>> includeRootPath() {
        return Optional.ofNullable(this.includeRootPath);
    }

    /**
     * Include folders
     * 
     */
    @Import(name="listFolders")
    private @Nullable Output<Boolean> listFolders;

    /**
     * @return Include folders
     * 
     */
    public Optional<Output<Boolean>> listFolders() {
        return Optional.ofNullable(this.listFolders);
    }

    /**
     * Include metadata timestamps
     * 
     */
    @Import(name="metadataTimestamps")
    private @Nullable Output<Boolean> metadataTimestamps;

    /**
     * @return Include metadata timestamps
     * 
     */
    public Optional<Output<Boolean>> metadataTimestamps() {
        return Optional.ofNullable(this.metadataTimestamps);
    }

    /**
     * Repository key
     * 
     */
    @Import(name="repositoryKey", required=true)
    private Output<String> repositoryKey;

    /**
     * @return Repository key
     * 
     */
    public Output<String> repositoryKey() {
        return this.repositoryKey;
    }

    private GetFileListArgs() {}

    private GetFileListArgs(GetFileListArgs $) {
        this.deepListing = $.deepListing;
        this.depth = $.depth;
        this.folderPath = $.folderPath;
        this.includeRootPath = $.includeRootPath;
        this.listFolders = $.listFolders;
        this.metadataTimestamps = $.metadataTimestamps;
        this.repositoryKey = $.repositoryKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFileListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFileListArgs $;

        public Builder() {
            $ = new GetFileListArgs();
        }

        public Builder(GetFileListArgs defaults) {
            $ = new GetFileListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deepListing Get deep listing
         * 
         * @return builder
         * 
         */
        public Builder deepListing(@Nullable Output<Boolean> deepListing) {
            $.deepListing = deepListing;
            return this;
        }

        /**
         * @param deepListing Get deep listing
         * 
         * @return builder
         * 
         */
        public Builder deepListing(Boolean deepListing) {
            return deepListing(Output.of(deepListing));
        }

        /**
         * @param depth Depth of the deep listing
         * 
         * @return builder
         * 
         */
        public Builder depth(@Nullable Output<Integer> depth) {
            $.depth = depth;
            return this;
        }

        /**
         * @param depth Depth of the deep listing
         * 
         * @return builder
         * 
         */
        public Builder depth(Integer depth) {
            return depth(Output.of(depth));
        }

        /**
         * @param folderPath Path of the folder
         * 
         * @return builder
         * 
         */
        public Builder folderPath(Output<String> folderPath) {
            $.folderPath = folderPath;
            return this;
        }

        /**
         * @param folderPath Path of the folder
         * 
         * @return builder
         * 
         */
        public Builder folderPath(String folderPath) {
            return folderPath(Output.of(folderPath));
        }

        /**
         * @param includeRootPath Include root path
         * 
         * @return builder
         * 
         */
        public Builder includeRootPath(@Nullable Output<Boolean> includeRootPath) {
            $.includeRootPath = includeRootPath;
            return this;
        }

        /**
         * @param includeRootPath Include root path
         * 
         * @return builder
         * 
         */
        public Builder includeRootPath(Boolean includeRootPath) {
            return includeRootPath(Output.of(includeRootPath));
        }

        /**
         * @param listFolders Include folders
         * 
         * @return builder
         * 
         */
        public Builder listFolders(@Nullable Output<Boolean> listFolders) {
            $.listFolders = listFolders;
            return this;
        }

        /**
         * @param listFolders Include folders
         * 
         * @return builder
         * 
         */
        public Builder listFolders(Boolean listFolders) {
            return listFolders(Output.of(listFolders));
        }

        /**
         * @param metadataTimestamps Include metadata timestamps
         * 
         * @return builder
         * 
         */
        public Builder metadataTimestamps(@Nullable Output<Boolean> metadataTimestamps) {
            $.metadataTimestamps = metadataTimestamps;
            return this;
        }

        /**
         * @param metadataTimestamps Include metadata timestamps
         * 
         * @return builder
         * 
         */
        public Builder metadataTimestamps(Boolean metadataTimestamps) {
            return metadataTimestamps(Output.of(metadataTimestamps));
        }

        /**
         * @param repositoryKey Repository key
         * 
         * @return builder
         * 
         */
        public Builder repositoryKey(Output<String> repositoryKey) {
            $.repositoryKey = repositoryKey;
            return this;
        }

        /**
         * @param repositoryKey Repository key
         * 
         * @return builder
         * 
         */
        public Builder repositoryKey(String repositoryKey) {
            return repositoryKey(Output.of(repositoryKey));
        }

        public GetFileListArgs build() {
            if ($.folderPath == null) {
                throw new MissingRequiredPropertyException("GetFileListArgs", "folderPath");
            }
            if ($.repositoryKey == null) {
                throw new MissingRequiredPropertyException("GetFileListArgs", "repositoryKey");
            }
            return $;
        }
    }

}
