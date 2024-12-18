// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ReleaseBundleV2SourceArtifactArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path for the artifact
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The SHA256 for the artifact
        /// </summary>
        [Input("sha256")]
        public Input<string>? Sha256 { get; set; }

        public ReleaseBundleV2SourceArtifactArgs()
        {
        }
        public static new ReleaseBundleV2SourceArtifactArgs Empty => new ReleaseBundleV2SourceArtifactArgs();
    }
}