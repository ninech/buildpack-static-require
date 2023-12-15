package require

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/nginx"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
		buffer     *bytes.Buffer
		indexFile  = `
<!DOCTYPE html>
<html></html>`
	)

	it.Before(func() {
		buffer = bytes.NewBuffer(nil)
	})

	context("when an index.html is present", func() {
		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp(t.TempDir(), "working-dir-*")
			Expect(err).NotTo(HaveOccurred())
			err = os.WriteFile(filepath.Join(workingDir, "index.html"), []byte(indexFile), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
		})

		it("requires nginx", func() {
			result, err := Detect(scribe.NewEmitter(buffer))(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(ContainElement(packit.BuildPlanRequirement{
				Name: nginx.NGINX, Metadata: BuildPlanMetadata{Launch: true},
			}))
		})
	})

	context("when an index.html is in the public directory", func() {
		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp(t.TempDir(), "working-dir-*")
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Mkdir(filepath.Join(workingDir, "public"), os.ModePerm)).NotTo(HaveOccurred())
			err = os.WriteFile(filepath.Join(workingDir, "public", "index.html"), []byte(indexFile), os.ModePerm)
			Expect(err).NotTo(HaveOccurred())
		})

		it("requires nginx", func() {
			result, err := Detect(scribe.NewEmitter(buffer))(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(ContainElement(packit.BuildPlanRequirement{
				Name: nginx.NGINX, Metadata: BuildPlanMetadata{Launch: true},
			}))
		})
	})

	context("when no index.html is present", func() {
		it.Before(func() {
		})

		it("fails detection", func() {
			_, err := Detect(scribe.NewEmitter(buffer))(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).To(MatchError(packit.Fail.WithMessage("no index.html found")))
		})
	})

	context("when index.html is a directory", func() {
		it.Before(func() {
			var err error
			workingDir, err = os.MkdirTemp(t.TempDir(), "working-dir-*")
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Mkdir(filepath.Join(workingDir, "index.html"), os.ModePerm)).NotTo(HaveOccurred())
		})

		it("fails detection", func() {
			_, err := Detect(scribe.NewEmitter(buffer))(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).To(MatchError(packit.Fail.WithMessage("no index.html found")))
		})
	})

	context("when webroot is set via env", func() {
		it.Before(func() {
			t.Setenv(webRootEnv, "custom")
		})

		it("requires nginx", func() {
			result, err := Detect(scribe.NewEmitter(buffer))(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan.Requires).To(ContainElement(packit.BuildPlanRequirement{
				Name: nginx.NGINX, Metadata: BuildPlanMetadata{Launch: true},
			}))
		})
	})
}
