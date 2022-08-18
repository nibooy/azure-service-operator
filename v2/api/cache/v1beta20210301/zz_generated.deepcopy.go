//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20210301

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterPropertiesARM) DeepCopyInto(out *ClusterPropertiesARM) {
	*out = *in
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterPropertiesMinimumTlsVersion)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterPropertiesARM.
func (in *ClusterPropertiesARM) DeepCopy() *ClusterPropertiesARM {
	if in == nil {
		return nil
	}
	out := new(ClusterPropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProperties_STATUSARM) DeepCopyInto(out *ClusterProperties_STATUSARM) {
	*out = *in
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterPropertiesSTATUSMinimumTlsVersion)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProperties_STATUSARM.
func (in *ClusterProperties_STATUSARM) DeepCopy() *ClusterProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(ClusterProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster_STATUS) DeepCopyInto(out *Cluster_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterPropertiesSTATUSMinimumTlsVersion)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointConnections != nil {
		in, out := &in.PrivateEndpointConnections, &out.PrivateEndpointConnections
		*out = make([]PrivateEndpointConnection_STATUS_SubResourceEmbedded, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.RedisVersion != nil {
		in, out := &in.RedisVersion, &out.RedisVersion
		*out = new(string)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster_STATUS.
func (in *Cluster_STATUS) DeepCopy() *Cluster_STATUS {
	if in == nil {
		return nil
	}
	out := new(Cluster_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster_STATUSARM) DeepCopyInto(out *Cluster_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ClusterProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster_STATUSARM.
func (in *Cluster_STATUSARM) DeepCopy() *Cluster_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(Cluster_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabasePropertiesARM) DeepCopyInto(out *DatabasePropertiesARM) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabasePropertiesClientProtocol)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabasePropertiesClusteringPolicy)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabasePropertiesEvictionPolicy)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]ModuleARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabasePropertiesARM.
func (in *DatabasePropertiesARM) DeepCopy() *DatabasePropertiesARM {
	if in == nil {
		return nil
	}
	out := new(DatabasePropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseProperties_STATUSARM) DeepCopyInto(out *DatabaseProperties_STATUSARM) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabasePropertiesSTATUSClientProtocol)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabasePropertiesSTATUSClusteringPolicy)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabasePropertiesSTATUSEvictionPolicy)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseProperties_STATUSARM.
func (in *DatabaseProperties_STATUSARM) DeepCopy() *DatabaseProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(DatabaseProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database_STATUS) DeepCopyInto(out *Database_STATUS) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabasePropertiesSTATUSClientProtocol)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabasePropertiesSTATUSClusteringPolicy)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabasePropertiesSTATUSEvictionPolicy)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.ResourceState != nil {
		in, out := &in.ResourceState, &out.ResourceState
		*out = new(ResourceState_STATUS)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database_STATUS.
func (in *Database_STATUS) DeepCopy() *Database_STATUS {
	if in == nil {
		return nil
	}
	out := new(Database_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database_STATUSARM) DeepCopyInto(out *Database_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(DatabaseProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database_STATUSARM.
func (in *Database_STATUSARM) DeepCopy() *Database_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(Database_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleARM) DeepCopyInto(out *ModuleARM) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleARM.
func (in *ModuleARM) DeepCopy() *ModuleARM {
	if in == nil {
		return nil
	}
	out := new(ModuleARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module_STATUS) DeepCopyInto(out *Module_STATUS) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module_STATUS.
func (in *Module_STATUS) DeepCopy() *Module_STATUS {
	if in == nil {
		return nil
	}
	out := new(Module_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module_STATUSARM) DeepCopyInto(out *Module_STATUSARM) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module_STATUSARM.
func (in *Module_STATUSARM) DeepCopy() *Module_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(Module_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence) DeepCopyInto(out *Persistence) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(PersistenceAofFrequency)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(PersistenceRdbFrequency)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence.
func (in *Persistence) DeepCopy() *Persistence {
	if in == nil {
		return nil
	}
	out := new(Persistence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceARM) DeepCopyInto(out *PersistenceARM) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(PersistenceAofFrequency)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(PersistenceRdbFrequency)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceARM.
func (in *PersistenceARM) DeepCopy() *PersistenceARM {
	if in == nil {
		return nil
	}
	out := new(PersistenceARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence_STATUS) DeepCopyInto(out *Persistence_STATUS) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(PersistenceSTATUSAofFrequency)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(PersistenceSTATUSRdbFrequency)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence_STATUS.
func (in *Persistence_STATUS) DeepCopy() *Persistence_STATUS {
	if in == nil {
		return nil
	}
	out := new(Persistence_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Persistence_STATUSARM) DeepCopyInto(out *Persistence_STATUSARM) {
	*out = *in
	if in.AofEnabled != nil {
		in, out := &in.AofEnabled, &out.AofEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AofFrequency != nil {
		in, out := &in.AofFrequency, &out.AofFrequency
		*out = new(PersistenceSTATUSAofFrequency)
		**out = **in
	}
	if in.RdbEnabled != nil {
		in, out := &in.RdbEnabled, &out.RdbEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RdbFrequency != nil {
		in, out := &in.RdbFrequency, &out.RdbFrequency
		*out = new(PersistenceSTATUSRdbFrequency)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Persistence_STATUSARM.
func (in *Persistence_STATUSARM) DeepCopy() *Persistence_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(Persistence_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS_SubResourceEmbedded) DeepCopyInto(out *PrivateEndpointConnection_STATUS_SubResourceEmbedded) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS_SubResourceEmbedded.
func (in *PrivateEndpointConnection_STATUS_SubResourceEmbedded) DeepCopy() *PrivateEndpointConnection_STATUS_SubResourceEmbedded {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS_SubResourceEmbedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM) DeepCopyInto(out *PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM.
func (in *PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM) DeepCopy() *PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointConnection_STATUS_SubResourceEmbeddedARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise) DeepCopyInto(out *RedisEnterprise) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise.
func (in *RedisEnterprise) DeepCopy() *RedisEnterprise {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterprise) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabase) DeepCopyInto(out *RedisEnterpriseDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabase.
func (in *RedisEnterpriseDatabase) DeepCopy() *RedisEnterpriseDatabase {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabaseList) DeepCopyInto(out *RedisEnterpriseDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterpriseDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabaseList.
func (in *RedisEnterpriseDatabaseList) DeepCopy() *RedisEnterpriseDatabaseList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabases_Spec) DeepCopyInto(out *RedisEnterpriseDatabases_Spec) {
	*out = *in
	if in.ClientProtocol != nil {
		in, out := &in.ClientProtocol, &out.ClientProtocol
		*out = new(DatabasePropertiesClientProtocol)
		**out = **in
	}
	if in.ClusteringPolicy != nil {
		in, out := &in.ClusteringPolicy, &out.ClusteringPolicy
		*out = new(DatabasePropertiesClusteringPolicy)
		**out = **in
	}
	if in.EvictionPolicy != nil {
		in, out := &in.EvictionPolicy, &out.EvictionPolicy
		*out = new(DatabasePropertiesEvictionPolicy)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(Persistence)
		(*in).DeepCopyInto(*out)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabases_Spec.
func (in *RedisEnterpriseDatabases_Spec) DeepCopy() *RedisEnterpriseDatabases_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabases_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseDatabases_SpecARM) DeepCopyInto(out *RedisEnterpriseDatabases_SpecARM) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(DatabasePropertiesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseDatabases_SpecARM.
func (in *RedisEnterpriseDatabases_SpecARM) DeepCopy() *RedisEnterpriseDatabases_SpecARM {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseDatabases_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterpriseList) DeepCopyInto(out *RedisEnterpriseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RedisEnterprise, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterpriseList.
func (in *RedisEnterpriseList) DeepCopy() *RedisEnterpriseList {
	if in == nil {
		return nil
	}
	out := new(RedisEnterpriseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RedisEnterpriseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_Spec) DeepCopyInto(out *RedisEnterprise_Spec) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MinimumTlsVersion != nil {
		in, out := &in.MinimumTlsVersion, &out.MinimumTlsVersion
		*out = new(ClusterPropertiesMinimumTlsVersion)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(Sku)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_Spec.
func (in *RedisEnterprise_Spec) DeepCopy() *RedisEnterprise_Spec {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisEnterprise_SpecARM) DeepCopyInto(out *RedisEnterprise_SpecARM) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(ClusterPropertiesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Sku != nil {
		in, out := &in.Sku, &out.Sku
		*out = new(SkuARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisEnterprise_SpecARM.
func (in *RedisEnterprise_SpecARM) DeepCopy() *RedisEnterprise_SpecARM {
	if in == nil {
		return nil
	}
	out := new(RedisEnterprise_SpecARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku) DeepCopyInto(out *Sku) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(SkuName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku.
func (in *Sku) DeepCopy() *Sku {
	if in == nil {
		return nil
	}
	out := new(Sku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkuARM) DeepCopyInto(out *SkuARM) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(SkuName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkuARM.
func (in *SkuARM) DeepCopy() *SkuARM {
	if in == nil {
		return nil
	}
	out := new(SkuARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUS) DeepCopyInto(out *Sku_STATUS) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(SkuSTATUSName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUS.
func (in *Sku_STATUS) DeepCopy() *Sku_STATUS {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sku_STATUSARM) DeepCopyInto(out *Sku_STATUSARM) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(SkuSTATUSName)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sku_STATUSARM.
func (in *Sku_STATUSARM) DeepCopy() *Sku_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(Sku_STATUSARM)
	in.DeepCopyInto(out)
	return out
}
