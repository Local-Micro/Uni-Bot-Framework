package core

import (
	"os"
	_path "path"
	"path/filepath"
	"runtime"
)

func GetBinaryDictory() (string, error) {
	tmp, err := getTmpDir()
	if err != nil {
		return nil, err
	}
	path, err := byExecutable()
	if err != nil {
		return nil, err
	}
	if strings.Contains(path, tmp) {
		path, err = byCaller()
		if err != nil {
			return nil, err
		}
	}
	return path, nil
}

func byCaller() (string, error) {
	var path string
	_, name, _, ok := runtime.Caller(0)
	if ok {
		path = _path.Dir(name)
	}
	return path, nil
}

func byExecutable() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return nil, err
	}
	path, err = filepath.EvalSymlinks(filepath.Dir(path))
	if err != nil {
		return nil, err
	}
	return path, nil
}

func getTmpDir() (string, error) {
	dir := os.Getenv("TMP")
	if dir == "" {
		dir = os.Getenv("TEMP")
		if dir == "" {
			dir = os.Getenv("TMPDIR")
		}
	}
	return dir, nil
}
