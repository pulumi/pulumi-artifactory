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
    public sealed class GetPermissionTargetReleaseBundleResult
    {
        public readonly Outputs.GetPermissionTargetReleaseBundleActionsResult? Actions;
        /// <summary>
        /// Pattern of artifacts to exclude.
        /// </summary>
        public readonly ImmutableArray<string> ExcludesPatterns;
        /// <summary>
        /// Pattern of artifacts to include.
        /// </summary>
        public readonly ImmutableArray<string> IncludesPatterns;
        /// <summary>
        /// List of repositories this permission target is applicable for. You can specify the
        /// name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote
        /// repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
        /// </summary>
        public readonly ImmutableArray<string> Repositories;

        [OutputConstructor]
        private GetPermissionTargetReleaseBundleResult(
            Outputs.GetPermissionTargetReleaseBundleActionsResult? actions,

            ImmutableArray<string> excludesPatterns,

            ImmutableArray<string> includesPatterns,

            ImmutableArray<string> repositories)
        {
            Actions = actions;
            ExcludesPatterns = excludesPatterns;
            IncludesPatterns = includesPatterns;
            Repositories = repositories;
        }
    }
}
