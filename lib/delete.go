package lib

import (
	"flag"
	"fmt"
	"os"
)

// deletePackage removes a package from the system using pacman.
// Uses the -Rns flags: -R (remove), -n (no backup), -s (remove dependencies).
//
// Parameters:
//   - pkg: Name of the package to delete
func deletePackage(pkg string) {
	// Run pacman with sudo to remove the package and its unused dependencies
	// -R: Remove package
	// -n: Don't save configuration files as .pacsave
	// -s: Remove unnecessary dependencies
	if err := RunInteractiveCommand("sudo", "pacman", "-Rns", pkg); err != nil {
		fmt.Println("Error deleting package:", err)
	}
}

// MainDeletePackage is the entry point for the delete command.
// Validates that a package name is provided and calls deletePackage.
func MainDeletePackage() {
	// Check if a package name was provided
	if flag.NArg() <= 0 {
		fmt.Println("Please provide an argument with " + os.Args[1])
		flag.Usage()
		os.Exit(1)
	}

	// Get the package name from command-line arguments
	argumentPkg := flag.Arg(0)
	deletePackage(argumentPkg)
}
