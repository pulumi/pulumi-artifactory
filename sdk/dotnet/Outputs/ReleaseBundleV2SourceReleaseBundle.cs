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
    public sealed class ReleaseBundleV2SourceReleaseBundle
    {
        /// <summary>
        /// The name of the release bundle.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Project key of the release bundle.
        /// </summary>
        public readonly string? ProjectKey;
        /// <summary>
        /// The key of the release bundle repository.
        /// </summary>
        public readonly string? RepositoryKey;
        /// <summary>
        /// The version of the release bundle.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private ReleaseBundleV2SourceReleaseBundle(
            string name,

            string? projectKey,

            string? repositoryKey,

            string version)
        {
            Name = name;
            ProjectKey = projectKey;
            RepositoryKey = repositoryKey;
            Version = version;
        }
    }
}
