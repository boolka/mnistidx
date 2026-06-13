package image_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx/image"
)

func TestImage(t *testing.T) {
	t.Parallel()

	idx := image.NewIDXImage(bytes.NewReader(mnistdb.TestImages))

	h, err := idx.ReadHeader()
	if err != nil {
		t.Fatal(err)
	}

	if h.MN != 2051 || h.ImagesCount != 10_000 || h.ImgCols != 28 || h.ImgRows != 28 {
		t.Fatal(h)
	}

	buf := make(image.ImageContent, h.ImgRows*h.ImgCols)

	imagesCount := 0

	for {
		err = idx.ReadImage(int(h.ImgRows), int(h.ImgCols), buf)

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatal(err)
		}

		imagesCount++
	}

	if imagesCount != 10_000 {
		t.Fatal(imagesCount)
	}
}
