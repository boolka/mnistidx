package mnistidx_test

import (
	"io"
	"os"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	userMnistDb "github.com/boolka/mnistidx/pkg/internal"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func BenchmarkImageIDX(b *testing.B) {
	mdb, err := userMnistDb.NewUserMnistDb()

	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		imagesFile, err := os.Open(mdb.GetDbPath(mnistdb.TrainImagesDb))

		if err != nil {
			b.Fatal(err)
		}

		labelsFile, err := os.Open(mdb.GetDbPath(mnistdb.TrainLabelsDb))

		if err != nil {
			b.Fatal(err)
		}

		i, err := mnistidx.NewIDX(imagesFile, labelsFile)

		if err != nil {
			b.Fatal(err)
		}

		buf := make([]byte, i.ImageBufSize())

		for {
			l, err := i.Read(buf)

			if err == io.EOF {
				break
			}

			if err != nil {
				b.Fatal(err)
			}

			if l < 0 || l > 9 {
				b.Fatal(l)
			}
		}
	}
}
