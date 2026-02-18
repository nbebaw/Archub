package lib

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// cleanup removes the Archub temporary directory and all its contents.
// Displays the amount of space being freed before deletion.
func cleanup() {
	// Calculate and display directory size before deletion
	getArchubDirSize()

	// Get path to the temporary directory
	archDir := GetArchubDirPath()

	// Remove the directory and all its contents
	err := os.RemoveAll(archDir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// convertBytes converts bytes to a human-readable format (B, kB, MB, GB, etc.).
//
// Parameters:
//   - s: Size in bytes
//   - base: Base for conversion (typically 1024 for binary, 1000 for decimal)
//
// Returns:
//   - string: Formatted size string (e.g., "1.23 MB", "456 kB")
func convertBytes(s float64, base float64) string {
	// Array of size unit suffixes
	var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	unitsLimit := len(sizes)
	i := 0

	// Divide by base until size is less than base or we reach the largest unit
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	// Use decimal places for sizes larger than kB
	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}

// getArchubDirSize calculates and displays the total size of the Archub temporary directory.
// Walks through all files in the directory tree and sums their sizes.
func getArchubDirSize() {
	archDir := GetArchubDirPath()
	var totalSize int64

	// Walk through directory tree and sum file sizes
	err := filepath.Walk(archDir, func(path string, info os.FileInfo, err error) error {
		// Handle walk errors
		if err != nil {
			return err
		}

		// Add file size to total (skip directories)
		if info != nil && !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error calculating directory size:", err)
		return
	}

	// Display the size that will be freed
	fmt.Printf("%s is now a free space\n", convertBytes(float64(totalSize), 1024.0))
}

// MainCleanUp is the entry point for the cleanup command.
// Validates arguments and calls cleanup to remove temporary files.
func MainCleanUp() {
	// Cleanup doesn't require any arguments
	if flag.NArg() <= 0 {
		cleanup()
		return
	}

	// Error: user provided unnecessary arguments
	fmt.Println(os.Args[1], "this flag do not need any argument")
	flag.Usage()
}
