// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210701

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

func Test_Images_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Images_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImagesSpecARM, ImagesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImagesSpecARM runs a test to see if a specific instance of Images_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImagesSpecARM(subject Images_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Images_SpecARM
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

// Generator of Images_SpecARM instances for property testing - lazily instantiated by ImagesSpecARMGenerator()
var imagesSpecARMGenerator gopter.Gen

// ImagesSpecARMGenerator returns a generator of Images_SpecARM instances for property testing.
// We first initialize imagesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImagesSpecARMGenerator() gopter.Gen {
	if imagesSpecARMGenerator != nil {
		return imagesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagesSpecARM(generators)
	imagesSpecARMGenerator = gen.Struct(reflect.TypeOf(Images_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagesSpecARM(generators)
	AddRelatedPropertyGeneratorsForImagesSpecARM(generators)
	imagesSpecARMGenerator = gen.Struct(reflect.TypeOf(Images_SpecARM{}), generators)

	return imagesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForImagesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImagesSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImagesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImagesSpecARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationARMGenerator())
	gens["Properties"] = gen.PtrOf(ImagePropertiesARMGenerator())
}

func Test_ExtendedLocationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExtendedLocationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExtendedLocationARM, ExtendedLocationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExtendedLocationARM runs a test to see if a specific instance of ExtendedLocationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExtendedLocationARM(subject ExtendedLocationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExtendedLocationARM
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

// Generator of ExtendedLocationARM instances for property testing - lazily instantiated by
// ExtendedLocationARMGenerator()
var extendedLocationARMGenerator gopter.Gen

// ExtendedLocationARMGenerator returns a generator of ExtendedLocationARM instances for property testing.
func ExtendedLocationARMGenerator() gopter.Gen {
	if extendedLocationARMGenerator != nil {
		return extendedLocationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExtendedLocationARM(generators)
	extendedLocationARMGenerator = gen.Struct(reflect.TypeOf(ExtendedLocationARM{}), generators)

	return extendedLocationARMGenerator
}

// AddIndependentPropertyGeneratorsForExtendedLocationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExtendedLocationARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ExtendedLocationType_EdgeZone))
}

func Test_ImagePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImagePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImagePropertiesARM, ImagePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImagePropertiesARM runs a test to see if a specific instance of ImagePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImagePropertiesARM(subject ImagePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImagePropertiesARM
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

// Generator of ImagePropertiesARM instances for property testing - lazily instantiated by ImagePropertiesARMGenerator()
var imagePropertiesARMGenerator gopter.Gen

// ImagePropertiesARMGenerator returns a generator of ImagePropertiesARM instances for property testing.
// We first initialize imagePropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImagePropertiesARMGenerator() gopter.Gen {
	if imagePropertiesARMGenerator != nil {
		return imagePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagePropertiesARM(generators)
	imagePropertiesARMGenerator = gen.Struct(reflect.TypeOf(ImagePropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImagePropertiesARM(generators)
	AddRelatedPropertyGeneratorsForImagePropertiesARM(generators)
	imagePropertiesARMGenerator = gen.Struct(reflect.TypeOf(ImagePropertiesARM{}), generators)

	return imagePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForImagePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImagePropertiesARM(gens map[string]gopter.Gen) {
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(ImagePropertiesHyperVGeneration_V1, ImagePropertiesHyperVGeneration_V2))
}

// AddRelatedPropertyGeneratorsForImagePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImagePropertiesARM(gens map[string]gopter.Gen) {
	gens["SourceVirtualMachine"] = gen.PtrOf(SubResourceARMGenerator())
	gens["StorageProfile"] = gen.PtrOf(ImageStorageProfileARMGenerator())
}

func Test_ImageStorageProfileARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageStorageProfileARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageStorageProfileARM, ImageStorageProfileARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageStorageProfileARM runs a test to see if a specific instance of ImageStorageProfileARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageStorageProfileARM(subject ImageStorageProfileARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageStorageProfileARM
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

// Generator of ImageStorageProfileARM instances for property testing - lazily instantiated by
// ImageStorageProfileARMGenerator()
var imageStorageProfileARMGenerator gopter.Gen

// ImageStorageProfileARMGenerator returns a generator of ImageStorageProfileARM instances for property testing.
// We first initialize imageStorageProfileARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageStorageProfileARMGenerator() gopter.Gen {
	if imageStorageProfileARMGenerator != nil {
		return imageStorageProfileARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfileARM(generators)
	imageStorageProfileARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfileARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageStorageProfileARM(generators)
	AddRelatedPropertyGeneratorsForImageStorageProfileARM(generators)
	imageStorageProfileARMGenerator = gen.Struct(reflect.TypeOf(ImageStorageProfileARM{}), generators)

	return imageStorageProfileARMGenerator
}

// AddIndependentPropertyGeneratorsForImageStorageProfileARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageStorageProfileARM(gens map[string]gopter.Gen) {
	gens["ZoneResilient"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForImageStorageProfileARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageStorageProfileARM(gens map[string]gopter.Gen) {
	gens["DataDisks"] = gen.SliceOf(ImageDataDiskARMGenerator())
	gens["OsDisk"] = gen.PtrOf(ImageOSDiskARMGenerator())
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

func Test_ImageDataDiskARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageDataDiskARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageDataDiskARM, ImageDataDiskARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageDataDiskARM runs a test to see if a specific instance of ImageDataDiskARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageDataDiskARM(subject ImageDataDiskARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageDataDiskARM
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

// Generator of ImageDataDiskARM instances for property testing - lazily instantiated by ImageDataDiskARMGenerator()
var imageDataDiskARMGenerator gopter.Gen

// ImageDataDiskARMGenerator returns a generator of ImageDataDiskARM instances for property testing.
// We first initialize imageDataDiskARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageDataDiskARMGenerator() gopter.Gen {
	if imageDataDiskARMGenerator != nil {
		return imageDataDiskARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDiskARM(generators)
	imageDataDiskARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDiskARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageDataDiskARM(generators)
	AddRelatedPropertyGeneratorsForImageDataDiskARM(generators)
	imageDataDiskARMGenerator = gen.Struct(reflect.TypeOf(ImageDataDiskARM{}), generators)

	return imageDataDiskARMGenerator
}

// AddIndependentPropertyGeneratorsForImageDataDiskARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageDataDiskARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageDataDiskCaching_None, ImageDataDiskCaching_ReadOnly, ImageDataDiskCaching_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["Lun"] = gen.PtrOf(gen.Int())
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		ImageDataDiskStorageAccountType_PremiumLRS,
		ImageDataDiskStorageAccountType_PremiumZRS,
		ImageDataDiskStorageAccountType_StandardLRS,
		ImageDataDiskStorageAccountType_StandardSSDLRS,
		ImageDataDiskStorageAccountType_StandardSSDZRS,
		ImageDataDiskStorageAccountType_UltraSSDLRS))
}

// AddRelatedPropertyGeneratorsForImageDataDiskARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageDataDiskARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(DiskEncryptionSetParametersARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceARMGenerator())
}

func Test_ImageOSDiskARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImageOSDiskARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImageOSDiskARM, ImageOSDiskARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImageOSDiskARM runs a test to see if a specific instance of ImageOSDiskARM round trips to JSON and back losslessly
func RunJSONSerializationTestForImageOSDiskARM(subject ImageOSDiskARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImageOSDiskARM
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

// Generator of ImageOSDiskARM instances for property testing - lazily instantiated by ImageOSDiskARMGenerator()
var imageOSDiskARMGenerator gopter.Gen

// ImageOSDiskARMGenerator returns a generator of ImageOSDiskARM instances for property testing.
// We first initialize imageOSDiskARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImageOSDiskARMGenerator() gopter.Gen {
	if imageOSDiskARMGenerator != nil {
		return imageOSDiskARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDiskARM(generators)
	imageOSDiskARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDiskARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImageOSDiskARM(generators)
	AddRelatedPropertyGeneratorsForImageOSDiskARM(generators)
	imageOSDiskARMGenerator = gen.Struct(reflect.TypeOf(ImageOSDiskARM{}), generators)

	return imageOSDiskARMGenerator
}

// AddIndependentPropertyGeneratorsForImageOSDiskARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImageOSDiskARM(gens map[string]gopter.Gen) {
	gens["BlobUri"] = gen.PtrOf(gen.AlphaString())
	gens["Caching"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskCaching_None, ImageOSDiskCaching_ReadOnly, ImageOSDiskCaching_ReadWrite))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsState"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskOsState_Generalized, ImageOSDiskOsState_Specialized))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(ImageOSDiskOsType_Linux, ImageOSDiskOsType_Windows))
	gens["StorageAccountType"] = gen.PtrOf(gen.OneConstOf(
		ImageOSDiskStorageAccountType_PremiumLRS,
		ImageOSDiskStorageAccountType_PremiumZRS,
		ImageOSDiskStorageAccountType_StandardLRS,
		ImageOSDiskStorageAccountType_StandardSSDLRS,
		ImageOSDiskStorageAccountType_StandardSSDZRS,
		ImageOSDiskStorageAccountType_UltraSSDLRS))
}

// AddRelatedPropertyGeneratorsForImageOSDiskARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImageOSDiskARM(gens map[string]gopter.Gen) {
	gens["DiskEncryptionSet"] = gen.PtrOf(DiskEncryptionSetParametersARMGenerator())
	gens["ManagedDisk"] = gen.PtrOf(SubResourceARMGenerator())
	gens["Snapshot"] = gen.PtrOf(SubResourceARMGenerator())
}

func Test_DiskEncryptionSetParametersARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskEncryptionSetParametersARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskEncryptionSetParametersARM, DiskEncryptionSetParametersARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskEncryptionSetParametersARM runs a test to see if a specific instance of DiskEncryptionSetParametersARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskEncryptionSetParametersARM(subject DiskEncryptionSetParametersARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskEncryptionSetParametersARM
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

// Generator of DiskEncryptionSetParametersARM instances for property testing - lazily instantiated by
// DiskEncryptionSetParametersARMGenerator()
var diskEncryptionSetParametersARMGenerator gopter.Gen

// DiskEncryptionSetParametersARMGenerator returns a generator of DiskEncryptionSetParametersARM instances for property testing.
func DiskEncryptionSetParametersARMGenerator() gopter.Gen {
	if diskEncryptionSetParametersARMGenerator != nil {
		return diskEncryptionSetParametersARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSetParametersARM(generators)
	diskEncryptionSetParametersARMGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSetParametersARM{}), generators)

	return diskEncryptionSetParametersARMGenerator
}

// AddIndependentPropertyGeneratorsForDiskEncryptionSetParametersARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskEncryptionSetParametersARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
