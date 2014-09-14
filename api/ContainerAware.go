// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// Interface contract for any services which needs to be supported by the Container
type ContainerAware interface {
	SetContainer(Container)
	ValidateWiring()
}
