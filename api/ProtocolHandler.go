// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type ProtocolHandler interface {
	OnStart()
	OnStop()
	GetName() string
	GetLabel() string
	OnRead(data ReadRequest)
	Write(*Thing, WriteRequest)

	GetChannel() (CommunicationChannel)
	SetChannel(CommunicationChannel)

	SetThingManager(ThingManager)
	SetFactory(Factory)
	SetEnvironment(Environment)
}
