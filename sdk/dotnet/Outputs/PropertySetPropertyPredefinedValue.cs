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
    public sealed class PropertySetPropertyPredefinedValue
    {
        /// <summary>
        /// Whether the value is selected by default in the UI.
        /// </summary>
        public readonly bool DefaultValue;
        /// <summary>
        /// Property set name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private PropertySetPropertyPredefinedValue(
            bool defaultValue,

            string name)
        {
            DefaultValue = defaultValue;
            Name = name;
        }
    }
}
