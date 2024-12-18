// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class VaultConfigurationConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("auth", required: true)]
        public Input<Inputs.VaultConfigurationConfigAuthArgs> Auth { get; set; } = null!;

        [Input("mounts", required: true)]
        private InputList<Inputs.VaultConfigurationConfigMountArgs>? _mounts;
        public InputList<Inputs.VaultConfigurationConfigMountArgs> Mounts
        {
            get => _mounts ?? (_mounts = new InputList<Inputs.VaultConfigurationConfigMountArgs>());
            set => _mounts = value;
        }

        /// <summary>
        /// The base URL of the Vault server.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public VaultConfigurationConfigArgs()
        {
        }
        public static new VaultConfigurationConfigArgs Empty => new VaultConfigurationConfigArgs();
    }
}