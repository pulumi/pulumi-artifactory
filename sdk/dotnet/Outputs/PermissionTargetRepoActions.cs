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
    public sealed class PermissionTargetRepoActions
    {
        /// <summary>
        /// Groups this permission applies for.
        /// </summary>
        public readonly ImmutableArray<Outputs.PermissionTargetRepoActionsGroup> Groups;
        /// <summary>
        /// Users this permission target applies for.
        /// </summary>
        public readonly ImmutableArray<Outputs.PermissionTargetRepoActionsUser> Users;

        [OutputConstructor]
        private PermissionTargetRepoActions(
            ImmutableArray<Outputs.PermissionTargetRepoActionsGroup> groups,

            ImmutableArray<Outputs.PermissionTargetRepoActionsUser> users)
        {
            Groups = groups;
            Users = users;
        }
    }
}
