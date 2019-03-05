package ghttp

import (
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Response ghttp response
type Response struct {
	*http.Response
	Url  string
	Body *Body
	req  *http.Request
}

// Body
type Body struct {
	reader           io.ReadCloser
	compressedReader io.ReadCloser
}

// Read
func (b *Body) Read(p []byte) (int, error) {
	if b.compressedReader != nil {
		return b.compressedReader.Read(p)
	}
	return b.reader.Read(p)
}

// Close
func (b *Body) Close() error {
	err := b.reader.Close()
	if b.compressedReader != nil {
		return b.compressedReader.Close()
	}
	return err
}

// FromToJson
func (b *Body) FromToJson(o interface{}) error {
	return json.NewDecoder(b).Decode(o)
}

// FromToString
func (b *Body) FromToString() (string, error) {
	body, err := ioutil.ReadAll(b)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Gzip
func Gzip() *compression {
	reader := func(buffer io.Reader) (io.ReadCloser, error) {
		return gzip.NewReader(buffer)
	}
	writer := func(buffer io.Writer) (io.WriteCloser, error) {
		return gzip.NewWriter(buffer), nil
	}
	return &compression{
		writer:          writer,
		reader:          reader,
		ContentEncoding: "gzip",
	}
}

// Deflate
func Deflate() *compression {
	reader := func(buffer io.Reader) (io.ReadCloser, error) {
		return zlib.NewReader(buffer)
	}
	writer := func(buffer io.Writer) (io.WriteCloser, error) {
		return zlib.NewWriter(buffer), nil
	}
	return &compression{
		writer:          writer,
		reader:          reader,
		ContentEncoding: "deflate",
	}
}

// Zlib
func Zlib() *compression {
	return Deflate()
}
