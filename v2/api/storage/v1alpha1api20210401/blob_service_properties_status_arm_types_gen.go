// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

// Deprecated version of BlobServiceProperties_STATUS. Use v1beta20210401.BlobServiceProperties_STATUS instead
type BlobServiceProperties_STATUSARM struct {
	Id         *string                                     `json:"id,omitempty"`
	Name       *string                                     `json:"name,omitempty"`
	Properties *BlobServiceProperties_STATUS_PropertiesARM `json:"properties,omitempty"`
	Sku        *Sku_STATUSARM                              `json:"sku,omitempty"`
	Type       *string                                     `json:"type,omitempty"`
}

// Deprecated version of BlobServiceProperties_STATUS_Properties. Use v1beta20210401.BlobServiceProperties_STATUS_Properties instead
type BlobServiceProperties_STATUS_PropertiesARM struct {
	AutomaticSnapshotPolicyEnabled *bool                                   `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed_STATUSARM                   `json:"changeFeed,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_STATUSARM        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules_STATUSARM                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                                 `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy_STATUSARM        `json:"deleteRetentionPolicy,omitempty"`
	IsVersioningEnabled            *bool                                   `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy_STATUSARM `json:"lastAccessTimeTrackingPolicy,omitempty"`
	RestorePolicy                  *RestorePolicyProperties_STATUSARM      `json:"restorePolicy,omitempty"`
}

// Deprecated version of Sku_STATUS. Use v1beta20210401.Sku_STATUS instead
type Sku_STATUSARM struct {
	Name *SkuName_STATUS `json:"name,omitempty"`
	Tier *Tier_STATUS    `json:"tier,omitempty"`
}

// Deprecated version of ChangeFeed_STATUS. Use v1beta20210401.ChangeFeed_STATUS instead
type ChangeFeed_STATUSARM struct {
	Enabled         *bool `json:"enabled,omitempty"`
	RetentionInDays *int  `json:"retentionInDays,omitempty"`
}

// Deprecated version of CorsRules_STATUS. Use v1beta20210401.CorsRules_STATUS instead
type CorsRules_STATUSARM struct {
	CorsRules []CorsRule_STATUSARM `json:"corsRules,omitempty"`
}

// Deprecated version of DeleteRetentionPolicy_STATUS. Use v1beta20210401.DeleteRetentionPolicy_STATUS instead
type DeleteRetentionPolicy_STATUSARM struct {
	Days    *int  `json:"days,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// Deprecated version of LastAccessTimeTrackingPolicy_STATUS. Use v1beta20210401.LastAccessTimeTrackingPolicy_STATUS instead
type LastAccessTimeTrackingPolicy_STATUSARM struct {
	BlobType                  []string                                `json:"blobType,omitempty"`
	Enable                    *bool                                   `json:"enable,omitempty"`
	Name                      *LastAccessTimeTrackingPolicySTATUSName `json:"name,omitempty"`
	TrackingGranularityInDays *int                                    `json:"trackingGranularityInDays,omitempty"`
}

// Deprecated version of RestorePolicyProperties_STATUS. Use v1beta20210401.RestorePolicyProperties_STATUS instead
type RestorePolicyProperties_STATUSARM struct {
	Days            *int    `json:"days,omitempty"`
	Enabled         *bool   `json:"enabled,omitempty"`
	LastEnabledTime *string `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string `json:"minRestoreTime,omitempty"`
}

// Deprecated version of SkuName_STATUS. Use v1beta20210401.SkuName_STATUS instead
type SkuName_STATUS string

const (
	SkuName_STATUS_PremiumLRS     = SkuName_STATUS("Premium_LRS")
	SkuName_STATUS_PremiumZRS     = SkuName_STATUS("Premium_ZRS")
	SkuName_STATUS_StandardGRS    = SkuName_STATUS("Standard_GRS")
	SkuName_STATUS_StandardGZRS   = SkuName_STATUS("Standard_GZRS")
	SkuName_STATUS_StandardLRS    = SkuName_STATUS("Standard_LRS")
	SkuName_STATUS_StandardRAGRS  = SkuName_STATUS("Standard_RAGRS")
	SkuName_STATUS_StandardRAGZRS = SkuName_STATUS("Standard_RAGZRS")
	SkuName_STATUS_StandardZRS    = SkuName_STATUS("Standard_ZRS")
)

// Deprecated version of Tier_STATUS. Use v1beta20210401.Tier_STATUS instead
type Tier_STATUS string

const (
	Tier_STATUS_Premium  = Tier_STATUS("Premium")
	Tier_STATUS_Standard = Tier_STATUS("Standard")
)

// Deprecated version of CorsRule_STATUS. Use v1beta20210401.CorsRule_STATUS instead
type CorsRule_STATUSARM struct {
	AllowedHeaders  []string                       `json:"allowedHeaders,omitempty"`
	AllowedMethods  []CorsRuleSTATUSAllowedMethods `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string                       `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string                       `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                           `json:"maxAgeInSeconds,omitempty"`
}

// Deprecated version of LastAccessTimeTrackingPolicySTATUSName. Use v1beta20210401.LastAccessTimeTrackingPolicySTATUSName
// instead
type LastAccessTimeTrackingPolicySTATUSName string

const LastAccessTimeTrackingPolicySTATUSName_AccessTimeTracking = LastAccessTimeTrackingPolicySTATUSName("AccessTimeTracking")

// Deprecated version of CorsRuleSTATUSAllowedMethods. Use v1beta20210401.CorsRuleSTATUSAllowedMethods instead
type CorsRuleSTATUSAllowedMethods string

const (
	CorsRuleSTATUSAllowedMethods_DELETE  = CorsRuleSTATUSAllowedMethods("DELETE")
	CorsRuleSTATUSAllowedMethods_GET     = CorsRuleSTATUSAllowedMethods("GET")
	CorsRuleSTATUSAllowedMethods_HEAD    = CorsRuleSTATUSAllowedMethods("HEAD")
	CorsRuleSTATUSAllowedMethods_MERGE   = CorsRuleSTATUSAllowedMethods("MERGE")
	CorsRuleSTATUSAllowedMethods_OPTIONS = CorsRuleSTATUSAllowedMethods("OPTIONS")
	CorsRuleSTATUSAllowedMethods_POST    = CorsRuleSTATUSAllowedMethods("POST")
	CorsRuleSTATUSAllowedMethods_PUT     = CorsRuleSTATUSAllowedMethods("PUT")
)
