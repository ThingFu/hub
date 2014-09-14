// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type ProtocolHandler interface {
	Start()
	Stop()
	IsEnabled() bool
	GetName() string
	GetLabel() string
	Handle(data interface{})

	SetProtocolConfiguration(ProtocolConfiguration)
	SetDeviceService(DeviceService)
	SetFactory(Factory)
	SetEnvironment(Environment)
}
