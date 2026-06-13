package label

import (
	"encoding/binary"
	"io"
)

type IDXLabel struct {
	io.Reader
}

func NewIDXLabel(r io.Reader) IDXLabel {
	return IDXLabel{
		Reader: r,
	}
}

type LabelHeader struct {
	MN          int32 // 2049
	LabelsCount int32
}

func (i *IDXLabel) ReadHeader() (*LabelHeader, error) {
	lh := new(LabelHeader)

	err := binary.Read(i.Reader, binary.BigEndian, lh)
	if err != nil {
		return nil, err
	}

	return lh, nil
}

type LabelContent int8

func (i *IDXLabel) ReadContent() (LabelContent, error) {
	label := make([]byte, 1)

	if _, err := i.Reader.Read(label); err != nil {
		return -1, err
	}

	return LabelContent(label[0]), nil
}
