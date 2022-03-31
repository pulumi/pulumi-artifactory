module github.com/pulumi/pulumi-artifactory/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
)

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1 // indirect
	// When this is merged, we can use a canonical module path: https://github.com/jfrog/terraform-provider-artifactory/pull/374
	github.com/jfrog/terraform-provider-artifactory/v2 v2.25.1-0.20220329001635-7669507f1a2a
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.27.0
)
