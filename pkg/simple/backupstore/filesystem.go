package backupstore

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func init() {
	RegisterProvider(&FilesystemStore{})
}

var _ BackupStore = (*FilesystemStore)(nil)

const (
	defaultRootDir = "/opt/kc/backups"
)

type FilesystemStore struct {
	RootDir string `json:"rootDir,omitempty" yaml:"rootDir,omitempty"` // root directory for storing backup files
}

func (fs *FilesystemStore) Type() string {
	return FSStorage
}

func (fs *FilesystemStore) Create() (BackupStore, error) {
	if fs.RootDir == "" {
		fs.RootDir = defaultRootDir
	}

	return fs, nil
}

func (fs *FilesystemStore) Save(ctx context.Context, r io.Reader, fileName string) (err error) {
	defer logProbe(ctx, fmt.Sprintf("save backup to %s", filepath.Join(fs.RootDir, fileName)), err)
	w, err := os.Create(filepath.Join(fs.RootDir, fileName))
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	return w.Sync()
}

func (fs *FilesystemStore) Delete(ctx context.Context, fileName string) (err error) {
	defer logProbe(ctx, fmt.Sprintf("delete backup from %s", filepath.Join(fs.RootDir, fileName)), err)
	err = os.Remove(filepath.Join(fs.RootDir, fileName))
	if err != nil && errors.Is(err, os.ErrNotExist) {
		// The target file is already deleted.
		return nil
	}
	return
}

func (fs *FilesystemStore) Download(ctx context.Context, fileName string, w io.Writer) (err error) {
	defer logProbe(ctx, fmt.Sprintf("download backup from %s", filepath.Join(fs.RootDir, fileName)), err)
	f, err := os.OpenFile(filepath.Join(fs.RootDir, fileName), os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	write := bufio.NewWriter(w)

	_, err = bufio.NewReader(f).WriteTo(write)
	if err != nil {
		return err
	}
	return write.Flush()
}
