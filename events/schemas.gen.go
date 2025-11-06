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

