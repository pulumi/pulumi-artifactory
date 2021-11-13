//go:build nodejs || all
// +build nodejs all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/artifactory",
		},
	})

	return baseJS
}
func TestAccLocalRepositoryTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "local-repository", "ts"),
		})

	integration.ProgramTest(t, &test)
}

