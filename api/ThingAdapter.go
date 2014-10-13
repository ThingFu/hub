// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import _ "net/http"

// Thing adapter to support any sensing or actuation
type ThingAdapter interface {
	Cycle(*Thing)
	OnRead(*Thing, *ThingService, ReadRequest, ProtocolHandler) (State)
	OnWrite(*Thing, string, WriteRequest, AppDB, ProtocolHandler)
	GetEventText(*Thing, *ThingService, State) (shortText string, longText string)
}
