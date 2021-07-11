// This file is not part of the project structure. This file is just a generator
// for `ayapingping-go` (this go module), and will not be generated when created new project.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type input struct {
	text string
	val string
}

type option map[string]*input

var version = "v1.0.0"

var baseProjectName = "ayapingping-go"

var baseLicenseProjectUrl = "<https://github.com/dalikewara/ayapingping-go/>"

var baseModulePath = "github.com/dalikewara/ayapingping-go"

func main() {
	options := initOptions()
	reader := bufio.NewReader(os.Stdin)
	options = runInputFlow(reader, options)
	if options["confirmation"].val == "y" {
		generator(options)
	}
}

func initOptions() option {
	opt := make(option)
	opt["welcomeMessage"] = &input{text: fmt.Sprintf("Welcome to AyaPingPing (Go) %s. Let's create a new project!", version)}
	opt["projectName"] = &input{text: "Project name (ex: my-project)"}
	opt["goModulePath"] = &input{text: "Go module path (ex: my-project, or example.com/user_example/my-project)"}
	opt["summary"] = &input{text: "We'll generate new project based on these options:"}
	opt["confirmation"] = &input{text: "Is that ok?"}
	return opt
}

func readStatementInput(r *bufio.Reader, text string, isWrong bool) string {
	if !isWrong {
		fmt.Print(fmt.Sprintf("%s... ", text))
	} else {
		fmt.Print(fmt.Sprintf("Wrong, %s%s... ", strings.ToLower(text[0:1]), text[1:]))
	}
	val, err := r.ReadString('\n')
	val = strings.TrimSpace(val)
	if err != nil {
		panic(err)
	}
	if val == "" {
		return readStatementInput(r, text, true)
	}
	return val
}

func readQuestionInput(r *bufio.Reader, text string, isWrong bool) string {
	if !isWrong {
		fmt.Print(fmt.Sprintf("%s (y/n)... ", text))
	} else {
		fmt.Print(fmt.Sprintf("Wrong, %s%s (y/n)... ", strings.ToLower(text[0:1]), text[1:]))
	}
	val, err := r.ReadString('\n')
	val = strings.TrimSpace(val)
	if err != nil {
		panic(err)
	}
	if val != "n" && val != "y" {
		return readQuestionInput(r, text, true)
	}
	return val
}

func writeFile(source, dest, replaceModule string) {
	fData, err := os.ReadFile(source)
	if err != nil {
		panic(err)
	}
	if replaceModule != "" {
		fData = bytes.Replace(fData, []byte(baseModulePath), []byte(replaceModule), -1)
	}
	if err = os.WriteFile(dest, fData, 0666); err != nil {
		panic(err)
	}
}

func writeFileInfo(info fs.FileInfo, name, filename, source, dest, replaceModule string) {
	if name == filename && !info.IsDir() {
		writeFile(source, dest, replaceModule)
		fmt.Println(fmt.Sprintf("Create %s... [ok]", dest))
	}
}

func writeFileInfoMatchFilename(info fs.FileInfo, filename, source, dest, replaceModule string) {
	matched, err := regexp.MatchString(filename, source)
	if err != nil {
		panic(err)
	}
	if matched {
		if info.IsDir() {
			if err = os.MkdirAll(dest, os.ModePerm); err != nil {
				panic(err)
			}
		} else {
			writeFile(source, dest, replaceModule)
		}
		fmt.Println(fmt.Sprintf("Create %s... [ok]", dest))
	}
}

func writeFileCustomLicense(name, filename, source, dest, projectName string) {
	if name == filename {
		fData, err := os.ReadFile(source)
		if err != nil {
			panic(err)
		}
		fData = bytes.Replace(fData, []byte("2021"), []byte(strconv.Itoa(time.Now().Year())), -1)
		fData = bytes.Replace(fData, []byte(" " + baseLicenseProjectUrl), []byte(""), -1)
		fData = bytes.Replace(fData, []byte(baseLicenseProjectUrl), []byte(""), -1)
		fData = bytes.Replace(fData, []byte(baseProjectName), []byte(projectName), -1)
		if err = os.WriteFile(dest, fData, 0666); err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Create %s... [ok]", dest))
	}
}

func writeFileEnv(name, filename, source, dest string) {
	if name == filename {
		fData, err := os.ReadFile(source)
		if err != nil {
			panic(err)
		}
		newDestDir, _ := path.Split(dest)
		newDest := path.Join(newDestDir, ".env")
		if err = os.WriteFile(newDest, fData, 0666); err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Create %s... [ok]", newDest))
	}
}

func mkDirInfoFilename(info fs.FileInfo, name, filename, dest string) {
	if name == filename && info.IsDir() {
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Create %s... [ok]", dest))
	}
}

func execGoModInit(projectName, goModulePath string) {
	cmd := exec.Command("go", "mod", "init", goModulePath)
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		fmt.Println("Something went wrong when execute `go mod init`")
		return
	}
	fmt.Println("Execute `go mod init`... [ok]")
}

func execGoModTidy(projectName string) {
	fmt.Println("Executing `go mod tidy`. Please wait...")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		fmt.Println("Something went wrong when execute `go mod tidy`")
		return
	}
	fmt.Println("Execute `go mod tidy`... [ok]")
}

func runInputFlow(r *bufio.Reader, opt option) option {
	fmt.Println(opt["welcomeMessage"].text)
	opt["projectName"].val = readStatementInput(r, opt["projectName"].text, false)
	opt["goModulePath"].val = readStatementInput(r, opt["goModulePath"].text, false)
	fmt.Println(fmt.Sprintf("\n%s\n", opt["summary"].text))
	fmt.Println(fmt.Sprintf("%s `%s`", opt["projectName"].text, opt["projectName"].val))
	fmt.Println(fmt.Sprintf("%s `%s`", opt["goModulePath"].text, opt["goModulePath"].val))
	fmt.Println("")
	opt["confirmation"].val = readQuestionInput(r, opt["confirmation"].text, false)
	return opt
}

func generator(opt option) {
	fmt.Println("")
	_, baseFile, _, runtimeOk := runtime.Caller(0)
	if !runtimeOk {
		panic("No package runtime found")
		return
	}
	baseDir := filepath.Dir(baseFile)
	projectName := opt["projectName"].val
	goModulePath := opt["goModulePath"].val
	if _, err := os.Stat(projectName); !os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("We can't create the project because a folder with name `%s` already exists", projectName))
		return
	}
	if err := os.Mkdir(projectName, os.ModePerm); err != nil {
		panic(err)
	}
	if err := filepath.Walk(baseDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		runtimePath := strings.TrimPrefix(strings.TrimPrefix(path, baseDir), string(os.PathSeparator))
		projectPath := filepath.Join(projectName, runtimePath)
		mkDirInfoFilename(info, "app", runtimePath, projectPath)
		writeFileInfoMatchFilename(info, "app/api", path, projectPath, goModulePath)
		mkDirInfoFilename(info, "domain", runtimePath, projectPath)
		writeFileInfoMatchFilename(info, "domain/user_example", path, projectPath, goModulePath)
		mkDirInfoFilename(info, "database", runtimePath, projectPath)
		writeFileInfoMatchFilename(info, "database/mysql", path, projectPath, goModulePath)
		mkDirInfoFilename(info, "config", runtimePath, projectPath)
		writeFileInfoMatchFilename(info, "config/env", path, projectPath, goModulePath)
		writeFileInfo(info, ".env.example", runtimePath, path, projectPath, goModulePath)
		writeFileEnv(".env.example", runtimePath, path, projectPath)
		writeFileInfo(info, ".gitignore", runtimePath, path, projectPath, goModulePath)
		writeFileCustomLicense("LICENSE", runtimePath, path, projectPath, projectName)
		writeFileInfo(info, "Makefile", runtimePath, path, projectPath, goModulePath)
		writeFileInfo(info, "README.md", runtimePath, path, projectPath, goModulePath)
		return nil
	}); err != nil {
		panic(err)
	}
	execGoModInit(projectName, goModulePath)
	execGoModTidy(projectName)
	fmt.Println("Done.")
}