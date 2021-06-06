package s3ow_test

import (
	"archive/zip"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hjr265/s3ow"
)

func ExampleWriter() {
	s, _ := session.NewSession()
	ow := s3ow.New(s, s3ow.Bucket(""), s3ow.Key(""))
	zw := zip.NewWriter(ow)
	// Add stuff to zip.
	zw.Close()
	ow.Close() // Returns nil after upload is complete, or error.
}

func TestWriter(t *testing.T) {
	// Start Minio and create a test bucket before running this test.
	// $ docker run -p 9000:9000 minio/minio:RELEASE.2021-05-27T22-06-31Z server /data
	// Bucket: s3ow

	s, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("minioadmin", "minioadmin", ""),
		Endpoint:         aws.String("http://localhost:9000"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		t.Fatal(err)
	}

	ts := time.Now()
	key := fmt.Sprintf("hello-%d.txt", ts.Unix())
	const hello = "Hello, world!"

	ow := s3ow.New(s, s3ow.ACL("private"), s3ow.Bucket("s3ow"), s3ow.Key(key))
	n, err := io.WriteString(ow, hello)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(hello) {
		t.Fatalf("want n == %d, got %d", len(hello), n)
	}
	err = ow.Close()
	if err != nil {
		t.Fatal(err)
	}

	out, err := s3.New(s).GetObject(&s3.GetObjectInput{
		Bucket: aws.String("s3ow"),
		Key:    aws.String(key),
	})
	if err != nil {
		t.Fatal(err)
	}
	b, err := io.ReadAll(out.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != hello {
		t.Fatalf("want b == %q, got %q", hello, string(b))
	}
}
