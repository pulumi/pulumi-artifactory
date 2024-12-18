// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFederatedAnsibleRepository
    {
        /// <summary>
        /// Retrieves a federated Ansible repository.
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
        ///     var federated_test_ansible_repo = Artifactory.GetFederatedAnsibleRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-ansible-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFederatedAnsibleRepositoryResult> InvokeAsync(GetFederatedAnsibleRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFederatedAnsibleRepositoryResult>("artifactory:index/getFederatedAnsibleRepository:getFederatedAnsibleRepository", args ?? new GetFederatedAnsibleRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated Ansible repository.
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
        ///     var federated_test_ansible_repo = Artifactory.GetFederatedAnsibleRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-ansible-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFederatedAnsibleRepositoryResult> Invoke(GetFederatedAnsibleRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedAnsibleRepositoryResult>("artifactory:index/getFederatedAnsibleRepository:getFederatedAnsibleRepository", args ?? new GetFederatedAnsibleRepositoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a federated Ansible repository.
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
        ///     var federated_test_ansible_repo = Artifactory.GetFederatedAnsibleRepository.Invoke(new()
        ///     {
        ///         Key = "federated-test-ansible-repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFederatedAnsibleRepositoryResult> Invoke(GetFederatedAnsibleRepositoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFederatedAnsibleRepositoryResult>("artifactory:index/getFederatedAnsibleRepository:getFederatedAnsibleRepository", args ?? new GetFederatedAnsibleRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFederatedAnsibleRepositoryArgs : global::Pulumi.InvokeArgs
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

        [Input("members")]
        private List<Inputs.GetFederatedAnsibleRepositoryMemberArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public List<Inputs.GetFederatedAnsibleRepositoryMemberArgs> Members
        {
            get => _members ?? (_members = new List<Inputs.GetFederatedAnsibleRepositoryMemberArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public string? Notes { get; set; }

        [Input("primaryKeypairRef")]
        public string? PrimaryKeypairRef { get; set; }

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

        [Input("xrayIndex")]
        public bool? XrayIndex { get; set; }

        public GetFederatedAnsibleRepositoryArgs()
        {
        }
        public static new GetFederatedAnsibleRepositoryArgs Empty => new GetFederatedAnsibleRepositoryArgs();
    }

    public sealed class GetFederatedAnsibleRepositoryInvokeArgs : global::Pulumi.InvokeArgs
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

        [Input("members")]
        private InputList<Inputs.GetFederatedAnsibleRepositoryMemberInputArgs>? _members;

        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public InputList<Inputs.GetFederatedAnsibleRepositoryMemberInputArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.GetFederatedAnsibleRepositoryMemberInputArgs>());
            set => _members = value;
        }

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("primaryKeypairRef")]
        public Input<string>? PrimaryKeypairRef { get; set; }

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

        [Input("xrayIndex")]
        public Input<bool>? XrayIndex { get; set; }

        public GetFederatedAnsibleRepositoryInvokeArgs()
        {
        }
        public static new GetFederatedAnsibleRepositoryInvokeArgs Empty => new GetFederatedAnsibleRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetFederatedAnsibleRepositoryResult
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
        /// <summary>
        /// The list of Federated members and must contain this repository URL (configured base URL
        /// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
        /// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
        /// to set up Federated repositories correctly.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFederatedAnsibleRepositoryMemberResult> Members;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly string? PrimaryKeypairRef;
        public readonly bool? PriorityResolution;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly ImmutableArray<string> PropertySets;
        /// <summary>
        /// Proxy key from Artifactory Proxies settings.
        /// </summary>
        public readonly string? Proxy;
        public readonly string? RepoLayoutRef;
        public readonly bool? XrayIndex;

        [OutputConstructor]
        private GetFederatedAnsibleRepositoryResult(
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

            ImmutableArray<Outputs.GetFederatedAnsibleRepositoryMemberResult> members,

            string? notes,

            string packageType,

            string? primaryKeypairRef,

            bool? priorityResolution,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            ImmutableArray<string> propertySets,

            string? proxy,

            string? repoLayoutRef,

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
            Members = members;
            Notes = notes;
            PackageType = packageType;
            PrimaryKeypairRef = primaryKeypairRef;
            PriorityResolution = priorityResolution;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            PropertySets = propertySets;
            Proxy = proxy;
            RepoLayoutRef = repoLayoutRef;
            XrayIndex = xrayIndex;
        }
    }
}