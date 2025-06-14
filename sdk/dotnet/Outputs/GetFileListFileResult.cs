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
    public sealed class GetFileListFileResult
    {
        /// <summary>
        /// Is this a folder
        /// </summary>
        public readonly bool Folder;
        /// <summary>
        /// Last modified time
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// File metadata
        /// </summary>
        public readonly Outputs.GetFileListFileMetadataTimestampsResult MetadataTimestamps;
        /// <summary>
        /// SHA-1 checksum
        /// </summary>
        public readonly string Sha1;
        /// <summary>
        /// SHA-256 checksum
        /// </summary>
        public readonly string Sha2;
        /// <summary>
        /// File size in bytes
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// URL to file
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GetFileListFileResult(
            bool folder,

            string lastModified,

            Outputs.GetFileListFileMetadataTimestampsResult metadataTimestamps,

            string sha1,

            string sha2,

            int size,

            string uri)
        {
            Folder = folder;
            LastModified = lastModified;
            MetadataTimestamps = metadataTimestamps;
            Sha1 = sha1;
            Sha2 = sha2;
            Size = size;
            Uri = uri;
        }
    }
}
