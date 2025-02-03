package mnistidx_test

import (
	"io"
	"os"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	userMnistDb "github.com/boolka/mnistidx/pkg/internal"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func TestNotMatchIDX(t *testing.T) {
	t.Parallel()

	mdb, err := userMnistDb.NewUserMnistDb()

	if err != nil {
		t.Fatal(err)
	}

	imagesFile, err := os.Open(mdb.GetDbPath(mnistdb.TestImagesDb))

	if err != nil {
		t.Fatal(err)
	}

	labelsFile, err := os.Open(mdb.GetDbPath(mnistdb.TrainLabelsDb))

	if err != nil {
		t.Fatal(err)
	}

	_, idxErr := mnistidx.NewIDX(imagesFile, labelsFile)

	if idxErr == nil {
		t.Fatal(err)
	}
}

func TestIDX(t *testing.T) {
	t.Parallel()

	mdb, err := userMnistDb.NewUserMnistDb()

	if err != nil {
		t.Fatal(err)
	}

	imagesFile, err := os.Open(mdb.GetDbPath(mnistdb.TrainImagesDb))

	if err != nil {
		t.Fatal(err)
	}

	labelsFile, err := os.Open(mdb.GetDbPath(mnistdb.TrainLabelsDb))

	if err != nil {
		t.Fatal(err)
	}

	i, err := mnistidx.NewIDX(imagesFile, labelsFile)

	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, i.ImageBufSize())

	for {
		l, err := i.Read(buf)

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatal(err)
		}

		if l < 0 || l > 9 {
			t.Fatal(l)
		}
	}
}
