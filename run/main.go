package main

import (
	"os"

	require "github.com/ninech/buildpack-static-require"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func main() {
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	packit.Run(require.Detect(logger), require.Build(logger))
}
