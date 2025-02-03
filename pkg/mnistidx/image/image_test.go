package image_test

import (
	"io"
	"os"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	userMnistDb "github.com/boolka/mnistidx/pkg/internal"
	"github.com/boolka/mnistidx/pkg/mnistidx/image"
)

func TestImage(t *testing.T) {
	t.Parallel()

	mdb, err := userMnistDb.NewUserMnistDb()

	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(mdb.GetDbPath(mnistdb.TestImagesDb))

	if err != nil {
		t.Fatal(err)
	}

	idx := image.NewIDXImage(f)

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
