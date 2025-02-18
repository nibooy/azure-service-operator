// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101storage

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

func Test_RouteTable_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable, RouteTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable runs a test to see if a specific instance of RouteTable round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable(subject RouteTable) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable
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

// Generator of RouteTable instances for property testing - lazily instantiated by RouteTableGenerator()
var routeTableGenerator gopter.Gen

// RouteTableGenerator returns a generator of RouteTable instances for property testing.
func RouteTableGenerator() gopter.Gen {
	if routeTableGenerator != nil {
		return routeTableGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRouteTable(generators)
	routeTableGenerator = gen.Struct(reflect.TypeOf(RouteTable{}), generators)

	return routeTableGenerator
}

// AddRelatedPropertyGeneratorsForRouteTable is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable(gens map[string]gopter.Gen) {
	gens["Spec"] = RouteTable_SpecGenerator()
	gens["Status"] = RouteTable_STATUSGenerator()
}

func Test_RouteTable_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_Spec, RouteTable_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_Spec runs a test to see if a specific instance of RouteTable_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_Spec(subject RouteTable_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_Spec
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

// Generator of RouteTable_Spec instances for property testing - lazily instantiated by RouteTable_SpecGenerator()
var routeTable_SpecGenerator gopter.Gen

// RouteTable_SpecGenerator returns a generator of RouteTable_Spec instances for property testing.
func RouteTable_SpecGenerator() gopter.Gen {
	if routeTable_SpecGenerator != nil {
		return routeTable_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec(generators)
	routeTable_SpecGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec{}), generators)

	return routeTable_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_RouteTable_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_STATUS, RouteTable_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_STATUS runs a test to see if a specific instance of RouteTable_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_STATUS(subject RouteTable_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_STATUS
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

// Generator of RouteTable_STATUS instances for property testing - lazily instantiated by RouteTable_STATUSGenerator()
var routeTable_STATUSGenerator gopter.Gen

// RouteTable_STATUSGenerator returns a generator of RouteTable_STATUS instances for property testing.
func RouteTable_STATUSGenerator() gopter.Gen {
	if routeTable_STATUSGenerator != nil {
		return routeTable_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_STATUS(generators)
	routeTable_STATUSGenerator = gen.Struct(reflect.TypeOf(RouteTable_STATUS{}), generators)

	return routeTable_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_STATUS(gens map[string]gopter.Gen) {
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
