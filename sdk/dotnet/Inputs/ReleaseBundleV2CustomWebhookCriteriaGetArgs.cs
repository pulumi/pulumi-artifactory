// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ReleaseBundleV2CustomWebhookCriteriaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Trigger on any release bundle.
        /// </summary>
        [Input("anyReleaseBundle", required: true)]
        public Input<bool> AnyReleaseBundle { get; set; } = null!;

        [Input("excludePatterns")]
        private InputList<string>? _excludePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: "org/apache/**".
        /// </summary>
        public InputList<string> ExcludePatterns
        {
            get => _excludePatterns ?? (_excludePatterns = new InputList<string>());
            set => _excludePatterns = value;
        }

        [Input("includePatterns")]
        private InputList<string>? _includePatterns;

        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: "org/apache/**".
        /// </summary>
        public InputList<string> IncludePatterns
        {
            get => _includePatterns ?? (_includePatterns = new InputList<string>());
            set => _includePatterns = value;
        }

        [Input("selectedReleaseBundles", required: true)]
        private InputList<string>? _selectedReleaseBundles;

        /// <summary>
        /// Trigger on this list of release bundle names.
        /// </summary>
        public InputList<string> SelectedReleaseBundles
        {
            get => _selectedReleaseBundles ?? (_selectedReleaseBundles = new InputList<string>());
            set => _selectedReleaseBundles = value;
        }

        public ReleaseBundleV2CustomWebhookCriteriaGetArgs()
        {
        }
        public static new ReleaseBundleV2CustomWebhookCriteriaGetArgs Empty => new ReleaseBundleV2CustomWebhookCriteriaGetArgs();
    }
}