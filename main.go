// This file is not part of the project structure. It is just a generator for `ayapingping-go` (this Go module) and will
// not be generated when creating a new project.

package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const name = "AyaPingPing (Go)"
const version = "v4.5.0"
const language = "Golang"
const generatorUrl = "https://raw.githubusercontent.com/dalikewara/ayapingping-sh/master/main_v4.sh"
const generatorFile = "main.sh"
const generatorFileTmp = "main_tmp.sh"
const baseStructureDir = "_base_structure"
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

	if err = checkGenerator(runtimeDir); err != nil {
		panic(err)
	}

	cmd := exec.Command(runtimeDir+pathSeparator+generatorFile, version, language, command, value, sourcePrefix, source)

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

func syncGenerator(runtimeDir string) {
	response, err := http.Get(generatorUrl)
	if err != nil {
		return
	}
	defer response.Body.Close()

	file, err := os.Create(runtimeDir + pathSeparator + generatorFileTmp)
	if err != nil {
		return
	}
	defer file.Close()

	if _, err = io.Copy(file, response.Body); err != nil {
		return
	}

	if !isFileValidSH(runtimeDir + pathSeparator + generatorFileTmp) {
		return
	}

	fileData, err := os.ReadFile(runtimeDir + pathSeparator + generatorFileTmp)
	if err != nil {
		return
	}

	if err = os.WriteFile(runtimeDir+pathSeparator+generatorFile, fileData, 0777); err != nil {
		return
	}

	return
}

func checkGenerator(runtimeDir string) error {
	chmod(runtimeDir)
	syncGenerator(runtimeDir)

	if !isFile(runtimeDir + pathSeparator + generatorFile) {
		return errors.New("no generator found, please connect to the internet and run the command again to synchronize")
	}

	if !isFileValidSH(runtimeDir + pathSeparator + generatorFile) {
		return errors.New("invalid generator file, please connect to the internet and run the command again to synchronize")
	}

	return nil
}

func chmod(runtimeDir string) {
	_ = os.Chmod(runtimeDir+pathSeparator+generatorFile, 0777)
	_ = os.Chmod(runtimeDir+pathSeparator+generatorFileTmp, 0777)
	_ = filepath.Walk(runtimeDir+pathSeparator+baseStructureDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(path, 0777)
	})
}

func isFile(path string) bool {
	if fileInfo, err := os.Stat(path); err == nil && !fileInfo.IsDir() {
		return true
	}

	return false
}

func isFileValidSH(path string) bool {
	content, err := os.ReadFile(path)
	if err != nil {
		return false
	}

	if len(content) >= 9 && string(content[:9]) == "#!/bin/sh" {
		return true
	}

	return false
}
