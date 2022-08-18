// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601storage

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

func Test_Configuration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfiguration, ConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfiguration runs a test to see if a specific instance of Configuration round trips to JSON and back losslessly
func RunJSONSerializationTestForConfiguration(subject Configuration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration
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

// Generator of Configuration instances for property testing - lazily instantiated by ConfigurationGenerator()
var configurationGenerator gopter.Gen

// ConfigurationGenerator returns a generator of Configuration instances for property testing.
func ConfigurationGenerator() gopter.Gen {
	if configurationGenerator != nil {
		return configurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForConfiguration(generators)
	configurationGenerator = gen.Struct(reflect.TypeOf(Configuration{}), generators)

	return configurationGenerator
}

// AddRelatedPropertyGeneratorsForConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersConfigurationsSpecGenerator()
	gens["Status"] = ConfigurationSTATUSGenerator()
}

func Test_Configuration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationSTATUS, ConfigurationSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationSTATUS runs a test to see if a specific instance of Configuration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationSTATUS(subject Configuration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration_STATUS
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

// Generator of Configuration_STATUS instances for property testing - lazily instantiated by
// ConfigurationSTATUSGenerator()
var configurationSTATUSGenerator gopter.Gen

// ConfigurationSTATUSGenerator returns a generator of Configuration_STATUS instances for property testing.
func ConfigurationSTATUSGenerator() gopter.Gen {
	if configurationSTATUSGenerator != nil {
		return configurationSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationSTATUS(generators)
	configurationSTATUSGenerator = gen.Struct(reflect.TypeOf(Configuration_STATUS{}), generators)

	return configurationSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationSTATUS(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServersConfigurations_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersConfigurations_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersConfigurationsSpec, ServersConfigurationsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersConfigurationsSpec runs a test to see if a specific instance of ServersConfigurations_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersConfigurationsSpec(subject ServersConfigurations_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersConfigurations_Spec
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

// Generator of ServersConfigurations_Spec instances for property testing - lazily instantiated by
// ServersConfigurationsSpecGenerator()
var serversConfigurationsSpecGenerator gopter.Gen

// ServersConfigurationsSpecGenerator returns a generator of ServersConfigurations_Spec instances for property testing.
func ServersConfigurationsSpecGenerator() gopter.Gen {
	if serversConfigurationsSpecGenerator != nil {
		return serversConfigurationsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersConfigurationsSpec(generators)
	serversConfigurationsSpecGenerator = gen.Struct(reflect.TypeOf(ServersConfigurations_Spec{}), generators)

	return serversConfigurationsSpecGenerator
}

// AddIndependentPropertyGeneratorsForServersConfigurationsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersConfigurationsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
