package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absPath string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(entry.absPath, className)
	data, err := ioutil.ReadFile(fileName)
	return data, entry, err
}
func (entry *DirEntry) String() string {
	return entry.absPath
}
