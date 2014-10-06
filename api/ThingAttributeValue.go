package api

type ThingAttributeValue struct {
	Name  string      `bson:"n"`
	Value interface{} `bson:"v"`
}

func NewThingAttributeValue(n string, v interface{}) ThingAttributeValue {
	attr := new(ThingAttributeValue)
	attr.Name = n
	attr.Value = v

	return *attr
}
