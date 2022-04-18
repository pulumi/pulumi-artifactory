[![Actions Status](https://github.com/pulumi/pulumi-artifactory/workflows/main/badge.svg)](https://github.com/pulumi/pulumi-artifactory/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fartifactory.svg)](https://www.npmjs.com/package/@pulumi/artifactory)
[![Python version](https://badge.fury.io/py/pulumi-artifactory.svg)](https://pypi.org/project/pulumi-artifactory)
[![NuGet version](https://badge.fury.io/nu/pulumi.artifactory.svg)](https://badge.fury.io/nu/pulumi.artifactory)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-artifactory/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-artifactory/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-artifactory/blob/main/LICENSE)

# Artifactory Resource Provider

The Artifactory Resource Provider lets you manage Artifactory resources.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/artifactory

or `yarn`:

    $ yarn add @pulumi/artifactory

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_artifactory

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-artifactory/sdk/v6

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Artifactory

## Configuration

The following configuration points are available:

- `artifactory:url` - (Required) URL of Artifactory. This can also be sourced from the `ARTIFACTORY_URL` environment variable.
- `artifactory:username` - (Optional) Username for basic auth. Requires password to be set. Conflicts with `apiKey`, 
  and `accessToken`. This can also be sourced from the `ARTIFACTORY_USERNAME` environment variable.
- `artifactory:password` - (Optional) Password for basic auth. Requires username to be set. Conflicts with `apiKey`, 
  and `accessToken`. This can also be sourced from the `ARTIFACTORY_PASSWORD` environment variable.
- `artifactory:apiKey` - (Optional) API key for api auth. Uses `X-JFrog-Art-Api` header. Conflicts with `username`, 
  `password`, and `accessToken`. This can also be sourced from the `ARTIFACTORY_API_KEY` environment variable.
- `artifactory:accessToken` - (Optional) API key for token auth. Uses `Authorization: Bearer` header. For xray 
  functionality, this is the only auth method accepted. Conflicts with `username` and `password`, and `apiKey`. This can
  also be sourced from the `ARTIFACTORY_ACCESS_TOKEN` environment variable.

## Reference

For further information, please visit [the Artifactory provider docs](https://www.pulumi.com/docs/intro/cloud-providers/artifactory)
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/artifactory).
