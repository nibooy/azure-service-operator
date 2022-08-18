// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_FirewallRule_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRule_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRuleSTATUSARM, FirewallRuleSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRuleSTATUSARM runs a test to see if a specific instance of FirewallRule_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRuleSTATUSARM(subject FirewallRule_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRule_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FirewallRule_STATUSARM instances for property testing - lazily instantiated by
// FirewallRuleSTATUSARMGenerator()
var firewallRuleSTATUSARMGenerator gopter.Gen

// FirewallRuleSTATUSARMGenerator returns a generator of FirewallRule_STATUSARM instances for property testing.
// We first initialize firewallRuleSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FirewallRuleSTATUSARMGenerator() gopter.Gen {
	if firewallRuleSTATUSARMGenerator != nil {
		return firewallRuleSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleSTATUSARM(generators)
	firewallRuleSTATUSARMGenerator = gen.Struct(reflect.TypeOf(FirewallRule_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForFirewallRuleSTATUSARM(generators)
	firewallRuleSTATUSARMGenerator = gen.Struct(reflect.TypeOf(FirewallRule_STATUSARM{}), generators)

	return firewallRuleSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRuleSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRuleSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFirewallRuleSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFirewallRuleSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FirewallRulePropertiesSTATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataSTATUSARMGenerator())
}

func Test_FirewallRuleProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRuleProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRulePropertiesSTATUSARM, FirewallRulePropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRulePropertiesSTATUSARM runs a test to see if a specific instance of FirewallRuleProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRulePropertiesSTATUSARM(subject FirewallRuleProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRuleProperties_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FirewallRuleProperties_STATUSARM instances for property testing - lazily instantiated by
// FirewallRulePropertiesSTATUSARMGenerator()
var firewallRulePropertiesSTATUSARMGenerator gopter.Gen

// FirewallRulePropertiesSTATUSARMGenerator returns a generator of FirewallRuleProperties_STATUSARM instances for property testing.
func FirewallRulePropertiesSTATUSARMGenerator() gopter.Gen {
	if firewallRulePropertiesSTATUSARMGenerator != nil {
		return firewallRulePropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRulePropertiesSTATUSARM(generators)
	firewallRulePropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(FirewallRuleProperties_STATUSARM{}), generators)

	return firewallRulePropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRulePropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRulePropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}
