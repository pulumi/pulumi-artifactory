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
    /// Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
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
    ///     var artifact_lifecycle_custom_webhook = new Artifactory.ArtifactLifecycleCustomWebhook("artifact-lifecycle-custom-webhook", new()
    ///     {
    ///         Key = "artifact-lifecycle-custom-webhook",
    ///         EventTypes = new[]
    ///         {
    ///             "archive",
    ///             "restore",
    ///         },
    ///         Handlers = new[]
    ///         {
    ///             new Artifactory.Inputs.ArtifactLifecycleCustomWebhookHandlerArgs
    ///             {
    ///                 Url = "https://tempurl.org",
    ///                 Method = "POST",
    ///                 Secrets = 
    ///                 {
    ///                     { "secretName1", "value1" },
    ///                     { "secretName2", "value2" },
    ///                 },
    ///                 HttpHeaders = 
    ///                 {
    ///                     { "headerName1", "value1" },
    ///                     { "headerName2", "value2" },
    ///                 },
    ///                 Payload = "{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/artifactLifecycleCustomWebhook:ArtifactLifecycleCustomWebhook")]
    public partial class ArtifactLifecycleCustomWebhook : global::Pulumi.CustomResource
    {
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
        /// List of event triggers for the Webhook. Allow values: `archive`, `restore`
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// At least one is required.
        /// </summary>
        [Output("handlers")]
        public Output<ImmutableArray<Outputs.ArtifactLifecycleCustomWebhookHandler>> Handlers { get; private set; } = null!;

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;


        /// <summary>
        /// Create a ArtifactLifecycleCustomWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ArtifactLifecycleCustomWebhook(string name, ArtifactLifecycleCustomWebhookArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/artifactLifecycleCustomWebhook:ArtifactLifecycleCustomWebhook", name, args ?? new ArtifactLifecycleCustomWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ArtifactLifecycleCustomWebhook(string name, Input<string> id, ArtifactLifecycleCustomWebhookState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/artifactLifecycleCustomWebhook:ArtifactLifecycleCustomWebhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ArtifactLifecycleCustomWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ArtifactLifecycleCustomWebhook Get(string name, Input<string> id, ArtifactLifecycleCustomWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new ArtifactLifecycleCustomWebhook(name, id, state, options);
        }
    }

    public sealed class ArtifactLifecycleCustomWebhookArgs : global::Pulumi.ResourceArgs
    {
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
        /// List of event triggers for the Webhook. Allow values: `archive`, `restore`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers")]
        private InputList<Inputs.ArtifactLifecycleCustomWebhookHandlerArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.ArtifactLifecycleCustomWebhookHandlerArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.ArtifactLifecycleCustomWebhookHandlerArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public ArtifactLifecycleCustomWebhookArgs()
        {
        }
        public static new ArtifactLifecycleCustomWebhookArgs Empty => new ArtifactLifecycleCustomWebhookArgs();
    }

    public sealed class ArtifactLifecycleCustomWebhookState : global::Pulumi.ResourceArgs
    {
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
        /// List of event triggers for the Webhook. Allow values: `archive`, `restore`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        [Input("handlers")]
        private InputList<Inputs.ArtifactLifecycleCustomWebhookHandlerGetArgs>? _handlers;

        /// <summary>
        /// At least one is required.
        /// </summary>
        public InputList<Inputs.ArtifactLifecycleCustomWebhookHandlerGetArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.ArtifactLifecycleCustomWebhookHandlerGetArgs>());
            set => _handlers = value;
        }

        /// <summary>
        /// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public ArtifactLifecycleCustomWebhookState()
        {
        }
        public static new ArtifactLifecycleCustomWebhookState Empty => new ArtifactLifecycleCustomWebhookState();
    }
}
