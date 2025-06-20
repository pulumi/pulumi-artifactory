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
    /// This resource can be used to manage Artifactory's OAuth SSO settings.
    /// 
    /// Only a single `artifactory.OauthSettings` resource is meant to be defined.
    /// 
    /// ~&gt;The `artifactory.OauthSettings` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Configure Artifactory OAuth SSO settings
    ///     var oauth = new Artifactory.OauthSettings("oauth", new()
    ///     {
    ///         Enable = true,
    ///         PersistUsers = true,
    ///         AllowUserToAccessProfile = true,
    ///         OauthProviders = new[]
    ///         {
    ///             new Artifactory.Inputs.OauthSettingsOauthProviderArgs
    ///             {
    ///                 Name = "okta",
    ///                 Enabled = false,
    ///                 Type = "openId",
    ///                 ClientId = "foo",
    ///                 ClientSecret = "bar",
    ///                 ApiUrl = "https://organization.okta.com/oauth2/v1/userinfo",
    ///                 AuthUrl = "https://organization.okta.com/oauth2/v1/authorize",
    ///                 TokenUrl = "https://organization.okta.com/oauth2/v1/token",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Current OAuth SSO settings can be imported using `oauth_settings` as the `ID`.
    /// If the resource is being imported, there will be a state drift, because `client_secret` can't be known. There are two options on how to approach this:
    /// 1) Don't set `client_secret` initially, import, then update the config with actual secret;
    /// 2) Accept that there is a drift initially and run `pulumi up` twice;
    /// 
    /// ```sh
    /// $ pulumi import artifactory:index/oauthSettings:OauthSettings oauth oauth_settings
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/oauthSettings:OauthSettings")]
    public partial class OauthSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allow persisted users to access their profile.  Default value is `false`.
        /// </summary>
        [Output("allowUserToAccessProfile")]
        public Output<bool?> AllowUserToAccessProfile { get; private set; } = null!;

        /// <summary>
        /// Enable OAuth SSO.  Default value is `true`.
        /// </summary>
        [Output("enable")]
        public Output<bool?> Enable { get; private set; } = null!;

        /// <summary>
        /// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        /// </summary>
        [Output("oauthProviders")]
        public Output<ImmutableArray<Outputs.OauthSettingsOauthProvider>> OauthProviders { get; private set; } = null!;

        /// <summary>
        /// Enable the creation of local Artifactory users.  Default value is `false`.
        /// </summary>
        [Output("persistUsers")]
        public Output<bool?> PersistUsers { get; private set; } = null!;


        /// <summary>
        /// Create a OauthSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OauthSettings(string name, OauthSettingsArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/oauthSettings:OauthSettings", name, args ?? new OauthSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OauthSettings(string name, Input<string> id, OauthSettingsState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/oauthSettings:OauthSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OauthSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OauthSettings Get(string name, Input<string> id, OauthSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new OauthSettings(name, id, state, options);
        }
    }

    public sealed class OauthSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow persisted users to access their profile.  Default value is `false`.
        /// </summary>
        [Input("allowUserToAccessProfile")]
        public Input<bool>? AllowUserToAccessProfile { get; set; }

        /// <summary>
        /// Enable OAuth SSO.  Default value is `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        [Input("oauthProviders", required: true)]
        private InputList<Inputs.OauthSettingsOauthProviderArgs>? _oauthProviders;

        /// <summary>
        /// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        /// </summary>
        public InputList<Inputs.OauthSettingsOauthProviderArgs> OauthProviders
        {
            get => _oauthProviders ?? (_oauthProviders = new InputList<Inputs.OauthSettingsOauthProviderArgs>());
            set => _oauthProviders = value;
        }

        /// <summary>
        /// Enable the creation of local Artifactory users.  Default value is `false`.
        /// </summary>
        [Input("persistUsers")]
        public Input<bool>? PersistUsers { get; set; }

        public OauthSettingsArgs()
        {
        }
        public static new OauthSettingsArgs Empty => new OauthSettingsArgs();
    }

    public sealed class OauthSettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow persisted users to access their profile.  Default value is `false`.
        /// </summary>
        [Input("allowUserToAccessProfile")]
        public Input<bool>? AllowUserToAccessProfile { get; set; }

        /// <summary>
        /// Enable OAuth SSO.  Default value is `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        [Input("oauthProviders")]
        private InputList<Inputs.OauthSettingsOauthProviderGetArgs>? _oauthProviders;

        /// <summary>
        /// OAuth provider settings block. Multiple blocks can be defined, at least one is required.
        /// </summary>
        public InputList<Inputs.OauthSettingsOauthProviderGetArgs> OauthProviders
        {
            get => _oauthProviders ?? (_oauthProviders = new InputList<Inputs.OauthSettingsOauthProviderGetArgs>());
            set => _oauthProviders = value;
        }

        /// <summary>
        /// Enable the creation of local Artifactory users.  Default value is `false`.
        /// </summary>
        [Input("persistUsers")]
        public Input<bool>? PersistUsers { get; set; }

        public OauthSettingsState()
        {
        }
        public static new OauthSettingsState Empty => new OauthSettingsState();
    }
}
