# mnistidx

[![Go Reference](https://pkg.go.dev/badge/github.com/boolka/mnistidx.svg)](https://pkg.go.dev/github.com/boolka/mnistidx)

A Go package for reading and processing MNIST dataset files in IDX format. This library provides a simple API to sequentially read handwritten digit images and their corresponding labels from MNIST database files.

## About MNIST

The [MNIST database](http://yann.lecun.com/exdb/mnist/) is a large database of handwritten digits commonly used for training and testing machine learning algorithms. It contains:

- **Training set**: 60,000 images with labels
- **Test set**: 10,000 images with labels
- **Image size**: 28×28 pixels (grayscale)
- **Labels**: Digits 0-9

## Features

- **Simple API**: Straightforward interface for reading image-label pairs sequentially
- **Efficient streaming**: Reads data on-demand without loading entire dataset into memory
- **Type safety**: Strongly-typed headers and content using Go types
- **Error handling**: Proper validation of data consistency (matching image/label counts)

## Installation

```bash
go get github.com/boolka/mnistidx
```

## Quick Start

```go
package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func main() {
	// Create IDX reader from mnistdb byte slices
	idx, err := mnistidx.NewIDX(
		bytes.NewReader(mnistdb.TrainImages),
		bytes.NewReader(mnistdb.TrainLabels),
	)
	if err != nil {
		panic(err)
	}

	// Allocate a buffer for a single image
	buf := make([]byte, idx.ImageBufSize())

	// Read and process images
	imageCount := 0
	for {
		label, err := idx.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// buf now contains the image data (784 bytes for 28×28 pixels)
		// label is the digit (0-9)
		fmt.Printf("Image %d: Label=%d\n", imageCount, label)
		imageCount++
	}

	fmt.Printf("Total images read: %d\n", imageCount)
}
```

## Usage Examples

### Example 1: Basic Image Reading

```go
package main

import (
	"bytes"
	"io"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func main() {
	idx, _ := mnistidx.NewIDX(
		bytes.NewReader(mnistdb.TrainImages),
		bytes.NewReader(mnistdb.TrainLabels),
	)

	buf := make([]byte, idx.ImageBufSize())

	// Read first 10 images
	for i := 0; i < 10; i++ {
		label, err := idx.Read(buf)
		if err == io.EOF {
			break
		}
		// Process image data in buf
		println("Label:", label)
	}
}
```

### Example 2: Image Statistics

```go
package main

import (
	"bytes"
	"io"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func main() {
	idx, _ := mnistidx.NewIDX(
		bytes.NewReader(mnistdb.TrainImages),
		bytes.NewReader(mnistdb.TrainLabels),
	)

	buf := make([]byte, idx.ImageBufSize())
	labelCounts := [10]int{} // Count of each digit 0-9

	for {
		label, err := idx.Read(buf)
		if err == io.EOF {
			break
		}
		labelCounts[label]++
	}

	// Print statistics
	for digit, count := range labelCounts {
		println("Digit", digit, "count:", count)
	}
}
```

### Example 3: Building a Dataset with Batches

```go
package main

import (
	"bytes"
	"io"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func main() {
	idx, _ := mnistidx.NewIDX(
		bytes.NewReader(mnistdb.TrainImages),
		bytes.NewReader(mnistdb.TrainLabels),
	)

	batchSize := 32
	batch := make([]byte, batchSize*int(idx.ImageBufSize()))
	imageBuf := make([]byte, idx.ImageBufSize())
	labels := make([]int8, batchSize)

	for batchIdx := 0; ; batchIdx++ {
		for i := 0; i < batchSize; i++ {
			label, err := idx.Read(imageBuf)
			if err == io.EOF {
				break
			}
			copy(batch[i*int(idx.ImageBufSize()):(i+1)*int(idx.ImageBufSize())], imageBuf)
			labels[i] = int8(label)
		}
	}
}
```

## API Reference

### MnistIDX

The main type for reading MNIST IDX format files.

#### Constructor

```go
func NewIDX(imagesReader, labelsReader io.Reader) (*MnistIDX, error)
```

Creates a new MNIST IDX reader from two io.Reader sources (one for images, one for labels).

- **Returns**: Pointer to MnistIDX or error if:
  - Image/label file format is invalid
  - Image and label counts don't match

#### Fields

```go
ImagesHeader image.ImageHeader  // Metadata about images
LabelsHeader label.LabelHeader  // Metadata about labels
```

**ImageHeader fields**:
- `MN`: Magic number (2051)
- `ImagesCount`: Number of images in the file
- `ImgRows`: Height of each image (28 for MNIST)
- `ImgCols`: Width of each image (28 for MNIST)

**LabelHeader fields**:
- `MN`: Magic number (2049)
- `LabelsCount`: Number of labels in the file

#### Methods

```go
func (i *MnistIDX) ImageBufSize() int32
```

Returns the size in bytes needed for a single image buffer. For MNIST, this is always 784 (28 × 28).

```go
func (i *MnistIDX) Read(buf []byte) (label int8, error)
```

Reads the next image and its corresponding label.

- **Parameters**:
  - `buf`: Byte slice with sufficient capacity (see `ImageBufSize()`)
- **Returns**: 
  - `label`: The digit label (0-9) for the image
  - `error`: `io.EOF` when all images are read, or other error if buffer is too small

## Requirements

- Go 1.23.4 or later
- `github.com/boolka/mnistdb` (for test data)

## License

MIT License - see [LICENSE](LICENSE) file for details.
