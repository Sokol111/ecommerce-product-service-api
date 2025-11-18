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

// Combined schemas with EventMetadata embedded for Avro serialization
// These resolve named type references like "com.ecommerce.events.EventMetadata"

//go:embed schemas/product_created_combined.json
var ProductCreatedCombinedSchema []byte // Use this for Serialize() - includes EventMetadata

//go:embed schemas/product_updated_combined.json
var ProductUpdatedCombinedSchema []byte // Use this for Serialize() - includes EventMetadata

