package require

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/nginx"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

const indexFile = "index.html"

type BuildPlanMetadata struct {
	Version       string `toml:"version,omitempty"`
	VersionSource string `toml:"version-source,omitempty"`
	Launch        bool   `toml:"launch"`
}

func Detect(logger scribe.Emitter) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		if _, err := WebRoot(context.WorkingDir); err != nil {
			return packit.DetectResult{}, err
		}

		result := packit.DetectResult{
			Plan: packit.BuildPlan{
				Requires: []packit.BuildPlanRequirement{
					{
						Name:     nginx.NGINX,
						Metadata: BuildPlanMetadata{Launch: true},
					},
				},
			},
		}

		return result, nil
	}
}

func WebRoot(workingDir string) (string, error) {
	paths := []string{"./", "./public"}
	publicDir := ""
	for _, path := range paths {
		f, err := os.Stat(filepath.Join(workingDir, path, indexFile))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return "", packit.Fail.WithMessage("failed to stat: %w", err)
		}
		if f.IsDir() {
			continue
		}

		publicDir = path
		break
	}

	if len(publicDir) == 0 {
		return "", packit.Fail.WithMessage("no index.html found")
	}

	return publicDir, nil
}
