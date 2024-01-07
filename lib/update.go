package lib

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func ListPackages() map[string]string {
	cmd := exec.Command("pacman", "-Qm")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running pacman -Qm:", err)
		return nil
	}

	packages := make(map[string]string)

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			packageName := fields[0]
			packageVersion := strings.Split(fields[1], "-")[0] // Retain only version part before '-'
			packages[packageName] = packageVersion
		}
	}

	return packages
}
func GetLatestVersionFromAUR(packageName string) (string, error) {
	cmd := exec.Command("curl", "-s", "https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h="+packageName)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output := out.String()
	versionIndex := strings.Index(output, "pkgver=")
	if versionIndex == -1 {
		return "", nil
	}

	version := strings.Split(output[versionIndex:], "\n")[0]
	version = strings.TrimPrefix(version, "pkgver=")

	return version, nil
}
