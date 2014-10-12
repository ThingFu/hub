package api

type WriteRequest struct {
	payload 	interface {}
	data 		map[string] interface {}
}

func NewWriteRequest(payload interface {}) WriteRequest {
	p := new(WriteRequest)
	p.payload = payload
	p.data = make(map[string] interface {})

	return *p
}
