package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func main() {
	var isMvi = true
	repoURL := "https://github.com/boryanz/DroidStarterTemplate.git"
	dest := filepath.Join(os.Getenv("HOME"), "Desktop", "droid-starter-new")

	os.RemoveAll(dest)
	cloneGithubRepo(dest, repoURL)

	var targetFile = "build.gradle.template.txt"
	var absolutePath = filepath.Join(dest, targetFile)
	checkIfFileExists(absolutePath)

	content, _ := os.ReadFile(absolutePath)
	fmt.Println("Content of the build.gradle.template.txt")
	fmt.Println(string(content))

	updatedBuildGradle := strings.ReplaceAll(string(content), "{{COMPILE_SDK_PLACEHOLDER}}", "35")
	updatedBuildGradle = strings.ReplaceAll(updatedBuildGradle, "{{PACKAGE_NAME_PLACEHOLDER}}", "\"com.boryans.droidstarter\"")
	updatedBuildGradle = strings.ReplaceAll(updatedBuildGradle, "{{TARGET_SDK_PLACEHOLDER}}", "35")
	updatedBuildGradle = strings.ReplaceAll(updatedBuildGradle, "{{MIN_SDK_PLACEHOLDER}}", "25")
	updatedBuildGradle = strings.ReplaceAll(updatedBuildGradle, "{{VERSION_NAME_PLACEHOLDER}}", "\"3.3.0\"")
	updatedBuildGradle = strings.ReplaceAll(updatedBuildGradle, "{{IS_MINIFIED_ENABLED_PLACEHOLDER}}", "true")
	fmt.Println(string(updatedBuildGradle))

	overwriteBuildGradleFile(dest, updatedBuildGradle)
	deleteBuildGradleTemplate(dest)

	//Second part
	if isMvi {
		//copy theme package and place into presentation
		var themeFilePath = filepath.Join(dest, "app", "src", "main", "java", "com", "android", "droidstartertemplatemvvm", "ui", "theme")
		var themeOutputPath = filepath.Join(dest, "app", "src", "main", "java", "com", "android", "droidstartermvi", "presentation", "theme")
		copyDir(themeFilePath, themeOutputPath)
		//delete mvvm package
		var mvvmPackagePath = filepath.Join(dest, "app", "src", "main", "java", "com", "android", "droidstartertemplatemvvm")
		os.RemoveAll(mvvmPackagePath)
	} else {
		//delete mvi whole package
		var mviPackagePath = filepath.Join(dest, "app", "src", "main", "java", "com", "android", "droidstartermvi")
		os.RemoveAll(mviPackagePath)
	}
}

func overwriteBuildGradleFile(dest string, updateBuildGradleFile string) {
	outputPath := filepath.Join(dest, "app", "build.gradle.kts")
	err := os.WriteFile(outputPath, []byte(updateBuildGradleFile), 0644)
	if err != nil {
		fmt.Println("Failed to write to a file ")
	}
}

func deleteBuildGradleTemplate(dest string) {
	fileToDelete := filepath.Join(dest, "build.gradle.template.txt")
	if err := os.Remove(fileToDelete); err != nil && !os.IsNotExist(err) {
		fmt.Println("Failed to delete build gradle template file")
	} else {
		fmt.Println("Successfully deleted build gradle template")
	}
}

func checkIfFileExists(absolutePath string) {
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		fmt.Printf("File not found")
	}
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

func copyDir(src string, dest string) error {
	// Walk through the source directory
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create the target path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			// Create directory
			return os.MkdirAll(targetPath, info.Mode())
		}

		// Copy file
		return copyFile(path, targetPath, info)
	})
}

func copyFile(srcFile string, destFile string, info os.FileInfo) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return err
}
