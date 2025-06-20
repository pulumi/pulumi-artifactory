// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class BuildWebhookCriteria
    {
        /// <summary>
        /// Trigger on any build.
        /// </summary>
        public readonly bool AnyBuild;
        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public readonly ImmutableArray<string> ExcludePatterns;
        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public readonly ImmutableArray<string> IncludePatterns;
        /// <summary>
        /// Trigger on this list of build names.
        /// </summary>
        public readonly ImmutableArray<string> SelectedBuilds;

        [OutputConstructor]
        private BuildWebhookCriteria(
            bool anyBuild,

            ImmutableArray<string> excludePatterns,

            ImmutableArray<string> includePatterns,

            ImmutableArray<string> selectedBuilds)
        {
            AnyBuild = anyBuild;
            ExcludePatterns = excludePatterns;
            IncludePatterns = includePatterns;
            SelectedBuilds = selectedBuilds;
        }
    }
}
