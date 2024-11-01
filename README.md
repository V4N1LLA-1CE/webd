# WebD

A fast WebP to PNG converter built in Go. WebD processes all WebP images in a given directory, converting them to PNG format while preserving the original directory structure.

## Demo

Coming soon

## Installation

```bash
go install github.com/v4n1lla-1ce/webd@latest
```

## Usage

Convert all WebP images in a directory:

```bash
webd /path/to/directory
```

Example:

```bash
# Convert all WebP images in current directory
webd .

# Convert all WebP images in specific directory
webd ~/Pictures/webp-images

# Show help
webd -h
```

## Features

- 🚀 Concurrent processing for faster conversion
- 🎯 Preserves original directory structure
- 💾 Keeps original files untouched
- 📁 Simple directory-based operation

---

# Case Study: Pipeline Concurrency Pattern in WebD

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
