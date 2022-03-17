// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// ## # Artifactory Virtual Helm Repository Resource
    /// 
    /// Provides an Artifactory virtual repository resource with Helm package type. This should be preferred over the original one-size-fits-all `artifactory.VirtualRepository`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo_helm_virtual = new Artifactory.VirtualHelmRepository("foo-helm-virtual", new Artifactory.VirtualHelmRepositoryArgs
    ///         {
    ///             Key = "foo-helm-virtual",
    ///             UseNamespaces = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Virtual repositories can be imported using their name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/virtualHelmRepository:VirtualHelmRepository foo foo
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/virtualHelmRepository:VirtualHelmRepository")]
    public partial class VirtualHelmRepository : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        /// another Artifactory instance.
        /// </summary>
        [Output("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Output<bool?> ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; private set; } = null!;

        /// <summary>
        /// Default repository to deploy artifacts.
        /// </summary>
        [Output("defaultDeploymentRepo")]
        public Output<string?> DefaultDeploymentRepo { get; private set; } = null!;

        /// <summary>
        /// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
        /// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Output("excludesPattern")]
        public Output<string?> ExcludesPattern { get; private set; } = null!;

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Output("includesPattern")]
        public Output<string?> IncludesPattern { get; private set; } = null!;

        /// <summary>
        /// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
        /// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
        /// 'libs-release-local').
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// A free text field to add additional notes about the repository. These are only visible to the administrator.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
        /// </summary>
        [Output("packageType")]
        public Output<string> PackageType { get; private set; } = null!;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        [Output("projectEnvironments")]
        public Output<ImmutableArray<string>> ProjectEnvironments { get; private set; } = null!;

        /// <summary>
        /// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
        /// repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Output("projectKey")]
        public Output<string?> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Output("repoLayoutRef")]
        public Output<string?> RepoLayoutRef { get; private set; } = null!;

        /// <summary>
        /// The effective list of actual repositories included in this virtual repository.
        /// </summary>
        [Output("repositories")]
        public Output<ImmutableArray<string>> Repositories { get; private set; } = null!;

        /// <summary>
        /// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
        /// repositories. A value of 0 indicates no caching.
        /// </summary>
        [Output("retrievalCachePeriodSeconds")]
        public Output<int?> RetrievalCachePeriodSeconds { get; private set; } = null!;

        /// <summary>
        /// From Artifactory 7.24.1 (SaaS Version), you can explicitly state a specific aggregated local or remote repository to fetch from a virtual by assigning namespaces to local and remote repositories. See https://www.jfrog.com/confluence/display/JFROG/Kubernetes+Helm+Chart+Repositories#KubernetesHelmChartRepositories-NamespaceSupportforHelmVirtualRepositories. Default to 'false'.
        /// </summary>
        [Output("useNamespaces")]
        public Output<bool?> UseNamespaces { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualHelmRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualHelmRepository(string name, VirtualHelmRepositoryArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/virtualHelmRepository:VirtualHelmRepository", name, args ?? new VirtualHelmRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualHelmRepository(string name, Input<string> id, VirtualHelmRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/virtualHelmRepository:VirtualHelmRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualHelmRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualHelmRepository Get(string name, Input<string> id, VirtualHelmRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualHelmRepository(name, id, state, options);
        }
    }

    public sealed class VirtualHelmRepositoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        /// another Artifactory instance.
        /// </summary>
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        /// <summary>
        /// Default repository to deploy artifacts.
        /// </summary>
        [Input("defaultDeploymentRepo")]
        public Input<string>? DefaultDeploymentRepo { get; set; }

        /// <summary>
        /// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
        /// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
        /// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
        /// 'libs-release-local').
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// A free text field to add additional notes about the repository. These are only visible to the administrator.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
        /// repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private InputList<string>? _repositories;

        /// <summary>
        /// The effective list of actual repositories included in this virtual repository.
        /// </summary>
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        /// <summary>
        /// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
        /// repositories. A value of 0 indicates no caching.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        /// <summary>
        /// From Artifactory 7.24.1 (SaaS Version), you can explicitly state a specific aggregated local or remote repository to fetch from a virtual by assigning namespaces to local and remote repositories. See https://www.jfrog.com/confluence/display/JFROG/Kubernetes+Helm+Chart+Repositories#KubernetesHelmChartRepositories-NamespaceSupportforHelmVirtualRepositories. Default to 'false'.
        /// </summary>
        [Input("useNamespaces")]
        public Input<bool>? UseNamespaces { get; set; }

        public VirtualHelmRepositoryArgs()
        {
        }
    }

    public sealed class VirtualHelmRepositoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
        /// another Artifactory instance.
        /// </summary>
        [Input("artifactoryRequestsCanRetrieveRemoteArtifacts")]
        public Input<bool>? ArtifactoryRequestsCanRetrieveRemoteArtifacts { get; set; }

        /// <summary>
        /// Default repository to deploy artifacts.
        /// </summary>
        [Input("defaultDeploymentRepo")]
        public Input<string>? DefaultDeploymentRepo { get; set; }

        /// <summary>
        /// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
        /// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
        /// artifacts are excluded.
        /// </summary>
        [Input("excludesPattern")]
        public Input<string>? ExcludesPattern { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
        /// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
        /// </summary>
        [Input("includesPattern")]
        public Input<string>? IncludesPattern { get; set; }

        /// <summary>
        /// The Repository Key. A mandatory identifier for the repository and must be unique. It cannot begin with a number or
        /// contain spaces or special characters. For local repositories, we recommend using a '-local' suffix (e.g.
        /// 'libs-release-local').
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// A free text field to add additional notes about the repository. These are only visible to the administrator.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        [Input("projectEnvironments")]
        private InputList<string>? _projectEnvironments;

        /// <summary>
        /// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
        /// </summary>
        public InputList<string> ProjectEnvironments
        {
            get => _projectEnvironments ?? (_projectEnvironments = new InputList<string>());
            set => _projectEnvironments = value;
        }

        /// <summary>
        /// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
        /// repository to a project, repository key must be prefixed with project key, separated by a dash.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Repository layout key for the local repository
        /// </summary>
        [Input("repoLayoutRef")]
        public Input<string>? RepoLayoutRef { get; set; }

        [Input("repositories")]
        private InputList<string>? _repositories;

        /// <summary>
        /// The effective list of actual repositories included in this virtual repository.
        /// </summary>
        public InputList<string> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<string>());
            set => _repositories = value;
        }

        /// <summary>
        /// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
        /// repositories. A value of 0 indicates no caching.
        /// </summary>
        [Input("retrievalCachePeriodSeconds")]
        public Input<int>? RetrievalCachePeriodSeconds { get; set; }

        /// <summary>
        /// From Artifactory 7.24.1 (SaaS Version), you can explicitly state a specific aggregated local or remote repository to fetch from a virtual by assigning namespaces to local and remote repositories. See https://www.jfrog.com/confluence/display/JFROG/Kubernetes+Helm+Chart+Repositories#KubernetesHelmChartRepositories-NamespaceSupportforHelmVirtualRepositories. Default to 'false'.
        /// </summary>
        [Input("useNamespaces")]
        public Input<bool>? UseNamespaces { get; set; }

        public VirtualHelmRepositoryState()
        {
        }
    }
}