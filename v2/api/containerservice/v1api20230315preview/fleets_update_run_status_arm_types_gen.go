// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230315preview

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

type Fleets_UpdateRun_STATUS_ARM struct {
	// ETag: If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.
	// Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in
	// the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header
	// fields.
	ETag *string `json:"eTag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The resource-specific properties for this resource.
	Properties *UpdateRunProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of the UpdateRun.
type UpdateRunProperties_STATUS_ARM struct {
	// ManagedClusterUpdate: The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be
	// modified until the run is started.
	ManagedClusterUpdate *ManagedClusterUpdate_STATUS_ARM `json:"managedClusterUpdate,omitempty"`

	// ProvisioningState: The provisioning state of the UpdateRun resource.
	ProvisioningState *UpdateRunProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Status: The status of the UpdateRun.
	Status *UpdateRunStatus_STATUS_ARM `json:"status,omitempty"`

	// Strategy: The strategy defines the order in which the clusters will be updated.
	// If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single
	// UpdateGroup targeting all members.
	// The strategy of the UpdateRun can be modified until the run is started.
	Strategy *UpdateRunStrategy_STATUS_ARM `json:"strategy,omitempty"`
}

// The update to be applied to the ManagedClusters.
type ManagedClusterUpdate_STATUS_ARM struct {
	// Upgrade: The upgrade to apply to the ManagedClusters.
	Upgrade *ManagedClusterUpgradeSpec_STATUS_ARM `json:"upgrade,omitempty"`
}

// The status of a UpdateRun.
type UpdateRunStatus_STATUS_ARM struct {
	// Stages: The stages composing an update run. Stages are run sequentially withing an UpdateRun.
	Stages []UpdateStageStatus_STATUS_ARM `json:"stages,omitempty"`

	// Status: The status of the UpdateRun.
	Status *UpdateStatus_STATUS_ARM `json:"status,omitempty"`
}

// The UpdateRunStrategy configures the sequence of Stages and Groups in which the clusters will be updated.
type UpdateRunStrategy_STATUS_ARM struct {
	// Stages: The list of stages that compose this update run.
	Stages []UpdateStage_STATUS_ARM `json:"stages,omitempty"`
}

// The upgrade to apply to a ManagedCluster.
type ManagedClusterUpgradeSpec_STATUS_ARM struct {
	// KubernetesVersion: The Kubernetes version to upgrade the member clusters to.
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`

	// Type: The upgrade type.
	// Full requires the KubernetesVersion property to be set.
	// NodeImageOnly requires the KubernetesVersion property not to be set.
	Type *ManagedClusterUpgradeType_STATUS `json:"type,omitempty"`
}

// Contains the groups to be updated by an UpdateRun.
// Update order:
// - Sequential between stages: Stages run sequentially.
// The previous stage must complete before the next one starts.
// - Parallel within a stage: Groups within a stage run in
// parallel.
// - Sequential within a group: Clusters within a group are updated sequentially.
type UpdateStage_STATUS_ARM struct {
	// AfterStageWaitInSeconds: The time in seconds to wait at the end of this stage before starting the next one. Defaults to
	// 0 seconds if unspecified.
	AfterStageWaitInSeconds *int `json:"afterStageWaitInSeconds,omitempty"`

	// Groups: A list of group names that compose the stage.
	// The groups will be updated in parallel. Each group name can only appear once in the UpdateRun.
	Groups []UpdateGroup_STATUS_ARM `json:"groups,omitempty"`

	// Name: The name of the stage. Must be unique within the UpdateRun.
	Name *string `json:"name,omitempty"`
}

// The status of a UpdateStage.
type UpdateStageStatus_STATUS_ARM struct {
	// AfterStageWaitStatus: The status of the wait period configured on the UpdateStage.
	AfterStageWaitStatus *WaitStatus_STATUS_ARM `json:"afterStageWaitStatus,omitempty"`

	// Groups: The list of groups to be updated as part of this UpdateStage.
	Groups []UpdateGroupStatus_STATUS_ARM `json:"groups,omitempty"`

	// Name: The name of the UpdateStage.
	Name *string `json:"name,omitempty"`

	// Status: The status of the UpdateStage.
	Status *UpdateStatus_STATUS_ARM `json:"status,omitempty"`
}

// The status for an operation or group of operations.
type UpdateStatus_STATUS_ARM struct {
	// CompletedTime: The time the operation or group was completed.
	CompletedTime *string `json:"completedTime,omitempty"`

	// Error: The error details when a failure is encountered.
	Error *ErrorDetail_STATUS_ARM `json:"error,omitempty"`

	// StartTime: The time the operation or group was started.
	StartTime *string `json:"startTime,omitempty"`

	// State: The State of the operation or group.
	State *UpdateState_STATUS `json:"state,omitempty"`
}

// The error detail.
type ErrorDetail_STATUS_ARM struct {
	// AdditionalInfo: The error additional info.
	AdditionalInfo []ErrorAdditionalInfo_STATUS_ARM `json:"additionalInfo,omitempty"`

	// Code: The error code.
	Code *string `json:"code,omitempty"`

	// Details: The error details.
	Details []ErrorDetail_STATUS_Unrolled_ARM `json:"details,omitempty"`

	// Message: The error message.
	Message *string `json:"message,omitempty"`

	// Target: The error target.
	Target *string `json:"target,omitempty"`
}

// A group to be updated.
type UpdateGroup_STATUS_ARM struct {
	// Name: The name of the Fleet member group to update.
	// It should match the name of an existing FleetMember group.
	// A group can only appear once across all UpdateStages in the UpdateRun.
	Name *string `json:"name,omitempty"`
}

// The status of a UpdateGroup.
type UpdateGroupStatus_STATUS_ARM struct {
	// Members: The list of member this UpdateGroup updates.
	Members []MemberUpdateStatus_STATUS_ARM `json:"members,omitempty"`

	// Name: The name of the UpdateGroup.
	Name *string `json:"name,omitempty"`

	// Status: The status of the UpdateGroup.
	Status *UpdateStatus_STATUS_ARM `json:"status,omitempty"`
}

// The status of the wait duration.
type WaitStatus_STATUS_ARM struct {
	// Status: The status of the wait duration.
	Status *UpdateStatus_STATUS_ARM `json:"status,omitempty"`

	// WaitDurationInSeconds: The wait duration configured in seconds.
	WaitDurationInSeconds *int `json:"waitDurationInSeconds,omitempty"`
}

// The resource management error additional info.
type ErrorAdditionalInfo_STATUS_ARM struct {
	// Info: The additional info.
	Info map[string]v1.JSON `json:"info,omitempty"`

	// Type: The additional info type.
	Type *string `json:"type,omitempty"`
}

type ErrorDetail_STATUS_Unrolled_ARM struct {
	// AdditionalInfo: The error additional info.
	AdditionalInfo []ErrorAdditionalInfo_STATUS_ARM `json:"additionalInfo,omitempty"`

	// Code: The error code.
	Code *string `json:"code,omitempty"`

	// Message: The error message.
	Message *string `json:"message,omitempty"`

	// Target: The error target.
	Target *string `json:"target,omitempty"`
}

// The status of a member update operation.
type MemberUpdateStatus_STATUS_ARM struct {
	// ClusterResourceId: The Azure resource id of the target Kubernetes cluster.
	ClusterResourceId *string `json:"clusterResourceId,omitempty"`

	// Name: The name of the FleetMember.
	Name *string `json:"name,omitempty"`

	// OperationId: The operation resource id of the latest attempt to perform the operation.
	OperationId *string `json:"operationId,omitempty"`

	// Status: The status of the MemberUpdate operation.
	Status *UpdateStatus_STATUS_ARM `json:"status,omitempty"`
}
