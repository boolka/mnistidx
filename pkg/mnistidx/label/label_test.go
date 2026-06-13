package label_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	"github.com/boolka/mnistidx/pkg/mnistidx/label"
)

func TestLabel(t *testing.T) {
	t.Parallel()

	l := label.NewIDXLabel(bytes.NewBuffer(mnistdb.TestLabels))

	h, err := l.ReadHeader()

	if err != nil {
		t.Fatal(err)
	}

	if h.MN != 2049 || h.LabelsCount != 10_000 {
		t.Fatal(h)
	}

	labelsCount := 0

	for {
		num, err := l.ReadContent()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatal(err)
		}

		if num > 9 || num < 0 {
			t.Fatal(num)
		}

		labelsCount++
	}

	if labelsCount != 10_000 {
		t.Fatal(labelsCount)
	}
}
