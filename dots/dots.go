package dots

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

type Dots struct {
	// Directory is where the doftiles and Dots configuration is stored.
	Directory string

	// Base directory. Most likely the user's home directory
	Base string
}

func (d *Dots) Init() error {
	if err := os.MkdirAll(d.Directory, 0o755); err != nil {
		return err
	}

	return nil
}

func (d *Dots) Add(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return errors.New("only regular files are supported")
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	rel, err := filepath.Rel(d.Base, abs)
	if err != nil {
		return err
	}

	targetPath := filepath.Join(d.Directory, rel)
	targetDir := filepath.Dir(targetPath)

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		if err := os.MkdirAll(targetDir, 0o755); err != nil {
			return err
		}
	}

	slog.Info("Copying file", "source", abs, "target", targetPath)
	if err := copyFile(abs, targetPath); err != nil {
		return err
	}

	return nil
}

func copyFile(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return errors.New("src must be file")
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}

func (d *Dots) Refresh() error {
	return errors.New("not implemented")
}
