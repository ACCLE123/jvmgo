package classpath

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(entry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, entry, nil
		}
	}
	return nil, nil, fmt.Errorf("class not found %s", className)
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}
