// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rules

import (
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/utils"
	"log"
	"time"
)

type DefaultRulesManager struct {
	rules        map[string]api.Rule
	thingManager api.ThingManager
	factory      api.Factory
	container    api.Container
}

func (r *DefaultRulesManager) SetFactory(factory api.Factory) {
	r.factory = factory
}

func (r DefaultRulesManager) GetRules() map[string]api.Rule {
	return r.rules
}

func (r DefaultRulesManager) GetRule(ruleId string) api.Rule {
	return r.rules[ruleId]
}

func (r DefaultRulesManager) Trigger(triggerType uint8, facts *api.RuleFacts) {
	exec := make(map[string]api.Rule)
	thing := facts.Thing

	if triggerType == api.TRIGGER_THING {
		thingDescriptor := thing.Descriptor
		target := facts.Target

		thingLastEvent := thing.LastEvent
		if utils.TimeWithinThreshold(thingLastEvent, thingDescriptor.EventUpdateBuffer, 5000) {
			thing.UpdateLastEvent(time.Now())

			r.thingManager.SaveThing(*thing)

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

func (r DefaultRulesManager) evaluateWhen(when *api.RuleWhen, facts *api.RuleFacts, rule *api.Rule) bool {
	condition := r.factory.CreateCondition(when.Event)

	return condition.Evaluate(when, facts, rule)
}

func (r DefaultRulesManager) executeConsequence(t api.RuleThen, d *api.Thing) {
	do := t.Do
	config := t.Config

	consequence := r.factory.CreateConsequence(do)

	if consequence != nil {
		consequence.Execute(config, r.GetContainer())
	}
}

func (r *DefaultRulesManager) RegisterRule(rule api.Rule) {
	r.rules[rule.Id] = rule
}

func NewDefaultRuleService(env api.Environment) DefaultRulesManager {
	svc := new(DefaultRulesManager)

	return *svc
}

func (s *DefaultRulesManager) SetThingManager(svc api.ThingManager) {
	s.thingManager = svc
}

func (d *DefaultRulesManager) GetContainer() api.Container {
	return d.container
}

func (s *DefaultRulesManager) SetContainer(c api.Container) {
	s.container = c
}

func (s *DefaultRulesManager) ValidateWiring() {
	if s.thingManager == nil {
		log.Fatal("thingManager not wired to DefaultRulesManager")
	}

	if s.factory == nil {
		log.Fatal("factory not wired to DefaultRulesManager")
	}
}

func NewRulesManager() *DefaultRulesManager {
	svc := new(DefaultRulesManager)
	svc.rules = make(map[string]api.Rule)

	return svc
}
