package networking

import "net/http"

type MiniResponseObject struct {
	RequestUrl     string
	RequestError   error
	HttpStatusCode int
	Headers        http.Header
	ResponseBody   []byte
}
