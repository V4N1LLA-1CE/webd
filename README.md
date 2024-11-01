# Webd - WebP to PNG Converter

A fast WebP to PNG converter built in Go. WebD processes all WebP images in a given directory, converting them to PNG format while preserving the original directory structure.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) version 1.22 or higher

### Quick Install

```bash
go install github.com/v4n1lla-1ce/webd/cmd/webd@latest
```

## Usage

```bash
webd <directory>
```

### Examples

```bash
# Help Menu (both works)
webd -h
webd

# Convert all WebP files in current directory
webd .

# Convert all WebP files in a specific directory whilst deleting all original files
webd -d /path/to/images

# Convert all WebP files in parent directory
webd ../
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

## Features

- 🚀 Concurrent processing for faster conversion
- 🎯 Preserves original directory structure
- 💾 Choose to keep original files untouched or delete them
- 📁 Simple directory-based operation

---

## Case Study: Pipeline Concurrency Pattern in WebD

This project demonstrates the pipeline concurrency pattern in Go. Here's how WebD processes multiple images efficiently using concurrent stages.

## The Problem

Converting multiple WebP images to PNG format presents several challenges:

1. CPU-intensive image decoding/encoding
2. I/O operations for reading/writing files
3. Processing multiple files efficiently

Traditional sequential processing would be slow, especially for large directories.

## The Solution: Pipeline Pattern

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

## Pipeline Architecture

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

## Implementation

The pipeline is constructed in stages:

```go
// This is a simplified version
files := LoadPipeline(directory)          // Find WebP files
decoded := NewPipeline(files, DecodeWebp) // Decode WebP
encoded := NewPipeline(decoded, EncodePng) // Encode PNG
saved := NewPipeline(encoded, SaveToDisk)  // Save files
```

## Benefits

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

## Visualisation

![pipeline](https://github.com/user-attachments/assets/95d7011a-b7bf-4ee8-8919-5e9af506a768)

## Limitations

- Only processes `.webp` files
- Converts exclusively to PNG format
- Requires write permissions in target directory

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
