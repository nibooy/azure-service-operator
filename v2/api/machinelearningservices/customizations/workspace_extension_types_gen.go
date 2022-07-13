// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210701 "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1beta20210701"
	v20210701s "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1beta20210701storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type WorkspaceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *WorkspaceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210701.Workspace{},
		&v20210701s.Workspace{}}
}