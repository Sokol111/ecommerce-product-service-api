package events

import "reflect"

// Event type constants - single source of truth for event names
// These match the schema names and should be used in EventMetadata.EventType
const (
	EventTypeProductCreated = "ProductCreatedEvent"
	EventTypeProductUpdated = "ProductUpdatedEvent"
)

// Event is a sealed interface that all event types must implement.
// This enables exhaustiveness checking when using type switches.
//
// When using exhaustive linter (golangci-lint), it will warn you if you
// don't handle all concrete Event types in a type switch.
type Event interface {
	isEvent() // unexported method seals the interface
}

// ProductCreatedEvent implements Event interface
func (*ProductCreatedEvent) isEvent() {}

// ProductUpdatedEvent implements Event interface
func (*ProductUpdatedEvent) isEvent() {}

// Topic constants - Kafka topics defined in AsyncAPI spec
const (
	TopicCatalogProductEvents = "catalog.product.events"
)

// Schema name constants - Avro schema full names (namespace.name)
const (
	SchemaNameProductCreated = "com.ecommerce.events.product.ProductCreatedEvent"
	SchemaNameProductUpdated = "com.ecommerce.events.product.ProductUpdatedEvent"
)

// DefaultTypeMapping is a pre-configured mapping of Avro schema full names to Go types.
// Use this with consumer.RegisterTypeMapping() to deserialize events.
//
// Example:
//   consumer.RegisterTypeMapping(events.DefaultTypeMapping)
var DefaultTypeMapping = map[string]reflect.Type{
	SchemaNameProductCreated: reflect.TypeOf(ProductCreatedEvent{}),  
	SchemaNameProductUpdated: reflect.TypeOf(ProductUpdatedEvent{}),  
}
