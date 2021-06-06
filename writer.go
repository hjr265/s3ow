package s3ow

import (
	"io"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Writer struct {
	client client.ConfigProvider
	input  s3manager.UploadInput
	output *s3manager.UploadOutput

	w *io.PipeWriter

	errch chan error
}

var _ io.WriteCloser = &Writer{}

// New returns an object writer that satisfies io.WriteCloser.
func New(c client.ConfigProvider, opts ...Option) *Writer {
	pr, pw := io.Pipe()

	w := Writer{
		client: c,
		input: s3manager.UploadInput{
			Body: pr,
		},
		w:     pw,
		errch: make(chan error),
	}
	for _, o := range opts {
		o.Apply(&w)
	}

	go func() {
		o, err := s3manager.NewUploader(w.client).Upload(&w.input)
		w.output = o
		w.errch <- err
	}()

	return &w
}

// Write passes b to s3manager to upload.
func (w *Writer) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

// Close marks the EOF for this upload, waits until the upload finishes, and returns error if any.
func (w *Writer) Close() error {
	err := w.w.Close()
	if err != nil {
		return err
	}
	return <-w.errch
}

// Output returns the output from the upload call to s3manager. This should be
// called after Close returns.
func (w *Writer) Output() *s3manager.UploadOutput {
	return w.output
}
