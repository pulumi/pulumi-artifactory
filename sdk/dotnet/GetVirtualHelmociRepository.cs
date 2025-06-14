// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetVirtualHelmociRepository
    {
        /// <summary>
        /// Retrieves a virtual Helm OCI repository.
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
        ///     var my_helmoci_virtual = Artifactory.GetVirtualHelmociRepository.Invoke(new()
        ///     {
        ///         Key = "my-helmoci-virtual",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVirtualHelmociRepositoryResult> InvokeAsync(GetVirtualHelmociRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualHelmociRepositoryResult>("artifactory:index/getVirtualHelmociRepository:getVirtualHelmociRepository", args ?? new GetVirtualHelmociRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Helm OCI repository.
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
        ///     var my_helmoci_virtual = Artifactory.GetVirtualHelmociRepository.Invoke(new()
        ///     {
        ///         Key = "my-helmoci-virtual",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualHelmociRepositoryResult> Invoke(GetVirtualHelmociRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualHelmociRepositoryResult>("artifactory:index/getVirtualHelmociRepository:getVirtualHelmociRepository", args ?? new GetVirtualHelmociRepositoryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a virtual Helm OCI repository.
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
        ///     var my_helmoci_virtual = Artifactory.GetVirtualHelmociRepository.Invoke(new()
        ///     {
        ///         Key = "my-helmoci-virtual",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVirtualHelmociRepositoryResult> Invoke(GetVirtualHelmociRepositoryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualHelmociRepositoryResult>("artifactory:index/getVirtualHelmociRepository:getVirtualHelmociRepository", args ?? new GetVirtualHelmociRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualHelmociRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        [Input("defaultDeploymentRepo")]
        public string? DefaultDeploymentRepo { get; set; }

        [Input("description")]
        public string? Description { get; set; }

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

        [Input("projectEnvironments")]
        private List<string>? _projectEnvironments;
        public List<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new List<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectKey")]
        public string? ProjectKey { get; set; }

        [Input("repoLayoutRef")]
        public string? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private List<string>? _repositories;
        public List<string> Repositories
        {
            get => _repositories ?? (_repositories = new List<string>());
            set => _repositories = value;
        }

        /// <summary>
        /// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
        /// </summary>
        [Input("resolveOciTagsByTimestamp")]
        public bool? ResolveOciTagsByTimestamp { get; set; }

        public GetVirtualHelmociRepositoryArgs()
        {
        }
        public static new GetVirtualHelmociRepositoryArgs Empty => new GetVirtualHelmociRepositoryArgs();
    }

    public sealed class GetVirtualHelmociRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        [Input("defaultDeploymentRepo")]
        public Input<string>? DefaultDeploymentRepo { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

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

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private InputList<string>? _repositories;
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        /// <summary>
        /// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
        /// </summary>
        [Input("resolveOciTagsByTimestamp")]
        public Input<bool>? ResolveOciTagsByTimestamp { get; set; }

        public GetVirtualHelmociRepositoryInvokeArgs()
        {
        }
        public static new GetVirtualHelmociRepositoryInvokeArgs Empty => new GetVirtualHelmociRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualHelmociRepositoryResult
    {
        public readonly bool? ArtifactoryRequestsCanRetrieveRemoteArtifacts;
        public readonly string? DefaultDeploymentRepo;
        public readonly string? Description;
        public readonly string? ExcludesPattern;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IncludesPattern;
        public readonly string Key;
        public readonly string? Notes;
        public readonly string PackageType;
        public readonly ImmutableArray<string> ProjectEnvironments;
        public readonly string? ProjectKey;
        public readonly string? RepoLayoutRef;
        public readonly ImmutableArray<string> Repositories;
        /// <summary>
        /// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
        /// </summary>
        public readonly bool? ResolveOciTagsByTimestamp;

        [OutputConstructor]
        private GetVirtualHelmociRepositoryResult(
            bool? artifactoryRequestsCanRetrieveRemoteArtifacts,

            string? defaultDeploymentRepo,

            string? description,

            string? excludesPattern,

            string id,

            string? includesPattern,

            string key,

            string? notes,

            string packageType,

            ImmutableArray<string> projectEnvironments,

            string? projectKey,

            string? repoLayoutRef,

            ImmutableArray<string> repositories,

            bool? resolveOciTagsByTimestamp)
        {
            ArtifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            DefaultDeploymentRepo = defaultDeploymentRepo;
            Description = description;
            ExcludesPattern = excludesPattern;
            Id = id;
            IncludesPattern = includesPattern;
            Key = key;
            Notes = notes;
            PackageType = packageType;
            ProjectEnvironments = projectEnvironments;
            ProjectKey = projectKey;
            RepoLayoutRef = repoLayoutRef;
            Repositories = repositories;
            ResolveOciTagsByTimestamp = resolveOciTagsByTimestamp;
        }
    }
}
