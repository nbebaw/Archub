package lib

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ListPackages returns a map of all AUR packages installed on the system.
// Uses 'pacman -Qm' to query foreign (non-repository) packages.
//
// Returns:
//   - map[string]string: Map where key is package name and value is version number
//   - Returns nil if pacman command fails
func ListPackages() map[string]string {
	// Run pacman with -Qm flag (Query foreign/AUR packages)
	cmd := exec.Command("pacman", "-Qm")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running pacman -Qm:", err)
		return nil
	}

	packages := make(map[string]string)

	// Parse pacman output line by line
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			packageName := fields[0]
			// Extract version number (strip release number after '-')
			// Example: "1.2.3-1" becomes "1.2.3"
			packageVersion := strings.Split(fields[1], "-")[0]
			packages[packageName] = packageVersion
		}
	}

	return packages
}

// GetLatestVersionFromAUR fetches the latest version of a package from AUR.
// Retrieves the PKGBUILD file and parses the pkgver variable.
//
// Parameters:
//   - packageName: Name of the AUR package
//
// Returns:
//   - string: Version string (e.g., "1.2.3"), empty string if not found
//   - error: nil on success, error if fetch fails
func GetLatestVersionFromAUR(packageName string) (string, error) {
	// Fetch PKGBUILD file from AUR git repository using curl
	cmd := exec.Command("curl", "-s", "https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h="+packageName)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Parse PKGBUILD to extract pkgver value
	output := out.String()
	versionIndex := strings.Index(output, "pkgver=")
	if versionIndex == -1 {
		// pkgver not found in PKGBUILD
		return "", nil
	}

	// Extract the version line and clean it up
	version := strings.Split(output[versionIndex:], "\n")[0]
	version = strings.TrimPrefix(version, "pkgver=")

	// Remove quotes if present (pkgver='1.2.3' or pkgver="1.2.3")
	if strings.Contains(version, "'") {
		version = strings.Trim(version, "'")
	}

	return version, nil
}

// CheckPackageUpdate checks if a package has an update available in AUR.
// Compares the current version with the latest version from AUR.
//
// Parameters:
//   - packageName: Name of the package to check
//   - currentVersion: Currently installed version
//
// Returns:
//   - bool: true if an update is available, false otherwise
//   - string: Latest version from AUR
//   - error: nil on success, error if version fetch fails
func CheckPackageUpdate(packageName, currentVersion string) (bool, string, error) {
	// Fetch latest version from AUR
	newVersion, err := GetLatestVersionFromAUR(packageName)
	if err != nil {
		return false, "", err
	}

	// Skip if version is empty or contains variable placeholders (e.g., ${pkgver})
	if len(newVersion) == 0 || strings.Contains(newVersion, "${") {
		return false, "", nil
	}

	// Compare versions
	if currentVersion != newVersion {
		return true, newVersion, nil
	}

	return false, newVersion, nil
}
