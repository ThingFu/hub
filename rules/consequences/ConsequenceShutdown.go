// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package consequences

import "log"

type Shutdown struct {
}

func (s Shutdown) Execute(config map[string]interface{}) {
	log.Println("Consequence: Shut Down")
}
