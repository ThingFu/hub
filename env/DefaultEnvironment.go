// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package env

import (
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/utils"
)

type DefaultEnvironment struct {
	startup   *utils.GoTime
	config    api.Configuration
	startedUp bool
	home      string
	container api.Container
}

func (e *DefaultEnvironment) GetUptime() string {
	return e.startup.Ago()
}

func (e *DefaultEnvironment) GetConfig() api.Configuration {
	return e.config
}

func (e *DefaultEnvironment) GetHome() string {
	return e.home
}

func (e *DefaultEnvironment) IsStartedUp() bool {
	return e.startedUp
}

func (d *DefaultEnvironment) GetContainer() api.Container {
	return d.container
}

func (d *DefaultEnvironment) SetContainer(c api.Container) {
	d.container = c
}

func (s *DefaultEnvironment) ValidateWiring() {

}

func NewEnvironment(home string, config api.Configuration) api.Environment {
	env := new(DefaultEnvironment)
	env.home = home
	env.config = config
	now := utils.Now()
	env.startup = now

	return env
}
