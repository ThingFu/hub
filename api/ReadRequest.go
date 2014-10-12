package api

type ReadRequest struct {
	payload 	interface {}
	data 		map[string] interface {}
}

func NewReadRequest(payload interface {}) ReadRequest {
	p := new(ReadRequest)
	p.payload = payload
	p.data = make(map[string] interface {})

	return *p
}


func (c ReadRequest) GetPayload() interface {} {
	return c.payload
}

func (c *ReadRequest) Put(key string, val interface {}) {
	c.data[key] = val
}

func (c *ReadRequest) Get(key string) (interface {}) {
	return c.data[key]
}

func (c *ReadRequest) GetAsString(key string) (string) {
	return c.Get(key).(string)
}

func (c *ReadRequest) GetAsInt(key string) (int) {
	return c.Get(key).(int)
}
