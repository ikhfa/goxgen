package main

import (
	"fmt"
	"os"

	"github.com/ikhfa/goxgen/utils"
)

const xgencFile = "xgenc.go"

func main() {
	if isFileExists(xgencFile) {
		fmt.Printf("run %s\n", xgencFile)
		err := utils.ExecCommand("./", "go", "run", "-mod=mod", "./"+xgencFile)
		if err != nil {
			fmt.Printf("run %s failed: %s\n", xgencFile, err)
		}
	} else {
		fmt.Printf("file %s not exists\n", xgencFile)
	}
}

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
