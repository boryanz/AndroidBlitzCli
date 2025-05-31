package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/droidstarter-cli/internal/androidops"
	"github.com/droidstarter-cli/internal/config"
	"github.com/droidstarter-cli/internal/fileops"
	"github.com/droidstarter-cli/internal/gitops"
)

func main() {
	repoURL := "https://github.com/boryanz/DroidStarterTemplate.git"
	rootPath := filepath.Join(os.Getenv("HOME"), "droidstarter")
	var relativePath = filepath.Join("app", "src", "main", "java", "com", "android")
	var config = config.ParseConfigJson()

	fileops.RemoveAll(rootPath)
	gitops.CloneGithubRepo(rootPath, repoURL)

	var buildGradleTemplatePath = filepath.Join(rootPath, "build.gradle.template.txt")
	fileops.EnsureFileExists(buildGradleTemplatePath)

	content, err := os.ReadFile(buildGradleTemplatePath)
	if err != nil {
		fmt.Printf("Error while reading build.gradle.template.txt: %v\n", err)

	}
	fmt.Println("Content of the build.gradle.template.txt")

	var buildGradleFileContent = string(content)
	androidops.ParseJsonAndReplaceBuildGradlePlaceholders(buildGradleFileContent)
	androidops.OverwriteBuildGradleFile(rootPath, buildGradleFileContent)
	androidops.DeleteBuildGradleTemplate(rootPath)

	//Second part - this needs to enum from config

	if !config.Architecture.IS_MVVM {
		//copy theme package and place into presentation
		var themeFilePath = filepath.Join(rootPath, relativePath, "theme")
		var themeOutputPath = filepath.Join(rootPath, relativePath, "droidstartermvi", "presentation", "theme")
		fileops.CopyDir(themeFilePath, themeOutputPath)
		//delete mvvm package
		var mvvmPackagePath = filepath.Join(rootPath, relativePath, "droidstartertemplatemvvm")
		fileops.RemoveAll(mvvmPackagePath)
		//delete theme package
		var themePath = filepath.Join(rootPath, relativePath, "theme")
		fileops.RemoveAll(themePath)
	} else {
		//delete mvi whole package
		var mviPackagePath = filepath.Join(rootPath, relativePath, "droidstartermvi")
		fileops.RemoveAll(mviPackagePath)
	}

	androidops.RemoveAllDisabledFeatures(config, rootPath, relativePath)

	//move features package into architecture (MVVM or MVI package)
	androidops.MoveEnabledFeaturesIntoPackages(config, rootPath, relativePath)
}
