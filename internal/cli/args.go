package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Arguments struct {
	DirPath      string
	DeleteOrigin bool
}

func GetArgs() Arguments {
	// define flags
	help := flag.Bool("h", false, "show help")
	deleteOrigin := flag.Bool("d", false, "delete original files after conversion")

	// parse flag
	flag.Parse()

	if *help {
		fmt.Println("\nUsage: webd [flags] <directory>")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	args := flag.Args()

	// make sure there is 1 directory argument
	if len(args) != 1 {
		fmt.Println("\nUsage: webd [flags] <directory>")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(0)
	}

	// get absolute path
	execDir, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		os.Exit(1)
	}

	// check if path exists and is accessable
	info, err := os.Stat(execDir)
	if os.IsNotExist(err) {
		fmt.Printf("%v does not exist\n\n", execDir)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error accessing path: %v\n\n", execDir)
		os.Exit(1)
	}

	// check if path is a directory
	if !info.IsDir() {
		fmt.Printf("%v is not a directory\n\n", execDir)
		os.Exit(1)
	}

	return Arguments{
		DirPath:      execDir,
		DeleteOrigin: *deleteOrigin,
	}
}
