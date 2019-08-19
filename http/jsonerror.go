package http

import (
	"encoding/json"
	h "net/http"
)

// JError works like http.Error but uses our response
// struct as the body of the response. Like http.Error
// you will still need to call a naked return in the http handler
func JError(w h.ResponseWriter, r *Response, httpcode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpcode)
	b, _ := json.Marshal(r)

	w.Write(b)
}
