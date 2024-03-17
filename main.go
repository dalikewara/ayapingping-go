// This file is not part of the project structure. It is just a generator for `ayapingping-go` (this Go module) and will
// not be generated when creating a new project.

package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const name = "AyaPingPing (Go)"
const version = "v4.4.5"
const language = "Golang"
const pathSeparator = string(os.PathSeparator)

func main() {
	var command, value, sourcePrefix, source, runtimeDir string

	argsLen := len(os.Args)

	if argsLen >= 2 {
		command = os.Args[1]
	}
	if argsLen >= 3 {
		value = os.Args[2]
	}
	if argsLen >= 4 {
		sourcePrefix = os.Args[3]
	}
	if argsLen >= 5 {
		source = os.Args[4]
	}

	runtimeDir, err := getRuntimeDir()
	if err != nil {
		panic(err)
	}

	_ = os.Chmod(runtimeDir+pathSeparator+"main_v4.sh", 0777)
	_ = os.Chmod(runtimeDir+pathSeparator+"main_v4_latest.sh", 0777)
	_ = filepath.Walk(runtimeDir+pathSeparator+"_base_structure", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(path, 0777)
	})

	cmd := exec.Command(runtimeDir+pathSeparator+"main_v4.sh", version, language, command, value, sourcePrefix, source)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	_ = cmd.Run()
}

func getRuntimeDir() (string, error) {
	_, runtimeFilepath, _, runtimeOk := runtime.Caller(0)
	if !runtimeOk {
		return "", errors.New("no package runtime found")
	}

	return filepath.Dir(runtimeFilepath), nil
}
