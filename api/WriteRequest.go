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

func (c *WriteRequest) Put(key string, val interface {}) {
	c.data[key] = val
}

func (c *WriteRequest) Get(key string) (interface {}) {
	return c.data[key]
}
