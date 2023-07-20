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
    /// Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
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
    ///     var distribution_custom_webhook = new Artifactory.DistributionCustomWebhook("distribution-custom-webhook", new()
    ///     {
    ///         Criteria = new Artifactory.Inputs.DistributionCustomWebhookCriteriaArgs
    ///         {
    ///             AnyReleaseBundle = false,
    ///             ExcludePatterns = new[]
    ///             {
    ///                 "bar/**",
    ///             },
    ///             IncludePatterns = new[]
    ///             {
    ///                 "foo/**",
    ///             },
    ///             RegisteredReleaseBundleNames = new[]
    ///             {
    ///                 "bundle-name",
    ///             },
    ///         },
    ///         EventTypes = new[]
    ///         {
    ///             "distribute_started",
    ///             "distribute_completed",
    ///             "distribute_aborted",
    ///             "distribute_failed",
    ///             "delete_started",
    ///             "delete_completed",
    ///             "delete_failed",
    ///         },
    ///         Handlers = new[]
    ///         {
    ///             new Artifactory.Inputs.DistributionCustomWebhookHandlerArgs
    ///             {
    ///                 HttpHeaders = 
    ///                 {
    ///                     { "headerName1", "value1" },
    ///                     { "headerName2", "value2" },
    ///                 },
    ///                 Payload = "{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }",
    ///                 Secrets = 
    ///                 {
    ///                     { "secretName1", "value1" },
    ///                     { "secretName2", "value2" },
    ///                 },
    ///                 Url = "https://tempurl.org",
    ///             },
    ///         },
    ///         Key = "distribution-custom-webhook",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/distributionCustomWebhook:DistributionCustomWebhook")]
    public partial class DistributionCustomWebhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.DistributionCustomWebhookCriteria> Criteria { get; private set; } = null!;

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Status of webhook. Default to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distribute_started`, `distribute_completed`, `distribute_aborted`, `distribute_failed, `delete_started`, `delete_completed`, `delete_failed`
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// At least one is required.
        /// </summary>
        [Output("handlers")]
        public Output<ImmutableArray<Outputs.DistributionCustomWebhookHandler>> Handlers { get; private set; } = null!;

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;


        /// <summary>
        /// Create a DistributionCustomWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DistributionCustomWebhook(string name, DistributionCustomWebhookArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/distributionCustomWebhook:DistributionCustomWebhook", name, args ?? new DistributionCustomWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DistributionCustomWebhook(string name, Input<string> id, DistributionCustomWebhookState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/distributionCustomWebhook:DistributionCustomWebhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DistributionCustomWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DistributionCustomWebhook Get(string name, Input<string> id, DistributionCustomWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new DistributionCustomWebhook(name, id, state, options);
        }
    }

    public sealed class DistributionCustomWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.DistributionCustomWebhookCriteriaArgs> Criteria { get; set; } = null!;

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes", required: true)]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distribute_started`, `distribute_completed`, `distribute_aborted`, `distribute_failed, `delete_started`, `delete_completed`, `delete_failed`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers", required: true)]
        private InputList<Inputs.DistributionCustomWebhookHandlerArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.DistributionCustomWebhookHandlerArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.DistributionCustomWebhookHandlerArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public DistributionCustomWebhookArgs()
        {
        }
        public static new DistributionCustomWebhookArgs Empty => new DistributionCustomWebhookArgs();
    }

    public sealed class DistributionCustomWebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies where the webhook will be applied on which repositories.
        /// </summary>
        [Input("criteria")]
        public Input<Inputs.DistributionCustomWebhookCriteriaGetArgs>? Criteria { get; set; }

        /// <summary>
        /// Webhook description. Max length 1000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of webhook. Default to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventTypes")]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distribute_started`, `distribute_completed`, `distribute_aborted`, `distribute_failed, `delete_started`, `delete_completed`, `delete_failed`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers")]
        private InputList<Inputs.DistributionCustomWebhookHandlerGetArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.DistributionCustomWebhookHandlerGetArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.DistributionCustomWebhookHandlerGetArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public DistributionCustomWebhookState()
        {
        }
        public static new DistributionCustomWebhookState Empty => new DistributionCustomWebhookState();
    }
}