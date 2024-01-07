package lib

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func deletePackage(pkg string) {
	cmd := exec.Command("sudo", "pacman", "-Rns", pkg)
	// redirect the output to terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // allowed interact with the command

	errStart := cmd.Start()
	if errStart != nil {
		fmt.Println("Error cmd start:", errStart)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error cmd wait:", err)
	}
}

func MainDeletePackage() {
	if flag.NArg() <= 0 {
		fmt.Println("Please provide an argument with " + os.Args[1])
		flag.Usage()
		os.Exit(1)
	} else {
		argumentPkg := flag.Arg(0)
		deletePackage(argumentPkg)
	}
}
