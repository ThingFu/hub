// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package consequences

import "fmt"

type SetEnvironmentVariable struct {
}

func (s SetEnvironmentVariable) Execute(config map[string]interface{}) {
	fmt.Println("Consequence: Set Environment Variable")
}