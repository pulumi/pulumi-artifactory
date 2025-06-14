// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ArchivePolicySearchCriteriaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The archive policy will archive packages based on how long ago they were created. For example, if this parameter is 2 then packages created more than 2 months ago will be archived as part of the policy.
        /// </summary>
        [Input("createdBeforeInMonths")]
        public Input<int>? CreatedBeforeInMonths { get; set; }

        [Input("excludedPackages")]
        private InputList<string>? _excludedPackages;

        /// <summary>
        /// Specify explicit package names that you want excluded from the policy. Only Name explicit names (and not patterns) are accepted.
        /// </summary>
        public InputList<string> ExcludedPackages
        {
            get => _excludedPackages ?? (_excludedPackages = new InputList<string>());
            set => _excludedPackages = value;
        }

        [Input("excludedRepos")]
        private InputList<string>? _excludedRepos;

        /// <summary>
        /// Specify patterns for repository names or explicit repository names that you want excluded from the archive policy.
        /// </summary>
        public InputList<string> ExcludedRepos
        {
            get => _excludedRepos ?? (_excludedRepos = new InputList<string>());
            set => _excludedRepos = value;
        }

        /// <summary>
        /// Set this value to `true` if you want the policy to run on all Artifactory projects. The default value is `false`.
        /// </summary>
        [Input("includeAllProjects")]
        public Input<bool>? IncludeAllProjects { get; set; }

        [Input("includedPackages", required: true)]
        private InputList<string>? _includedPackages;

        /// <summary>
        /// Specify a pattern for a package name or an explicit package name. It accept only single element which can be specific package or pattern, and for including all packages use `**`. Example: `included_packages = ["**"]`
        /// </summary>
        public InputList<string> IncludedPackages
        {
            get => _includedPackages ?? (_includedPackages = new InputList<string>());
            set => _includedPackages = value;
        }

        [Input("includedProjects")]
        private InputList<string>? _includedProjects;

        /// <summary>
        /// List of projects on which you want this policy to run. To include repositories that are not assigned to any project, enter the project key `default`.
        /// 
        /// ~&gt;This setting is relevant only on the global level, for Platform Admins.
        /// </summary>
        public InputList<string> IncludedProjects
        {
            get => _includedProjects ?? (_includedProjects = new InputList<string>());
            set => _includedProjects = value;
        }

        /// <summary>
        /// Set a value for the number of latest versions to keep. The archive policy will remove all versions before the number you select here. The latest version is always excluded.
        /// 
        /// ~&gt;Versions are determined by creation date.
        /// 
        /// ~&gt;Not all package types support this condition. If you include a package type in your policy that is not compatible with this condition, a validation error (400) is returned. For information on which package types support this condition, see here.
        /// </summary>
        [Input("keepLastNVersions")]
        public Input<int>? KeepLastNVersions { get; set; }

        /// <summary>
        /// The archive policy will archive packages based on how long ago they were downloaded. For example, if this parameter is 5 then packages downloaded more than 5 months ago will be archived as part of the policy.
        /// 
        /// ~&gt;JFrog recommends using the `last_downloaded_before_in_months` condition to ensure that packages currently in use are not archived.
        /// </summary>
        [Input("lastDownloadedBeforeInMonths")]
        public Input<int>? LastDownloadedBeforeInMonths { get; set; }

        [Input("packageTypes", required: true)]
        private InputList<string>? _packageTypes;
        public InputList<string> PackageTypes
        {
            get => _packageTypes ?? (_packageTypes = new InputList<string>());
            set => _packageTypes = value;
        }

        [Input("repos", required: true)]
        private InputList<string>? _repos;

        /// <summary>
        /// Specify one or more patterns for the repository name(s) on which you want the archive policy to run. You can also specify explicit repository names. Specifying at least one pattern or explicit name is required. Only packages in repositories that match the pattern or explicit name will be archived. For including all repos use `**`. Example: `repos = ["**"]`
        /// </summary>
        public InputList<string> Repos
        {
            get => _repos ?? (_repos = new InputList<string>());
            set => _repos = value;
        }

        public ArchivePolicySearchCriteriaArgs()
        {
        }
        public static new ArchivePolicySearchCriteriaArgs Empty => new ArchivePolicySearchCriteriaArgs();
    }
}
