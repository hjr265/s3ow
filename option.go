package s3ow

import "github.com/aws/aws-sdk-go/aws"

type Option interface {
	Apply(*Writer)
}

type OptionFunc func(*Writer)

func (o OptionFunc) Apply(w *Writer) {
	o(w)
}

func ACL(acl string) Option {
	return OptionFunc(func(w *Writer) {
		w.input.ACL = aws.String(acl)
	})
}

func Bucket(bucket string) Option {
	return OptionFunc(func(w *Writer) {
		w.input.Bucket = aws.String(bucket)
	})
}

func Key(key string) Option {
	return OptionFunc(func(w *Writer) {
		w.input.Key = aws.String(key)
	})
}

func ContentDisposition(value string) Option {
	return OptionFunc(func(w *Writer) {
		w.input.ContentDisposition = aws.String(value)
	})
}
