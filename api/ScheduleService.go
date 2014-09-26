// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import "time"

const (
	SCHEDULE_TICKER_INTERVAL = 15 * time.Second
)

// Service for managing scheduled tasks, e.g. the Cycle for things
type ScheduleService interface {
	Start()
	Stop()
	ContainerAware

	SetRulesService(RulesService)
	SetThingService(ThingService)
}
