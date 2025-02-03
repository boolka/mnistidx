package mnistdb

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/boolka/mnistdb/pkg/mnistdb"
)

var userRootDir = os.Getenv("HOME")

func NewUserMnistDb() (*mnistdb.MnistDb, error) {
	dir := filepath.Join(userRootDir, ".mnistdb")

	if f, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)

		if err != nil {
			return nil, err
		}
	} else {
		if !f.IsDir() {
			return nil, errors.New(dir + " is not directory")
		}
	}

	mdb, err := mnistdb.NewMnistDb(dir)

	if err != nil {
		return nil, err
	}

	err = mdb.UploadMnistDbs()

	if err != nil {
		return nil, err
	}

	return mdb, err
}
