# Generated Messaging API

This package is auto-generated from AsyncAPI and Avro schemas.

## Usage

```go
import "github.com/Sokol111/ecommerce-product-service-api/events"

// Initialize serializer
serializer, err := events.NewAvroSerializer()
if err != nil {
    log.Fatal(err)
}

// Register schemas (schemas are embedded)
// Example: Register any event schema
err = serializer.RegisterSchema("YourEventName", string(events.YourEventSchema))
if err != nil {
    log.Fatal(err)
}

// Serialize event
event := &events.YourEventName{
    Metadata: events.EventMetadata{
        EventId:   "unique-id",
        EventType: "YourEventName",
        // ... other metadata fields
    },
    Payload: events.YourEventPayload{
        // ... your payload fields
    },
}
data, err := serializer.Serialize("YourEventName", event)
if err != nil {
    log.Fatal(err)
}

// Deserialize event
var receivedEvent events.YourEventName
err = serializer.Deserialize("YourEventName", data, &receivedEvent)
if err != nil {
    log.Fatal(err)
}
```

## Embedded Resources

- AsyncAPI specification: `events.AsyncAPISpec`
- Avro schemas: `events.*Schema` variables

## Available Schemas

The following Avro schemas are available:

- `events.EventMetadataSchema` - EventMetadata
- `events.ProductCreatedSchema` - ProductCreatedEvent
- `events.ProductUpdatedSchema` - ProductUpdatedEvent
