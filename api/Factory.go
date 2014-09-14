// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type Factory interface {
	CreateCondition(string) Condition
	CreateConsequence(string) Consequence
	CreateProtocolHandler(string, ProtocolConfiguration) ProtocolHandler
	CreateDeviceAdapter(string) DeviceAdapter
	ContainerAware
}
