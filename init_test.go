package require

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitBuildpackStatic(t *testing.T) {
	suite := spec.New("buildpack-static-require", spec.Report(report.Terminal{}), spec.Parallel())
	suite("Detect", testDetect)
	suite.Run(t)
}
