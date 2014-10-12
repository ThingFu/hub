// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import _ "net/http"

// Thing adapter to support any sensing or actuation
type ThingAdapter interface {
	Cycle(*Thing)
	OnRead(*Thing, *ThingService, ReadRequest) (state map[string]interface{})
	OnWrite(*Thing, string, WriteRequest, AppDB)
	GetEventText(Thing *Thing, service *ThingService, state map[string]interface{}) (shortText string, longText string)
	// HandleGet(req http.Request, res http.Response)
}
