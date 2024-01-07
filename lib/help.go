package lib

import (
	"flag"
	"fmt"
	"os"
)

func GetHelp() {
	fmt.Println("Usage:")
	fmt.Println(" archub [options]")
	fmt.Println("\nOptions:")
	fmt.Printf(" -s, --search [options] \t : Search a package\n")
	fmt.Printf(" -i, --install [options] \t : Install a package\n")
	fmt.Printf(" -c, --clean \t\t\t : Clean up\n")
	fmt.Printf(" -l, --list \t\t\t : List all aur installed Packages\n")
	fmt.Printf(" -u, --update [package] \t : Update a Package\n")
	fmt.Printf(" -u --all, --update --all \t : Update all Packages\n")
	fmt.Printf(" -d, --delete [package] \t : Delete a Package\n")
	fmt.Printf(" --check \t\t\t : Check for updates\n")
	fmt.Printf(" --help \t\t\t : How to use Archub\n")
	fmt.Printf(" --version \t\t\t : Version of Archub\n")
}

func MainHelp() {
	if flag.NArg() <= 0 {
		flag.Usage()
		return
	}

	fmt.Println(os.Args[1], "this flag do not need any argument")
	flag.Usage()
}
