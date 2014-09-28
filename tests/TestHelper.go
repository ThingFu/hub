package tests

import "github.com/go-home/hub/api"

func NewMockContainer() api.Container {
	c := new(MockContainer)
	c.thingService = new(MockThingService)

	return c
}
