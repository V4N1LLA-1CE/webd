package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

type Arguments struct {
	DirPath string
}

func GetArgs() Arguments {
	// make sure there is an argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: webd <directory>")
		os.Exit(1)
	}

	// get absolute path
	execDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		os.Exit(1)
	}

	// check if path exists and is accessable
	info, err := os.Stat(execDir)
	if os.IsNotExist(err) {
		fmt.Printf("%v does not exist\n", execDir)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error accessing path: %v\n", execDir)
		os.Exit(1)
	}

	// check if path is a directory
	if !info.IsDir() {
		fmt.Printf("%v is not a directory\n", execDir)
		os.Exit(1)
	}

	return Arguments{
		DirPath: execDir,
	}
}
