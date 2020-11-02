package ingester

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

type Mapping struct {
	Column     string                 `json:"column,omitempty"`
	Datatype   string                 `json:"datatype,omitempty"`
	Properties map[string]interface{} `json:"Properties,omitempty"`
}
