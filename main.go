// This file is not part of the project structure. It is just a generator for `ayapingping-go` (this Go module) and will not be generated when creating a new project.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const name = "AyaPingPing (Go)"
const version = "v4.0.0"
const pathSeparator = string(os.PathSeparator)

type feature struct {
	name              string
	path              string
	projectPath       string
	goModule          string
	featureDependency *featureDependency
}

type features []*feature

type featureDependency struct {
	Domains  []string `json:"domains"`
	Features []string `json:"features"`
	Commons  []string `json:"commons"`
}

func main() {
	fmt.Println(name + " " + version)

	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "addFeature" {
			if err := addFeature(os.Args[i+1:]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return
		}
	}

	if err := createNewProject(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createNewProject() error {
	reader := bufio.NewReader(os.Stdin)

	projectName, err := readInput(reader, "Enter project name (ex: my-project)... ", false)
	if err != nil {
		return err
	}

	goModule, err := readInput(reader, "Enter go module (ex: my-project, or example.com/user_example/my-project)... ", false)
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Project name: " + projectName)
	fmt.Println("Go module: " + goModule)
	fmt.Println("")

	confirmation, err := readInput(reader, "Type `y` to confirm, otherwise will abort the process... ", false)
	if err != nil {
		return err
	}

	if confirmation != "y" {
		return errors.New("process aborted")
	}

	fmt.Println("")

	if projectExist := checkFileOrDirExist(projectName); projectExist {
		return errors.New("project `" + projectName + "` already exist, so aborted")
	}

	runtimeDir, err := getRuntimeDir()
	if err != nil {
		return err
	}

	currentGoModule, err := getGoModuleFromProject(runtimeDir)
	if err != nil {
		_ = removeProject(projectName)

		return err
	}
	currentGoModule += "/structure"

	runtimeDirContents, err := getRuntimeDirContents()
	if err != nil {
		return err
	}

	if err = createDir(projectName); err != nil {
		return err
	}

	for _, path := range runtimeDirContents {
		pathSplit := strings.Split(path, pathSeparator+"structure"+pathSeparator)
		lenPathSplit := len(pathSplit)

		if isFile(path) && lenPathSplit == 1 {
			destinationPath := projectName + pathSeparator + filepath.Base(pathSplit[0])

			if err = copyFile(path, destinationPath, currentGoModule, goModule); err != nil {
				_ = removeProject(projectName)

				return err
			}
		}
		if isDir(path) && lenPathSplit >= 2 {
			if err = copyDir(runtimeDir, path, projectName+pathSeparator+pathSplit[lenPathSplit-1], currentGoModule, goModule); err != nil {
				_ = removeProject(projectName)

				return err
			}
		}
	}

	if err = execGoModInit(projectName, goModule); err != nil {
		_ = removeProject(projectName)

		return err
	}

	if err = execGoModTidy(projectName); err != nil {
		_ = removeProject(projectName)

		return err
	}

	if err = execGoModVendor(projectName); err != nil {
		_ = removeProject(projectName)

		return err
	}

	fmt.Println("")
	fmt.Println("Project created!")

	return nil
}

func addFeature(args []string) error {
	if len(args) != 3 || args[1] != "from" {
		return errors.New("invalid `addFeature` arguments, please follow: addFeature [feature1,feature2,...] from [/local/project or https://example.com/user/project.git or git@example.com:user/project.git]")
	}
	if len(args[0]) < 1 || args[0] == " " {
		return errors.New("feature name cannot be empty or blank space")
	}
	if len(args[2]) < 1 || args[2] == " " {
		return errors.New("argument `from` cannot be empty or blank space")
	}

	fromPath := args[2]

	if isFromGit(fromPath) {

	}

	fmt.Println("Checking features... [RUNNING]")

	feats := collectFeaturesFromArgument(args[0], fromPath)

	fmt.Println("Checking features... [OK]")

	if len(feats) > 0 {
		fmt.Println("")

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Features to be added from `" + fromPath + "`: ")

		for _, feat := range feats {
			if feat == nil {
				continue
			}

			fmt.Println("----------------------------------------------------")
			fmt.Println("Name: " + feat.name)

			if feat.featureDependency == nil {
				fmt.Println("Dependency: none (possible missing packages)")
			} else {
				fmt.Println(fmt.Sprintf("Dependency: Domains: %v", feat.featureDependency.Domains))
				fmt.Println(fmt.Sprintf("            Commons: %v", feat.featureDependency.Commons))
				fmt.Println(fmt.Sprintf("            Other Features: %v", feat.featureDependency.Features))
			}
		}

		fmt.Println("")

		confirmation, err := readInput(reader, "Type `y` to confirm, otherwise will abort the process... ", false)
		if err != nil {
			return err
		}

		if confirmation != "y" {
			return errors.New("process aborted")
		}

		fmt.Println("")
	} else {
		return errors.New("no feature found to be added")
	}

	fmt.Println("Adding features... [RUNNING]")

	currentGoModule, err := getGoModuleFromProject("")
	if err != nil {
		return errors.New("no go module found from the current project")
	}

	var totalAdded int

	for _, feat := range feats {
		if feat == nil {
			continue
		}

		if err = copyDir(fromPath, feat.path, "features"+pathSeparator+feat.name, feat.goModule, currentGoModule); err != nil {
			continue
		}

		totalAdded += 1

		if feat.featureDependency != nil {
			for _, domainFilepath := range feat.featureDependency.Domains {
				if err = copyFile(feat.projectPath+pathSeparator+"domain"+pathSeparator+domainFilepath, "domain"+pathSeparator+domainFilepath, feat.goModule, currentGoModule); err != nil {
					continue
				}
			}

			for _, commonFilepath := range feat.featureDependency.Commons {
				if err = copyFile(feat.projectPath+pathSeparator+"commons"+pathSeparator+commonFilepath, "commons"+pathSeparator+commonFilepath, feat.goModule, currentGoModule); err != nil {
					continue
				}
			}
		}
	}

	fmt.Println("Adding features... [OK]")

	if err = execGoModTidy(""); err != nil {
		return err
	}

	if err = execGoModVendor(""); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%v feature(s) added", totalAdded))

	return nil
}

func readInput(reader *bufio.Reader, text string, isWrong bool) (string, error) {
	if isWrong {
		fmt.Print(fmt.Sprintf("Wrong, %s", strings.ToLower(text)))
	} else {
		fmt.Print(text)
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return readInput(reader, text, true)
	}

	return input, nil
}

func getRuntimeDir() (string, error) {
	_, runtimeFilepath, _, runtimeOk := runtime.Caller(0)
	if !runtimeOk {
		return "", errors.New("no package runtime found")
	}

	return filepath.Dir(runtimeFilepath), nil
}

func getRuntimeDirContents() ([]string, error) {
	var list []string

	runtimeDir, err := getRuntimeDir()
	if err != nil {
		return nil, err
	}
	lenRuntimeDir := len(runtimeDir)

	if err = filepath.Walk(runtimeDir, func(path string, info os.FileInfo, err error) error {
		pathCut := path[lenRuntimeDir:]

		if pathCut == pathSeparator+"LICENSE" || pathCut == pathSeparator+"README.md" || pathCut == pathSeparator+".gitignore" || pathCut == pathSeparator+"structure" {
			list = append(list, path)

			return nil
		}
		if len(pathCut) >= 11 && pathCut[:11] == pathSeparator+"structure"+pathSeparator {
			list = append(list, path)

			return nil
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return list, nil
}

func getDirContents(dirPath string) ([]string, error) {
	var list []string

	if err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		list = append(list, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return list, nil
}

func checkFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

func isDir(path string) bool {
	if fileInfo, err := os.Stat(path); err == nil && fileInfo.IsDir() {
		return true
	}

	return false
}

func isFile(path string) bool {
	if fileInfo, err := os.Stat(path); err == nil && !fileInfo.IsDir() {
		return true
	}

	return false
}

func createDir(dirPath string) error {
	fmt.Println("Creating `" + dirPath + "`... [RUNNING]")

	if checkFileOrDirExist(dirPath) {
		fmt.Println("         `" + dirPath + "`... [EXIST], so skipping...")

		return nil
	}

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println("         `" + dirPath + "`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	fmt.Println("         `" + dirPath + "`... [OK]")

	return nil
}

func copyFile(sourcePath string, destinationPath string, sourceGoModule string, destinationGoModule string) error {
	fmt.Println("Copying `" + sourcePath + "`... [RUNNING]")

	if err := createDir(filepath.Dir(destinationPath)); err != nil {
		fmt.Println("        `" + sourcePath + "`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	if checkFileOrDirExist(destinationPath) {
		fmt.Println("        `" + sourcePath + "`... [EXIST], so skipping...")

		return nil
	}

	fileData, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Println("        `" + sourcePath + "`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	if err = os.WriteFile(destinationPath, fileData, os.ModePerm); err != nil {
		fmt.Println("        `" + sourcePath + "`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	if err = replaceGoModule(destinationPath, sourceGoModule, destinationGoModule); err != nil {
		fmt.Println("        `" + sourcePath + "`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	fmt.Println("        `" + sourcePath + "`... [OK]")

	return nil
}

func copyDir(sourceProjectBasePath string, sourcePath string, destinationPath string, sourceGoModule string, destinationGoModule string) error {
	if err := createDir(destinationPath); err != nil {
		return err
	}

	dirContents, err := getDirContents(sourcePath)
	if err != nil {
		return err
	}

	for _, path := range dirContents {
		subDestinationPath := strings.ReplaceAll(path, sourceProjectBasePath, "")

		if len(subDestinationPath) > 0 && subDestinationPath[:1] == pathSeparator {
			subDestinationPath = subDestinationPath[1:]
		}

		if isDir(path) {
			if err = createDir(subDestinationPath); err != nil {
				return err
			}
		}

		if isFile(path) {
			if err = copyFile(path, subDestinationPath, sourceGoModule, destinationGoModule); err != nil {
				return err
			}
		}
	}

	return nil
}

func getGoModuleFromProject(projectPath string) (string, error) {
	cmd := exec.Command("go", "list", "-m")
	cmd.Dir = projectPath
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func replaceGoModule(filepath string, oldGoModule string, newModule string) error {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	fileData = bytes.Replace(fileData, []byte(oldGoModule), []byte(newModule), -1)
	if err = os.WriteFile(filepath, fileData, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func removeProject(dirPath string) error {
	return os.RemoveAll(dirPath)
}

func execGoModInit(projectDir string, goModule string) error {
	fmt.Println("Executing `go mod init`... [RUNNING]")

	cmd := exec.Command("go", "mod", "init", goModule)
	cmd.Dir = projectDir

	if err := cmd.Run(); err != nil {
		fmt.Println("          `go mod init`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	fmt.Println("          `go mod init`... [OK]")

	return nil
}

func execGoModTidy(projectDir string) error {
	fmt.Println("Executing `go mod tidy`... [RUNNING]")

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir

	if err := cmd.Run(); err != nil {
		fmt.Println("          `go mod tidy`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	fmt.Println("          `go mod tidy`... [OK]")

	return nil
}

func execGoModVendor(projectDir string) error {
	fmt.Println("Executing `go mod vendor`... [RUNNING]")

	cmd := exec.Command("go", "mod", "vendor")
	cmd.Dir = projectDir

	if err := cmd.Run(); err != nil {
		fmt.Println("          `go mod vendor`... [ERROR], so skipping...")
		fmt.Println(err)

		return err
	}

	fmt.Println("          `go mod vendor`... [OK]")

	return nil
}

func isFromGit(from string) bool {
	if strings.Contains(from, ".git") && (strings.Contains(from, "https://") || strings.Contains(from, "http://") || strings.Contains(from, "git@")) {
		return true
	}

	return false
}

func readDependencyFile(filepath string) (*featureDependency, error) {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var dependency featureDependency

	if err = json.Unmarshal(fileData, &dependency); err != nil {
		return nil, err
	}

	return &dependency, nil
}

func collectFeaturesFromArgument(feaArg string, fromPath string) features {
	fromGoModule, errGoModule := getGoModuleFromProject(fromPath)
	if errGoModule != nil {
		fmt.Println("no go module found from `" + fromPath + "`, so skipping...")
		return nil
	}

	var feats features

	for _, fea := range strings.Split(strings.ReplaceAll(feaArg, " ", ""), ",") {
		feaPath := fromPath + pathSeparator + "features" + pathSeparator + fea

		if !checkFileOrDirExist(feaPath) {
			fmt.Println("`" + fea + "` feature directory not exist, so skipping...")
			continue
		}

		var featureDependencyData *featureDependency

		featureDependencyFilepath := feaPath + pathSeparator + "dependency.json"

		if !checkFileOrDirExist(featureDependencyFilepath) {
			fmt.Println("WARNING, `" + fea + "` feature `dependency.json` file not exist")
		} else {
			data, err := readDependencyFile(featureDependencyFilepath)
			if err != nil {
				fmt.Println("WARNING, error when parsing `" + fea + "` feature `dependency.json` file: " + err.Error())
			} else {
				featureDependencyData = data
			}
		}

		feat := &feature{
			name:              fea,
			path:              feaPath,
			projectPath:       fromPath,
			featureDependency: featureDependencyData,
			goModule:          fromGoModule,
		}

		var exist bool

		for _, f := range feats {
			if f == nil {
				continue
			}
			if f.name == feat.name {
				exist = true
				break
			}
		}

		if !exist {
			feats = append(feats, feat)
		}

		if feat.featureDependency != nil {
			for _, otherFea := range feat.featureDependency.Features {
				otherFeats := collectFeaturesFromArgument(otherFea, fromPath)
				for _, of := range otherFeats {
					if of == nil {
						continue
					}

					exist = false

					for _, f := range feats {
						if f == nil {
							continue
						}
						if f.name == of.name {
							exist = true
							break
						}
					}

					if !exist {
						feats = append(feats, of)
					}
				}
			}
		}
	}

	return feats
}
