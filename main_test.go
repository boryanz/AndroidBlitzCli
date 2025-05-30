package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Helper to create a dummy file for testing
func createDummyFile(t *testing.T, path string) {
	err := os.WriteFile(path, []byte("dummy content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create dummy file for test: %v", err)
	}
}

func TestDeleteBuildGradleTemplate_Success(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "test_delete_template_success")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	// Create the dummy template file inside the temp directory
	templatePath := filepath.Join(tempDir, "build.gradle.template.txt")
	createDummyFile(t, templatePath)

	// Call the function under test
	deleteBuildGradleTemplate(tempDir) // Your function expects the root path

	// Assert that the file is deleted
	if _, err := os.Stat(templatePath); !os.IsNotExist(err) {
		t.Errorf("File %s was not deleted, or an unexpected error occurred: %v", templatePath, err)
	}
}

// Test case for file not existing (should not error)
func TestDeleteBuildGradleTemplate_FileNotExists(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_delete_template_not_exists")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Call the function under test on a path where the file doesn't exist
	deleteBuildGradleTemplate(tempDir)

	// No assertion needed other than no panic/crash, as the function should handle it gracefully
	// You could add logging check if you captured output, but for this specific function, just not crashing is enough.
}
