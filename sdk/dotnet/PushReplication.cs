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
    /// &gt; This resource is deprecated and replaced by `artifactory.LocalRepositoryMultiReplication` for clarity.
    /// 
    /// Provides an Artifactory push replication resource. This can be used to create and manage Artifactory push replications using [Multi-push Replication API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-CreateorReplaceLocalMulti-pushReplication).
    /// 
    /// Push replication is used to synchronize Local Repositories, and is implemented by the Artifactory server on the near
    /// end invoking a synchronization of artifacts to the far end.
    /// 
    /// See the [Official Documentation](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-PushReplication).
    /// 
    /// &gt; This resource requires Artifactory Enterprise license.
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
    ///     var config = new Config();
    ///     // The base URL of the Artifactory deployment
    ///     var artifactoryUrl = config.Require("artifactoryUrl");
    ///     // The username for the Artifactory
    ///     var artifactoryUsername = config.Require("artifactoryUsername");
    ///     // The password for the Artifactory
    ///     var artifactoryPassword = config.Require("artifactoryPassword");
    ///     // Create a replication between two artifactory local repositories
    ///     var providerTestSource = new Artifactory.LocalMavenRepository("provider_test_source", new()
    ///     {
    ///         Key = "provider_test_source",
    ///     });
    /// 
    ///     var providerTestDest = new Artifactory.LocalMavenRepository("provider_test_dest", new()
    ///     {
    ///         Key = "provider_test_dest",
    ///     });
    /// 
    ///     var foo_rep = new Artifactory.PushReplication("foo-rep", new()
    ///     {
    ///         RepoKey = providerTestSource.Key,
    ///         CronExp = "0 0 * * * ?",
    ///         EnableEventReplication = true,
    ///         Replications = new[]
    ///         {
    ///             new Artifactory.Inputs.PushReplicationReplicationArgs
    ///             {
    ///                 Url = providerTestDest.Key.Apply(key =&gt; $"{artifactoryUrl}/{key}"),
    ///                 Username = "$var.artifactory_username",
    ///                 Password = "$var.artifactory_password",
    ///                 Enabled = true,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Push replication configs can be imported using their repo key, e.g.
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/pushReplication:PushReplication foo-rep provider_test_source
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/pushReplication:PushReplication")]
    public partial class PushReplication : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Output("cronExp")]
        public Output<string> CronExp { get; private set; } = null!;

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
        /// </summary>
        [Output("enableEventReplication")]
        public Output<bool?> EnableEventReplication { get; private set; } = null!;

        [Output("replications")]
        public Output<ImmutableArray<Outputs.PushReplicationReplication>> Replications { get; private set; } = null!;

        /// <summary>
        /// Repository name.
        /// </summary>
        [Output("repoKey")]
        public Output<string> RepoKey { get; private set; } = null!;


        /// <summary>
        /// Create a PushReplication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PushReplication(string name, PushReplicationArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/pushReplication:PushReplication", name, args ?? new PushReplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PushReplication(string name, Input<string> id, PushReplicationState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/pushReplication:PushReplication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PushReplication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PushReplication Get(string name, Input<string> id, PushReplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new PushReplication(name, id, state, options);
        }
    }

    public sealed class PushReplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Input("cronExp", required: true)]
        public Input<string> CronExp { get; set; } = null!;

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
        /// </summary>
        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        [Input("replications")]
        private InputList<Inputs.PushReplicationReplicationArgs>? _replications;
        public InputList<Inputs.PushReplicationReplicationArgs> Replications
        {
            get => _replications ?? (_replications = new InputList<Inputs.PushReplicationReplicationArgs>());
            set => _replications = value;
        }

        /// <summary>
        /// Repository name.
        /// </summary>
        [Input("repoKey", required: true)]
        public Input<string> RepoKey { get; set; } = null!;

        public PushReplicationArgs()
        {
        }
        public static new PushReplicationArgs Empty => new PushReplicationArgs();
    }

    public sealed class PushReplicationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A valid CRON expression that you can use to control replication frequency. Eg: `0 0 12 * * ? *`, `0 0 2 ? * MON-SAT *`. Note: use 6 or 7 parts format - Seconds, Minutes Hours, Day Of Month, Month, Day Of Week, Year (optional). Specifying both a day-of-week AND a day-of-month parameter is not supported. One of them should be replaced by `?`. Incorrect: `* 5,7,9 14/2 * * WED,SAT *`, correct: `* 5,7,9 14/2 ? * WED,SAT *`. See details in [Cron Trigger Tutorial](https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html).
        /// </summary>
        [Input("cronExp")]
        public Input<string>? CronExp { get; set; }

        /// <summary>
        /// When set, each event will trigger replication of the artifacts changed in this event. This can be any type of event on artifact, e.g. added, deleted or property change.
        /// </summary>
        [Input("enableEventReplication")]
        public Input<bool>? EnableEventReplication { get; set; }

        [Input("replications")]
        private InputList<Inputs.PushReplicationReplicationGetArgs>? _replications;
        public InputList<Inputs.PushReplicationReplicationGetArgs> Replications
        {
            get => _replications ?? (_replications = new InputList<Inputs.PushReplicationReplicationGetArgs>());
            set => _replications = value;
        }

        /// <summary>
        /// Repository name.
        /// </summary>
        [Input("repoKey")]
        public Input<string>? RepoKey { get; set; }

        public PushReplicationState()
        {
        }
        public static new PushReplicationState Empty => new PushReplicationState();
    }
}
