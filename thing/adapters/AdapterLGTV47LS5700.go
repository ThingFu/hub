// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/thingfu/hub/api"
	"log"
	"net/http"
	"strings"
)

type AdapterLGTV47LS5700 struct {
}

func (d *AdapterLGTV47LS5700) OnActuate(t *api.Thing, op string, params map[string]interface{}, db api.AppDB) {

}

func (tv *AdapterLGTV47LS5700) Cycle(dev *api.Thing) {
	tv.initSession(dev)
}

func (tv *AdapterLGTV47LS5700) OnSense(dev *api.Thing, data api.ThingData) (state map[string]interface{}) {
	return nil
}

func (tv *AdapterLGTV47LS5700) initSession(dev *api.Thing) {
	host := dev.GetAttribute("host").Value.(string)
	url := "http://" + host + ":8080/roap/api/auth"
	pairingKey := dev.GetAttribute("pairingKey").Value.(string)
	payload := "<auth><type>AuthReq</type><value>" + pairingKey + "</value></auth>"

	client := &http.Client{}
	post_data := strings.NewReader(payload)
	req, _ := http.NewRequest("POST", url, post_data)
	req.Header.Add("Content-Type", "application/atom+xml")

	_, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}
}

func (d *AdapterLGTV47LS5700) GetEventText(thing *api.Thing, sensor *api.ThingService, state map[string]interface{}) (shortText string, longText string) {
	return "", ""
}
