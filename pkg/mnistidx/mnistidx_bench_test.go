package mnistidx_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func BenchmarkImageIDX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idx, err := mnistidx.NewIDX(bytes.NewReader(mnistdb.TrainImages), bytes.NewReader(mnistdb.TrainLabels))
		if err != nil {
			b.Fatal(err)
		}

		buf := make([]byte, idx.ImageBufSize())

		for {
			l, err := idx.Read(buf)

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
