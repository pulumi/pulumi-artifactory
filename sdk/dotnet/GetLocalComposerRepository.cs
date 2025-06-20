// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetLocalComposerRepository
    {
        /// <summary>
        /// Retrieves a local composer repository.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var local_test_composer_repo = Artifactory.GetLocalComposerRepository.Invoke(new()
        ///     {
        ///         Key = "local-test-composer-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLocalComposerRepositoryResult> InvokeAsync(GetLocalComposerRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocalComposerRepositoryResult>("artifactory:index/getLocalComposerRepository:getLocalComposerRepository", args ?? new GetLocalComposerRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a local composer repository.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var local_test_composer_repo = Artifactory.GetLocalComposerRepository.Invoke(new()
        ///     {
        ///         Key = "local-test-composer-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLocalComposerRepositoryResult> Invoke(GetLocalComposerRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalComposerRepositoryResult>("artifactory:index/getLocalComposerRepository:getLocalComposerRepository", args ?? new GetLocalComposerRepositoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a local composer repository.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var local_test_composer_repo = Artifactory.GetLocalComposerRepository.Invoke(new()
        ///     {
        ///         Key = "local-test-composer-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLocalComposerRepositoryResult> Invoke(GetLocalComposerRepositoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocalComposerRepositoryResult>("artifactory:index/getLocalComposerRepository:getLocalComposerRepository", args ?? new GetLocalComposerRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocalComposerRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("downloadDirect")]
        public bool? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public string? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public string? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("notes")]
        public string? Notes { get; set; }

        [Input("priorityResolution")]
        public bool? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private List<string>? _projectEnvironments;
        public List<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new List<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectKey")]
        public string? ProjectKey { get; set; }

        [Input("propertySets")]
        private List<string>? _propertySets;
        public List<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new List<string>());
            set => _propertySets = value;
        }

        [Input("repoLayoutRef")]
        public string? RepoLayoutRef { get; set; }

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetLocalComposerRepositoryArgs()
        {
        }
        public static new GetLocalComposerRepositoryArgs Empty => new GetLocalComposerRepositoryArgs();
    }

    public sealed class GetLocalComposerRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("downloadDirect")]
        public Input<bool>? DownloadDirect { get; set; }

        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("priorityResolution")]
        public Input<bool>? PriorityResolution { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("propertySets")]
        private InputList<string>? _propertySets;
        public InputList<string> PropertySets
        {
            get => _propertySets ?? (_propertySets = new InputList<string>());
            set => _propertySets = value;
        }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetLocalComposerRepositoryInvokeArgs()
        {
        }
        public static new GetLocalComposerRepositoryInvokeArgs Empty => new GetLocalComposerRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocalComposerRepositoryResult
    {
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly bool? BlackedOut;
        public readonly bool? CdnRedirect;
        public readonly string? Description;
        public readonly bool? DownloadDirect;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        /// <summary>
        /// the identity key of the repo.
        /// </summary>
        public readonly string Key;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        public readonly string? RepoLayoutRef;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetLocalComposerRepositoryResult(
            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool? cdnRedirect,

            string? description,

            bool? downloadDirect,

            string? excludesPattern,

            string id,

            string? includesPattern,

            string key,

            string? notes,

            string packageType,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? repoLayoutRef,

            bool? xrayIndex)
        {
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            CdnRedirect = cdnRedirect;
            Description = description;
            DownloadDirect = downloadDirect;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            Notes = notes;
            PackageType = packageType;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            RepoLayoutRef = repoLayoutRef;
            XrayIndex = xrayIndex;
        }
    }
}
