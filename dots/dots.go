package dots

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
)

type Dots struct {
	// Directory is where the doftiles and Dots configuration is stored.
	Directory string

	// RelativeTo directory. Most likely the user's home directory
	RelativeTo string
}

func (d *Dots) repoAndWorktree() (*git.Repository, *git.Worktree, error) {
	r, err := git.PlainOpen(d.Directory)
	if err != nil {
		return nil, nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, nil, err
	}

	return r, w, nil
}

func (d *Dots) Status() error {
	_, w, err := d.repoAndWorktree()
	if err != nil {
		return err
	}

	s, err := w.Status()
	if err != nil {
		return err
	}

	for k, v := range s {
		switch {
		case v.Staging == git.Untracked:
			fmt.Printf("Uncommited new file: %s\n", k)
		case v.Staging == git.Added:
			fmt.Printf("Uncommited file: %s\n", k)
		case v.Staging == git.Deleted:
			fmt.Printf("Uncommited delete: %s\n", k)
		case v.Staging == git.Modified:
			fmt.Printf("Uncommited modification: %s\n", k)
		case v.Worktree == git.Untracked:
			fmt.Printf("Uncommited && unstaged new file: %s\n", k)
		case v.Worktree == git.Added:
			fmt.Printf("Uncommited && unstaged file: %s\n", k)
		case v.Worktree == git.Deleted:
			fmt.Printf("Uncommited && unstaged delete: %s\n", k)
		case v.Worktree == git.Modified:
			fmt.Printf("Uncommited && unstaged modification: %s\n", k)

		default:
			fmt.Println(k, v)
		}
	}

	return nil
}

func (d *Dots) Commit(glob, msg string) error {
	r, w, err := d.repoAndWorktree()
	if err != nil {
		return err
	}

	if err := w.AddGlob(glob); err != nil {
		return err
	}

	s, err := w.Status()
	if err != nil {
		return err
	}

	hasChanges := false
	for _, f := range s {
		if f.Staging != git.Unmodified {
			hasChanges = true
			break
		}
	}

	if hasChanges {
		h, err := w.Commit(msg, &git.CommitOptions{AllowEmptyCommits: false})
		if err != nil {
			return err
		}
		slog.Info("Created commit", "hash", h)
	} else {
		slog.Info("Skipping commit as no files have changed")
	}

	err = r.Push(&git.PushOptions{Progress: os.Stdout})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return err
	}

	return nil
}

func (d *Dots) Init() error {
	if err := os.MkdirAll(d.Directory, 0o755); err != nil {
		return err
	}

	return nil
}

func (d *Dots) Apply() error {
	if err := filepath.WalkDir(d.Directory, func(path string, e fs.DirEntry, err error) error {
		relative, err := filepath.Rel(d.Directory, path)
		if err != nil {
			return err
		}

		if strings.HasPrefix(relative, ".git/") {
			// ignore the .git directory
			return nil
		}

		if e.IsDir() {
			return nil
		}

		target := filepath.Join(d.RelativeTo, relative)

		slog.Info("Copying dotfile", "src", path, "target", target)
		return copyFile(path, target)
	}); err != nil {
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

	rel, err := filepath.Rel(d.RelativeTo, abs)
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
