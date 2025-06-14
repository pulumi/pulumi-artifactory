// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class PermissionTargetBuildActions
    {
        /// <summary>
        /// Groups this permission applies for.
        /// </summary>
        public readonly ImmutableArray<Outputs.PermissionTargetBuildActionsGroup> Groups;
        /// <summary>
        /// Users this permission target applies for.
        /// </summary>
        public readonly ImmutableArray<Outputs.PermissionTargetBuildActionsUser> Users;

        [OutputConstructor]
        private PermissionTargetBuildActions(
            ImmutableArray<Outputs.PermissionTargetBuildActionsGroup> groups,

            ImmutableArray<Outputs.PermissionTargetBuildActionsUser> users)
        {
            Groups = groups;
            Users = users;
        }
    }
}
