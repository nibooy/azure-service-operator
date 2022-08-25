// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

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

func Test_Server_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Server_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServer_STATUSARM, Server_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServer_STATUSARM runs a test to see if a specific instance of Server_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServer_STATUSARM(subject Server_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Server_STATUSARM
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

// Generator of Server_STATUSARM instances for property testing - lazily instantiated by Server_STATUSARMGenerator()
var server_STATUSARMGenerator gopter.Gen

// Server_STATUSARMGenerator returns a generator of Server_STATUSARM instances for property testing.
// We first initialize server_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Server_STATUSARMGenerator() gopter.Gen {
	if server_STATUSARMGenerator != nil {
		return server_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_STATUSARM(generators)
	server_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Server_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServer_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForServer_STATUSARM(generators)
	server_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Server_STATUSARM{}), generators)

	return server_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServer_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServer_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServer_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServer_STATUSARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(Identity_STATUSARMGenerator())
	gens["Properties"] = gen.PtrOf(ServerProperties_STATUSARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSARMGenerator())
}

func Test_Identity_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity_STATUSARM, Identity_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity_STATUSARM runs a test to see if a specific instance of Identity_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity_STATUSARM(subject Identity_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_STATUSARM
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

// Generator of Identity_STATUSARM instances for property testing - lazily instantiated by Identity_STATUSARMGenerator()
var identity_STATUSARMGenerator gopter.Gen

// Identity_STATUSARMGenerator returns a generator of Identity_STATUSARM instances for property testing.
func Identity_STATUSARMGenerator() gopter.Gen {
	if identity_STATUSARMGenerator != nil {
		return identity_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity_STATUSARM(generators)
	identity_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Identity_STATUSARM{}), generators)

	return identity_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForIdentity_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity_STATUSARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(Identity_STATUS_Type_UserAssigned))
}

func Test_ServerProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerProperties_STATUSARM, ServerProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerProperties_STATUSARM runs a test to see if a specific instance of ServerProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServerProperties_STATUSARM(subject ServerProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerProperties_STATUSARM
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

// Generator of ServerProperties_STATUSARM instances for property testing - lazily instantiated by
// ServerProperties_STATUSARMGenerator()
var serverProperties_STATUSARMGenerator gopter.Gen

// ServerProperties_STATUSARMGenerator returns a generator of ServerProperties_STATUSARM instances for property testing.
// We first initialize serverProperties_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServerProperties_STATUSARMGenerator() gopter.Gen {
	if serverProperties_STATUSARMGenerator != nil {
		return serverProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUSARM(generators)
	serverProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerProperties_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForServerProperties_STATUSARM(generators)
	serverProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(ServerProperties_STATUSARM{}), generators)

	return serverProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForServerProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["AdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["AvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_STATUS_CreateMode_Default,
		ServerProperties_STATUS_CreateMode_GeoRestore,
		ServerProperties_STATUS_CreateMode_PointInTimeRestore,
		ServerProperties_STATUS_CreateMode_Replica))
	gens["FullyQualifiedDomainName"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicaCapacity"] = gen.PtrOf(gen.Int())
	gens["ReplicationRole"] = gen.PtrOf(gen.OneConstOf(ReplicationRole_STATUS_None, ReplicationRole_STATUS_Replica, ReplicationRole_STATUS_Source))
	gens["RestorePointInTime"] = gen.PtrOf(gen.AlphaString())
	gens["SourceServerResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		ServerProperties_STATUS_State_Disabled,
		ServerProperties_STATUS_State_Dropping,
		ServerProperties_STATUS_State_Ready,
		ServerProperties_STATUS_State_Starting,
		ServerProperties_STATUS_State_Stopped,
		ServerProperties_STATUS_State_Stopping,
		ServerProperties_STATUS_State_Updating))
	gens["Version"] = gen.PtrOf(gen.OneConstOf(ServerVersion_STATUS_57, ServerVersion_STATUS_8021))
}

// AddRelatedPropertyGeneratorsForServerProperties_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServerProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Backup"] = gen.PtrOf(Backup_STATUSARMGenerator())
	gens["DataEncryption"] = gen.PtrOf(DataEncryption_STATUSARMGenerator())
	gens["HighAvailability"] = gen.PtrOf(HighAvailability_STATUSARMGenerator())
	gens["MaintenanceWindow"] = gen.PtrOf(MaintenanceWindow_STATUSARMGenerator())
	gens["Network"] = gen.PtrOf(Network_STATUSARMGenerator())
	gens["Storage"] = gen.PtrOf(Storage_STATUSARMGenerator())
}

func Test_Sku_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_STATUSARM, Sku_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_STATUSARM runs a test to see if a specific instance of Sku_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_STATUSARM(subject Sku_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUSARM
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

// Generator of Sku_STATUSARM instances for property testing - lazily instantiated by Sku_STATUSARMGenerator()
var sku_STATUSARMGenerator gopter.Gen

// Sku_STATUSARMGenerator returns a generator of Sku_STATUSARM instances for property testing.
func Sku_STATUSARMGenerator() gopter.Gen {
	if sku_STATUSARMGenerator != nil {
		return sku_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_STATUSARM(generators)
	sku_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Sku_STATUSARM{}), generators)

	return sku_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_STATUSARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Sku_STATUS_Tier_Burstable, Sku_STATUS_Tier_GeneralPurpose, Sku_STATUS_Tier_MemoryOptimized))
}

func Test_Backup_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backup_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackup_STATUSARM, Backup_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackup_STATUSARM runs a test to see if a specific instance of Backup_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBackup_STATUSARM(subject Backup_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backup_STATUSARM
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

// Generator of Backup_STATUSARM instances for property testing - lazily instantiated by Backup_STATUSARMGenerator()
var backup_STATUSARMGenerator gopter.Gen

// Backup_STATUSARMGenerator returns a generator of Backup_STATUSARM instances for property testing.
func Backup_STATUSARMGenerator() gopter.Gen {
	if backup_STATUSARMGenerator != nil {
		return backup_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackup_STATUSARM(generators)
	backup_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Backup_STATUSARM{}), generators)

	return backup_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForBackup_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackup_STATUSARM(gens map[string]gopter.Gen) {
	gens["BackupRetentionDays"] = gen.PtrOf(gen.Int())
	gens["EarliestRestoreDate"] = gen.PtrOf(gen.AlphaString())
	gens["GeoRedundantBackup"] = gen.PtrOf(gen.OneConstOf(EnableStatusEnum_STATUS_Disabled, EnableStatusEnum_STATUS_Enabled))
}

func Test_DataEncryption_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataEncryption_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataEncryption_STATUSARM, DataEncryption_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataEncryption_STATUSARM runs a test to see if a specific instance of DataEncryption_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDataEncryption_STATUSARM(subject DataEncryption_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataEncryption_STATUSARM
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

// Generator of DataEncryption_STATUSARM instances for property testing - lazily instantiated by
// DataEncryption_STATUSARMGenerator()
var dataEncryption_STATUSARMGenerator gopter.Gen

// DataEncryption_STATUSARMGenerator returns a generator of DataEncryption_STATUSARM instances for property testing.
func DataEncryption_STATUSARMGenerator() gopter.Gen {
	if dataEncryption_STATUSARMGenerator != nil {
		return dataEncryption_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDataEncryption_STATUSARM(generators)
	dataEncryption_STATUSARMGenerator = gen.Struct(reflect.TypeOf(DataEncryption_STATUSARM{}), generators)

	return dataEncryption_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForDataEncryption_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDataEncryption_STATUSARM(gens map[string]gopter.Gen) {
	gens["GeoBackupKeyUri"] = gen.PtrOf(gen.AlphaString())
	gens["GeoBackupUserAssignedIdentityId"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryKeyUri"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentityId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(DataEncryption_STATUS_Type_AzureKeyVault, DataEncryption_STATUS_Type_SystemManaged))
}

func Test_HighAvailability_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of HighAvailability_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForHighAvailability_STATUSARM, HighAvailability_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForHighAvailability_STATUSARM runs a test to see if a specific instance of HighAvailability_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForHighAvailability_STATUSARM(subject HighAvailability_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual HighAvailability_STATUSARM
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

// Generator of HighAvailability_STATUSARM instances for property testing - lazily instantiated by
// HighAvailability_STATUSARMGenerator()
var highAvailability_STATUSARMGenerator gopter.Gen

// HighAvailability_STATUSARMGenerator returns a generator of HighAvailability_STATUSARM instances for property testing.
func HighAvailability_STATUSARMGenerator() gopter.Gen {
	if highAvailability_STATUSARMGenerator != nil {
		return highAvailability_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForHighAvailability_STATUSARM(generators)
	highAvailability_STATUSARMGenerator = gen.Struct(reflect.TypeOf(HighAvailability_STATUSARM{}), generators)

	return highAvailability_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForHighAvailability_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForHighAvailability_STATUSARM(gens map[string]gopter.Gen) {
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(HighAvailability_STATUS_Mode_Disabled, HighAvailability_STATUS_Mode_SameZone, HighAvailability_STATUS_Mode_ZoneRedundant))
	gens["StandbyAvailabilityZone"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		HighAvailability_STATUS_State_CreatingStandby,
		HighAvailability_STATUS_State_FailingOver,
		HighAvailability_STATUS_State_Healthy,
		HighAvailability_STATUS_State_NotEnabled,
		HighAvailability_STATUS_State_RemovingStandby))
}

func Test_MaintenanceWindow_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MaintenanceWindow_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMaintenanceWindow_STATUSARM, MaintenanceWindow_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMaintenanceWindow_STATUSARM runs a test to see if a specific instance of MaintenanceWindow_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMaintenanceWindow_STATUSARM(subject MaintenanceWindow_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MaintenanceWindow_STATUSARM
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

// Generator of MaintenanceWindow_STATUSARM instances for property testing - lazily instantiated by
// MaintenanceWindow_STATUSARMGenerator()
var maintenanceWindow_STATUSARMGenerator gopter.Gen

// MaintenanceWindow_STATUSARMGenerator returns a generator of MaintenanceWindow_STATUSARM instances for property testing.
func MaintenanceWindow_STATUSARMGenerator() gopter.Gen {
	if maintenanceWindow_STATUSARMGenerator != nil {
		return maintenanceWindow_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUSARM(generators)
	maintenanceWindow_STATUSARMGenerator = gen.Struct(reflect.TypeOf(MaintenanceWindow_STATUSARM{}), generators)

	return maintenanceWindow_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMaintenanceWindow_STATUSARM(gens map[string]gopter.Gen) {
	gens["CustomWindow"] = gen.PtrOf(gen.AlphaString())
	gens["DayOfWeek"] = gen.PtrOf(gen.Int())
	gens["StartHour"] = gen.PtrOf(gen.Int())
	gens["StartMinute"] = gen.PtrOf(gen.Int())
}

func Test_Network_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Network_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetwork_STATUSARM, Network_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetwork_STATUSARM runs a test to see if a specific instance of Network_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetwork_STATUSARM(subject Network_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Network_STATUSARM
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

// Generator of Network_STATUSARM instances for property testing - lazily instantiated by Network_STATUSARMGenerator()
var network_STATUSARMGenerator gopter.Gen

// Network_STATUSARMGenerator returns a generator of Network_STATUSARM instances for property testing.
func Network_STATUSARMGenerator() gopter.Gen {
	if network_STATUSARMGenerator != nil {
		return network_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetwork_STATUSARM(generators)
	network_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Network_STATUSARM{}), generators)

	return network_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForNetwork_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetwork_STATUSARM(gens map[string]gopter.Gen) {
	gens["DelegatedSubnetResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateDnsZoneResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(EnableStatusEnum_STATUS_Disabled, EnableStatusEnum_STATUS_Enabled))
}

func Test_Storage_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Storage_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorage_STATUSARM, Storage_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorage_STATUSARM runs a test to see if a specific instance of Storage_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorage_STATUSARM(subject Storage_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Storage_STATUSARM
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

// Generator of Storage_STATUSARM instances for property testing - lazily instantiated by Storage_STATUSARMGenerator()
var storage_STATUSARMGenerator gopter.Gen

// Storage_STATUSARMGenerator returns a generator of Storage_STATUSARM instances for property testing.
func Storage_STATUSARMGenerator() gopter.Gen {
	if storage_STATUSARMGenerator != nil {
		return storage_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorage_STATUSARM(generators)
	storage_STATUSARMGenerator = gen.Struct(reflect.TypeOf(Storage_STATUSARM{}), generators)

	return storage_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForStorage_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorage_STATUSARM(gens map[string]gopter.Gen) {
	gens["AutoGrow"] = gen.PtrOf(gen.OneConstOf(EnableStatusEnum_STATUS_Disabled, EnableStatusEnum_STATUS_Enabled))
	gens["Iops"] = gen.PtrOf(gen.Int())
	gens["StorageSizeGB"] = gen.PtrOf(gen.Int())
	gens["StorageSku"] = gen.PtrOf(gen.AlphaString())
}