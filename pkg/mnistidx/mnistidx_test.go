package mnistidx_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx"
)

func TestNotMatchIDX(t *testing.T) {
	t.Parallel()

	if _, err := mnistidx.NewIDX(bytes.NewReader(mnistdb.TestImages), bytes.NewReader(mnistdb.TrainLabels)); err == nil {
		t.Fatal(err)
	}
}

func TestIDX(t *testing.T) {
	t.Parallel()

	idx, err := mnistidx.NewIDX(bytes.NewReader(mnistdb.TrainImages), bytes.NewReader(mnistdb.TrainLabels))
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, idx.ImageBufSize())

	i := 0

	for {
		l, err := idx.Read(buf)

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatal(err)
		}

		if l < 0 || l > 9 {
			t.Fatal(l)
		}

		i++
	}

	if i != 60_000 {
		t.Fatal(i)
	}
}
