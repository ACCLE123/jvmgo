package main

import (
	"fmt"
	"jvmgo/classpath"
	"strings"
)

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:% X\n", classData)
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		fmt.Println("xjre " + cmd.XjreOption)
		fmt.Println("classpath " + cmd.cpOption)
		fmt.Println("class " + cmd.class)
		startJVM(cmd)
	}
}
