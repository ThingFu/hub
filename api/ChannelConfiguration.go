// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type ChannelConfiguration struct {
	Type 		string
	Enabled 	bool
	Port    	string
	Baud    	uint32
	Protocols	[] string
}
