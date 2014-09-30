// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"fmt"
	"github.com/thingfu/hub/api"
)

type AdapterWebAppUtilities struct {
}

func (d *AdapterWebAppUtilities) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {
	db.Put(op, params)

	fmt.Println("AdapterWebAppUtilities -- Actuate")
}

func (d *AdapterWebAppUtilities) Cycle(dev *api.Thing) {
	fmt.Println("AdapterWebAppUtilities -- Cycle")
}

func (d *AdapterWebAppUtilities) OnSense(dev *api.Thing, data api.ThingData) (state map[string]interface{}) {
	return nil
}

func (d *AdapterWebAppUtilities) GetEventText(dev *api.Thing, sensor *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	return
}
