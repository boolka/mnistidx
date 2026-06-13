package image

import (
	"encoding/binary"
	"io"
)

type IDXImage struct {
	io.Reader
}

func NewIDXImage(r io.Reader) IDXImage {
	return IDXImage{
		Reader: r,
	}
}

type ImageHeader struct {
	MN          int32 // 2051
	ImagesCount int32
	ImgRows     int32
	ImgCols     int32
}

func (i *IDXImage) ReadHeader() (*ImageHeader, error) {
	ih := new(ImageHeader)

	err := binary.Read(i.Reader, binary.BigEndian, ih)
	if err != nil {
		return nil, err
	}

	return ih, nil
}

type ImageContent []byte

func (i *IDXImage) ReadImage(w, h int, buf ImageContent) error {
	n, err := i.Reader.Read(buf)
	if err != nil || n != w*h {
		return err
	}

	return nil
}
