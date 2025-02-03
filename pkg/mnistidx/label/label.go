package label

import (
	"encoding/binary"
	"io"
)

type IDXLabel struct {
	r io.Reader
}

func NewIDXLabel(r io.Reader) IDXLabel {
	return IDXLabel{
		r: r,
	}
}

type LabelHeader struct {
	MN          int32 // 2049
	LabelsCount int32
}

func (i *IDXLabel) ReadHeader() (*LabelHeader, error) {
	l := new(LabelHeader)

	err := binary.Read(i.r, binary.BigEndian, l)

	if err != nil {
		return nil, err
	}

	return l, nil
}

type LabelContent int8

func (i *IDXLabel) ReadContent() (LabelContent, error) {
	label := make([]byte, 1)
	_, err := i.r.Read(label)

	if err != nil {
		return -1, err
	}

	return LabelContent(label[0]), nil
}
