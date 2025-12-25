// Dummy class for go mod tidy to keep the dependencies in go.mod file
package main

import (
	"os"

	"github.com/mikefarah/yq/v4/pkg/yqlib"
)

func main() {
	exitCode := 0
	if yqlib.NewYamlDecoder(yqlib.YamlPreferences{}) == nil {
		exitCode = 1
	}
	if yqlib.NewYamlEncoder(yqlib.YamlPreferences{}) == nil {
		exitCode = 1
	}
	os.Exit(exitCode)
}
