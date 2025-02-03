package label_test

import (
	"io"
	"os"
	"testing"

	"github.com/boolka/mnistdb/pkg/mnistdb"
	userMnistDb "github.com/boolka/mnistidx/pkg/internal"
	"github.com/boolka/mnistidx/pkg/mnistidx/label"
)

func TestLabel(t *testing.T) {
	t.Parallel()

	mdb, err := userMnistDb.NewUserMnistDb()

	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(mdb.GetDbPath(mnistdb.TestLabelsDb))

	if err != nil {
		t.Fatal(err)
	}

	l := label.NewIDXLabel(f)

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
