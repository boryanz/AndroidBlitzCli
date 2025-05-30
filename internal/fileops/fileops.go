package fileops

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// EnsureFileExists checks if a file exists at the given path.
func EnsureFileExists(absolutePath string) error {
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", absolutePath)
	}
	return nil
}

func CopyDir(src string, dest string) error {
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
		return CopyFile(path, targetPath, info)
	})
}

func CopyFile(srcFile string, destFile string, info os.FileInfo) error {
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

// RemoveAll removes a directory or file recursively.
func RemoveAll(path string) error {
	fmt.Printf("Attempting to remove: %s\n", path)
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("failed to remove %s: %w", path, err)
	}
	fmt.Printf("✅ Successfully removed: %s\n", path)
	return nil
}
