package storagemigration

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Commit finalises the migration by deleting leftover data.
func Commit(root string) error {
	logrus.WithField("storage_root", root).Info("committing changes")

	// remove aufs layer data
	err := removeDirIfExists(aufsRoot(root))
	if err != nil {
		return err
	}

	// remove images
	aufsImageDir := filepath.Join(root, "image", "aufs")
	err = removeDirIfExists(aufsImageDir)
	if err != nil {
		return err
	}

	return nil
}
