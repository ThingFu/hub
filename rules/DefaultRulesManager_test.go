package rules

import (
	"github.com/thingfu/hub/api"
	"github.com/thingfu/hub/rules/conditions"
	"github.com/thingfu/hub/tests"
	"testing"
)

func TestDailyCondition(t *testing.T) {
	d := new(conditions.Daily)

	when := new(api.RuleWhen)
	facts := new(api.RuleFacts)
	rule := new(api.Rule)

	if !d.Evaluate(when, facts, rule) {
		t.Errorf("Daily Condition: Evaluation failed")
	}

	rule.LastRun = tests.NewTime("now")
	if d.Evaluate(when, facts, rule) {
		t.Errorf("Daily Condition: Evaluation failed")
	}

	rule.LastRun = tests.NewTime("yesterday")
	if !d.Evaluate(when, facts, rule) {
		t.Errorf("Daily Condition: Evaluation failed")
	}
}

func TestHourlyCondition(t *testing.T) {
	d := new(conditions.Hourly)

	when := new(api.RuleWhen)
	facts := new(api.RuleFacts)
	rule := new(api.Rule)

	if !d.Evaluate(when, facts, rule) {
		t.Errorf("Hourly Condition: Evaluation failed")
	}

	rule.LastRun = tests.NewTime("now")
	if d.Evaluate(when, facts, rule) {
		t.Errorf("Hourly Condition: Evaluation failed")
	}

	rule.LastRun = tests.NewTime("a_minute_Ago")
	if d.Evaluate(when, facts, rule) {
		t.Errorf("Hourly Condition: Evaluation failed")
	}

	rule.LastRun = tests.NewTime("hour_ago")
	if !d.Evaluate(when, facts, rule) {
		t.Errorf("Hourly Condition: Evaluation failed")
	}

	rule.LastRun = tests.NewTime("2hours_ago")
	if !d.Evaluate(when, facts, rule) {
		t.Errorf("Hourly Condition: Evaluation failed")
	}
}

func TestMonthlyCondition(t *testing.T)   {}
func TestQuarterlyCondition(t *testing.T) {}
func TestTriggerCondition(t *testing.T)   {}
func TestWeeklyCondition(t *testing.T)    {}
func TestYearlyCondition(t *testing.T)    {}
