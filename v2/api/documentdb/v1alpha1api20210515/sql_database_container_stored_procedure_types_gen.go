// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	alpha20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of SqlDatabaseContainerStoredProcedure. Use v1beta20210515.SqlDatabaseContainerStoredProcedure instead
type SqlDatabaseContainerStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec `json:"spec,omitempty"`
	Status            SqlStoredProcedureGetResults_STATUS                         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerStoredProcedure{}

// GetConditions returns the conditions of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetConditions() conditions.Conditions {
	return procedure.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (procedure *SqlDatabaseContainerStoredProcedure) SetConditions(conditions conditions.Conditions) {
	procedure.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerStoredProcedure{}

// ConvertFrom populates our SqlDatabaseContainerStoredProcedure from the provided hub SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210515s.SqlDatabaseContainerStoredProcedure

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = procedure.AssignPropertiesFromSqlDatabaseContainerStoredProcedure(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to procedure")
	}

	return nil
}

// ConvertTo populates the provided hub SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210515s.SqlDatabaseContainerStoredProcedure
	err := procedure.AssignPropertiesToSqlDatabaseContainerStoredProcedure(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from procedure")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerstoredprocedure,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &SqlDatabaseContainerStoredProcedure{}

// Default applies defaults to the SqlDatabaseContainerStoredProcedure resource
func (procedure *SqlDatabaseContainerStoredProcedure) Default() {
	procedure.defaultImpl()
	var temp interface{} = procedure
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (procedure *SqlDatabaseContainerStoredProcedure) defaultAzureName() {
	if procedure.Spec.AzureName == "" {
		procedure.Spec.AzureName = procedure.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseContainerStoredProcedure resource
func (procedure *SqlDatabaseContainerStoredProcedure) defaultImpl() { procedure.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabaseContainerStoredProcedure{}

// AzureName returns the Azure name of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) AzureName() string {
	return procedure.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedure SqlDatabaseContainerStoredProcedure) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetSpec() genruntime.ConvertibleSpec {
	return &procedure.Spec
}

// GetStatus returns the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetStatus() genruntime.ConvertibleStatus {
	return &procedure.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedure *SqlDatabaseContainerStoredProcedure) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

// NewEmptyStatus returns a new empty (blank) status
func (procedure *SqlDatabaseContainerStoredProcedure) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlStoredProcedureGetResults_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (procedure *SqlDatabaseContainerStoredProcedure) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(procedure.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  procedure.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlStoredProcedureGetResults_STATUS); ok {
		procedure.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlStoredProcedureGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	procedure.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerstoredprocedure,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &SqlDatabaseContainerStoredProcedure{}

// ValidateCreate validates the creation of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateCreate() error {
	validations := procedure.createValidations()
	var temp interface{} = procedure
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateDelete() error {
	validations := procedure.deleteValidations()
	var temp interface{} = procedure
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateUpdate(old runtime.Object) error {
	validations := procedure.updateValidations()
	var temp interface{} = procedure
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) createValidations() []func() error {
	return []func() error{procedure.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return procedure.validateResourceReferences()
		},
		procedure.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (procedure *SqlDatabaseContainerStoredProcedure) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&procedure.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (procedure *SqlDatabaseContainerStoredProcedure) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*SqlDatabaseContainerStoredProcedure)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, procedure)
}

// AssignPropertiesFromSqlDatabaseContainerStoredProcedure populates our SqlDatabaseContainerStoredProcedure from the provided source SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesFromSqlDatabaseContainerStoredProcedure(source *alpha20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	procedure.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec() to populate field Spec")
	}
	procedure.Spec = spec

	// Status
	var status SqlStoredProcedureGetResults_STATUS
	err = status.AssignPropertiesFromSqlStoredProcedureGetResultsSTATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureGetResultsSTATUS() to populate field Status")
	}
	procedure.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerStoredProcedure populates the provided destination SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesToSqlDatabaseContainerStoredProcedure(destination *alpha20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	destination.ObjectMeta = *procedure.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
	err := procedure.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210515s.SqlStoredProcedureGetResults_STATUS
	err = procedure.Status.AssignPropertiesToSqlStoredProcedureGetResultsSTATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureGetResultsSTATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (procedure *SqlDatabaseContainerStoredProcedure) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: procedure.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerStoredProcedure",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of SqlDatabaseContainerStoredProcedure. Use v1beta20210515.SqlDatabaseContainerStoredProcedure instead
type SqlDatabaseContainerStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerStoredProcedure `json:"items"`
}

type DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string               `json:"azureName,omitempty"`
	Location  *string              `json:"location,omitempty"`
	Options   *CreateUpdateOptions `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`

	// +kubebuilder:validation:Required
	Resource *SqlStoredProcedureResource `json:"resource,omitempty"`
	Tags     map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if procedures == nil {
		return nil, nil
	}
	result := &DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM{}

	// Set property ‘Location’:
	if procedures.Location != nil {
		location := *procedures.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if procedures.Options != nil || procedures.Resource != nil {
		result.Properties = &SqlStoredProcedureCreateUpdatePropertiesARM{}
	}
	if procedures.Options != nil {
		optionsARM, err := (*procedures.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := *optionsARM.(*CreateUpdateOptionsARM)
		result.Properties.Options = &options
	}
	if procedures.Resource != nil {
		resourceARM, err := (*procedures.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resourceARM.(*SqlStoredProcedureResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if procedures.Tags != nil {
		result.Tags = make(map[string]string, len(procedures.Tags))
		for key, value := range procedures.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	procedures.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		procedures.Location = &location
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 CreateUpdateOptions
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			procedures.Options = &options
		}
	}

	// Set property ‘Owner’:
	procedures.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlStoredProcedureResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			procedures.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		procedures.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			procedures.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec from the provided source
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec)
	if ok {
		// Populate our instance from source
		return procedures.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = procedures.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec)
	if ok {
		// Populate destination from our instance
		return procedures.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}
	err := procedures.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec populates our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec from the provided source DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(source *alpha20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) error {

	// AzureName
	procedures.AzureName = source.AzureName

	// Location
	procedures.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		procedures.Options = &option
	} else {
		procedures.Options = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		procedures.Owner = &owner
	} else {
		procedures.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureResource
		err := resource.AssignPropertiesFromSqlStoredProcedureResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureResource() to populate field Resource")
		}
		procedures.Resource = &resource
	} else {
		procedures.Resource = nil
	}

	// Tags
	procedures.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec populates the provided destination DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec from our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(destination *alpha20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = procedures.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(procedures.Location)

	// Options
	if procedures.Options != nil {
		var option alpha20210515s.CreateUpdateOptions
		err := procedures.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = procedures.OriginalVersion()

	// Owner
	if procedures.Owner != nil {
		owner := procedures.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if procedures.Resource != nil {
		var resource alpha20210515s.SqlStoredProcedureResource
		err := procedures.Resource.AssignPropertiesToSqlStoredProcedureResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(procedures.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) SetAzureName(azureName string) {
	procedures.AzureName = azureName
}

// Deprecated version of SqlStoredProcedureGetResults_STATUS. Use v1beta20210515.SqlStoredProcedureGetResults_STATUS instead
type SqlStoredProcedureGetResults_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition                           `json:"conditions,omitempty"`
	Id         *string                                          `json:"id,omitempty"`
	Location   *string                                          `json:"location,omitempty"`
	Name       *string                                          `json:"name,omitempty"`
	Resource   *SqlStoredProcedureGetProperties_STATUS_Resource `json:"resource,omitempty"`
	Tags       map[string]string                                `json:"tags,omitempty"`
	Type       *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlStoredProcedureGetResults_STATUS{}

// ConvertStatusFrom populates our SqlStoredProcedureGetResults_STATUS from the provided source
func (results *SqlStoredProcedureGetResults_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210515s.SqlStoredProcedureGetResults_STATUS)
	if ok {
		// Populate our instance from source
		return results.AssignPropertiesFromSqlStoredProcedureGetResultsSTATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.SqlStoredProcedureGetResults_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignPropertiesFromSqlStoredProcedureGetResultsSTATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlStoredProcedureGetResults_STATUS
func (results *SqlStoredProcedureGetResults_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210515s.SqlStoredProcedureGetResults_STATUS)
	if ok {
		// Populate destination from our instance
		return results.AssignPropertiesToSqlStoredProcedureGetResultsSTATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.SqlStoredProcedureGetResults_STATUS{}
	err := results.AssignPropertiesToSqlStoredProcedureGetResultsSTATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &SqlStoredProcedureGetResults_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (results *SqlStoredProcedureGetResults_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureGetResults_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (results *SqlStoredProcedureGetResults_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureGetResults_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureGetResults_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		results.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		results.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		results.Name = &name
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlStoredProcedureGetProperties_STATUS_Resource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			results.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		results.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			results.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		results.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureGetResultsSTATUS populates our SqlStoredProcedureGetResults_STATUS from the provided source SqlStoredProcedureGetResults_STATUS
func (results *SqlStoredProcedureGetResults_STATUS) AssignPropertiesFromSqlStoredProcedureGetResultsSTATUS(source *alpha20210515s.SqlStoredProcedureGetResults_STATUS) error {

	// Conditions
	results.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	results.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	results.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	results.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureGetProperties_STATUS_Resource
		err := resource.AssignPropertiesFromSqlStoredProcedureGetPropertiesSTATUSResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureGetPropertiesSTATUSResource() to populate field Resource")
		}
		results.Resource = &resource
	} else {
		results.Resource = nil
	}

	// Tags
	results.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	results.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureGetResultsSTATUS populates the provided destination SqlStoredProcedureGetResults_STATUS from our SqlStoredProcedureGetResults_STATUS
func (results *SqlStoredProcedureGetResults_STATUS) AssignPropertiesToSqlStoredProcedureGetResultsSTATUS(destination *alpha20210515s.SqlStoredProcedureGetResults_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(results.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(results.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(results.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(results.Name)

	// Resource
	if results.Resource != nil {
		var resource alpha20210515s.SqlStoredProcedureGetProperties_STATUS_Resource
		err := results.Resource.AssignPropertiesToSqlStoredProcedureGetPropertiesSTATUSResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureGetPropertiesSTATUSResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(results.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(results.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of SqlStoredProcedureGetProperties_STATUS_Resource. Use v1beta20210515.SqlStoredProcedureGetProperties_STATUS_Resource instead
type SqlStoredProcedureGetProperties_STATUS_Resource struct {
	Body *string  `json:"body,omitempty"`
	Etag *string  `json:"_etag,omitempty"`
	Id   *string  `json:"id,omitempty"`
	Rid  *string  `json:"_rid,omitempty"`
	Ts   *float64 `json:"_ts,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlStoredProcedureGetProperties_STATUS_Resource{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureGetProperties_STATUS_Resource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureGetProperties_STATUS_ResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureGetProperties_STATUS_Resource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureGetProperties_STATUS_ResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureGetProperties_STATUS_ResourceARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		resource.Etag = &etag
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// Set property ‘Rid’:
	if typedInput.Rid != nil {
		rid := *typedInput.Rid
		resource.Rid = &rid
	}

	// Set property ‘Ts’:
	if typedInput.Ts != nil {
		ts := *typedInput.Ts
		resource.Ts = &ts
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureGetPropertiesSTATUSResource populates our SqlStoredProcedureGetProperties_STATUS_Resource from the provided source SqlStoredProcedureGetProperties_STATUS_Resource
func (resource *SqlStoredProcedureGetProperties_STATUS_Resource) AssignPropertiesFromSqlStoredProcedureGetPropertiesSTATUSResource(source *alpha20210515s.SqlStoredProcedureGetProperties_STATUS_Resource) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Rid
	resource.Rid = genruntime.ClonePointerToString(source.Rid)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		resource.Ts = &t
	} else {
		resource.Ts = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureGetPropertiesSTATUSResource populates the provided destination SqlStoredProcedureGetProperties_STATUS_Resource from our SqlStoredProcedureGetProperties_STATUS_Resource
func (resource *SqlStoredProcedureGetProperties_STATUS_Resource) AssignPropertiesToSqlStoredProcedureGetPropertiesSTATUSResource(destination *alpha20210515s.SqlStoredProcedureGetProperties_STATUS_Resource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Rid
	destination.Rid = genruntime.ClonePointerToString(resource.Rid)

	// Ts
	if resource.Ts != nil {
		t := *resource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of SqlStoredProcedureResource. Use v1beta20210515.SqlStoredProcedureResource instead
type SqlStoredProcedureResource struct {
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	Id *string `json:"id,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlStoredProcedureResource{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlStoredProcedureResource) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	result := &SqlStoredProcedureResourceARM{}

	// Set property ‘Body’:
	if resource.Body != nil {
		body := *resource.Body
		result.Body = &body
	}

	// Set property ‘Id’:
	if resource.Id != nil {
		id := *resource.Id
		result.Id = &id
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureResource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureResource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureResourceARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureResource populates our SqlStoredProcedureResource from the provided source SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignPropertiesFromSqlStoredProcedureResource(source *alpha20210515s.SqlStoredProcedureResource) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureResource populates the provided destination SqlStoredProcedureResource from our SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignPropertiesToSqlStoredProcedureResource(destination *alpha20210515s.SqlStoredProcedureResource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerStoredProcedure{}, &SqlDatabaseContainerStoredProcedureList{})
}
