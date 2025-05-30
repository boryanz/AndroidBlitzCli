package androidops

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/droidstarter-cli/internal/config"
	"github.com/droidstarter-cli/internal/fileops"
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

func MoveEnabledFeaturesIntoPackages(config config.Config, rootPath string, relativePath string) {
	if config.Architecture.IS_MVVM {
		var featuresPath = filepath.Join(rootPath, relativePath, "features")
		var featuresOutput = filepath.Join(rootPath, relativePath, "droidstartermvvm", "features")
		fileops.CopyDir(featuresPath, featuresOutput)
		fileops.RemoveAll(featuresPath)
	} else {
		var featuresPath = filepath.Join(rootPath, relativePath, "features")
		var featuresOutput = filepath.Join(rootPath, relativePath, "droidstartermvi", "features")
		fileops.CopyDir(featuresPath, featuresOutput)
		fileops.RemoveAll(featuresPath)
	}
}

func RemoveAllDisabledFeatures(config config.Config, rootPath string, relativePath string) {
	if !config.NotificationFeature.ENABLED {
		var notificationsPackage = filepath.Join(rootPath, relativePath, "droidstartermvi", "features", "notifications")
		fileops.RemoveAll(notificationsPackage)
	}
	if config.FirebaseAuthFeature.ENABLED {
		var firebasePackage = filepath.Join(rootPath, relativePath, "droidstartermvi", "features", "firebaseauth")
		fileops.RemoveAll(firebasePackage)
	}
	if config.RoomFeature.ENABLED {
		var roomPackage = filepath.Join(rootPath, relativePath, "droidstartermvi", "features")
		fileops.RemoveAll(roomPackage)
	}
	if config.RetrofitFeature.ENABLED {
		var retrofitPackage = filepath.Join(rootPath, relativePath, "droidstartermvi", "features", "retrofit")
		fileops.RemoveAll(retrofitPackage)
	}
}
