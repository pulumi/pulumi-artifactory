// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// This is a access token that can be given to you by your admin under `Identity and Access`. If not set, the 'api_key'
// attribute value will be used.
func GetAccessToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "artifactory:accessToken")
}

// API token. Projects functionality will not work with any auth method other than access tokens
//
// Deprecated: An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
// In September 2022, the option to block the usage/creation of API Keys will be enabled by default, with the option for admins to change it back to enable API Keys.
// In January 2023, API Keys will be deprecated all together and the option to use them will no longer be available.
func GetApiKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "artifactory:apiKey")
}

// Toggle for pre-flight checking of Artifactory Pro and Enterprise license. Default to `true`.
func GetCheckLicense(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "artifactory:checkLicense")
	if err == nil {
		return v
	}
	return false
}
func GetUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "artifactory:url")
}
