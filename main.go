package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nbebaw/archub/lib"
)

// Version is set at build time using ldflags.
// Example: go build -ldflags "-X main.Version=v1.0.0"
var Version = "No Version Provided"

func main() {
	// Initialize Archub temporary directory at startup
	lib.CreateArchubDir()

	// Get path to temporary directory for package operations
	archDir := lib.GetArchubDirPath()

	// Define and parse command-line flags
	lib.DefineFlags()

	// Route to appropriate handler based on flags
	switch {
	// Display system information (CPU, memory, host)
	case lib.InfoOfSystem:
		lib.CpuInfo()
		lib.HostInfo()
		lib.MemoryInfo()
		return

	// Delete a package from the system
	case lib.DeletePkg || lib.DeletePkgAlias:
		lib.MainDeletePackage()
		return

	// Display version information
	case lib.ShowVersion:
		if len(os.Args) > 2 {
			fmt.Println("Error: Illegal option or argument.")
			flag.Usage()
			return
		}
		if flag.NArg() <= 0 {
			fmt.Printf("archub v%s\n", Version)
			return
		}

		fmt.Println(os.Args[1], "this flag do not need any argument")
		flag.Usage()
		return

	// Display help information
	case lib.Help:
		if len(os.Args) > 2 {
			fmt.Println("Error: Illegal option or argument.")
			flag.Usage()
			return
		}
		lib.MainHelp()
		return

	// Clean up temporary directory
	case lib.CleanUp || lib.CleanUpAlias:
		if len(os.Args) > 2 {
			fmt.Println("Error: Illegal option or argument.")
			flag.Usage()
			return
		}
		lib.MainCleanUp()
		return

	// Search for a package in AUR
	case lib.SearchPkg || lib.SearchPkgAlias:
		if flag.NArg() > 0 {
			argumentPkg := flag.Arg(0)
			lib.MainSearch(argumentPkg)
		} else {
			fmt.Println("Please provide an argument with " + os.Args[1])
			flag.Usage()
			os.Exit(1)
		}
		return

	// Show detailed package information from AUR
	case lib.PkgInfo || lib.PkgInfoAlias:
		if flag.NArg() > 0 {
			argumentPkg := flag.Arg(0)
			lib.MainPkgInfo(argumentPkg)
		} else {
			fmt.Println("Please provide an argument with " + os.Args[1])
			flag.Usage()
			os.Exit(1)
		}
		return

	// Install a package from AUR
	case lib.InstallPkg || lib.InstallPkgAlias:
		if flag.NArg() > 0 {
			argumentPkg := flag.Arg(0)
			lib.BuildAndInstall(argumentPkg, archDir)
		} else {
			fmt.Println("Please provide an argument with " + os.Args[1])
			flag.Usage()
			os.Exit(1)
		}
		return

	// Check for available updates without installing
	case lib.CheckUpdatePkg:
		if len(os.Args) > 2 {
			fmt.Println("Options do not allowed together")
			flag.Usage()
			return
		}
		if flag.NArg() <= 0 {
			updateExists := false
			// Get all installed AUR packages
			packageMap := lib.ListPackages()

			// Check each package for updates
			for packageName, packageVersion := range packageMap {
				hasUpdate, newVersion, err := lib.CheckPackageUpdate(packageName, packageVersion)
				if err != nil {
					fmt.Println("Error checking update for", packageName, ":", err)
					continue
				}

				// Display packages with available updates
				if hasUpdate {
					if !updateExists {
						fmt.Println("There is an update for the following packages:")
					}
					fmt.Printf("%s%s%s %s -> %s\n", lib.ColorLightRed, packageName, lib.ColorNone, packageVersion, newVersion)
					updateExists = true
				}
			}

			// No updates available
			if !updateExists {
				fmt.Println("There is nothing to do.")
				fmt.Println("Have a good day!")
			}
			return
		}

		fmt.Println(os.Args[1], "this flag do not need any argument")
		flag.Usage()
		return

	// Update a package or all packages
	case lib.UpdatePkg || lib.UpdatePkgAlias:
		if lib.UpdateAllPkg {
			updateExists := false
			var response string
			// If both -u and --all flags are provided without an argument
			if flag.NArg() == 0 {
				packageMap := lib.ListPackages()
				for packageName, packageVersion := range packageMap {
					hasUpdate, newVersion, err := lib.CheckPackageUpdate(packageName, packageVersion)
					if err != nil {
						fmt.Println("Error checking update for", packageName, ":", err)
						continue
					}
					if hasUpdate {
						if !updateExists {
							fmt.Println("There is an update for the following packages:")
						}
						fmt.Printf("%s%s%s %s -> %s\n", lib.ColorRed, packageName, lib.ColorNone, packageVersion, newVersion)
						if len(response) == 0 {
							fmt.Print("Do you want to update all? (y/n) ")
							fmt.Scan(&response)
							if response == "y" || response == "Y" {
								lib.BuildAndInstall(packageName, archDir)
							}
						} else {
							lib.BuildAndInstall(packageName, archDir)
						}

						updateExists = true
					}
				}
				if !updateExists {
					fmt.Println("There is no update")
				}
			} else {
				fmt.Println("Please provide -u --all without an argument")
				flag.Usage()
				os.Exit(1)
			}
		} else {
			// If only -u flag is provided or -u flag is provided with an argument
			if flag.NArg() > 0 {
				argumentPkg := flag.Arg(0)
				lib.BuildAndInstall(argumentPkg, archDir)
			} else {
				fmt.Println("Please provide an argument with -u or use -u --all")
				os.Exit(1)
			}
		}
		return
	// -l || --list
	case lib.ListPkg || lib.ListPkgAlias:
		if flag.NArg() <= 0 {
			packageMap := lib.ListPackages()
			for packageName, packageVersion := range packageMap {
				fmt.Printf("%sPackage:%s %s, %sVersion:%s %s\n", lib.ColorGreen, lib.ColorNone, packageName, lib.ColorLightRed, lib.ColorNone, packageVersion)
			}
			return
		}

		fmt.Println(os.Args[1], "this flag do not need any argument")
		flag.Usage()
		return
	}

	// Show help when user give only the main options
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	// Show error and help when user write an illegal arguments/options
	if flag.NArg() > 0 {
		fmt.Println("Error: Illegal option or argument.")
		flag.Usage()
		return
	}
}
