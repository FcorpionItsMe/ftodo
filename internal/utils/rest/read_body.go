package rest

import (
	"errors"
	"io"
	"net/http"
)

func ReadBody(r *http.Request) ([]byte, error) {
	bodyLength := r.ContentLength
	body := make([]byte, bodyLength)
	n, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		return nil, err
	}
	if n == 0 {
		return nil, errors.New("body is empty")
	}
	return body, nil
}
