// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class LocalRepositoryMultiReplicationReplicationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
        /// </summary>
        [Input("checkBinaryExistenceInFilestore")]
        public Input<bool>? CheckBinaryExistenceInFilestore { get; set; }

        /// <summary>
        /// When set to `true`, the `proxy` attribute will be ignored (from version 7.41.7). The default value is `false`.
        /// </summary>
        [Input("disableProxy")]
        public Input<bool>? DisableProxy { get; set; }

        /// <summary>
        /// When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`. By default, no artifacts are excluded.
        /// </summary>
        [Input("excludePathPrefixPattern")]
        public Input<string>? ExcludePathPrefixPattern { get; set; }

        /// <summary>
        /// List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**/*)`.
        /// </summary>
        [Input("includePathPrefixPattern")]
        public Input<string>? IncludePathPrefixPattern { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Replication ID, the value is unknown until the resource is created. Can't be set or updated.
        /// </summary>
        [Input("replicationKey")]
        public Input<string>? ReplicationKey { get; set; }

        /// <summary>
        /// The network timeout in milliseconds to use for remote operations. Default value is `15000`.
        /// </summary>
        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        /// <summary>
        /// When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
        /// </summary>
        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        /// <summary>
        /// When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
        /// </summary>
        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        /// <summary>
        /// When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
        /// </summary>
        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        /// <summary>
        /// The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// Username on the remote Artifactory instance.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public LocalRepositoryMultiReplicationReplicationGetArgs()
        {
        }
        public static new LocalRepositoryMultiReplicationReplicationGetArgs Empty => new LocalRepositoryMultiReplicationReplicationGetArgs();
    }
}
