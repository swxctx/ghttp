package ghttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// reqtimeout
type reqtimeout interface {
	Timeout() bool
}

// Error
type Error struct {
	timeout bool
	Err     error
}

// compression
type compression struct {
	writer          func(buffer io.Writer) (io.WriteCloser, error)
	reader          func(buffer io.Reader) (io.ReadCloser, error)
	ContentEncoding string
}

// Timeout
func (e *Error) Timeout() bool {
	return e.timeout
}

// Error
func (e *Error) Error() string {
	return e.Err.Error()
}

// addHeaders
func (r Request) addHeaders(headersMap http.Header) {
	if len(r.UserAgent) > 0 {
		headersMap.Add("User-Agent", r.UserAgent)
	}
	if r.Accept != "" {
		headersMap.Add("Accept", r.Accept)
	}
	if r.ContentType != "" {
		headersMap.Add("Content-Type", r.ContentType)
	}
	if r.XForwardedFor != "" {
		headersMap.Add("X-Forwarded-For", r.XForwardedFor)
	}
}

// Return value if nonempty, def otherwise.
func valueOrDefault(value, def string) string {
	if value != "" {
		return value
	}
	return def
}

// paramParse
func paramParse(query interface{}) (string, error) {
	switch query.(type) {
	case url.Values:
		return query.(url.Values).Encode(), nil
	case *url.Values:
		return query.(*url.Values).Encode(), nil
	default:
		var (
			v = &url.Values{}
		)
		err := paramParseStruct(v, query)
		return v.Encode(), err
	}
}

// paramParseStruct
func paramParseStruct(v *url.Values, query interface{}) error {
	var (
		va = reflect.ValueOf(query)
		t  = reflect.TypeOf(query)
	)
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		va = va.Elem()
		t = va.Type()
	}

	if t.Kind() != reflect.Struct {
		return errors.New("Can't parse Query.")
	}

	for i := 0; i < t.NumField(); i++ {
		var (
			name string
		)

		field := va.Field(i)
		typeField := t.Field(i)

		if !field.CanInterface() {
			continue
		}

		urlTag := typeField.Tag.Get("json")
		if urlTag == "-" {
			continue
		}

		name, opts := parseTag(urlTag)

		var (
			omitEmpty, squash bool
		)
		omitEmpty = opts.Contains("omitempty")
		squash = opts.Contains("squash")

		if squash {
			err := paramParseStruct(v, field.Interface())
			if err != nil {
				return err
			}
			continue
		}

		if urlTag == "" {
			name = strings.ToLower(typeField.Name)
		}

		if val := fmt.Sprintf("%v", field.Interface()); !(omitEmpty && len(val) == 0) {
			v.Add(name, val)
		}
	}
	return nil
}

func prepareRequestBody(b interface{}) (io.Reader, error) {
	switch b.(type) {
	case string:
		return strings.NewReader(b.(string)), nil
	case io.Reader:
		return b.(io.Reader), nil
	case []byte:
		return bytes.NewReader(b.([]byte)), nil
	case url.Values:
		return strings.NewReader(b.(url.Values).Encode()), nil
	case *url.Values:
		return strings.NewReader(b.(*url.Values).Encode()), nil
	case nil:
		return nil, nil
	default:
		j, err := json.Marshal(b)
		if err == nil {
			return bytes.NewReader(j), nil
		}
		return nil, err
	}
}

// GetTlsConfig 生成tls config
func GetTlsConfig(cerPath, keyPath string) *tls.Config {
	// get certs
	certs, err := tls.LoadX509KeyPair(cerPath, keyPath)
	if err != nil {
		return &tls.Config{InsecureSkipVerify: true}
	}

	return &tls.Config{
		Certificates: []tls.Certificate{certs},
	}
}
