package lib

import (
	"fmt"
	"os"
	"os/exec"
)

// Terminal color codes for formatted output
const (
	ColorRed      = "\033[0;31m"  // Red color for errors and warnings
	ColorGreen    = "\033[1;32m"  // Green color for success messages and labels
	ColorYellow   = "\033[1;33m"  // Yellow color for prompts and highlights
	ColorLightRed = "\033[1;91m"  // Light red color for package names
	ColorNone     = "\033[0m"     // Reset color to default
)

// Global flag variables to track which command-line options were provided
// Each flag has both a short form (e.g., -s) and a long form (e.g., --search)
var (
	Help            bool // --help: Display help message
	ShowVersion     bool // --version: Display version information
	CleanUp         bool // -c: Clean up temporary directories
	CleanUpAlias    bool // --clean: Long form of CleanUp
	ListPkg         bool // -l: List all AUR packages
	ListPkgAlias    bool // --list: Long form of ListPkg
	CheckUpdatePkg  bool // --check: Check for available updates without installing
	UpdateAllPkg    bool // --all: Update all packages (used with -u)
	UpdatePkg       bool // -u: Update a specific package
	UpdatePkgAlias  bool // --update: Long form of UpdatePkg
	DeletePkg       bool // -d: Delete a package
	DeletePkgAlias  bool // --delete: Long form of DeletePkg
	SearchPkg       bool // -s: Search for a package in AUR
	SearchPkgAlias  bool // --search: Long form of SearchPkg
	InstallPkg      bool // -i: Install a package from AUR
	InstallPkgAlias bool // --install: Long form of InstallPkg
	InfoOfSystem    bool // --info: Display system information (CPU, memory, host)
	PkgInfo         bool // -Si: Show detailed package information from AUR
	PkgInfoAlias    bool // --pkginfo: Long form of PkgInfo
)

// RunInteractiveCommand executes a command with stdin/stdout/stderr connected to the terminal.
// This allows the command to interact with the user (e.g., prompts, password input).
//
// Parameters:
//   - name: The command to execute (e.g., "makepkg", "sudo")
//   - args: Variable number of command arguments
//
// Returns:
//   - error: nil if command succeeds, wrapped error if command fails
//
// Example:
//
//	err := RunInteractiveCommand("makepkg", "-si")
func RunInteractiveCommand(name string, args ...string) error {
	// Create the command with arguments
	cmd := exec.Command(name, args...)

	// Connect command's streams to terminal for user interaction
	cmd.Stdout = os.Stdout // Show command output
	cmd.Stderr = os.Stderr // Show command errors
	cmd.Stdin = os.Stdin   // Allow user input (e.g., passwords, confirmations)

	// Start the command execution
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting command: %w", err)
	}

	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for command: %w", err)
	}

	return nil
}
