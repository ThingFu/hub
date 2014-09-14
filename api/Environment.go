// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

type Environment interface {
	GetUptime() string
	GetConfig() Configuration
	GetHome() string
	IsStartedUp() bool
	ContainerAware
}
