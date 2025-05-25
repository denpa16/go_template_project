package http

import (
	"bytes"
	"net/http"
)

func GetResponse(
	w http.ResponseWriter,
	handlerName string,
	err error,
	statusCode int,
	body *[]byte,
) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	if body != nil {
		_, _ = w.Write(*body)
	}
	if err != nil {
		buf := bytes.NewBufferString(handlerName)
		buf.WriteString(": ")
		buf.WriteString(err.Error())
		buf.WriteString("\n")
		_, _ = w.Write(buf.Bytes())
	}
}
