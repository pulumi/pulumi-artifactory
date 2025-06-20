// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class BuildWebhookCriteriaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger on any build.
        /// </summary>
        [Input("anyBuild", required: true)]
        public Input<bool> AnyBuild { get; set; } = null!;

        [Input("excludePatterns")]
        private InputList<string>? _excludePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public InputList<string> ExcludePatterns
        {
            get => _excludePatterns ?? (_excludePatterns = new InputList<string>());
            set => _excludePatterns = value;
        }

        [Input("includePatterns")]
        private InputList<string>? _includePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: `org/apache/**`.
        /// </summary>
        public InputList<string> IncludePatterns
        {
            get => _includePatterns ?? (_includePatterns = new InputList<string>());
            set => _includePatterns = value;
        }

        [Input("selectedBuilds", required: true)]
        private InputList<string>? _selectedBuilds;

        /// <summary>
        /// Trigger on this list of build names.
        /// </summary>
        public InputList<string> SelectedBuilds
        {
            get => _selectedBuilds ?? (_selectedBuilds = new InputList<string>());
            set => _selectedBuilds = value;
        }

        public BuildWebhookCriteriaGetArgs()
        {
        }
        public static new BuildWebhookCriteriaGetArgs Empty => new BuildWebhookCriteriaGetArgs();
    }
}
