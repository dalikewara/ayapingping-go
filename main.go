// This file is not part of the project structure. This file is just a generator
// for `ayapingping-go` (this go module), and will not be generated when creating new project.

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
	"strings"
)

type input struct {
	text string
	val  string
}

type option struct {
	welcomeMessage *input
	projectName    *input
	goModulePath   *input
	summary        *input
	confirmation   *input
}

var version = "v1.2.1"
var baseModulePath = "github.com/dalikewara/ayapingping-go"
var opt = &option{
	welcomeMessage: &input{
		text: fmt.Sprintf("Welcome to AyaPingPing (Go) %s. Let's create a new project!", version),
	},
	projectName: &input{
		text: "Project name (ex: my-project)",
	},
	goModulePath: &input{
		text: "Go module path (ex: my-project, or example.com/user_example/my-project)",
	},
	summary: &input{
		text: "We'll generate new project based on these options:",
	},
	confirmation: &input{
		text: "Is that ok?",
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(opt.welcomeMessage.text)
	opt.projectName.val = readStatementInput(reader, opt.projectName.text, false)
	opt.goModulePath.val = readStatementInput(reader, opt.goModulePath.text, false)
	fmt.Println(fmt.Sprintf("\n%s\n", opt.summary.text))
	fmt.Println(fmt.Sprintf("%s `%s`", opt.projectName.text, opt.projectName.val))
	fmt.Println(fmt.Sprintf("%s `%s`", opt.goModulePath.text, opt.goModulePath.val))
	fmt.Println("")
	opt.confirmation.val = readQuestionInput(reader, opt.confirmation.text, false)
	if opt.confirmation.val == "y" {
		generator(opt)
	} else {
		fmt.Println("You just refuse to confirm, so we'll not generate your project")
	}
}

func readStatementInput(reader *bufio.Reader, text string, isWrong bool) string {
	if !isWrong {
		fmt.Print(fmt.Sprintf("%s... ", text))
	} else {
		fmt.Print(fmt.Sprintf("Wrong, %s%s... ", strings.ToLower(text[0:1]), text[1:]))
	}
	val, err := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	if err != nil {
		panic(err)
	}
	if val == "" {
		return readStatementInput(reader, text, true)
	}
	return val
}

func readQuestionInput(reader *bufio.Reader, text string, isWrong bool) string {
	if !isWrong {
		fmt.Print(fmt.Sprintf("%s (y/n)... ", text))
	} else {
		fmt.Print(fmt.Sprintf("Wrong, %s%s (y/n)... ", strings.ToLower(text[0:1]), text[1:]))
	}
	val, err := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	if err != nil {
		panic(err)
	}
	if val != "n" && val != "y" {
		return readQuestionInput(reader, text, true)
	}
	return val
}

func generator(option *option) {
	fmt.Println("")
	_, baseFile, _, runtimeOk := runtime.Caller(0)
	if !runtimeOk {
		panic("No package runtime found")
		return
	}
	baseDir := filepath.Dir(baseFile)
	projectName := option.projectName.val
	goModulePath := option.goModulePath.val
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
		walkCreateDir(info, "infra", runtimePath, projectPath)
		walkCreateFileAndReplaceModule(info, "infra", path, projectPath, "")
		walkCreateDir(info, "src", runtimePath, projectPath)
		walkCreateFileAndReplaceModule(info, "src", path, projectPath, goModulePath)
		walkWriteFile(info, ".env.example", runtimePath, path, projectPath, goModulePath)
		walkWriteEnvFromExample(".env.example", runtimePath, path, projectPath)
		walkWriteFile(info, ".gitignore", runtimePath, path, projectPath, goModulePath)
		walkWriteFile(info, "LICENSE", runtimePath, path, projectPath, goModulePath)
		walkWriteFile(info, "Makefile", runtimePath, path, projectPath, goModulePath)
		walkWriteFile(info, "README.md", runtimePath, path, projectPath, goModulePath)
		return nil
	}); err != nil {
		panic(err)
	}
	execGoModInit(projectName, goModulePath)
	execGoModTidy(projectName)
	execGoModVendor(projectName)
	fmt.Println("Done.")
}

func execGoModInit(projectName, goModulePath string) {
	cmd := exec.Command("go", "mod", "init", goModulePath)
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		fmt.Println("Something went wrong when executing `go mod init`")
		return
	}
	fmt.Println("Execute `go mod init`... [ok]")
}

func execGoModTidy(projectName string) {
	fmt.Println("Executing `go mod tidy`. Please wait...")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		fmt.Println("Something went wrong when executing `go mod tidy`")
		return
	}
	fmt.Println("Execute `go mod tidy`... [ok]")
}

func execGoModVendor(projectName string) {
	fmt.Println("Executing `go mod vendor`. Please wait...")
	cmd := exec.Command("go", "mod", "vendor")
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		fmt.Println("Something went wrong when executing `go mod vendor`")
		return
	}
	fmt.Println("Execute `go mod vendor`... [ok]")
}

func writeFile(sourcePath, destinationPath, newModulePath string) {
	fData, err := os.ReadFile(sourcePath)
	if err != nil {
		panic(err)
	}
	if newModulePath != "" {
		fData = bytes.Replace(fData, []byte(baseModulePath), []byte(newModulePath), -1)
	}
	if err = os.WriteFile(destinationPath, fData, 0666); err != nil {
		panic(err)
	}
}

func walkCreateDir(fileInfo fs.FileInfo, folderName, fileName, destinationPath string) {
	if folderName == fileName && fileInfo.IsDir() {
		if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Create %s... [ok]", destinationPath))
	}
}

func walkCreateFileAndReplaceModule(fileInfo fs.FileInfo, fileName, sourcePath, destinationPath,
	newModulePath string) {
	matched, err := regexp.MatchString(fileName, sourcePath)
	gitKeepMatched, err2 := regexp.MatchString(".gitkeep", sourcePath)
	if err != nil || err2 != nil {
		panic(err)
	}
	if matched {
		if fileInfo.IsDir() {
			if err = os.MkdirAll(destinationPath, os.ModePerm); err != nil {
				panic(err)
			}
		} else if !gitKeepMatched {
			writeFile(sourcePath, destinationPath, newModulePath)
		}
		fmt.Println(fmt.Sprintf("Create %s... [ok]", destinationPath))
	}
}

func walkWriteFile(fileInfo fs.FileInfo, folderName, fileName, sourcePath, destinationPath,
	newModulePath string) {
	if folderName == fileName && !fileInfo.IsDir() {
		writeFile(sourcePath, destinationPath, newModulePath)
		fmt.Println(fmt.Sprintf("Create %s... [ok]", destinationPath))
	}
}

func walkWriteEnvFromExample(folderName, fileName, sourcePath, destinationPath string) {
	if folderName == fileName {
		fData, err := os.ReadFile(sourcePath)
		if err != nil {
			panic(err)
		}
		newDestDir, _ := path.Split(destinationPath)
		newDest := path.Join(newDestDir, ".env")
		if err = os.WriteFile(newDest, fData, 0666); err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("Create %s... [ok]", newDest))
	}
}
