# mnistidx

Mnistidx module provides api to work with mnist databases.

## Usage

```go
package main

import (
	"os"
)

func main() {
	imagesFile, err := os.Open("filepath/to/images/database/file")

	if err != nil {
		panic(err)
	}

	labelsFile, err := os.Open("filepath/to/labels/database/file")

	if err != nil {
		panic(err)
	}

	i, err := mnistidx.NewIDX(imagesFile, labelsFile)

	if err != nil {
		panic(err)
	}

	buf := make([]byte, i.ImageBufSize())

	for {
		l, err := i.Read(buf)

		if err == io.EOF {
			break
		}

		// buf contains image
		// l represent current label
	}
}
```

## Tests

If you run tests mnist databases automatically downloaded to `~/.mnistdb` directory 

## Info

For additional info look at [mnistdb](https://yann.lecun.com/exdb/mnist/)