package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/droidstarter-cli/internal/androidops"
	"github.com/droidstarter-cli/internal/config"
	"github.com/droidstarter-cli/internal/fileops"

	"github.com/go-git/go-git/v5"
)

func main() {

	config := config.ParseConfigJson()
	replacements := map[string]string{
		"{{COMPILE_SDK_PLACEHOLDER}}":         fmt.Sprintf("%d", config.AppBuildGradle.COMPILE_SDK),
		"{{PACKAGE_NAME_PLACEHOLDER}}":        fmt.Sprintf("\"%s\"", config.AppBuildGradle.PACKAGE_NAME),
		"{{TARGET_SDK_PLACEHOLDER}}":          fmt.Sprintf("%d", config.AppBuildGradle.TARGET_SDK),
		"{{MIN_SDK_PLACEHOLDER}}":             fmt.Sprintf("%d", config.AppBuildGradle.MINIMUM_SDK),
		"{{VERSION_NAME_PLACEHOLDER}}":        fmt.Sprintf("\"%s\"", config.AppBuildGradle.APP_VERSION),
		"{{IS_MINIFIED_ENABLED_PLACEHOLDER}}": fmt.Sprintf("%t", config.AppBuildGradle.IS_MINIFIED_ENABLED),
	}

	repoURL := "https://github.com/boryanz/DroidStarterTemplate.git"
	rootPath := filepath.Join(os.Getenv("HOME"), "droidstarter")
	var relativePath = filepath.Join("app", "src", "main", "java", "com", "android")

	fileops.RemoveAll(rootPath)
	cloneGithubRepo(rootPath, repoURL)

	var targetFile = "build.gradle.template.txt"
	var absolutePath = filepath.Join(rootPath, targetFile)
	fileops.EnsureFileExists(absolutePath)

	content, _ := os.ReadFile(absolutePath)
	fmt.Println("Content of the build.gradle.template.txt")
	fmt.Println(string(content))

	updatedBuildGradle := string(content)
	for placeholder, value := range replacements {
		updatedBuildGradle = strings.ReplaceAll(updatedBuildGradle, placeholder, value)
		fmt.Println(value)
	}

	androidops.OverwriteBuildGradleFile(rootPath, updatedBuildGradle)
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

	/**
	Remove all features which are set false in the config file.
	**/
	androidops.RemoveAllDisabledFeatures(config, rootPath, relativePath)

	//move features package into architecture (MVVM or MVI package)
	androidops.MoveEnabledFeaturesIntoPackages(config, rootPath, relativePath)
}

func cloneGithubRepo(dest string, repoUrl string) {
	fmt.Println("Cloning into:", dest)
	_, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Printf("Clone failed: %v", err)
	}

	fmt.Println("✅ Cloned successfully!")
}
