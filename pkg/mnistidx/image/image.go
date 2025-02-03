package image

import (
	"encoding/binary"
	"io"
)

type IDXImage struct {
	r io.Reader
}

func NewIDXImage(r io.Reader) IDXImage {
	return IDXImage{
		r: r,
	}
}

type ImageHeader struct {
	MN          int32 // 2051
	ImagesCount int32
	ImgRows     int32
	ImgCols     int32
}

func (i *IDXImage) ReadHeader() (*ImageHeader, error) {
	h := new(ImageHeader)
	err := binary.Read(i.r, binary.BigEndian, h)

	if err != nil {
		return nil, err
	}

	return h, nil
}

type ImageContent []byte

func (i *IDXImage) ReadImage(w, h int, buf ImageContent) error {
	s := w * h

	n, err := i.r.Read(buf)

	if err != nil || n != s {
		return err
	}

	return nil
}
