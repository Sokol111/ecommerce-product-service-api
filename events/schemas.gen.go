package events

import _ "embed"

//go:embed asyncapi.yaml
var AsyncAPISpec []byte

//go:embed schemas/event_metadata.avsc
var EventMetadataSchema []byte

//go:embed schemas/product_created.avsc
var ProductCreatedSchema []byte

//go:embed schemas/product_updated.avsc
var ProductUpdatedSchema []byte

// Combined schemas with EventMetadata inlined for Avro serialization
// These schemas have the EventMetadata type definition embedded inline

//go:embed schemas/product_created_combined.json
var ProductCreatedCombinedSchema []byte // Use this for Serialize() - has EventMetadata inlined

//go:embed schemas/product_updated_combined.json
var ProductUpdatedCombinedSchema []byte // Use this for Serialize() - has EventMetadata inlined

