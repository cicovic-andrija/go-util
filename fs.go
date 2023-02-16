package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func FileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return !stat.IsDir(), nil
}

func DirectoryExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

func OpenFile(path string) (*os.File, error) {
	if _, statErr := os.Stat(path); errors.Is(statErr, fs.ErrNotExist) {
		return nil, fs.ErrNotExist
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %q: %v", path, err)
	}

	return file, nil
}
