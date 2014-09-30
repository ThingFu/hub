package source

import "github.com/thingfu/hub/api"

type MongoAppDB struct {
	thingId string
}

func NewAppDB(id string) api.AppDB {
	db := new(MongoAppDB)
	db.thingId = id

	return db
}

func (m *MongoAppDB) Put(string, interface{}) {

}

func (m *MongoAppDB) Get(string) interface{} {
	return nil
}

func (m *MongoAppDB) Delete(string) {

}

func (m *MongoAppDB) Purge() {

}
