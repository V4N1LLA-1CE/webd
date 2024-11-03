package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const Version = "v0.1.4"

type Arguments struct {
	DirPath      string
	DeleteOrigin bool
	SourceExt    string
	TargetExt    string
	Verbosity    bool
	Mode         string
}

type ConversionMode struct {
	modeName  string
	enabled   bool
	sourceExt string
	targetExt string
}

func init() {
	// override the default flag.CommandLine
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// custom error handler
	flag.CommandLine.Usage = func() {
		fmt.Println("Help Menu: webd -h")
		fmt.Println()
	}
}

func GetArgs() Arguments {
	// define flags
	help := flag.Bool("h", false, "show help")
	deleteOrigin := flag.Bool("d", false, "delete original files after conversion")
	version := flag.Bool("v", false, "show webd version")
	verbosity := flag.Bool("verbose", false, "show logs when converting")

	// conversion mode flags
	webp2png := flag.Bool("webp2png", false, "convert WebP to PNG")
	png2webp := flag.Bool("png2webp", false, "convert PNG to WebP")

	// parse flag
	flag.Parse()

	if *help {
		fmt.Println("\nUsage: webd [flags] <directory>")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		displayVersion()
		os.Exit(0)
	}

	// check mutually exclusive coversion modes
	modes := []ConversionMode{
		{enabled: *webp2png, modeName: "webp2png", sourceExt: "webp", targetExt: "png"},
		{enabled: *png2webp, modeName: "png2webp", sourceExt: "png", targetExt: "webp"},
	}

	enabledModes := 0
	var selectedMode ConversionMode

	for _, mode := range modes {
		if mode.enabled {
			enabledModes++
			selectedMode = mode
		}
	}

	// allow only 1 mode
	if enabledModes != 1 {
		if enabledModes > 1 {
			fmt.Println("\nToo many conversion modes specified! Please specify only one.")
		} else {
			fmt.Println("\nPlease specify a conversion mode!\n")
		}
		os.Exit(1)
	}

	args := flag.Args()

	// make sure there is 1 directory argument
	if len(args) != 1 {
		fmt.Println("\nYou must specify a directory!\n")
		os.Exit(1)
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
		SourceExt:    selectedMode.sourceExt,
		TargetExt:    selectedMode.targetExt,
		Verbosity:    *verbosity,
		Mode:         selectedMode.modeName,
	}
}

func displayVersion() {
	fmt.Printf("webd version %v\n", Version)
}
