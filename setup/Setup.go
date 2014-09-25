// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package setup

import (
	"encoding/json"
	"fmt"
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/container"
	"github.com/go-home/hub/utils"
	"github.com/go-home/hub/web"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Setup(env api.Environment) {
	// setupChan := make(chan bool, 6)
	go loadDeviceTypes(env)
	loadDevices(env)
	loadRules(env)
	go setupProtocols()
	startScheduler()
	setupWebApplication(env)
	setupCompleted()
}

func startScheduler() {
	log.Println("[INFO] Whipping Hamsters into a Frenzy..")
	scheduleService := container.Instance().ScheduleService()
	scheduleService.Start()
}

func loadRules(env api.Environment) {
	log.Println("[INFO] Dividing by Zero")
	rulesService := container.Instance().RulesService()
	filepath.Walk(env.GetHome()+"/rules", func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			var rule api.Rule
			err = json.Unmarshal(content, &rule)

			if err != nil {
				fmt.Printf("Error Parsing Rule %s\n", path)
			}
			rule.Id = utils.RandomString(7)
			rule.Path = path

			log.Printf("[INFO] Registering rule file: %s\n", rule.Name)

			rulesService.RegisterRule(rule)
		}
		return nil
	})
}

func setupCompleted() {
	log.Println("[INFO] Node Started")
}

func loadDevices(env api.Environment) {
	log.Println("[INFO] Finding Waldo..")
	dataSource := container.Instance().DataSource()
	deviceManager := container.Instance().DeviceService()

	devices := dataSource.GetDevices()
	for _, device := range devices {
		deviceManager.RegisterDevice(device)
	}
}

func loadDeviceTypes(env api.Environment) {
	log.Println("[INFO] Pressing the Any Key...")
	deviceService := container.Instance().DeviceService()
	home := env.GetHome()
	root := home + "/devices"

	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, "device.json") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			var deviceType api.DeviceType
			err = json.Unmarshal(content, &deviceType)

			if err != nil {
				log.Println("error: ", err)
			}
			deviceType.Path = strings.Replace(path, "/device.json", "", -1)
			deviceService.RegisterDeviceType(deviceType)
		}
		return nil
	})
}

func setupWebApplication(env api.Environment) {
	web.NewWebApplication(env.GetConfig().ServerPort)
}

func setupProtocols() {
	log.Println("[INFO] Register Protocols")
	c := container.Instance()

	handlers := c.ProtocolHandlers()
	for _, handler := range handlers {
		if handler.IsEnabled() {
			handler.Start()
		}
	}
}
