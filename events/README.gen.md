# Generated Messaging API

This package is auto-generated from AsyncAPI and Avro schemas.

## Usage

### Typed API (Recommended)

```go
import "github.com/Sokol111/ecommerce-product-service-api/events"

// Serialize event (no registration needed!)
event := &events.YourEventName{
    Metadata: &events.EventMetadata{
        EventId:       "unique-id",
        EventType:     "YourEventName",
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
✅ **Zero runtime overhead** - schemas parsed once using sync.Once  
✅ **Embedded resources** - all schemas included in the package  

## Embedded Resources

- AsyncAPI specification: `events.AsyncAPISpec`
- Avro schemas: `events.*Schema` variables

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
