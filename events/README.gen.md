# Generated Messaging API

This package is auto-generated from AsyncAPI and Avro schemas.

## 📦 What's Included

- **Go Types**: Generated from Avro schemas using `avrogen`
- **Event Constants**: Type-safe constants for event types and topics
- **Schema Embeddings**: Avro schemas embedded as `[]byte` for runtime access
- **Sealed Interface**: Exhaustiveness checking for event handlers

## 🔧 Usage with Schema Registry

This package is designed to work with Confluent Schema Registry for production use.

### Producer Example

```go
import (
    "events"
    "github.com/Sokol111/ecommerce-commons/pkg/messaging/kafka/schemaregistry"
)

// Create Schema Registry serializer
serializer, err := schemaregistry.NewSerializer(schemaregistry.Config{
    URL:                 "http://localhost:8081",
    AutoRegisterSchemas: true,
})
if err != nil {
    log.Fatal(err)
}
defer serializer.Close()

// Create event
event := &events.YourEventName{
    Metadata: events.EventMetadata{
        EventID:       uuid.New().String(),
        EventType:     events.EventTypeYourEvent, // Use constant!
        Source:        "your-service",
        Timestamp:     time.Now().UnixMilli(),
        CorrelationID: correlationID,
        TraceID:       traceID,
    },
    Payload: events.YourEventPayload{
        // ... your payload fields
    },
}

// Serialize with Schema Registry (auto-registers schema)
bytes, err := serializer.Serialize("your.topic-value", event)
if err != nil {
    log.Fatal(err)
}

// Send to Kafka
producer.Produce(&kafka.Message{
    TopicPartition: kafka.TopicPartition{
        Topic:     &events.TopicYourTopic,
        Partition: kafka.PartitionAny,
    },
    Value: bytes, // Contains [0x00][schema_id][avro_data]
    Key:   []byte(entityID),
})
```

### Consumer Example

```go
import (
    "events"
    "github.com/Sokol111/ecommerce-commons/pkg/messaging/kafka/schemaregistry"
)

// Create Schema Registry deserializer
deserializer, err := schemaregistry.NewDeserializer(schemaregistry.Config{
    URL: "http://localhost:8081",
})
if err != nil {
    log.Fatal(err)
}
defer deserializer.Close()

// Receive from Kafka
msg := <-consumer.Events()

// Deserialize (schema fetched from Registry by ID)
event, err := deserializer.Deserialize("your.topic-value", msg.Value)
if err != nil {
    log.Fatal(err)
}

// Type switch with exhaustiveness checking
switch e := event.(type) {
case *events.YourEventName:
    handleYourEvent(e)
// Add cases for all event types - linter will warn if you miss one
default:
    log.Printf("unhandled event type: %T", event)
}
```

## 🎯 Event Type Constants

Always use the generated constants for `EventMetadata.EventType`:

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

## Available Schemas" >> events/README.gen.md
echo "" >> events/README.gen.md
for schema in avro/*.avsc; do
  basename=$(basename "$schema" .avsc)
  varname=$(echo "$basename" | sed -E 's/(^|_)([a-z])/\U\2/g')
  schema_type=$(jq -r '.name' "$schema" 2>/dev/null || echo "Unknown")
  echo "- \`events.${varname}Schema\` - $schema_type (embedded Avro schema)" >> events/README.gen.md
done

cat >> events/README.gen.md << 'EOF'

## 📖 Schema Registry Benefits

### Why Schema Registry?

- ✅ **Centralized schema storage** - Single source of truth
- ✅ **Automatic versioning** - Track schema evolution
- ✅ **Compatibility checking** - Prevent breaking changes
- ✅ **Schema ID in messages** - Efficient wire format
- ✅ **No code regeneration** - Schema changes don't require service redeployment

### Wire Format

Messages are serialized as: `[0x00][schema_id (4 bytes)][avro_data]`

This allows consumers to fetch the exact schema used for serialization from the Registry.

### Compatibility Modes

- **BACKWARD** (recommended): New consumers can read old data
- **FORWARD**: Old consumers can read new data
- **FULL**: Both BACKWARD and FORWARD
- **NONE**: No compatibility checks

## 🔧 Configuration

Add to your service config:

```yaml
kafka:
  brokers: "kafka:29092"
  schema-registry:
    url: "http://schema-registry:8081"
    auto-register-schemas: true
    cache-capacity: 1000
```

