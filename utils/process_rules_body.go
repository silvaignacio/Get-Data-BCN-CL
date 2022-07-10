package utils

import (
	"encoding/xml"
	"example.com/main/model/all"
	"example.com/main/model/by"
	"io"
	"io/ioutil"
)

var (
	r  all.ChileanRulesResponse
	rr by.ChileanRuleResponse
)

func ProcessRulesBody(body io.ReadCloser) all.RulesResult {
	data, _ := ioutil.ReadAll(body)
	result := all.RulesResult{}

	_ = xml.Unmarshal([]byte(data), &result)
	return result
}

func ProcessRuleBody(body io.ReadCloser) by.RuleResult {
	data, _ := ioutil.ReadAll(body)
	result := by.RuleResult{}

	_ = xml.Unmarshal([]byte(data), &result)
	return result
}
