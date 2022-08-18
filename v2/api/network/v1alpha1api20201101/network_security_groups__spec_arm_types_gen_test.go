// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_NetworkSecurityGroups_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroups_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSpecARM, NetworkSecurityGroupsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSpecARM runs a test to see if a specific instance of NetworkSecurityGroups_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSpecARM(subject NetworkSecurityGroups_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroups_SpecARM
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

// Generator of NetworkSecurityGroups_SpecARM instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSpecARMGenerator()
var networkSecurityGroupsSpecARMGenerator gopter.Gen

// NetworkSecurityGroupsSpecARMGenerator returns a generator of NetworkSecurityGroups_SpecARM instances for property testing.
func NetworkSecurityGroupsSpecARMGenerator() gopter.Gen {
	if networkSecurityGroupsSpecARMGenerator != nil {
		return networkSecurityGroupsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSpecARM(generators)
	networkSecurityGroupsSpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroups_SpecARM{}), generators)

	return networkSecurityGroupsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
