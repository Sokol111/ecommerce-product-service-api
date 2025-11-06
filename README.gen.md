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
err = serializer.RegisterSchema("ProductCreatedEvent", string(events.ProductCreatedSchema))

// Serialize event
event := &events.ProductCreatedEvent{...}
data, err := serializer.Serialize("ProductCreatedEvent", event)
```

## Embedded Resources

- AsyncAPI specification: `events.AsyncAPISpec`
- Avro schemas: `events.*Schema` variables
