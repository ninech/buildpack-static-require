package require

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func Build(logger scribe.Emitter) packit.BuildFunc {
	// build is a noop for this buildpack
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		return packit.BuildResult{}, nil
	}
}
