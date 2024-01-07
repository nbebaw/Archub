package lib

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func GetArchhubDirPath() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	homeDir := currentUser.HomeDir
	DirPath := filepath.Join(homeDir, ".config/archub_tmp")
	return DirPath
}

func CreateArchubDir() {
	// Check if the directory already exists
	archubDir := GetArchhubDirPath()
	if _, err := os.Stat(archubDir); os.IsNotExist(err) {
		// Create the hidden directory
		if err := os.Mkdir(archubDir, 0700); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	} else if err != nil {
		fmt.Println("Error checking directory existence:", err)
		return
	}
}
