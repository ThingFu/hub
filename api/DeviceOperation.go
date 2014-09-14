// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type DeviceOperation struct {
	Name        string
	Description string
	config      DeviceOperationConfiguration
}

/*
{ "name": "Power Off", "description": "", "config": { "command": 1 }, "params": [] },
*/
