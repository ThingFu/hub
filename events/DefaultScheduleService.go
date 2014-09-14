// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"github.com/go-home/hub/api"
	"log"
	"time"
)

type DefaultScheduleService struct {
	ruleTicker    *time.Ticker
	rulesService  api.RulesService
	deviceService api.DeviceService
	container     api.Container
}

func (d *DefaultScheduleService) GetContainer() api.Container {
	return d.container
}

func (d *DefaultScheduleService) SetContainer(c api.Container) {
	d.container = c
}

func (s DefaultScheduleService) Start() {
	s.ruleTicker = time.NewTicker(api.SCHEDULE_TICKER_INTERVAL)

	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-s.ruleTicker.C:
				s.rulesService.Trigger(0, new(api.RuleFacts))
				s.deviceService.Cycle()

			case <-quit:
				s.ruleTicker.Stop()
				return
			}
		}
	}()
}

func (d *DefaultScheduleService) Stop() {

}

func (d *DefaultScheduleService) ValidateWiring() {
	if d.rulesService == nil {
		log.Fatal("rulesService not wired to DefaultScheduleService")
	}

	if d.deviceService == nil {
		log.Fatal("deviceService not wired to DefaultScheduleService")
	}
}

func (d *DefaultScheduleService) SetRulesService(s api.RulesService) {
	d.rulesService = s
}

func (d *DefaultScheduleService) SetDeviceService(s api.DeviceService) {
	d.deviceService = s
}
