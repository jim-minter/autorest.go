//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import "encoding/json"

func unmarshalFirewallPolicyRuleClassification(rawMsg json.RawMessage) (FirewallPolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FirewallPolicyRuleClassification
	switch m["ruleType"] {
	case string(FirewallPolicyRuleTypeFirewallPolicyFilterRule):
		b = &FirewallPolicyFilterRule{}
	case string(FirewallPolicyRuleTypeFirewallPolicyNatRule):
		b = &FirewallPolicyNatRule{}
	default:
		b = &FirewallPolicyRule{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFirewallPolicyRuleClassificationArray(rawMsg json.RawMessage) ([]FirewallPolicyRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FirewallPolicyRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFirewallPolicyRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFirewallPolicyRuleConditionClassification(rawMsg json.RawMessage) (FirewallPolicyRuleConditionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FirewallPolicyRuleConditionClassification
	switch m["ruleConditionType"] {
	case string(FirewallPolicyRuleConditionTypeApplicationRuleCondition):
		b = &ApplicationRuleCondition{}
	case string(FirewallPolicyRuleConditionTypeNatRuleCondition):
		b = &NatRuleCondition{}
	case string(FirewallPolicyRuleConditionTypeNetworkRuleCondition):
		b = &RuleCondition{}
	default:
		b = &FirewallPolicyRuleCondition{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFirewallPolicyRuleConditionClassificationArray(rawMsg json.RawMessage) ([]FirewallPolicyRuleConditionClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FirewallPolicyRuleConditionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFirewallPolicyRuleConditionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}