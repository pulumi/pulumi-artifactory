// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Creates a local RPM repository.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// using Std = Pulumi.Std;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var some_keypair_gpg_1 = new Artifactory.Keypair("some-keypair-gpg-1", new()
    ///     {
    ///         PairName = $"some-keypair{randid.Id}",
    ///         PairType = "GPG",
    ///         Alias = "foo-alias1",
    ///         PrivateKey = Std.File.Invoke(new()
    ///         {
    ///             Input = "samples/gpg.priv",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         PublicKey = Std.File.Invoke(new()
    ///         {
    ///             Input = "samples/gpg.pub",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///     });
    /// 
    ///     var some_keypair_gpg_2 = new Artifactory.Keypair("some-keypair-gpg-2", new()
    ///     {
    ///         PairName = $"some-keypair{randid.Id}",
    ///         PairType = "GPG",
    ///         Alias = "foo-alias2",
    ///         PrivateKey = Std.File.Invoke(new()
    ///         {
    ///             Input = "samples/gpg.priv",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         PublicKey = Std.File.Invoke(new()
    ///         {
    ///             Input = "samples/gpg.pub",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///     });
    /// 
    ///     var terraform_local_test_rpm_repo_basic = new Artifactory.LocalRpmRepository("terraform-local-test-rpm-repo-basic", new()
    ///     {
    ///         Key = "terraform-local-test-rpm-repo-basic",
    ///         YumRootDepth = 5,
    ///         CalculateYumMetadata = true,
    ///         EnableFileListsIndexing = true,
    ///         YumGroupFileNames = "file-1.xml,file-2.xml",
    ///         PrimaryKeypairRef = some_keypairGPG1.PairName,
    ///         SecondaryKeypairRef = some_keypairGPG2.PairName,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             some_keypair_gpg_1,
    ///             some_keypair_gpg_2,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Local repositories can be imported using their name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/localRpmRepository:LocalRpmRepository terraform-local-test-rpm-repo-basic terraform-local-test-rpm-repo-basic
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/localRpmRepository:LocalRpmRepository")]
    public partial class LocalRpmRepository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Output("archiveBrowsingEnabled")]
        public Output<bool> ArchiveBrowsingEnabled { get; private set; } = null!;

        /// <summary>
        /// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
        /// </summary>
        [Output("blackedOut")]
        public Output<bool> BlackedOut { get; private set; } = null!;

        /// <summary>
        /// Default: `false`.
        /// </summary>
        [Output("calculateYumMetadata")]
        public Output<bool> CalculateYumMetadata { get; private set; } = null!;

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Output("cdnRedirect")]
        public Output<bool> CdnRedirect { get; private set; } = null!;

        /// <summary>
        /// Public description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
        /// storage provider. Available in Enterprise+ and Edge licenses only.
        /// </summary>
        [Output("downloadDirect")]
        public Output<bool> DownloadDirect { get; private set; } = null!;

        /// <summary>
        /// Default: `false`.
        /// </summary>
        [Output("enableFileListsIndexing")]
        public Output<bool> EnableFileListsIndexing { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
        /// artifacts are excluded.
        /// </summary>
        [Output("excludesPattern")]
        public Output<string> ExcludesPattern { get; private set; } = null!;

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        /// </summary>
        [Output("includesPattern")]
        public Output<string> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Internal description.
        /// </summary>
        [Output("notes")]
        public Output<string> Notes { get; private set; } = null!;

        /// <summary>
        /// The primary GPG key to be used to sign packages.
        /// </summary>
        [Output("primaryKeypairRef")]
        public Output<string> PrimaryKeypairRef { get; private set; } = null!;

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Output("priorityResolution")]
        public Output<bool> PriorityResolution { get; private set; } = null!;

        [Output("projectEnvironments")]
        public Output<ImmutableArray<string>> ProjectEnvironments { get; private set; } = null!;

        /// <summary>
        /// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Output("projectKey")]
        public Output<string> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// List of property set name
        /// </summary>
        [Output("propertySets")]
        public Output<ImmutableArray<string>> PropertySets { get; private set; } = null!;

        /// <summary>
        /// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
        /// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
        /// </summary>
        [Output("repoLayoutRef")]
        public Output<string> RepoLayoutRef { get; private set; } = null!;

        /// <summary>
        /// The secondary GPG key to be used to sign packages.
        /// </summary>
        [Output("secondaryKeypairRef")]
        public Output<string> SecondaryKeypairRef { get; private set; } = null!;

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Output("xrayIndex")]
        public Output<bool> XrayIndex { get; private set; } = null!;

        /// <summary>
        /// A comma separated list of XML file names containing RPM group component definitions. 
        /// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
        /// generating a gzipped version of the group files, if required. Default is empty string.
        /// </summary>
        [Output("yumGroupFileNames")]
        public Output<string> YumGroupFileNames { get; private set; } = null!;

        /// <summary>
        /// The depth, relative to the repository's root folder, where RPM metadata is created. 
        /// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
        /// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
        /// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
        /// snapshots are not cleaned up.
        /// </summary>
        [Output("yumRootDepth")]
        public Output<int> YumRootDepth { get; private set; } = null!;


        /// <summary>
        /// Create a LocalRpmRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocalRpmRepository(string name, LocalRpmRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/localRpmRepository:LocalRpmRepository", name, args ?? new LocalRpmRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocalRpmRepository(string name, Input<string> id, LocalRpmRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/localRpmRepository:LocalRpmRepository", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LocalRpmRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocalRpmRepository Get(string name, Input<string> id, LocalRpmRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new LocalRpmRepository(name, id, state, options);
        }
    }

    public sealed class LocalRpmRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        /// <summary>
        /// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
        /// </summary>
        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        /// <summary>
        /// Default: `false`.
        /// </summary>
        [Input("calculateYumMetadata")]
        public Input<bool>? CalculateYumMetadata { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        /// <summary>
        /// Public description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
        /// storage provider. Available in Enterprise+ and Edge licenses only.
        /// </summary>
        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        /// <summary>
        /// Default: `false`.
        /// </summary>
        [Input("enableFileListsIndexing")]
        public Input<bool>? EnableFileListsIndexing { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The primary GPG key to be used to sign packages.
        /// </summary>
        [Input("primaryKeypairRef")]
        public Input<string>? PrimaryKeypairRef { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;

        /// <summary>
        /// List of property set name
        /// </summary>
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
        /// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// The secondary GPG key to be used to sign packages.
        /// </summary>
        [Input("secondaryKeypairRef")]
        public Input<string>? SecondaryKeypairRef { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        /// <summary>
        /// A comma separated list of XML file names containing RPM group component definitions. 
        /// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
        /// generating a gzipped version of the group files, if required. Default is empty string.
        /// </summary>
        [Input("yumGroupFileNames")]
        public Input<string>? YumGroupFileNames { get; set; }

        /// <summary>
        /// The depth, relative to the repository's root folder, where RPM metadata is created. 
        /// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
        /// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
        /// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
        /// snapshots are not cleaned up.
        /// </summary>
        [Input("yumRootDepth")]
        public Input<int>? YumRootDepth { get; set; }

        public LocalRpmRepositoryArgs()
        {
        }
        public static new LocalRpmRepositoryArgs Empty => new LocalRpmRepositoryArgs();
    }

    public sealed class LocalRpmRepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        /// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        /// security (e.g., cross-site scripting attacks).
        /// </summary>
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        /// <summary>
        /// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
        /// </summary>
        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        /// <summary>
        /// Default: `false`.
        /// </summary>
        [Input("calculateYumMetadata")]
        public Input<bool>? CalculateYumMetadata { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
        /// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
        /// </summary>
        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        /// <summary>
        /// Public description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
        /// storage provider. Available in Enterprise+ and Edge licenses only.
        /// </summary>
        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        /// <summary>
        /// Default: `false`.
        /// </summary>
        [Input("enableFileListsIndexing")]
        public Input<bool>? EnableFileListsIndexing { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
        /// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Internal description.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The primary GPG key to be used to sign packages.
        /// </summary>
        [Input("primaryKeypairRef")]
        public Input<string>? PrimaryKeypairRef { get; set; }

        /// <summary>
        /// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
        /// </summary>
        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
        /// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;

        /// <summary>
        /// List of property set name
        /// </summary>
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        /// <summary>
        /// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
        /// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        /// <summary>
        /// The secondary GPG key to be used to sign packages.
        /// </summary>
        [Input("secondaryKeypairRef")]
        public Input<string>? SecondaryKeypairRef { get; set; }

        /// <summary>
        /// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
        /// Xray settings.
        /// </summary>
        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        /// <summary>
        /// A comma separated list of XML file names containing RPM group component definitions. 
        /// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
        /// generating a gzipped version of the group files, if required. Default is empty string.
        /// </summary>
        [Input("yumGroupFileNames")]
        public Input<string>? YumGroupFileNames { get; set; }

        /// <summary>
        /// The depth, relative to the repository's root folder, where RPM metadata is created. 
        /// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
        /// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
        /// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
        /// snapshots are not cleaned up.
        /// </summary>
        [Input("yumRootDepth")]
        public Input<int>? YumRootDepth { get; set; }

        public LocalRpmRepositoryState()
        {
        }
        public static new LocalRpmRepositoryState Empty => new LocalRpmRepositoryState();
    }
}
