package fileops

import (
	"fmt"
	"os"
	"path/filepath"
)

func OverwriteBuildGradleFile(destRoot string, updatedContent string) error {
	outputPath := filepath.Join(destRoot, "app", "build.gradle.kts")
	err := os.WriteFile(outputPath, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to build.gradle.kts: %w", err)
	}
	fmt.Printf("✅ Updated %s\n", outputPath)
	return nil
}

func DeleteBuildGradleTemplate(destRoot string) error {
	fileToDelete := filepath.Join(destRoot, "build.gradle.template.txt")
	if err := os.Remove(fileToDelete); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete build gradle template file: %w", err)
	} else if err == nil {
		fmt.Printf("✅ Successfully deleted %s\n", fileToDelete)
	}
	return nil
}
