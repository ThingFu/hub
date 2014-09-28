// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rules

import (
	"github.com/go-home/hub/api"
	"github.com/go-home/hub/utils"
	"log"
	"time"
)

type DefaultRulesService struct {
	rules        map[string]api.Rule
	thingService api.ThingService
	factory      api.Factory
	container    api.Container
}

func (r *DefaultRulesService) SetFactory(factory api.Factory) {
	r.factory = factory
}

func (r DefaultRulesService) GetRules() map[string]api.Rule {
	return r.rules
}

func (r DefaultRulesService) GetRule(ruleId string) api.Rule {
	return r.rules[ruleId]
}

func (r DefaultRulesService) Trigger(triggerType uint8, facts *api.RuleFacts) {
	exec := make(map[string]api.Rule)
	thing := facts.Thing

	if triggerType == api.TRIGGER_THING {
		thingDescriptor := thing.Descriptor
		target := facts.Target

		thingLastEvent := thing.LastEvent
		if utils.TimeWithinThreshold(thingLastEvent, thingDescriptor.EventUpdateBuffer, 5000) {
			thing.UpdateLastEvent(time.Now())

			r.thingService.SaveThing(*thing)

			for idx, rule := range r.rules {
				targets := rule.Targets
				if len(targets) > 0 {
					contains := false
					for _, t := range targets {
						if t == target {
							contains = true
						}
					}

					if contains {
						exec[idx] = rule
					}
				}
			}
		}
	} else {
		exec = r.rules
	}

	for key_rule, rule := range exec {
		if !rule.Enabled {
			continue
		}

		if !utils.TimeWithinThreshold(rule.LastRun, rule.Buffer, 5000) {
			continue
		}

		whens := rule.When
		doAction := true
		for _, when := range whens {
			result := r.evaluateWhen(&when, facts, &rule)
			if result == false {
				doAction = false
				break
			}
		}

		if doAction {
			rule.LastRun = time.Now()
			thens := rule.Then
			for _, o := range thens {
				if rule.Async {
					go r.executeConsequence(o, thing)
				} else {
					r.executeConsequence(o, thing)
				}
			}
			r.rules[key_rule] = rule
		}
	}
}

func (r DefaultRulesService) evaluateWhen(when *api.RuleWhen, facts *api.RuleFacts, rule *api.Rule) bool {
	condition := r.factory.CreateCondition(when.Trigger)

	return condition.Evaluate(when, facts, rule)
}

func (r DefaultRulesService) executeConsequence(t api.RuleThen, d *api.Thing) {
	do := t.Do
	config := t.Config

	consequence := r.factory.CreateConsequence(do)

	if consequence != nil {
		consequence.Execute(config, r.GetContainer())
	}
}

func (r *DefaultRulesService) RegisterRule(rule api.Rule) {
	r.rules[rule.Id] = rule
}

func NewDefaultRuleService(env api.Environment) DefaultRulesService {
	svc := new(DefaultRulesService)

	return *svc
}

func (s *DefaultRulesService) SetThingService(svc api.ThingService) {
	s.thingService = svc
}

func (d *DefaultRulesService) GetContainer() api.Container {
	return d.container
}

func (s *DefaultRulesService) SetContainer(c api.Container) {
	s.container = c
}

func (s *DefaultRulesService) ValidateWiring() {
	if s.thingService == nil {
		log.Fatal("thingService not wired to DefaultRulesService")
	}

	if s.factory == nil {
		log.Fatal("factory not wired to DefaultRulesService")
	}
}

func NewRulesService() *DefaultRulesService {
	svc := new(DefaultRulesService)
	svc.rules = make(map[string]api.Rule)

	return svc
}
