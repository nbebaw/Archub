package lib

import "flag"

func DefineFlags() {
	// Allowed main options
	flag.BoolVar(&ShowVersion, "version", false, "Show version information") // Version
	flag.BoolVar(&Help, "help", false, "Show help")                          // Help
	flag.BoolVar(&CheckUpdatePkg, "check", false, "Check for updates")       // Check for updates

	flag.BoolVar(&CleanUp, "c", false, "Clean up all tmp folders")          // Clean up
	flag.BoolVar(&CleanUpAlias, "clean", false, "Clean up all tmp folders") // Clean up

	flag.BoolVar(&ListPkg, "l", false, "List all aur installed Packages")         // List
	flag.BoolVar(&ListPkgAlias, "list", false, "List all aur installed Packages") // List

	flag.BoolVar(&UpdatePkg, "u", false, "Update a Package")           // Update a Package
	flag.BoolVar(&UpdatePkgAlias, "update", false, "Update a Package") // Update a Package
	flag.BoolVar(&UpdateAllPkg, "all", false, "Update all packages")   // Update all packages

	flag.BoolVar(&DeletePkg, "d", false, "Delete a package")           // Delete a Package
	flag.BoolVar(&DeletePkgAlias, "delete", false, "Delete a package") // Delete a Package

	flag.BoolVar(&SearchPkg, "s", false, "Search a package")           // Search a package
	flag.BoolVar(&SearchPkgAlias, "search", false, "Search a package") // Search a package

	flag.BoolVar(&InstallPkg, "i", false, "Install a Package")            // Install a package
	flag.BoolVar(&InstallPkgAlias, "install", false, "Install a Package") // Install a package

	// Usage
	flag.Usage = GetHelp
	flag.Parse()
}
