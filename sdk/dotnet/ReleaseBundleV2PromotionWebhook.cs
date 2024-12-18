// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
    /// 
    /// ## Example Usage
    /// 
    /// .
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var release_bundle_v2_promotion_webhook = new Artifactory.ReleaseBundleV2PromotionWebhook("release-bundle-v2-promotion-webhook", new()
    ///     {
    ///         Key = "release-bundle-v2-promotion-webhook",
    ///         EventTypes = new[]
    ///         {
    ///             "release_bundle_v2_promotion_completed",
    ///             "release_bundle_v2_promotion_failed",
    ///             "release_bundle_v2_promotion_started",
    ///         },
    ///         Criteria = new Artifactory.Inputs.ReleaseBundleV2PromotionWebhookCriteriaArgs
    ///         {
    ///             SelectedEnvironments = new[]
    ///             {
    ///                 "PROD",
    ///                 "DEV",
    ///             },
    ///         },
    ///         Handlers = new[]
    ///         {
    ///             new Artifactory.Inputs.ReleaseBundleV2PromotionWebhookHandlerArgs
    ///             {
    ///                 Url = "http://tempurl.org/webhook",
    ///                 Secret = "some-secret",
    ///                 Proxy = "proxy-key",
    ///                 CustomHttpHeaders = 
    ///                 {
    ///                     { "header-1", "value-1" },
    ///                     { "header-2", "value-2" },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/releaseBundleV2PromotionWebhook:ReleaseBundleV2PromotionWebhook")]
    public partial class ReleaseBundleV2PromotionWebhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which enviroments.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.ReleaseBundleV2PromotionWebhookCriteria?> Criteria { get; private set; } = null!;

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Status of webhook. Default to `true`
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// At least one is required.
        /// </summary>
        [Output("handlers")]
        public Output<ImmutableArray<Outputs.ReleaseBundleV2PromotionWebhookHandler>> Handlers { get; private set; } = null!;

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;


        /// <summary>
        /// Create a ReleaseBundleV2PromotionWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReleaseBundleV2PromotionWebhook(string name, ReleaseBundleV2PromotionWebhookArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/releaseBundleV2PromotionWebhook:ReleaseBundleV2PromotionWebhook", name, args ?? new ReleaseBundleV2PromotionWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReleaseBundleV2PromotionWebhook(string name, Input<string> id, ReleaseBundleV2PromotionWebhookState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/releaseBundleV2PromotionWebhook:ReleaseBundleV2PromotionWebhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReleaseBundleV2PromotionWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReleaseBundleV2PromotionWebhook Get(string name, Input<string> id, ReleaseBundleV2PromotionWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new ReleaseBundleV2PromotionWebhook(name, id, state, options);
        }
    }

    public sealed class ReleaseBundleV2PromotionWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which enviroments.
        /// </summary>
        [Input("criteria")]
        public Input<Inputs.ReleaseBundleV2PromotionWebhookCriteriaArgs>? Criteria { get; set; }

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to `true`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes", required: true)]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers")]
        private InputList<Inputs.ReleaseBundleV2PromotionWebhookHandlerArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.ReleaseBundleV2PromotionWebhookHandlerArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.ReleaseBundleV2PromotionWebhookHandlerArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public ReleaseBundleV2PromotionWebhookArgs()
        {
        }
        public static new ReleaseBundleV2PromotionWebhookArgs Empty => new ReleaseBundleV2PromotionWebhookArgs();
    }

    public sealed class ReleaseBundleV2PromotionWebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which enviroments.
        /// </summary>
        [Input("criteria")]
        public Input<Inputs.ReleaseBundleV2PromotionWebhookCriteriaGetArgs>? Criteria { get; set; }

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to `true`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes")]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of event triggers for the Webhook. Allow values: `release_bundle_v2_promotion_started`, `release_bundle_v2_promotion_completed`, `release_bundle_v2_promotion_failed`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers")]
        private InputList<Inputs.ReleaseBundleV2PromotionWebhookHandlerGetArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.ReleaseBundleV2PromotionWebhookHandlerGetArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.ReleaseBundleV2PromotionWebhookHandlerGetArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public ReleaseBundleV2PromotionWebhookState()
        {
        }
        public static new ReleaseBundleV2PromotionWebhookState Empty => new ReleaseBundleV2PromotionWebhookState();
    }
}