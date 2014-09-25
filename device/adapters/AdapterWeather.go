// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adapters

import (
	"github.com/go-home/hub/api"
	/*
		"log"
		"fmt"
		"net/http"
		"io/ioutil"
	*/)

type AdapterWeather struct {
}

func (d *AdapterWeather) Cycle(dev *api.Device) {
	/*
		city := dev.GetAttribute("city").Value.(string)
		url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s", city)

		client := &http.Client{}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Not able to get weather data")

			return
		}

		response, err := client.Do(request)
		defer response.Body.Close()
		if err != nil {
			log.Println("Not able to get weather data")

			return
		}

		content, err := ioutil.ReadAll(response.Body)


		fmt.Println(string(content[:len(content)]))

		log.Println("Adapter Weather Cycle")
	*/
}

func (d *AdapterWeather) OnSense(dev *api.Device, data api.DeviceData) (state map[string]interface{}) {
	return nil
}

func (d *AdapterWeather) GetEventText(dev *api.Device, sensor *api.Sensor, state map[string]interface{}) (shortText string, longText string) {
	return
}
