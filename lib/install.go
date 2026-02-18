package lib

import (
	"fmt"
	"os"
	"os/exec"
)

// CloneAURPackage clones an AUR package repository from the AUR git server.
// If the package directory already exists, it skips cloning.
//
// Parameters:
//   - packageName: Name of the AUR package to clone
//   - DirPath: Directory path where the package should be cloned
//
// Returns:
//   - error: nil if successful, error if cloning fails
func CloneAURPackage(packageName, DirPath string) error {
	// Change working directory to the temporary directory
	if err := os.Chdir(DirPath); err != nil {
		return err
	}

	// Check if package directory already exists
	if _, err := os.Stat(packageName); !os.IsNotExist(err) {
		// Package already cloned, skip cloning
		return nil
	}

	// Clone the AUR package repository
	cmd := exec.Command("git", "clone", "https://aur.archlinux.org/"+packageName+".git")
	cmd.Stdout = os.Stdout // Show git output
	cmd.Stderr = os.Stderr // Show git errors

	return cmd.Run()
}

// BuildAndInstall clones an AUR package, builds it using makepkg, and installs it.
// The makepkg -si flags mean: -s (sync dependencies), -i (install after build).
//
// Parameters:
//   - packageName: Name of the AUR package to build and install
//   - DirPath: Directory path where the package will be cloned
func BuildAndInstall(packageName, DirPath string) {
	// First, clone the package repository
	if err := CloneAURPackage(packageName, DirPath); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Change into the cloned package directory
	if err := os.Chdir(packageName); err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Build and install the package using makepkg
	// -s: Install missing dependencies using pacman
	// -i: Install the package after building
	if err := RunInteractiveCommand("makepkg", "-si"); err != nil {
		fmt.Println("Error building package:", err)
	}
}
