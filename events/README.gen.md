# Generated Messaging API

This package is auto-generated from AsyncAPI and Avro schemas.

## Usage

### Typed API (Recommended)

```go
import "github.com/Sokol111/ecommerce-product-service-api/events"

// Serialize event (no registration needed!)
// Use event type constants for type safety
event := &events.YourEventName{
    Metadata: &events.EventMetadata{
        EventId:       "unique-id",
        EventType:     events.EventTypeYourEvent, // ← Type-safe constant
        Source:        "your-service",
        Timestamp:     time.Now().UnixMilli(),
        SchemaVersion: 1,
    },
    Payload: &events.YourEventPayload{
        // ... your payload fields
    },
}

// Simply call the typed marshal function
data, err := events.MarshalYourEventName(event)
if err != nil {
    log.Fatal(err)
}

// Deserialize with typed function
receivedEvent, err := events.UnmarshalYourEventName(data)
if err != nil {
    log.Fatal(err)
}
```

### Generic API (for dynamic use cases)

```go
// Marshal any event by schema name
data, err := events.Marshal("YourEventName", event)
if err != nil {
    log.Fatal(err)
}

// Unmarshal any event by schema name
var event events.YourEventName
err = events.Unmarshal("YourEventName", data, &event)
if err != nil {
    log.Fatal(err)
}
```

## Features

✅ **No manual schema registration** - schemas are pre-compiled at package init  
✅ **Type-safe helpers** - dedicated functions for each event type  
✅ **Type-safe event constants** - use `EventType*` constants instead of strings  
✅ **Zero runtime overhead** - schemas parsed once using sync.Once  
✅ **Embedded resources** - all schemas included in the package  

## Event Type Constants

Always use the generated constants for `EventMetadata.EventType` to ensure consistency:

- `events.EventTypeProductCreated` = "ProductCreatedEvent"
- `events.EventTypeProductUpdated` = "ProductUpdatedEvent"

## Embedded Resources

- AsyncAPI specification: `events.AsyncAPISpec`
- Avro schemas: `events.*Schema` variables

## Exhaustiveness Checking

All event types implement the sealed `Event` interface:

```go
type Event interface {
    isEvent() // unexported method - seals the interface
}
```

This enables exhaustiveness checking in type switches. When using `golangci-lint` with 
the `exhaustive` linter, you'll get compile-time warnings if you don't handle all event types:

```go
func handleEvent(event events.Event) error {
    switch e := event.(type) {
    case *events.ProductCreatedEvent:
        // handle created
    case *events.ProductUpdatedEvent:
        // handle updated
    // exhaustive linter will warn if you forget any Event type!
    default:
        return fmt.Errorf("unhandled event: %T", event)
    }
}
```

**Configure golangci-lint:**

```yaml
linters-settings:
  exhaustive:
    default-signifies-exhaustive: false
    explicit-exhaustive-switch: true
```

## Available Event Types

The following event types have typed serialization helpers:

- **ProductCreatedEvent**
  - `MarshalProductCreatedEvent(event *ProductCreatedEvent) ([]byte, error)`
  - `UnmarshalProductCreatedEvent(data []byte) (*ProductCreatedEvent, error)`

- **ProductUpdatedEvent**
  - `MarshalProductUpdatedEvent(event *ProductUpdatedEvent) ([]byte, error)`
  - `UnmarshalProductUpdatedEvent(data []byte) (*ProductUpdatedEvent, error)`

## Available Schemas

- `events.EventMetadataSchema` - EventMetadata (embedded Avro schema)
- `events.ProductCreatedSchema` - ProductCreatedEvent (embedded Avro schema)
- `events.ProductUpdatedSchema` - ProductUpdatedEvent (embedded Avro schema)
