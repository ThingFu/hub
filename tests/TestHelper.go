package tests

import "github.com/thingfu/hub/api"

func NewMockContainer() api.Container {
	c := new(MockContainer)
	c.thingManager = new(MockThingManager)

	return c
}
