# S3 Object Writer

[![Go Reference](https://pkg.go.dev/badge/github.com/hjr265/s3ow.svg)](https://pkg.go.dev/github.com/hjr265/s3ow)

Write objects directly to AWS S3-like storages without requiring any intermediary files. 

## Installation

Install s3ow using the go get command:

```
$ go get github.com/hjr265/s3ow
```

## Usage

Creating and uploading a ZIP to S3:

```golang
s, _ := session.NewSession(&aws.Config{ /* ... */ })
ow := s3ow.New(s, s3ow.Bucket(""), s3ow.Key(""))
zw := zip.NewWriter(ow)
// Add stuff to zip.
zw.Close()
ow.Close() // Returns nil after upload is complete, or error.
```

## Contributing

Contributions are welcome.

## License

This Go package (s3ow) is available under the [BSD (3-Clause) License](https://opensource.org/licenses/BSD-3-Clause).
