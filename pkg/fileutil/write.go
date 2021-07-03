package fileutil

import (
	"io/fs"
	"os"

	"github.com/pkg/errors"
)

const (
	newFilePermDefault fs.FileMode = 0755
)

func WriteFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, newFilePermDefault)
	if err != nil {
		return errors.Wrapf(err, "error creating file: %s", path)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return err
	}
	return nil
}
