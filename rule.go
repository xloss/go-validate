package go_validate

import "strings"

type Rule struct {
	Name   string
	Params []string
}

func (rule *Rule) Setup(text string) {
	ruleData := strings.SplitN(text, ":", 2)

	rule.Name = strings.ToLower(ruleData[0])

	if len(ruleData) > 1 {
		rule.Params = strings.Split(ruleData[1], ",")
	}
}
