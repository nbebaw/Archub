package lib

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func cleanup() {
	getArchubDirSize()
	archDir := GetArchhubDirPath()
	err := os.RemoveAll(archDir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func convertBytes(s float64, base float64) string {
	var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	unitsLimit := len(sizes)
	i := 0
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}

func getArchubDirSize() {
	archDir := GetArchhubDirPath()
	var totalSize int64
	err := filepath.Walk(archDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s will be free\n", convertBytes(float64(totalSize), 1024.0))
}

func MainCleanUp() {
	if flag.NArg() <= 0 {
		cleanup()
		return
	}

	fmt.Println(os.Args[1], "this flag do not need any argument")
	flag.Usage()
}
