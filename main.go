package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nbebaw/archub/lib"
)

var Version = "No Version Provided"

func main() {
	lib.CreateArchubDir()

	archDir := lib.GetArchhubDirPath()
	lib.DefineFlags()

	switch {
	// -d, --delete
	case lib.DeletePkg || lib.DeletePkgAlias:
		lib.MainDeletePackage()
		return
	// -version || --version
	case lib.ShowVersion:
		if flag.NArg() <= 0 {
			fmt.Printf("archub v%s\n", Version)
			return
		}

		fmt.Println(os.Args[1], "this flag do not need any argument")
		flag.Usage()
		return
	// -help || --help
	case lib.Help:
		lib.MainHelp()
		return

	// -c || --c
	case lib.CleanUp || lib.CleanUpAlias:
		lib.MainCleanUp()
		return

	// -s || --s
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
	// -i || --i
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
	// -check || --check
	case lib.CheckUpdatePkg:
		if flag.NArg() <= 0 {
			updateExists := false
			packageMap := lib.ListPackages()
			for packageName, packageVersion := range packageMap {
				newVersion, err := lib.GetLatestVersionFromAUR(packageName)
				if err != nil {
					fmt.Println(err)
				}
				if packageVersion != newVersion {
					fmt.Println("There is an update for the following packages:")
					fmt.Printf("%s%s%s %s -> %s\n", lib.ColorRed, packageName, lib.ColorNone, packageVersion, newVersion)
				}
			}
			if !updateExists {
				fmt.Println("There is nothing to do.")
				fmt.Println("Have a good day!")
			}
			return
		}

		fmt.Println(os.Args[1], "this flag do not need any argument")
		flag.Usage()
		return
	// -u || --u [package]
	case lib.UpdatePkg || lib.UpdatePkgAlias:
		if lib.UpdateAllPkg {
			updateExists := false
			// If both -u and --all flags are provided without an argument
			if flag.NArg() == 0 {
				packageMap := lib.ListPackages()
				for packageName, packageVersion := range packageMap {
					newVersion, err := lib.GetLatestVersionFromAUR(packageName)
					if err != nil {
						fmt.Println(err)
					}
					if packageVersion != newVersion {
						var response string
						fmt.Println("There is an update for the following packages:")
						fmt.Printf("%s%s%s %s -> %s\n", lib.ColorRed, packageName, lib.ColorNone, packageVersion, newVersion)
						fmt.Print("Do you want to update all? (y/n)")
						fmt.Scan(&response)
						if response == "y" || response == "Y" {
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
				os.Exit(1) // Exit the program or handle the case based on your application's requirements
			}
		} else {
			// If only -u flag is provided or -u flag is provided with an argument
			if flag.NArg() > 0 {
				argumentPkg := flag.Arg(0)
				lib.BuildAndInstall(argumentPkg, archDir)
			} else {
				fmt.Println("Please provide an argument with -u or use -u --all")
				os.Exit(1) // Exit the program or handle the case based on your application's requirements
			}
		}
		return
	// -l || --l
	case lib.ListPkg || lib.ListPkgAlias:
		if flag.NArg() <= 0 {
			packageMap := lib.ListPackages()
			for packageName, packageVersion := range packageMap {
				fmt.Printf("Package: %s, Version: %s\n", packageName, packageVersion)
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
