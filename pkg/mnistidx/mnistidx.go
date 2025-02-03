// # mnistidx
//
// package provide types to work with mnist database file contents
package mnistidx

import (
	"errors"
	"io"

	"github.com/boolka/mnistidx/pkg/mnistidx/image"
	"github.com/boolka/mnistidx/pkg/mnistidx/label"
)

// MnistIDX structure contain both images & labels database info
type MnistIDX struct {
	ImagesHeader image.ImageHeader
	LabelsHeader label.LabelHeader
	images       image.IDXImage
	labels       label.IDXLabel
}

// NewIDX to create new mnistidx instance
//
// Accepts images & labels [io.Reader] arguments and return pointer to [github.com/boolka/mnistidx/pkg/mnistidx.MnistIDX] if successful
//
// Returns error if corresponding image & label databases contains not equal items count
func NewIDX(imagesReader, labelsReader io.Reader) (*MnistIDX, error) {
	idxImages := image.NewIDXImage(imagesReader)
	idxLabels := label.NewIDXLabel(labelsReader)

	ih, err := idxImages.ReadHeader()

	if err != nil {
		return nil, err
	}

	lh, err := idxLabels.ReadHeader()

	if err != nil {
		return nil, err
	}

	if ih.ImagesCount != lh.LabelsCount {
		return nil, errors.New("images & labels count are not match")
	}

	return &MnistIDX{
		ImagesHeader: *ih,
		LabelsHeader: *lh,
		images:       idxImages,
		labels:       idxLabels,
	}, nil
}

// ImageBufSize returns image buffer size to allocate for single image
func (i *MnistIDX) ImageBufSize() int32 {
	return i.ImagesHeader.ImgRows * i.ImagesHeader.ImgCols
}

// Read is writes next image to buf and returns correspond label.
// Use [github.com/boolka/mnistidx/pkg/mnistidx.MnistIDX.ImageBufSize] to allocate enough space buffer.
//
// Returns an error if there is not enough space to fit the image in the buffer
func (i *MnistIDX) Read(buf image.ImageContent) (label.LabelContent, error) {
	if len(buf) < int(i.ImageBufSize()) {
		return -1, errors.New("not enough buffer space")
	}

	err := i.images.ReadImage(int(i.ImagesHeader.ImgCols), int(i.ImagesHeader.ImgRows), buf)

	if err != nil {
		return -1, err
	}

	l, err := i.labels.ReadContent()

	if err != nil {
		return -1, err
	}

	return l, nil
}
