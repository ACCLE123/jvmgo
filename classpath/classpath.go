package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (c *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := c.userClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	panic("there is no such class: " + className)
}

func (c *ClassPath) String() string {
	return c.userClasspath.String()
}

func (c *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

func (c *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("java home dir not found")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return false
		}
	}
	return true
}
