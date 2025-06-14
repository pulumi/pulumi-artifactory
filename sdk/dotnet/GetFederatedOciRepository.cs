// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFederatedOciRepository
    {
        /// <summary>
        /// Retrieves a federated OCI repository.
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
        ///     var my_oci_federated = Artifactory.GetFederatedOciRepository.Invoke(new()
        ///     {
        ///         Key = "my-oci-federated",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFederatedOciRepositoryResult> InvokeAsync(GetFederatedOciRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFederatedOciRepositoryResult>("artifactory:index/getFederatedOciRepository:getFederatedOciRepository", args ?? new GetFederatedOciRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated OCI repository.
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
        ///     var my_oci_federated = Artifactory.GetFederatedOciRepository.Invoke(new()
        ///     {
        ///         Key = "my-oci-federated",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFederatedOciRepositoryResult> Invoke(GetFederatedOciRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedOciRepositoryResult>("artifactory:index/getFederatedOciRepository:getFederatedOciRepository", args ?? new GetFederatedOciRepositoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated OCI repository.
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
        ///     var my_oci_federated = Artifactory.GetFederatedOciRepository.Invoke(new()
        ///     {
        ///         Key = "my-oci-federated",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFederatedOciRepositoryResult> Invoke(GetFederatedOciRepositoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedOciRepositoryResult>("artifactory:index/getFederatedOciRepository:getFederatedOciRepository", args ?? new GetFederatedOciRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFederatedOciRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public bool? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public bool? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public bool? CdnRedirect { get; set; }

        [Input("cleanupOnDelete")]
        public bool? CleanupOnDelete { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
        /// </summary>
        [Input("disableProxy")]
        public bool? DisableProxy { get; set; }

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

        [Input("maxUniqueTags")]
        public int? MaxUniqueTags { get; set; }

        [Input("members")]
        private List<Inputs.GetFederatedOciRepositoryMemberArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public List<Inputs.GetFederatedOciRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new List<Inputs.GetFederatedOciRepositoryMemberArgs>());
            set => _members = value;
        }

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

        /// <summary>
        /// Proxy key from Artifactory Proxies settings.
        /// </summary>
        [Input("proxy")]
        public string? Proxy { get; set; }

        [Input("repoLayoutRef")]
        public string? RepoLayoutRef { get; set; }

        [Input("tagRetention")]
        public int? TagRetention { get; set; }

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetFederatedOciRepositoryArgs()
        {
        }
        public static new GetFederatedOciRepositoryArgs Empty => new GetFederatedOciRepositoryArgs();
    }

    public sealed class GetFederatedOciRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("archiveBrowsingEnabled")]
        public Input<bool>? ArchiveBrowsingEnabled { get; set; }

        [Input("blackedOut")]
        public Input<bool>? BlackedOut { get; set; }

        [Input("cdnRedirect")]
        public Input<bool>? CdnRedirect { get; set; }

        [Input("cleanupOnDelete")]
        public Input<bool>? CleanupOnDelete { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
        /// </summary>
        [Input("disableProxy")]
        public Input<bool>? DisableProxy { get; set; }

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

        [Input("maxUniqueTags")]
        public Input<int>? MaxUniqueTags { get; set; }

        [Input("members")]
        private InputList<Inputs.GetFederatedOciRepositoryMemberInputArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.GetFederatedOciRepositoryMemberInputArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GetFederatedOciRepositoryMemberInputArgs>());
            set => _members = value;
        }

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

        /// <summary>
        /// Proxy key from Artifactory Proxies settings.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("tagRetention")]
        public Input<int>? TagRetention { get; set; }

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetFederatedOciRepositoryInvokeArgs()
        {
        }
        public static new GetFederatedOciRepositoryInvokeArgs Empty => new GetFederatedOciRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetFederatedOciRepositoryResult
    {
        public readonly bool? ArchiveBrowsingEnabled;
        public readonly bool? BlackedOut;
        public readonly bool? CdnRedirect;
        public readonly bool? CleanupOnDelete;
        public readonly string? Description;
        /// <summary>
        /// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
        /// </summary>
        public readonly bool? DisableProxy;
        public readonly bool? DownloadDirect;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly int? MaxUniqueTags;
        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedOciRepositoryMemberResult> Members;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        /// <summary>
        /// Proxy key from Artifactory Proxies settings.
        /// </summary>
        public readonly string? Proxy;
        public readonly string? RepoLayoutRef;
        public readonly int? TagRetention;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetFederatedOciRepositoryResult(
            bool? archiveBrowsingEnabled,

            bool? blackedOut,

            bool? cdnRedirect,

            bool? cleanupOnDelete,

            string? description,

            bool? disableProxy,

            bool? downloadDirect,

            string? excludesPattern,

            string id,

            string? includesPattern,

            string key,

            int? maxUniqueTags,

            ImmutableArray<Outputs.GetFederatedOciRepositoryMemberResult> members,

            string? notes,

            string packageType,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? proxy,

            string? repoLayoutRef,

            int? tagRetention,

            bool? xrayIndex)
        {
            ArchiveBrowsingEnabled = archiveBrowsingEnabled;
            BlackedOut = blackedOut;
            CdnRedirect = cdnRedirect;
            CleanupOnDelete = cleanupOnDelete;
            Description = description;
            DisableProxy = disableProxy;
            DownloadDirect = downloadDirect;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            MaxUniqueTags = maxUniqueTags;
            Members = members;
            Notes = notes;
            PackageType = packageType;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            Proxy = proxy;
            RepoLayoutRef = repoLayoutRef;
            TagRetention = tagRetention;
            XrayIndex = xrayIndex;
        }
    }
}
