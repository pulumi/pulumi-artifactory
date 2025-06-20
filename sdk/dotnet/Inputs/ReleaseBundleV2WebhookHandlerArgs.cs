// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class ReleaseBundleV2WebhookHandlerArgs : global::Pulumi.ResourceArgs
    {
        [Input("customHttpHeaders")]
        private InputMap<string>? _customHttpHeaders;

        /// <summary>
        /// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
        /// </summary>
        public InputMap<string> CustomHttpHeaders
        {
            get => _customHttpHeaders ?? (_customHttpHeaders = new InputMap<string>());
            set => _customHttpHeaders = value;
        }

        /// <summary>
        /// Proxy key from Artifactory UI (Administration &gt; Proxies &gt; Configuration).
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// Secret authentication token that will be sent to the configured URL. The value will be sent as `x-jfrog-event-auth` header.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// When set to `true`, the secret will be used to sign the event payload, allowing the target to validate that the payload content has not been changed and will not be passed as part of the event. If left unset or set to `false`, the secret is passed through the `X-JFrog-Event-Auth` HTTP header.
        /// </summary>
        [Input("useSecretForSigning")]
        public Input<bool>? UseSecretForSigning { get; set; }

        public ReleaseBundleV2WebhookHandlerArgs()
        {
        }
        public static new ReleaseBundleV2WebhookHandlerArgs Empty => new ReleaseBundleV2WebhookHandlerArgs();
    }
}
