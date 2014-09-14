// Copyright 2014 Zubair Hamed. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conditions

import "github.com/go-home/hub/api"

type Sense struct {
}

func (s Sense) Evaluate(when *api.RuleWhen, facts *api.RuleFacts, rule *api.Rule) bool {
	if when.Target != facts.Target {
		return false
	}
	return true
}

/*
exports.name = "sense";

exports.fn = function ($when, $fact, $rule) {
    var $target = $fact.target;
    if ($when.target !== $target) {
        return false;
    }
    return true;
}
*/
