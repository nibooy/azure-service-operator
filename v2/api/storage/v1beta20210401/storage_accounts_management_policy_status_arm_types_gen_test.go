// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

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

func Test_StorageAccounts_ManagementPolicy_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_ManagementPolicy_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_ManagementPolicy_STATUS_ARM, StorageAccounts_ManagementPolicy_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_ManagementPolicy_STATUS_ARM runs a test to see if a specific instance of StorageAccounts_ManagementPolicy_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_ManagementPolicy_STATUS_ARM(subject StorageAccounts_ManagementPolicy_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_ManagementPolicy_STATUS_ARM
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

// Generator of StorageAccounts_ManagementPolicy_STATUS_ARM instances for property testing - lazily instantiated by
// StorageAccounts_ManagementPolicy_STATUS_ARMGenerator()
var storageAccounts_ManagementPolicy_STATUS_ARMGenerator gopter.Gen

// StorageAccounts_ManagementPolicy_STATUS_ARMGenerator returns a generator of StorageAccounts_ManagementPolicy_STATUS_ARM instances for property testing.
// We first initialize storageAccounts_ManagementPolicy_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccounts_ManagementPolicy_STATUS_ARMGenerator() gopter.Gen {
	if storageAccounts_ManagementPolicy_STATUS_ARMGenerator != nil {
		return storageAccounts_ManagementPolicy_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_ManagementPolicy_STATUS_ARM(generators)
	storageAccounts_ManagementPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_ManagementPolicy_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccounts_ManagementPolicy_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccounts_ManagementPolicy_STATUS_ARM(generators)
	storageAccounts_ManagementPolicy_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_ManagementPolicy_STATUS_ARM{}), generators)

	return storageAccounts_ManagementPolicy_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccounts_ManagementPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccounts_ManagementPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccounts_ManagementPolicy_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_ManagementPolicy_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagementPolicyProperties_STATUS_ARMGenerator())
}

func Test_ManagementPolicyProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyProperties_STATUS_ARM, ManagementPolicyProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyProperties_STATUS_ARM runs a test to see if a specific instance of ManagementPolicyProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyProperties_STATUS_ARM(subject ManagementPolicyProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyProperties_STATUS_ARM
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

// Generator of ManagementPolicyProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicyProperties_STATUS_ARMGenerator()
var managementPolicyProperties_STATUS_ARMGenerator gopter.Gen

// ManagementPolicyProperties_STATUS_ARMGenerator returns a generator of ManagementPolicyProperties_STATUS_ARM instances for property testing.
// We first initialize managementPolicyProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyProperties_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicyProperties_STATUS_ARMGenerator != nil {
		return managementPolicyProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyProperties_STATUS_ARM(generators)
	managementPolicyProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyProperties_STATUS_ARM(generators)
	managementPolicyProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyProperties_STATUS_ARM{}), generators)

	return managementPolicyProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagementPolicyProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Policy"] = gen.PtrOf(ManagementPolicySchema_STATUS_ARMGenerator())
}

func Test_ManagementPolicySchema_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicySchema_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicySchema_STATUS_ARM, ManagementPolicySchema_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicySchema_STATUS_ARM runs a test to see if a specific instance of ManagementPolicySchema_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicySchema_STATUS_ARM(subject ManagementPolicySchema_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicySchema_STATUS_ARM
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

// Generator of ManagementPolicySchema_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicySchema_STATUS_ARMGenerator()
var managementPolicySchema_STATUS_ARMGenerator gopter.Gen

// ManagementPolicySchema_STATUS_ARMGenerator returns a generator of ManagementPolicySchema_STATUS_ARM instances for property testing.
func ManagementPolicySchema_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicySchema_STATUS_ARMGenerator != nil {
		return managementPolicySchema_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicySchema_STATUS_ARM(generators)
	managementPolicySchema_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicySchema_STATUS_ARM{}), generators)

	return managementPolicySchema_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicySchema_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicySchema_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Rules"] = gen.SliceOf(ManagementPolicyRule_STATUS_ARMGenerator())
}

func Test_ManagementPolicyRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyRule_STATUS_ARM, ManagementPolicyRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyRule_STATUS_ARM runs a test to see if a specific instance of ManagementPolicyRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyRule_STATUS_ARM(subject ManagementPolicyRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyRule_STATUS_ARM
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

// Generator of ManagementPolicyRule_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicyRule_STATUS_ARMGenerator()
var managementPolicyRule_STATUS_ARMGenerator gopter.Gen

// ManagementPolicyRule_STATUS_ARMGenerator returns a generator of ManagementPolicyRule_STATUS_ARM instances for property testing.
// We first initialize managementPolicyRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyRule_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicyRule_STATUS_ARMGenerator != nil {
		return managementPolicyRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyRule_STATUS_ARM(generators)
	managementPolicyRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyRule_STATUS_ARM(generators)
	managementPolicyRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyRule_STATUS_ARM{}), generators)

	return managementPolicyRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ManagementPolicyRule_Type_STATUS_Lifecycle))
}

// AddRelatedPropertyGeneratorsForManagementPolicyRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Definition"] = gen.PtrOf(ManagementPolicyDefinition_STATUS_ARMGenerator())
}

func Test_ManagementPolicyDefinition_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyDefinition_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyDefinition_STATUS_ARM, ManagementPolicyDefinition_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyDefinition_STATUS_ARM runs a test to see if a specific instance of ManagementPolicyDefinition_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyDefinition_STATUS_ARM(subject ManagementPolicyDefinition_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyDefinition_STATUS_ARM
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

// Generator of ManagementPolicyDefinition_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicyDefinition_STATUS_ARMGenerator()
var managementPolicyDefinition_STATUS_ARMGenerator gopter.Gen

// ManagementPolicyDefinition_STATUS_ARMGenerator returns a generator of ManagementPolicyDefinition_STATUS_ARM instances for property testing.
func ManagementPolicyDefinition_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicyDefinition_STATUS_ARMGenerator != nil {
		return managementPolicyDefinition_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyDefinition_STATUS_ARM(generators)
	managementPolicyDefinition_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyDefinition_STATUS_ARM{}), generators)

	return managementPolicyDefinition_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyDefinition_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyDefinition_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.PtrOf(ManagementPolicyAction_STATUS_ARMGenerator())
	gens["Filters"] = gen.PtrOf(ManagementPolicyFilter_STATUS_ARMGenerator())
}

func Test_ManagementPolicyAction_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyAction_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyAction_STATUS_ARM, ManagementPolicyAction_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyAction_STATUS_ARM runs a test to see if a specific instance of ManagementPolicyAction_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyAction_STATUS_ARM(subject ManagementPolicyAction_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyAction_STATUS_ARM
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

// Generator of ManagementPolicyAction_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicyAction_STATUS_ARMGenerator()
var managementPolicyAction_STATUS_ARMGenerator gopter.Gen

// ManagementPolicyAction_STATUS_ARMGenerator returns a generator of ManagementPolicyAction_STATUS_ARM instances for property testing.
func ManagementPolicyAction_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicyAction_STATUS_ARMGenerator != nil {
		return managementPolicyAction_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyAction_STATUS_ARM(generators)
	managementPolicyAction_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyAction_STATUS_ARM{}), generators)

	return managementPolicyAction_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyAction_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyAction_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BaseBlob"] = gen.PtrOf(ManagementPolicyBaseBlob_STATUS_ARMGenerator())
	gens["Snapshot"] = gen.PtrOf(ManagementPolicySnapShot_STATUS_ARMGenerator())
	gens["Version"] = gen.PtrOf(ManagementPolicyVersion_STATUS_ARMGenerator())
}

func Test_ManagementPolicyFilter_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyFilter_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyFilter_STATUS_ARM, ManagementPolicyFilter_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyFilter_STATUS_ARM runs a test to see if a specific instance of ManagementPolicyFilter_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyFilter_STATUS_ARM(subject ManagementPolicyFilter_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyFilter_STATUS_ARM
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

// Generator of ManagementPolicyFilter_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicyFilter_STATUS_ARMGenerator()
var managementPolicyFilter_STATUS_ARMGenerator gopter.Gen

// ManagementPolicyFilter_STATUS_ARMGenerator returns a generator of ManagementPolicyFilter_STATUS_ARM instances for property testing.
// We first initialize managementPolicyFilter_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyFilter_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicyFilter_STATUS_ARMGenerator != nil {
		return managementPolicyFilter_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyFilter_STATUS_ARM(generators)
	managementPolicyFilter_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyFilter_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyFilter_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyFilter_STATUS_ARM(generators)
	managementPolicyFilter_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyFilter_STATUS_ARM{}), generators)

	return managementPolicyFilter_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyFilter_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyFilter_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BlobTypes"] = gen.SliceOf(gen.AlphaString())
	gens["PrefixMatch"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagementPolicyFilter_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyFilter_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["BlobIndexMatch"] = gen.SliceOf(TagFilter_STATUS_ARMGenerator())
}

func Test_ManagementPolicyBaseBlob_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyBaseBlob_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyBaseBlob_STATUS_ARM, ManagementPolicyBaseBlob_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyBaseBlob_STATUS_ARM runs a test to see if a specific instance of ManagementPolicyBaseBlob_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyBaseBlob_STATUS_ARM(subject ManagementPolicyBaseBlob_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyBaseBlob_STATUS_ARM
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

// Generator of ManagementPolicyBaseBlob_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicyBaseBlob_STATUS_ARMGenerator()
var managementPolicyBaseBlob_STATUS_ARMGenerator gopter.Gen

// ManagementPolicyBaseBlob_STATUS_ARMGenerator returns a generator of ManagementPolicyBaseBlob_STATUS_ARM instances for property testing.
// We first initialize managementPolicyBaseBlob_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagementPolicyBaseBlob_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicyBaseBlob_STATUS_ARMGenerator != nil {
		return managementPolicyBaseBlob_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyBaseBlob_STATUS_ARM(generators)
	managementPolicyBaseBlob_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyBaseBlob_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagementPolicyBaseBlob_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForManagementPolicyBaseBlob_STATUS_ARM(generators)
	managementPolicyBaseBlob_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyBaseBlob_STATUS_ARM{}), generators)

	return managementPolicyBaseBlob_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagementPolicyBaseBlob_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagementPolicyBaseBlob_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EnableAutoTierToHotFromCool"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForManagementPolicyBaseBlob_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyBaseBlob_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterModification_STATUS_ARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterModification_STATUS_ARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterModification_STATUS_ARMGenerator())
}

func Test_ManagementPolicySnapShot_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicySnapShot_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicySnapShot_STATUS_ARM, ManagementPolicySnapShot_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicySnapShot_STATUS_ARM runs a test to see if a specific instance of ManagementPolicySnapShot_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicySnapShot_STATUS_ARM(subject ManagementPolicySnapShot_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicySnapShot_STATUS_ARM
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

// Generator of ManagementPolicySnapShot_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicySnapShot_STATUS_ARMGenerator()
var managementPolicySnapShot_STATUS_ARMGenerator gopter.Gen

// ManagementPolicySnapShot_STATUS_ARMGenerator returns a generator of ManagementPolicySnapShot_STATUS_ARM instances for property testing.
func ManagementPolicySnapShot_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicySnapShot_STATUS_ARMGenerator != nil {
		return managementPolicySnapShot_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicySnapShot_STATUS_ARM(generators)
	managementPolicySnapShot_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicySnapShot_STATUS_ARM{}), generators)

	return managementPolicySnapShot_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicySnapShot_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicySnapShot_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterCreation_STATUS_ARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterCreation_STATUS_ARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterCreation_STATUS_ARMGenerator())
}

func Test_ManagementPolicyVersion_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagementPolicyVersion_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagementPolicyVersion_STATUS_ARM, ManagementPolicyVersion_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagementPolicyVersion_STATUS_ARM runs a test to see if a specific instance of ManagementPolicyVersion_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagementPolicyVersion_STATUS_ARM(subject ManagementPolicyVersion_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagementPolicyVersion_STATUS_ARM
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

// Generator of ManagementPolicyVersion_STATUS_ARM instances for property testing - lazily instantiated by
// ManagementPolicyVersion_STATUS_ARMGenerator()
var managementPolicyVersion_STATUS_ARMGenerator gopter.Gen

// ManagementPolicyVersion_STATUS_ARMGenerator returns a generator of ManagementPolicyVersion_STATUS_ARM instances for property testing.
func ManagementPolicyVersion_STATUS_ARMGenerator() gopter.Gen {
	if managementPolicyVersion_STATUS_ARMGenerator != nil {
		return managementPolicyVersion_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForManagementPolicyVersion_STATUS_ARM(generators)
	managementPolicyVersion_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ManagementPolicyVersion_STATUS_ARM{}), generators)

	return managementPolicyVersion_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForManagementPolicyVersion_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagementPolicyVersion_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Delete"] = gen.PtrOf(DateAfterCreation_STATUS_ARMGenerator())
	gens["TierToArchive"] = gen.PtrOf(DateAfterCreation_STATUS_ARMGenerator())
	gens["TierToCool"] = gen.PtrOf(DateAfterCreation_STATUS_ARMGenerator())
}

func Test_TagFilter_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagFilter_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagFilter_STATUS_ARM, TagFilter_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagFilter_STATUS_ARM runs a test to see if a specific instance of TagFilter_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTagFilter_STATUS_ARM(subject TagFilter_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagFilter_STATUS_ARM
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

// Generator of TagFilter_STATUS_ARM instances for property testing - lazily instantiated by
// TagFilter_STATUS_ARMGenerator()
var tagFilter_STATUS_ARMGenerator gopter.Gen

// TagFilter_STATUS_ARMGenerator returns a generator of TagFilter_STATUS_ARM instances for property testing.
func TagFilter_STATUS_ARMGenerator() gopter.Gen {
	if tagFilter_STATUS_ARMGenerator != nil {
		return tagFilter_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagFilter_STATUS_ARM(generators)
	tagFilter_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(TagFilter_STATUS_ARM{}), generators)

	return tagFilter_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTagFilter_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagFilter_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Op"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_DateAfterCreation_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DateAfterCreation_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDateAfterCreation_STATUS_ARM, DateAfterCreation_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDateAfterCreation_STATUS_ARM runs a test to see if a specific instance of DateAfterCreation_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDateAfterCreation_STATUS_ARM(subject DateAfterCreation_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DateAfterCreation_STATUS_ARM
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

// Generator of DateAfterCreation_STATUS_ARM instances for property testing - lazily instantiated by
// DateAfterCreation_STATUS_ARMGenerator()
var dateAfterCreation_STATUS_ARMGenerator gopter.Gen

// DateAfterCreation_STATUS_ARMGenerator returns a generator of DateAfterCreation_STATUS_ARM instances for property testing.
func DateAfterCreation_STATUS_ARMGenerator() gopter.Gen {
	if dateAfterCreation_STATUS_ARMGenerator != nil {
		return dateAfterCreation_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDateAfterCreation_STATUS_ARM(generators)
	dateAfterCreation_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DateAfterCreation_STATUS_ARM{}), generators)

	return dateAfterCreation_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDateAfterCreation_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDateAfterCreation_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DaysAfterCreationGreaterThan"] = gen.PtrOf(gen.Float64())
}

func Test_DateAfterModification_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DateAfterModification_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDateAfterModification_STATUS_ARM, DateAfterModification_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDateAfterModification_STATUS_ARM runs a test to see if a specific instance of DateAfterModification_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDateAfterModification_STATUS_ARM(subject DateAfterModification_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DateAfterModification_STATUS_ARM
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

// Generator of DateAfterModification_STATUS_ARM instances for property testing - lazily instantiated by
// DateAfterModification_STATUS_ARMGenerator()
var dateAfterModification_STATUS_ARMGenerator gopter.Gen

// DateAfterModification_STATUS_ARMGenerator returns a generator of DateAfterModification_STATUS_ARM instances for property testing.
func DateAfterModification_STATUS_ARMGenerator() gopter.Gen {
	if dateAfterModification_STATUS_ARMGenerator != nil {
		return dateAfterModification_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDateAfterModification_STATUS_ARM(generators)
	dateAfterModification_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DateAfterModification_STATUS_ARM{}), generators)

	return dateAfterModification_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDateAfterModification_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDateAfterModification_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DaysAfterLastAccessTimeGreaterThan"] = gen.PtrOf(gen.Float64())
	gens["DaysAfterModificationGreaterThan"] = gen.PtrOf(gen.Float64())
}
