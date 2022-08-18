// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

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

func Test_ManagedClustersAgentPools_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClustersAgentPools_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClustersAgentPoolsSpecARM, ManagedClustersAgentPoolsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClustersAgentPoolsSpecARM runs a test to see if a specific instance of ManagedClustersAgentPools_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClustersAgentPoolsSpecARM(subject ManagedClustersAgentPools_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClustersAgentPools_SpecARM
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

// Generator of ManagedClustersAgentPools_SpecARM instances for property testing - lazily instantiated by
// ManagedClustersAgentPoolsSpecARMGenerator()
var managedClustersAgentPoolsSpecARMGenerator gopter.Gen

// ManagedClustersAgentPoolsSpecARMGenerator returns a generator of ManagedClustersAgentPools_SpecARM instances for property testing.
// We first initialize managedClustersAgentPoolsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClustersAgentPoolsSpecARMGenerator() gopter.Gen {
	if managedClustersAgentPoolsSpecARMGenerator != nil {
		return managedClustersAgentPoolsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpecARM(generators)
	managedClustersAgentPoolsSpecARMGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPools_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpecARM(generators)
	AddRelatedPropertyGeneratorsForManagedClustersAgentPoolsSpecARM(generators)
	managedClustersAgentPoolsSpecARMGenerator = gen.Struct(reflect.TypeOf(ManagedClustersAgentPools_SpecARM{}), generators)

	return managedClustersAgentPoolsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClustersAgentPoolsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForManagedClustersAgentPoolsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClustersAgentPoolsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagedClusterAgentPoolProfilePropertiesARMGenerator())
}

func Test_ManagedClusterAgentPoolProfilePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusterAgentPoolProfilePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusterAgentPoolProfilePropertiesARM, ManagedClusterAgentPoolProfilePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusterAgentPoolProfilePropertiesARM runs a test to see if a specific instance of ManagedClusterAgentPoolProfilePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusterAgentPoolProfilePropertiesARM(subject ManagedClusterAgentPoolProfilePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusterAgentPoolProfilePropertiesARM
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

// Generator of ManagedClusterAgentPoolProfilePropertiesARM instances for property testing - lazily instantiated by
// ManagedClusterAgentPoolProfilePropertiesARMGenerator()
var managedClusterAgentPoolProfilePropertiesARMGenerator gopter.Gen

// ManagedClusterAgentPoolProfilePropertiesARMGenerator returns a generator of ManagedClusterAgentPoolProfilePropertiesARM instances for property testing.
// We first initialize managedClusterAgentPoolProfilePropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusterAgentPoolProfilePropertiesARMGenerator() gopter.Gen {
	if managedClusterAgentPoolProfilePropertiesARMGenerator != nil {
		return managedClusterAgentPoolProfilePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesARM(generators)
	managedClusterAgentPoolProfilePropertiesARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfilePropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesARM(generators)
	managedClusterAgentPoolProfilePropertiesARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfilePropertiesARM{}), generators)

	return managedClusterAgentPoolProfilePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesARM(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.OneConstOf(
		ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile_MIG1G,
		ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile_MIG2G,
		ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile_MIG3G,
		ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile_MIG4G,
		ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile_MIG7G))
	gens["KubeletDiskType"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesKubeletDiskType_OS, ManagedClusterAgentPoolProfilePropertiesKubeletDiskType_Temporary))
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesMode_System, ManagedClusterAgentPoolProfilePropertiesMode_User))
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesOsDiskType_Ephemeral, ManagedClusterAgentPoolProfilePropertiesOsDiskType_Managed))
	gens["OsSKU"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesOsSKU_CBLMariner, ManagedClusterAgentPoolProfilePropertiesOsSKU_Ubuntu))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesOsType_Linux, ManagedClusterAgentPoolProfilePropertiesOsType_Windows))
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy_Deallocate, ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy_Delete))
	gens["ScaleSetPriority"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesScaleSetPriority_Regular, ManagedClusterAgentPoolProfilePropertiesScaleSetPriority_Spot))
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ManagedClusterAgentPoolProfilePropertiesType_AvailabilitySet, ManagedClusterAgentPoolProfilePropertiesType_VirtualMachineScaleSets))
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesARM(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfigARMGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfigARMGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettingsARMGenerator())
}

func Test_AgentPoolUpgradeSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettingsARM, AgentPoolUpgradeSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettingsARM runs a test to see if a specific instance of AgentPoolUpgradeSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettingsARM(subject AgentPoolUpgradeSettingsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettingsARM
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

// Generator of AgentPoolUpgradeSettingsARM instances for property testing - lazily instantiated by
// AgentPoolUpgradeSettingsARMGenerator()
var agentPoolUpgradeSettingsARMGenerator gopter.Gen

// AgentPoolUpgradeSettingsARMGenerator returns a generator of AgentPoolUpgradeSettingsARM instances for property testing.
func AgentPoolUpgradeSettingsARMGenerator() gopter.Gen {
	if agentPoolUpgradeSettingsARMGenerator != nil {
		return agentPoolUpgradeSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsARM(generators)
	agentPoolUpgradeSettingsARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettingsARM{}), generators)

	return agentPoolUpgradeSettingsARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsARM(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfigARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfigARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfigARM, KubeletConfigARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfigARM runs a test to see if a specific instance of KubeletConfigARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfigARM(subject KubeletConfigARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfigARM
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

// Generator of KubeletConfigARM instances for property testing - lazily instantiated by KubeletConfigARMGenerator()
var kubeletConfigARMGenerator gopter.Gen

// KubeletConfigARMGenerator returns a generator of KubeletConfigARM instances for property testing.
func KubeletConfigARMGenerator() gopter.Gen {
	if kubeletConfigARMGenerator != nil {
		return kubeletConfigARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfigARM(generators)
	kubeletConfigARMGenerator = gen.Struct(reflect.TypeOf(KubeletConfigARM{}), generators)

	return kubeletConfigARMGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfigARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfigARM(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_LinuxOSConfigARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfigARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfigARM, LinuxOSConfigARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfigARM runs a test to see if a specific instance of LinuxOSConfigARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfigARM(subject LinuxOSConfigARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfigARM
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

// Generator of LinuxOSConfigARM instances for property testing - lazily instantiated by LinuxOSConfigARMGenerator()
var linuxOSConfigARMGenerator gopter.Gen

// LinuxOSConfigARMGenerator returns a generator of LinuxOSConfigARM instances for property testing.
// We first initialize linuxOSConfigARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfigARMGenerator() gopter.Gen {
	if linuxOSConfigARMGenerator != nil {
		return linuxOSConfigARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfigARM(generators)
	linuxOSConfigARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfigARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfigARM(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfigARM(generators)
	linuxOSConfigARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfigARM{}), generators)

	return linuxOSConfigARMGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfigARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfigARM(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfigARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfigARM(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfigARMGenerator())
}

func Test_SysctlConfigARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfigARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfigARM, SysctlConfigARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfigARM runs a test to see if a specific instance of SysctlConfigARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfigARM(subject SysctlConfigARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfigARM
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

// Generator of SysctlConfigARM instances for property testing - lazily instantiated by SysctlConfigARMGenerator()
var sysctlConfigARMGenerator gopter.Gen

// SysctlConfigARMGenerator returns a generator of SysctlConfigARM instances for property testing.
func SysctlConfigARMGenerator() gopter.Gen {
	if sysctlConfigARMGenerator != nil {
		return sysctlConfigARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfigARM(generators)
	sysctlConfigARMGenerator = gen.Struct(reflect.TypeOf(SysctlConfigARM{}), generators)

	return sysctlConfigARMGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfigARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfigARM(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}
