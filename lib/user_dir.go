package lib

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// GetArchhubDirPath returns the path to Archub's temporary directory.
// The directory is located at ~/.config/archub_tmp and is used to store
// cloned AUR package repositories during the build process.
//
// Returns:
//   - string: Full path to the archub temporary directory, or empty string on error
func GetArchubDirPath() string {
	// Get current user information to find home directory
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Build path: /home/username/.config/archub_tmp
	homeDir := currentUser.HomeDir
	DirPath := filepath.Join(homeDir, ".config/archub_tmp")
	return DirPath
}

// CreateArchubDir creates the Archub temporary directory if it doesn't exist.
// The directory is created with 0700 permissions (owner read/write/execute only).
// Should be called at application startup to ensure the directory exists.
func CreateArchubDir() {
	// Get the path to the archub directory
	archubDir := GetArchubDirPath()

	// Check if the directory already exists
	if _, err := os.Stat(archubDir); os.IsNotExist(err) {
		// Directory doesn't exist, create it with restricted permissions
		// 0700 = owner can read, write, and execute; others have no access
		if err := os.Mkdir(archubDir, 0700); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	} else if err != nil {
		// Error occurred while checking directory existence
		fmt.Println("Error checking directory existence:", err)
		return
	}
	// If directory already exists, do nothing
}
