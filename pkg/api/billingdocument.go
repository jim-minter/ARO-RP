package api

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

// BillingDocuments represents billing documents.
// pkg/database/cosmosdb requires its definition.
type BillingDocuments struct {
	Count            int                `json:"_count,omitempty"`
	ResourceID       string             `json:"_rid,omitempty"`
	BillingDocuments []*BillingDocument `json:"Documents,omitempty"`
}

func (c *BillingDocuments) String() string {
	return encodeJSON(c)
}

// BillingDocument represents a billing document.
// pkg/database/cosmosdb requires its definition.
type BillingDocument struct {
	MissingFields

	ID          string                 `json:"id,omitempty"`
	ResourceID  string                 `json:"_rid,omitempty"`
	Timestamp   int                    `json:"_ts,omitempty"`
	Self        string                 `json:"_self,omitempty"`
	ETag        string                 `json:"_etag,omitempty" deep:"-"`
	Attachments string                 `json:"_attachments,omitempty"`
	LSN         int                    `json:"_lsn,omitempty"`
	Metadata    map[string]interface{} `json:"_metadata,omitempty"`

	Billing *Billing `json:"billing,omitempty"`

	Key                       string `json:"key,omitempty"`
	ClusterResourceGroupIDKey string `json:"clusterResourceGroupIDKey,omitempty"`
	InfraID                   string `json:"infraId,omitempty"`
}

func (c *BillingDocument) String() string {
	return encodeJSON(c)
}
