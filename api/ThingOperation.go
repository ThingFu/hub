// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Operations/Actuables for a Thing
type ThingOperation struct {
	Name        string
	Description string
	config      ThingOperationConfiguration
}

/*
{ "name": "Power Off", "description": "", "config": { "command": 1 }, "params": [] },
*/
