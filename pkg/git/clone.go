package git

import (
	"path/filepath"

	clog "github.com/cockroachdb/cockroach/util/log"
	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/util"
)

type Clone struct {
	Git
	util.FileSystem
}

// Download downloads the application source code from the GIT repository
// and checkout the Ref specified in the config.
func (c *Clone) Download(config *api.Config) error {
	targetSourceDir := filepath.Join(config.WorkingDir, "upload", "src")

	if c.ValidCloneSpec(config.Source) {

		if len(config.ContextDir) > 0 {
			targetSourceDir = filepath.Join(config.WorkingDir, "upload", "tmp")
		}
		if clog.V(2) {
			clog.Infof("Cloning into %s", targetSourceDir)
		}
		if err := c.Clone(config.Source, targetSourceDir); err != nil {
			if clog.V(1) {
				clog.Infof("Git clone failed: %+v", err)
			}
			return err
		}

		if config.Ref != "" {
			if clog.V(1) {
				clog.Infof("Checking out ref %s", config.Ref)
			}

			if err := c.Checkout(targetSourceDir, config.Ref); err != nil {
				return err
			}
		}

		if len(config.ContextDir) > 0 {
			originalTargetDir := filepath.Join(config.WorkingDir, "upload", "src")
			c.RemoveDirectory(originalTargetDir)
			// we want to copy entire dir contents, thus we need to use dir/. construct
			path := filepath.Join(targetSourceDir, config.ContextDir) + string(filepath.Separator) + "."
			err := c.Copy(path, originalTargetDir)
			if err != nil {
				return err
			}
			c.RemoveDirectory(targetSourceDir)
		}

		return nil
	}

	// we want to copy entire dir contents, thus we need to use dir/. construct
	path := filepath.Join(config.Source, config.ContextDir) + string(filepath.Separator) + "."
	return c.Copy(path, targetSourceDir)
}
