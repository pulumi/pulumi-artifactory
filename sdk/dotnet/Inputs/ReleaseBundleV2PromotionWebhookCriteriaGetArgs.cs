// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ReleaseBundleV2PromotionWebhookCriteriaGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludePatterns")]
        private InputList<string>? _excludePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
        /// </summary>
        public InputList<string> ExcludePatterns
        {
            get => _excludePatterns ?? (_excludePatterns = new InputList<string>());
            set => _excludePatterns = value;
        }

        [Input("includePatterns")]
        private InputList<string>? _includePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\nAnt-style path expressions are supported (*, **, ?).\nFor example: "org/apache/**"
        /// </summary>
        public InputList<string> IncludePatterns
        {
            get => _includePatterns ?? (_includePatterns = new InputList<string>());
            set => _includePatterns = value;
        }

        [Input("selectedEnvironments", required: true)]
        private InputList<string>? _selectedEnvironments;

        /// <summary>
        /// Trigger on this list of environment names.
        /// </summary>
        public InputList<string> SelectedEnvironments
        {
            get => _selectedEnvironments ?? (_selectedEnvironments = new InputList<string>());
            set => _selectedEnvironments = value;
        }

        public ReleaseBundleV2PromotionWebhookCriteriaGetArgs()
        {
        }
        public static new ReleaseBundleV2PromotionWebhookCriteriaGetArgs Empty => new ReleaseBundleV2PromotionWebhookCriteriaGetArgs();
    }
}
