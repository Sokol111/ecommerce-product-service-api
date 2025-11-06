package events

// Event type constants - single source of truth for event names
// These match the schema names and should be used in EventMetadata.EventType
const (
	EventTypeProductCreated = "ProductCreatedEvent"
	EventTypeProductUpdated = "ProductUpdatedEvent"
)
