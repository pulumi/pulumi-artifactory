// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class GetFederatedDockerRepositoryMemberResult
    {
        public readonly bool Enabled;
        public readonly string Url;

        [OutputConstructor]
        private GetFederatedDockerRepositoryMemberResult(
            bool enabled,

            string url)
        {
            Enabled = enabled;
            Url = url;
        }
    }
}