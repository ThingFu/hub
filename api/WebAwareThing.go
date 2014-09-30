package api

import "net/http"

// Interface contract for any thing which is able to respond
// to http requests to /api/thing/{id}
type WebAwareThing interface {
	HandleRequest(w http.ResponseWriter, req *http.Request)
	HandleOperation(w http.ResponseWriter, req *http.Request)
}
