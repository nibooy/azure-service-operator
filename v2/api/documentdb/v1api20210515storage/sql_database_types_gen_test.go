// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515storage

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

func Test_SqlDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabase, SqlDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabase runs a test to see if a specific instance of SqlDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabase(subject SqlDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabase
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

// Generator of SqlDatabase instances for property testing - lazily instantiated by SqlDatabaseGenerator()
var sqlDatabaseGenerator gopter.Gen

// SqlDatabaseGenerator returns a generator of SqlDatabase instances for property testing.
func SqlDatabaseGenerator() gopter.Gen {
	if sqlDatabaseGenerator != nil {
		return sqlDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabase(generators)
	sqlDatabaseGenerator = gen.Struct(reflect.TypeOf(SqlDatabase{}), generators)

	return sqlDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_SqlDatabase_SpecGenerator()
	gens["Status"] = DatabaseAccounts_SqlDatabase_STATUSGenerator()
}

func Test_DatabaseAccounts_SqlDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabase_Spec, DatabaseAccounts_SqlDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabase_Spec runs a test to see if a specific instance of DatabaseAccounts_SqlDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabase_Spec(subject DatabaseAccounts_SqlDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabase_Spec
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

// Generator of DatabaseAccounts_SqlDatabase_Spec instances for property testing - lazily instantiated by
// DatabaseAccounts_SqlDatabase_SpecGenerator()
var databaseAccounts_SqlDatabase_SpecGenerator gopter.Gen

// DatabaseAccounts_SqlDatabase_SpecGenerator returns a generator of DatabaseAccounts_SqlDatabase_Spec instances for property testing.
// We first initialize databaseAccounts_SqlDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabase_SpecGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabase_SpecGenerator != nil {
		return databaseAccounts_SqlDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_Spec(generators)
	databaseAccounts_SqlDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabase_Spec(generators)
	databaseAccounts_SqlDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabase_Spec{}), generators)

	return databaseAccounts_SqlDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseResourceGenerator())
}

func Test_DatabaseAccounts_SqlDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_SqlDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_SqlDatabase_STATUS, DatabaseAccounts_SqlDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_SqlDatabase_STATUS runs a test to see if a specific instance of DatabaseAccounts_SqlDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_SqlDatabase_STATUS(subject DatabaseAccounts_SqlDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_SqlDatabase_STATUS
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

// Generator of DatabaseAccounts_SqlDatabase_STATUS instances for property testing - lazily instantiated by
// DatabaseAccounts_SqlDatabase_STATUSGenerator()
var databaseAccounts_SqlDatabase_STATUSGenerator gopter.Gen

// DatabaseAccounts_SqlDatabase_STATUSGenerator returns a generator of DatabaseAccounts_SqlDatabase_STATUS instances for property testing.
// We first initialize databaseAccounts_SqlDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_SqlDatabase_STATUSGenerator() gopter.Gen {
	if databaseAccounts_SqlDatabase_STATUSGenerator != nil {
		return databaseAccounts_SqlDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_STATUS(generators)
	databaseAccounts_SqlDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabase_STATUS(generators)
	databaseAccounts_SqlDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_SqlDatabase_STATUS{}), generators)

	return databaseAccounts_SqlDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_SqlDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_SqlDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUSGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseGetProperties_Resource_STATUSGenerator())
}

func Test_SqlDatabaseGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseGetProperties_Resource_STATUS, SqlDatabaseGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseGetProperties_Resource_STATUS runs a test to see if a specific instance of SqlDatabaseGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseGetProperties_Resource_STATUS(subject SqlDatabaseGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseGetProperties_Resource_STATUS
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

// Generator of SqlDatabaseGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// SqlDatabaseGetProperties_Resource_STATUSGenerator()
var sqlDatabaseGetProperties_Resource_STATUSGenerator gopter.Gen

// SqlDatabaseGetProperties_Resource_STATUSGenerator returns a generator of SqlDatabaseGetProperties_Resource_STATUS instances for property testing.
func SqlDatabaseGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if sqlDatabaseGetProperties_Resource_STATUSGenerator != nil {
		return sqlDatabaseGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS(generators)
	sqlDatabaseGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetProperties_Resource_STATUS{}), generators)

	return sqlDatabaseGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Colls"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
	gens["Users"] = gen.PtrOf(gen.AlphaString())
}

func Test_SqlDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseResource, SqlDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseResource runs a test to see if a specific instance of SqlDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseResource(subject SqlDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseResource
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

// Generator of SqlDatabaseResource instances for property testing - lazily instantiated by
// SqlDatabaseResourceGenerator()
var sqlDatabaseResourceGenerator gopter.Gen

// SqlDatabaseResourceGenerator returns a generator of SqlDatabaseResource instances for property testing.
func SqlDatabaseResourceGenerator() gopter.Gen {
	if sqlDatabaseResourceGenerator != nil {
		return sqlDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseResource(generators)
	sqlDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseResource{}), generators)

	return sqlDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
