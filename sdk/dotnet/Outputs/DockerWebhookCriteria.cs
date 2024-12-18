// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class DockerWebhookCriteria
    {
        /// <summary>
        /// Trigger on any federated repo.
        /// </summary>
        public readonly bool AnyFederated;
        /// <summary>
        /// Trigger on any local repo.
        /// </summary>
        public readonly bool AnyLocal;
        /// <summary>
        /// Trigger on any remote repo.
        /// </summary>
        public readonly bool AnyRemote;
        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public readonly ImmutableArray<string> ExcludePatterns;
        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public readonly ImmutableArray<string> IncludePatterns;
        /// <summary>
        /// Trigger on this list of repo keys.
        /// </summary>
        public readonly ImmutableArray<string> RepoKeys;

        [OutputConstructor]
        private DockerWebhookCriteria(
            bool anyFederated,

            bool anyLocal,

            bool anyRemote,

            ImmutableArray<string> excludePatterns,

            ImmutableArray<string> includePatterns,

            ImmutableArray<string> repoKeys)
        {
            AnyFederated = anyFederated;
            AnyLocal = anyLocal;
            AnyRemote = anyRemote;
            ExcludePatterns = excludePatterns;
            IncludePatterns = includePatterns;
            RepoKeys = repoKeys;
        }
    }
}