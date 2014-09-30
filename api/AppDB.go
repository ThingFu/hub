package api

type AppDB interface {
	Put(string, interface{})
	Get(string) interface{}
	Delete(string)
	Purge()
}
