// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import (
	"encoding/json"
	alpha20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
	v20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401storage"
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

func Test_StorageAccountsQueueServicesQueue_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueue to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsQueueServicesQueue, StorageAccountsQueueServicesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsQueueServicesQueue tests if a specific instance of StorageAccountsQueueServicesQueue round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsQueueServicesQueue(subject StorageAccountsQueueServicesQueue) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20210401s.StorageAccountsQueueServicesQueue
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsQueueServicesQueue
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsQueueServicesQueue_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueue to StorageAccountsQueueServicesQueue via AssignPropertiesToStorageAccountsQueueServicesQueue & AssignPropertiesFromStorageAccountsQueueServicesQueue returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue, StorageAccountsQueueServicesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue tests if a specific instance of StorageAccountsQueueServicesQueue can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesQueue(subject StorageAccountsQueueServicesQueue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210401s.StorageAccountsQueueServicesQueue
	err := copied.AssignPropertiesToStorageAccountsQueueServicesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueServicesQueue
	err = actual.AssignPropertiesFromStorageAccountsQueueServicesQueue(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsQueueServicesQueue_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueue via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueue, StorageAccountsQueueServicesQueueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueue runs a test to see if a specific instance of StorageAccountsQueueServicesQueue round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueue(subject StorageAccountsQueueServicesQueue) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueue
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

// Generator of StorageAccountsQueueServicesQueue instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueueGenerator()
var storageAccountsQueueServicesQueueGenerator gopter.Gen

// StorageAccountsQueueServicesQueueGenerator returns a generator of StorageAccountsQueueServicesQueue instances for property testing.
func StorageAccountsQueueServicesQueueGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueueGenerator != nil {
		return storageAccountsQueueServicesQueueGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue(generators)
	storageAccountsQueueServicesQueueGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueue{}), generators)

	return storageAccountsQueueServicesQueueGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsQueueServicesQueue(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsQueueServicesQueuesSpecGenerator()
	gens["Status"] = StorageQueueSTATUSGenerator()
}

func Test_StorageAccountsQueueServicesQueues_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsQueueServicesQueues_Spec to StorageAccountsQueueServicesQueues_Spec via AssignPropertiesToStorageAccountsQueueServicesQueuesSpec & AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsQueueServicesQueuesSpec, StorageAccountsQueueServicesQueuesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsQueueServicesQueuesSpec tests if a specific instance of StorageAccountsQueueServicesQueues_Spec can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsQueueServicesQueuesSpec(subject StorageAccountsQueueServicesQueues_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210401s.StorageAccountsQueueServicesQueues_Spec
	err := copied.AssignPropertiesToStorageAccountsQueueServicesQueuesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsQueueServicesQueues_Spec
	err = actual.AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsQueueServicesQueues_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsQueueServicesQueues_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSpec, StorageAccountsQueueServicesQueuesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSpec runs a test to see if a specific instance of StorageAccountsQueueServicesQueues_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsQueueServicesQueuesSpec(subject StorageAccountsQueueServicesQueues_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsQueueServicesQueues_Spec
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

// Generator of StorageAccountsQueueServicesQueues_Spec instances for property testing - lazily instantiated by
// StorageAccountsQueueServicesQueuesSpecGenerator()
var storageAccountsQueueServicesQueuesSpecGenerator gopter.Gen

// StorageAccountsQueueServicesQueuesSpecGenerator returns a generator of StorageAccountsQueueServicesQueues_Spec instances for property testing.
func StorageAccountsQueueServicesQueuesSpecGenerator() gopter.Gen {
	if storageAccountsQueueServicesQueuesSpecGenerator != nil {
		return storageAccountsQueueServicesQueuesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpec(generators)
	storageAccountsQueueServicesQueuesSpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsQueueServicesQueues_Spec{}), generators)

	return storageAccountsQueueServicesQueuesSpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsQueueServicesQueuesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

func Test_StorageQueue_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageQueue_STATUS to StorageQueue_STATUS via AssignPropertiesToStorageQueueSTATUS & AssignPropertiesFromStorageQueueSTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageQueueSTATUS, StorageQueueSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageQueueSTATUS tests if a specific instance of StorageQueue_STATUS can be assigned to v1alpha1api20210401storage and back losslessly
func RunPropertyAssignmentTestForStorageQueueSTATUS(subject StorageQueue_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other alpha20210401s.StorageQueue_STATUS
	err := copied.AssignPropertiesToStorageQueueSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageQueue_STATUS
	err = actual.AssignPropertiesFromStorageQueueSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageQueue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageQueue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageQueueSTATUS, StorageQueueSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageQueueSTATUS runs a test to see if a specific instance of StorageQueue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageQueueSTATUS(subject StorageQueue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageQueue_STATUS
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

// Generator of StorageQueue_STATUS instances for property testing - lazily instantiated by StorageQueueSTATUSGenerator()
var storageQueueSTATUSGenerator gopter.Gen

// StorageQueueSTATUSGenerator returns a generator of StorageQueue_STATUS instances for property testing.
func StorageQueueSTATUSGenerator() gopter.Gen {
	if storageQueueSTATUSGenerator != nil {
		return storageQueueSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageQueueSTATUS(generators)
	storageQueueSTATUSGenerator = gen.Struct(reflect.TypeOf(StorageQueue_STATUS{}), generators)

	return storageQueueSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageQueueSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageQueueSTATUS(gens map[string]gopter.Gen) {
	gens["ApproximateMessageCount"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
