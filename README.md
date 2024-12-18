# Webd - CLI Tool for image conversion

A fast image converter built in Go. **Webd** processes all images in a given directory, converting only targeted formats to desired formats. This tool is helpful for bulk conversions and can be potentially automated using scripts. Support for more formats coming soon!

## Demo (v0.1.2)
![Demo GIF Optimizer](https://github.com/user-attachments/assets/cacbc04c-ec69-4641-a786-783877359a5f)


## Features

- **Fast conversions powered by concurrency**
- PNG ⇔ WebP
- JPEG ⇔ PNG
- JPEG ⇔ WEBP
- **Optional cleanup** - Delete/replace original files after conversion
- **Verbose logging** - Track conversion timing and operations
- **Targeted conversion** - Only processes specified file types, leaving others untouched (no need to worry about having other file types mixed in)

## Table of Contents

- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Quick Install](#quick-install)
- [Usage](#usage)
  - [Examples](#examples)
  - [Troubleshooting](#troubleshooting)
- [Documentation](#documentation)
  - [Command Line Options](#command-line-options)
- [Case Study: Pipeline Concurrency Pattern](#case-study-pipeline-concurrency-pattern-in-webd)
  - [The Problem](#the-problem)
  - [The Solution: Pipeline Pattern](#the-solution-pipeline-pattern)
  - [Pipeline Architecture](#pipeline-architecture)
  - [Implementation](#implementation)
  - [Benefits](#benefits)
  - [Visualization](#visualisation)
- [License](#license)

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) version 1.22 or higher

### Quick Install

```bash
go install github.com/v4n1lla-1ce/webd/cmd/webd@latest
```

## Usage

```bash
# you MUST specify a conversion flag i.e. -webp2png
webd <flags> <directory>
```

### Examples

```bash
# Help Menu (both works)
webd -h
webd

# Convert all WebP files to PNG in current directory
webd --webp2png .

# Convert all WebP files to PNG in a specific directory whilst deleting all original files
# Setting --verbose will turn on logs
webd --webp2png -d --verbose /path/to/images

# Convert all PNG files to WebP in parent directory
webd --png2webp ../
```

### Troubleshooting

If you get a "command not found" error after installation, you need to add Go's bin directory to your PATH:

#### Linux/macOS

Add this line to your `~/.bashrc`, `~/.zshrc`, or equivalent:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Then restart your terminal or run:

```bash
source ~/.bashrc  # or source ~/.zshrc
```

#### Windows

1. Open System Properties → Advanced → Environment Variables
2. Under "User variables", find "Path"
3. Click "Edit" and add: `%USERPROFILE%\go\bin`
4. Click "OK" to save
5. Restart your terminal

## Documentation

### Command Line Options

| Options      | Description                                                                                |
| ------------ | ------------------------------------------------------------------------------------------ |
| `-h`         | Display help menu                                                                          |
| `-d`         | Delete original files after conversion                                                   |
| `-v`         | Display version information.                                                                |
| `--verbose`  | Logs the files converted/deleted depending on options selected                             |
| `--webp2png` | Converts WEBP files to PNG in the directory specified |
| `--webp2jpg` | Converts WEBP files to JPEG in the directory specified |
| `--png2webp` | Converts PNG files to WEBP in the directory specified |
| `--png2jpg` | Converts PNG files to JPEG in the directory specified |
| `--jpg2png` | Converts JPEG files to PNG in the directory specified |
| `--jpg2webp` | Converts JPEG files to WEBP in the directory specified |


<br>

---

<br>

## Case Study: Pipeline Concurrency Pattern in WebD

This project demonstrates the pipeline concurrency pattern in Go. Here's how WebD processes multiple images efficiently using concurrent stages.

### The Problem

Converting multiple WebP images to PNG format presents several challenges:

1. CPU-intensive image decoding/encoding
2. I/O operations for reading/writing files
3. Processing multiple files efficiently

Traditional sequential processing would be slow, especially for large directories.

### The Solution: Pipeline Pattern

WebD implements a concurrent pipeline where each processing stage runs independently:

```go
func NewPipeline[I any, O any](input <-chan I, process func(I) O) <-chan O {
    out := make(chan O)
    go func() {
        defer close(out)
        for in := range input {
            out <- process(in)
        }
    }()
    return out
}
```

### Pipeline Architecture

```
Load Files → Decode WebP → Encode PNG → Save to Disk
   ↓            ↓            ↓            ↓
[webp imgs] → [images] → [png bytes] → [saved files]
```

Data flows through the pipeline using a custom structure:

```go
type PipelineData struct {
    Value      any    // Current processing data
    SourcePath string // Original file path
    Directory  string // Working directory
    BaseName   string // Original filename
}
```

### Implementation

The pipeline is constructed in stages:

```go
// This is a simplified version
files := LoadPipeline(directory)          // Find WebP files
decoded := NewPipeline(files, DecodeWebp) // Decode WebP
encoded := NewPipeline(decoded, EncodePng) // Encode PNG
saved := NewPipeline(encoded, SaveToDisk)  // Save files
```

### Benefits

1. **Concurrent Processing**

   - Multiple images processed simultaneously
   - Efficient CPU utilization
   - Reduced total processing time

2. **Memory Efficiency**

   - Works like a conveyor belt, processes one image at a time
   - Only loads a few images at once, not your entire folder
   - Automatically slows down if your computer gets busy

3. **Maintainable Code**
   - Clear separation of concerns
   - Easy to add new processing stages

### Visualisation

![pipeline](https://github.com/user-attachments/assets/95d7011a-b7bf-4ee8-8919-5e9af506a768)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
