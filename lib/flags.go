package lib

import "flag"

// DefineFlags configures all command-line flags for the Archub application.
// Each flag is defined with both a short form (e.g., -s) and a long form (e.g., --search).
// After defining all flags, it sets the custom usage function and parses the command-line arguments.
func DefineFlags() {
	// Information flags (no package argument required)
	flag.BoolVar(&ShowVersion, "version", false, "Show version information")
	flag.BoolVar(&Help, "help", false, "Show help")
	flag.BoolVar(&CheckUpdatePkg, "check", false, "Check for updates")
	flag.BoolVar(&InfoOfSystem, "info", false, "info about your system")

	// Cleanup flags (no package argument required)
	flag.BoolVar(&CleanUp, "c", false, "Clean up all tmp folders")
	flag.BoolVar(&CleanUpAlias, "clean", false, "Clean up all tmp folders")

	// List flags (no package argument required)
	flag.BoolVar(&ListPkg, "l", false, "List all aur installed Packages")
	flag.BoolVar(&ListPkgAlias, "list", false, "List all aur installed Packages")

	// Update flags (may require package argument)
	flag.BoolVar(&UpdatePkg, "u", false, "Update a Package")
	flag.BoolVar(&UpdatePkgAlias, "update", false, "Update a Package")
	flag.BoolVar(&UpdateAllPkg, "all", false, "Update all packages")

	// Delete flags (requires package argument)
	flag.BoolVar(&DeletePkg, "d", false, "Delete a package")
	flag.BoolVar(&DeletePkgAlias, "delete", false, "Delete a package")

	// Search flags (requires package argument)
	flag.BoolVar(&SearchPkg, "s", false, "Search a package")
	flag.BoolVar(&SearchPkgAlias, "search", false, "Search a package")

	// Install flags (requires package argument)
	flag.BoolVar(&InstallPkg, "i", false, "Install a Package")
	flag.BoolVar(&InstallPkgAlias, "install", false, "Install a Package")

	// Package info flags (requires package argument)
	flag.BoolVar(&PkgInfo, "Si", false, "Show package information")
	flag.BoolVar(&PkgInfoAlias, "pkginfo", false, "Show package information")

	// Set custom usage function and parse command-line arguments
	flag.Usage = GetHelp
	flag.Parse()
}
