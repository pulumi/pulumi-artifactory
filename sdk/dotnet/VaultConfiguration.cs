// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// This resource enables you to configure an external vault connector to use as a centralized secret management tool for the keys used to sign packages. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/vault).
    /// This feature is supported with Enterprise X and Enterprise+ licenses.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// using Std = Pulumi.Std;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_vault_config_app_role = new Artifactory.VaultConfiguration("my-vault-config-app-role", new()
    ///     {
    ///         Name = "my-vault-config-app-role",
    ///         Config = new Artifactory.Inputs.VaultConfigurationConfigArgs
    ///         {
    ///             Url = "http://127.0.0.1:8200",
    ///             Auth = new Artifactory.Inputs.VaultConfigurationConfigAuthArgs
    ///             {
    ///                 Type = "AppRole",
    ///                 RoleId = "1b62ff05...",
    ///                 SecretId = "acbd6657...",
    ///             },
    ///             Mounts = new[]
    ///             {
    ///                 new Artifactory.Inputs.VaultConfigurationConfigMountArgs
    ///                 {
    ///                     Path = "secret",
    ///                     Type = "KV2",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var my_vault_config_cert = new Artifactory.VaultConfiguration("my-vault-config-cert", new()
    ///     {
    ///         Name = "my-vault-config-cert",
    ///         Config = new Artifactory.Inputs.VaultConfigurationConfigArgs
    ///         {
    ///             Url = "http://127.0.0.1:8200",
    ///             Auth = new Artifactory.Inputs.VaultConfigurationConfigAuthArgs
    ///             {
    ///                 Type = "Certificate",
    ///                 Certificate = Std.File.Invoke(new()
    ///                 {
    ///                     Input = "samples/public.pem",
    ///                 }).Apply(invoke =&gt; invoke.Result),
    ///                 CertificateKey = Std.File.Invoke(new()
    ///                 {
    ///                     Input = "samples/private.pem",
    ///                 }).Apply(invoke =&gt; invoke.Result),
    ///             },
    ///             Mounts = new[]
    ///             {
    ///                 new Artifactory.Inputs.VaultConfigurationConfigMountArgs
    ///                 {
    ///                     Path = "secret",
    ///                     Type = "KV2",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/vaultConfiguration:VaultConfiguration my-vault-config my-vault-config
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/vaultConfiguration:VaultConfiguration")]
    public partial class VaultConfiguration : global::Pulumi.CustomResource
    {
        [Output("config")]
        public Output<Outputs.VaultConfigurationConfig> Config { get; private set; } = null!;

        /// <summary>
        /// Name of the Vault configuration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a VaultConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VaultConfiguration(string name, VaultConfigurationArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/vaultConfiguration:VaultConfiguration", name, args ?? new VaultConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VaultConfiguration(string name, Input<string> id, VaultConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/vaultConfiguration:VaultConfiguration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VaultConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VaultConfiguration Get(string name, Input<string> id, VaultConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new VaultConfiguration(name, id, state, options);
        }
    }

    public sealed class VaultConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("config", required: true)]
        public Input<Inputs.VaultConfigurationConfigArgs> Config { get; set; } = null!;

        /// <summary>
        /// Name of the Vault configuration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VaultConfigurationArgs()
        {
        }
        public static new VaultConfigurationArgs Empty => new VaultConfigurationArgs();
    }

    public sealed class VaultConfigurationState : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<Inputs.VaultConfigurationConfigGetArgs>? Config { get; set; }

        /// <summary>
        /// Name of the Vault configuration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public VaultConfigurationState()
        {
        }
        public static new VaultConfigurationState Empty => new VaultConfigurationState();
    }
}
