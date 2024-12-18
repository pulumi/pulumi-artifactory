// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetBuildActionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<Inputs.PermissionTargetBuildActionsGroupArgs>? _groups;

        /// <summary>
        /// Groups this permission applies for.
        /// </summary>
        public InputList<Inputs.PermissionTargetBuildActionsGroupArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.PermissionTargetBuildActionsGroupArgs>());
            set => _groups = value;
        }

        [Input("users")]
        private InputList<Inputs.PermissionTargetBuildActionsUserArgs>? _users;

        /// <summary>
        /// Users this permission target applies for.
        /// </summary>
        public InputList<Inputs.PermissionTargetBuildActionsUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.PermissionTargetBuildActionsUserArgs>());
            set => _users = value;
        }

        public PermissionTargetBuildActionsArgs()
        {
        }
        public static new PermissionTargetBuildActionsArgs Empty => new PermissionTargetBuildActionsArgs();
    }
}