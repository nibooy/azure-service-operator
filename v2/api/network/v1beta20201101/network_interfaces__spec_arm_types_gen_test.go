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

func Test_NetworkInterfaces_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaces_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfacesSpecARM, NetworkInterfacesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfacesSpecARM runs a test to see if a specific instance of NetworkInterfaces_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfacesSpecARM(subject NetworkInterfaces_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaces_SpecARM
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

// Generator of NetworkInterfaces_SpecARM instances for property testing - lazily instantiated by
// NetworkInterfacesSpecARMGenerator()
var networkInterfacesSpecARMGenerator gopter.Gen

// NetworkInterfacesSpecARMGenerator returns a generator of NetworkInterfaces_SpecARM instances for property testing.
// We first initialize networkInterfacesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfacesSpecARMGenerator() gopter.Gen {
	if networkInterfacesSpecARMGenerator != nil {
		return networkInterfacesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfacesSpecARM(generators)
	networkInterfacesSpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaces_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfacesSpecARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfacesSpecARM(generators)
	networkInterfacesSpecARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaces_SpecARM{}), generators)

	return networkInterfacesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfacesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfacesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterfacesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfacesSpecARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationARMGenerator())
	gens["Properties"] = gen.PtrOf(NetworkInterfacesSpecPropertiesARMGenerator())
}

func Test_NetworkInterfaces_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaces_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfacesSpecPropertiesARM, NetworkInterfacesSpecPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfacesSpecPropertiesARM runs a test to see if a specific instance of NetworkInterfaces_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfacesSpecPropertiesARM(subject NetworkInterfaces_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaces_Spec_PropertiesARM
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

// Generator of NetworkInterfaces_Spec_PropertiesARM instances for property testing - lazily instantiated by
// NetworkInterfacesSpecPropertiesARMGenerator()
var networkInterfacesSpecPropertiesARMGenerator gopter.Gen

// NetworkInterfacesSpecPropertiesARMGenerator returns a generator of NetworkInterfaces_Spec_PropertiesARM instances for property testing.
// We first initialize networkInterfacesSpecPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfacesSpecPropertiesARMGenerator() gopter.Gen {
	if networkInterfacesSpecPropertiesARMGenerator != nil {
		return networkInterfacesSpecPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesARM(generators)
	networkInterfacesSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaces_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfacesSpecPropertiesARM(generators)
	networkInterfacesSpecPropertiesARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaces_Spec_PropertiesARM{}), generators)

	return networkInterfacesSpecPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["EnableAcceleratedNetworking"] = gen.PtrOf(gen.Bool())
	gens["EnableIPForwarding"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNetworkInterfacesSpecPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfacesSpecPropertiesARM(gens map[string]gopter.Gen) {
	gens["DnsSettings"] = gen.PtrOf(NetworkInterfaceDnsSettingsARMGenerator())
	gens["IpConfigurations"] = gen.SliceOf(NetworkInterfacesSpecPropertiesIpConfigurationsARMGenerator())
	gens["NetworkSecurityGroup"] = gen.PtrOf(SubResourceARMGenerator())
}

func Test_NetworkInterfaceDnsSettingsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaceDnsSettingsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceDnsSettingsARM, NetworkInterfaceDnsSettingsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceDnsSettingsARM runs a test to see if a specific instance of NetworkInterfaceDnsSettingsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceDnsSettingsARM(subject NetworkInterfaceDnsSettingsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaceDnsSettingsARM
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

// Generator of NetworkInterfaceDnsSettingsARM instances for property testing - lazily instantiated by
// NetworkInterfaceDnsSettingsARMGenerator()
var networkInterfaceDnsSettingsARMGenerator gopter.Gen

// NetworkInterfaceDnsSettingsARMGenerator returns a generator of NetworkInterfaceDnsSettingsARM instances for property testing.
func NetworkInterfaceDnsSettingsARMGenerator() gopter.Gen {
	if networkInterfaceDnsSettingsARMGenerator != nil {
		return networkInterfaceDnsSettingsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceDnsSettingsARM(generators)
	networkInterfaceDnsSettingsARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceDnsSettingsARM{}), generators)

	return networkInterfaceDnsSettingsARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceDnsSettingsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceDnsSettingsARM(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
	gens["InternalDnsNameLabel"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterfaces_Spec_Properties_IpConfigurationsARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaces_Spec_Properties_IpConfigurationsARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfacesSpecPropertiesIpConfigurationsARM, NetworkInterfacesSpecPropertiesIpConfigurationsARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfacesSpecPropertiesIpConfigurationsARM runs a test to see if a specific instance of NetworkInterfaces_Spec_Properties_IpConfigurationsARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfacesSpecPropertiesIpConfigurationsARM(subject NetworkInterfaces_Spec_Properties_IpConfigurationsARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaces_Spec_Properties_IpConfigurationsARM
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

// Generator of NetworkInterfaces_Spec_Properties_IpConfigurationsARM instances for property testing - lazily
// instantiated by NetworkInterfacesSpecPropertiesIpConfigurationsARMGenerator()
var networkInterfacesSpecPropertiesIpConfigurationsARMGenerator gopter.Gen

// NetworkInterfacesSpecPropertiesIpConfigurationsARMGenerator returns a generator of NetworkInterfaces_Spec_Properties_IpConfigurationsARM instances for property testing.
// We first initialize networkInterfacesSpecPropertiesIpConfigurationsARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfacesSpecPropertiesIpConfigurationsARMGenerator() gopter.Gen {
	if networkInterfacesSpecPropertiesIpConfigurationsARMGenerator != nil {
		return networkInterfacesSpecPropertiesIpConfigurationsARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesIpConfigurationsARM(generators)
	networkInterfacesSpecPropertiesIpConfigurationsARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaces_Spec_Properties_IpConfigurationsARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesIpConfigurationsARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfacesSpecPropertiesIpConfigurationsARM(generators)
	networkInterfacesSpecPropertiesIpConfigurationsARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaces_Spec_Properties_IpConfigurationsARM{}), generators)

	return networkInterfacesSpecPropertiesIpConfigurationsARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesIpConfigurationsARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfacesSpecPropertiesIpConfigurationsARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterfacesSpecPropertiesIpConfigurationsARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfacesSpecPropertiesIpConfigurationsARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkInterfaceIPConfigurationPropertiesFormatARMGenerator())
}

func Test_SubResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResourceARM, SubResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResourceARM runs a test to see if a specific instance of SubResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResourceARM(subject SubResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResourceARM
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

// Generator of SubResourceARM instances for property testing - lazily instantiated by SubResourceARMGenerator()
var subResourceARMGenerator gopter.Gen

// SubResourceARMGenerator returns a generator of SubResourceARM instances for property testing.
func SubResourceARMGenerator() gopter.Gen {
	if subResourceARMGenerator != nil {
		return subResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResourceARM(generators)
	subResourceARMGenerator = gen.Struct(reflect.TypeOf(SubResourceARM{}), generators)

	return subResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResourceARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterfaceIPConfigurationPropertiesFormatARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterfaceIPConfigurationPropertiesFormatARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceIPConfigurationPropertiesFormatARM, NetworkInterfaceIPConfigurationPropertiesFormatARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceIPConfigurationPropertiesFormatARM runs a test to see if a specific instance of NetworkInterfaceIPConfigurationPropertiesFormatARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceIPConfigurationPropertiesFormatARM(subject NetworkInterfaceIPConfigurationPropertiesFormatARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterfaceIPConfigurationPropertiesFormatARM
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

// Generator of NetworkInterfaceIPConfigurationPropertiesFormatARM instances for property testing - lazily instantiated
// by NetworkInterfaceIPConfigurationPropertiesFormatARMGenerator()
var networkInterfaceIPConfigurationPropertiesFormatARMGenerator gopter.Gen

// NetworkInterfaceIPConfigurationPropertiesFormatARMGenerator returns a generator of NetworkInterfaceIPConfigurationPropertiesFormatARM instances for property testing.
// We first initialize networkInterfaceIPConfigurationPropertiesFormatARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfaceIPConfigurationPropertiesFormatARMGenerator() gopter.Gen {
	if networkInterfaceIPConfigurationPropertiesFormatARMGenerator != nil {
		return networkInterfaceIPConfigurationPropertiesFormatARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormatARM(generators)
	networkInterfaceIPConfigurationPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceIPConfigurationPropertiesFormatARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormatARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormatARM(generators)
	networkInterfaceIPConfigurationPropertiesFormatARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterfaceIPConfigurationPropertiesFormatARM{}), generators)

	return networkInterfaceIPConfigurationPropertiesFormatARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormatARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["Primary"] = gen.PtrOf(gen.Bool())
	gens["PrivateIPAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIPAddressVersion"] = gen.PtrOf(gen.OneConstOf(NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion_IPv4, NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion_IPv6))
	gens["PrivateIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod_Dynamic, NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod_Static))
}

// AddRelatedPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormatARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfaceIPConfigurationPropertiesFormatARM(gens map[string]gopter.Gen) {
	gens["ApplicationGatewayBackendAddressPools"] = gen.SliceOf(SubResourceARMGenerator())
	gens["ApplicationSecurityGroups"] = gen.SliceOf(SubResourceARMGenerator())
	gens["LoadBalancerBackendAddressPools"] = gen.SliceOf(SubResourceARMGenerator())
	gens["LoadBalancerInboundNatRules"] = gen.SliceOf(SubResourceARMGenerator())
	gens["PublicIPAddress"] = gen.PtrOf(SubResourceARMGenerator())
	gens["Subnet"] = gen.PtrOf(SubResourceARMGenerator())
	gens["VirtualNetworkTaps"] = gen.SliceOf(SubResourceARMGenerator())
}
