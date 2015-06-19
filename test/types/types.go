package test

import (
	"io"
	"mime/multipart"
	"time"

	"github.com/flynn/flynn/pkg/iotool"
)

type BuildLog struct {
	mw *multipart.Writer
}

func NewBuildLog(w io.Writer) *BuildLog {
	return &BuildLog{multipart.NewWriter(w)}
}

func (b *BuildLog) NewFile(name string) (io.Writer, error) {
	return b.mw.CreateFormFile(name, name)
}

func (b *BuildLog) NewFileWithTimeout(name string, timeout time.Duration) (io.Writer, error) {
	w, err := b.NewFile(name)
	if err != nil {
		return nil, err
	}
	return iotool.NewTimeoutWriter(w, timeout), nil
}

func (b *BuildLog) Boundary() string {
	return b.mw.Boundary()
}

func (b *BuildLog) Close() error {
	return b.mw.Close()
}
