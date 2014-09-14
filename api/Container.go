// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Container handles all the services used as well as any cross-dependencies
type Container interface {
	RulesService() RulesService
	DeviceService() DeviceService
	Env() Environment
	ScheduleService() ScheduleService
	DataSource() DataSource
	Factory() Factory
	ProtocolHandlers() map[string]ProtocolHandler
	ProtocolHandler(string) ProtocolHandler
}


