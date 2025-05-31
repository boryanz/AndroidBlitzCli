package gitops

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func CloneGithubRepo(dest string, repoUrl string) {
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
