package s3ow_test

import (
	"archive/zip"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hjr265/s3ow"
)

func ExampleWriter() {
	s, _ := session.NewSession(&aws.Config{ /* ... */ })
	ow := s3ow.New(s, s3ow.Bucket(""), s3ow.Key(""))
	zw := zip.NewWriter(ow)
	// Add stuff to zip.
	zw.Close()
	ow.Close() // Returns nil after upload is complete, or error.
}
