// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
	"log"
	"net/http"
	"strings"
)

type AdapterLGTV47LS5700 struct {
}

func (tv *AdapterLGTV47LS5700) Cycle(dev *api.Device) {
	tv.initSession(dev)
}

func (tv *AdapterLGTV47LS5700) OnSense(dev *api.Device, data api.DeviceData) {

}

func (tv *AdapterLGTV47LS5700) initSession(dev *api.Device) {
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

func (d *AdapterLGTV47LS5700) GetEventText(*api.Device, *api.Sensor) (shortText string, longText string) {
	return "", ""
}
