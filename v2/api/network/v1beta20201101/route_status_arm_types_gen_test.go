// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

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

func Test_Route_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Route_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoute_STATUSARM, Route_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoute_STATUSARM runs a test to see if a specific instance of Route_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoute_STATUSARM(subject Route_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Route_STATUSARM
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

// Generator of Route_STATUSARM instances for property testing - lazily instantiated by Route_STATUSARMGenerator()
var route_STATUSARMGenerator gopter.Gen

// Route_STATUSARMGenerator returns a generator of Route_STATUSARM instances for property testing.
// We first initialize route_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Route_STATUSARMGenerator() gopter.Gen {
	if route_STATUSARMGenerator != nil {
		return route_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoute_STATUSARM(generators)
	route_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Route_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoute_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForRoute_STATUSARM(generators)
	route_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Route_STATUSARM{}), generators)

	return route_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRoute_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoute_STATUSARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoute_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoute_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoutePropertiesFormat_STATUSARMGenerator())
}

func Test_RoutePropertiesFormat_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoutePropertiesFormat_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoutePropertiesFormat_STATUSARM, RoutePropertiesFormat_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoutePropertiesFormat_STATUSARM runs a test to see if a specific instance of RoutePropertiesFormat_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoutePropertiesFormat_STATUSARM(subject RoutePropertiesFormat_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoutePropertiesFormat_STATUSARM
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

// Generator of RoutePropertiesFormat_STATUSARM instances for property testing - lazily instantiated by
// RoutePropertiesFormat_STATUSARMGenerator()
var routePropertiesFormat_STATUSARMGenerator gopter.Gen

// RoutePropertiesFormat_STATUSARMGenerator returns a generator of RoutePropertiesFormat_STATUSARM instances for property testing.
func RoutePropertiesFormat_STATUSARMGenerator() gopter.Gen {
	if routePropertiesFormat_STATUSARMGenerator != nil {
		return routePropertiesFormat_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoutePropertiesFormat_STATUSARM(generators)
	routePropertiesFormat_STATUSARMGenerator = gen.Struct(reflect.TypeOf(RoutePropertiesFormat_STATUSARM{}), generators)

	return routePropertiesFormat_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForRoutePropertiesFormat_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoutePropertiesFormat_STATUSARM(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["HasBgpOverride"] = gen.PtrOf(gen.Bool())
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.OneConstOf(
		RouteNextHopType_STATUS_Internet,
		RouteNextHopType_STATUS_None,
		RouteNextHopType_STATUS_VirtualAppliance,
		RouteNextHopType_STATUS_VirtualNetworkGateway,
		RouteNextHopType_STATUS_VnetLocal))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
}