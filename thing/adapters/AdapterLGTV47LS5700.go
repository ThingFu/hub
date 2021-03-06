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

func (d *AdapterLGTV47LS5700) OnWrite(t *api.Thing, op string, params api.WriteRequest, db api.AppDB, handler api.ProtocolHandler) {

}

func (tv *AdapterLGTV47LS5700) Cycle(dev *api.Thing) {
	tv.initSession(dev)
}

func (tv *AdapterLGTV47LS5700) OnRead(dev *api.Thing, service *api.ThingService, data api.ReadRequest, handler api.ProtocolHandler) (state api.State) {
	return nil
}

func (tv *AdapterLGTV47LS5700) initSession(dev *api.Thing) {
	log.Printf("initSession")
	host := dev.GetAttributeValue("host").Value.(string)
	url := "http://" + host + ":8080/roap/api/auth"
	pairingKey := dev.GetAttributeValue("pairingKey").Value.(string)
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

func (d *AdapterLGTV47LS5700) GetEventText(thing *api.Thing, sensor *api.ThingService, state api.State) (shortText string, longText string) {
	return "", ""
}
