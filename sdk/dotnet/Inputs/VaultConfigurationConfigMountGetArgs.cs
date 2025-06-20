// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class VaultConfigurationConfigMountGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Vault secret engine path
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Vault supports several secret engines, each one has different capabilities. The supported secret engine types are: `KV1` and `KV2`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public VaultConfigurationConfigMountGetArgs()
        {
        }
        public static new VaultConfigurationConfigMountGetArgs Empty => new VaultConfigurationConfigMountGetArgs();
    }
}
