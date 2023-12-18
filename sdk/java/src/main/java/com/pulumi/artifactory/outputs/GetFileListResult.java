// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetFileListFile;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFileListResult {
    /**
     * @return Creation time
     * 
     */
    private String created;
    /**
     * @return Get deep listing
     * 
     */
    private @Nullable Boolean deepListing;
    /**
     * @return Depth of the deep listing
     * 
     */
    private @Nullable Integer depth;
    /**
     * @return A list of files.
     * 
     */
    private List<GetFileListFile> files;
    /**
     * @return Path of the folder
     * 
     */
    private String folderPath;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Include root path
     * 
     */
    private @Nullable Boolean includeRootPath;
    /**
     * @return Include folders
     * 
     */
    private @Nullable Boolean listFolders;
    /**
     * @return Include metadata timestamps
     * 
     */
    private @Nullable Boolean metadataTimestamps;
    /**
     * @return Repository key
     * 
     */
    private String repositoryKey;
    /**
     * @return URL to file/path
     * 
     */
    private String uri;

    private GetFileListResult() {}
    /**
     * @return Creation time
     * 
     */
    public String created() {
        return this.created;
    }
    /**
     * @return Get deep listing
     * 
     */
    public Optional<Boolean> deepListing() {
        return Optional.ofNullable(this.deepListing);
    }
    /**
     * @return Depth of the deep listing
     * 
     */
    public Optional<Integer> depth() {
        return Optional.ofNullable(this.depth);
    }
    /**
     * @return A list of files.
     * 
     */
    public List<GetFileListFile> files() {
        return this.files;
    }
    /**
     * @return Path of the folder
     * 
     */
    public String folderPath() {
        return this.folderPath;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Include root path
     * 
     */
    public Optional<Boolean> includeRootPath() {
        return Optional.ofNullable(this.includeRootPath);
    }
    /**
     * @return Include folders
     * 
     */
    public Optional<Boolean> listFolders() {
        return Optional.ofNullable(this.listFolders);
    }
    /**
     * @return Include metadata timestamps
     * 
     */
    public Optional<Boolean> metadataTimestamps() {
        return Optional.ofNullable(this.metadataTimestamps);
    }
    /**
     * @return Repository key
     * 
     */
    public String repositoryKey() {
        return this.repositoryKey;
    }
    /**
     * @return URL to file/path
     * 
     */
    public String uri() {
        return this.uri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFileListResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String created;
        private @Nullable Boolean deepListing;
        private @Nullable Integer depth;
        private List<GetFileListFile> files;
        private String folderPath;
        private String id;
        private @Nullable Boolean includeRootPath;
        private @Nullable Boolean listFolders;
        private @Nullable Boolean metadataTimestamps;
        private String repositoryKey;
        private String uri;
        public Builder() {}
        public Builder(GetFileListResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.created = defaults.created;
    	      this.deepListing = defaults.deepListing;
    	      this.depth = defaults.depth;
    	      this.files = defaults.files;
    	      this.folderPath = defaults.folderPath;
    	      this.id = defaults.id;
    	      this.includeRootPath = defaults.includeRootPath;
    	      this.listFolders = defaults.listFolders;
    	      this.metadataTimestamps = defaults.metadataTimestamps;
    	      this.repositoryKey = defaults.repositoryKey;
    	      this.uri = defaults.uri;
        }

        @CustomType.Setter
        public Builder created(String created) {
            this.created = Objects.requireNonNull(created);
            return this;
        }
        @CustomType.Setter
        public Builder deepListing(@Nullable Boolean deepListing) {
            this.deepListing = deepListing;
            return this;
        }
        @CustomType.Setter
        public Builder depth(@Nullable Integer depth) {
            this.depth = depth;
            return this;
        }
        @CustomType.Setter
        public Builder files(List<GetFileListFile> files) {
            this.files = Objects.requireNonNull(files);
            return this;
        }
        public Builder files(GetFileListFile... files) {
            return files(List.of(files));
        }
        @CustomType.Setter
        public Builder folderPath(String folderPath) {
            this.folderPath = Objects.requireNonNull(folderPath);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder includeRootPath(@Nullable Boolean includeRootPath) {
            this.includeRootPath = includeRootPath;
            return this;
        }
        @CustomType.Setter
        public Builder listFolders(@Nullable Boolean listFolders) {
            this.listFolders = listFolders;
            return this;
        }
        @CustomType.Setter
        public Builder metadataTimestamps(@Nullable Boolean metadataTimestamps) {
            this.metadataTimestamps = metadataTimestamps;
            return this;
        }
        @CustomType.Setter
        public Builder repositoryKey(String repositoryKey) {
            this.repositoryKey = Objects.requireNonNull(repositoryKey);
            return this;
        }
        @CustomType.Setter
        public Builder uri(String uri) {
            this.uri = Objects.requireNonNull(uri);
            return this;
        }
        public GetFileListResult build() {
            final var _resultValue = new GetFileListResult();
            _resultValue.created = created;
            _resultValue.deepListing = deepListing;
            _resultValue.depth = depth;
            _resultValue.files = files;
            _resultValue.folderPath = folderPath;
            _resultValue.id = id;
            _resultValue.includeRootPath = includeRootPath;
            _resultValue.listFolders = listFolders;
            _resultValue.metadataTimestamps = metadataTimestamps;
            _resultValue.repositoryKey = repositoryKey;
            _resultValue.uri = uri;
            return _resultValue;
        }
    }
}