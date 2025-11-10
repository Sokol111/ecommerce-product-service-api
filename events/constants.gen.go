package events

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
