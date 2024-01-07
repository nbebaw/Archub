package lib

import (
	"fmt"
	"os"
	"os/exec"
)

func CloneAURPackage(packageName, DirPath string) error {
	// Change working directory to the hidden directory
	if err := os.Chdir(DirPath); err != nil {
		return err
	}
	if _, err := os.Stat(packageName); !os.IsNotExist(err) {
		// Folder is already exist
		return nil
	} else {
		cmd := exec.Command("git", "clone", "https://aur.archlinux.org/"+packageName+".git")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}
}

func BuildAndInstall(packageName, DirPath string) {
	err := CloneAURPackage(packageName, DirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Change directory to the cloned package directory
	err = os.Chdir(packageName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	cmd := exec.Command("makepkg", "-si")
	// redirect the output to terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // allowed interact with the command

	errStart := cmd.Start()
	if errStart != nil {
		fmt.Println("Error cmd start:", errStart)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error cmd wait:", err)
	}
}
