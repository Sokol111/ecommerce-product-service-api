# ecommerce-product-service-api

API definitions for Product service including REST API (OpenAPI) and Messaging API (AsyncAPI + Avro).

## Architecture Approach

This repository follows a **Schema-First approach with Avro as the source of truth**:

- 🎯 **Avro schemas** (`avro/`) - **Source of Truth** for event data structures
  - Used for runtime serialization/deserialization
  - Registered in Confluent Schema Registry
  - Support schema evolution and compatibility checks
  
- 📖 **AsyncAPI** (`asyncapi/`) - Documentation layer
  - Describes channels, operations, and message flows
  - References Avro schemas (no schema duplication)
  - Provides human-readable API documentation
  
- 🔧 **OpenAPI** (`openapi/`) - REST API specification
  - Defines HTTP endpoints and request/response formats

This approach **eliminates schema duplication** and ensures Avro schemas remain the single source of truth for event structures.

## Structure

```
.
├── avro/                        # 🎯 Source of Truth for event schemas
│   ├── event_metadata.avsc      # Common EventMetadata type (reusable)
│   ├── product_created.avsc     # ProductCreatedEvent schema
│   └── product_updated.avsc     # ProductUpdatedEvent schema
├── asyncapi/                    # 📖 Event flow documentation
│   └── asyncapi.yaml            # References Avro schemas
├── openapi/                     # 🔧 REST API specification
│   └── openapi.yml              # OpenAPI/REST API spec
├── api/                         # Generated Go REST API (git-ignored)
├── events/                      # Generated Go events API (git-ignored)
└── js-client/                   # Generated JS client (git-ignored)
```

**Key principle**: When modifying event structures, edit Avro schemas only. AsyncAPI references them automatically.

### Event Structure (Metadata + Payload Pattern)

All events follow a consistent two-field structure:

```json
{
  "metadata": {
    "event_id": "uuid",
    "event_type": "ProductCreated",
    "source": "ecommerce-product-service",
    "timestamp": 1699123456789,
    "schema_version": 1,
    "trace_id": "optional-trace-id",
    "correlation_id": "optional-correlation-id"
  },
  "payload": {
    // Business-specific data (varies per event type)
  }
}
```

**Benefits**:
- ✅ Clear separation: technical metadata vs business data
- ✅ Reusable `EventMetadata` type across all events
- ✅ All data in Avro payload (not Kafka headers) for reliability
- ✅ Better for event sourcing and debugging

## REST API

REST API is automatically generated from OpenAPI specification.

See [openapi/openapi.yml](openapi/openapi.yml) for full API documentation.

## Messaging API (Kafka Events)

### Topic Naming

All product events are published to a single topic: **`catalog.product.events`**

**Why single topic?**
- ✅ Guaranteed event ordering per product (using product_id as partition key)
- ✅ Simpler consumer subscription (one topic instead of multiple)
- ✅ Better for event sourcing (complete event stream in order)
- ✅ Schema evolution via Schema Registry (not in topic name)

### Events

#### ProductCreated

Published when a new product is created.

**Topic:** `catalog.product.events`  
**Event Type:** `ProductCreated` (in metadata.event_type)  
**Schema:** [avro/product_created.avsc](avro/product_created.avsc)

#### ProductUpdated

Published when a product is updated.

**Topic:** `catalog.product.events`  
**Event Type:** `ProductUpdated` (in metadata.event_type)  
**Schema:** [avro/product_updated.avsc](avro/product_updated.avsc)

### Usage in Services

Generated code is available as artifacts from GitHub Actions releases.

```go
import "github.com/Sokol111/ecommerce-product-service-api/events"

// Initialize serializer
serializer, err := events.NewAvroSerializer()
serializer.RegisterSchema("ProductCreatedEvent", string(events.ProductCreatedSchema))

// Create event with metadata + payload structure
event := &events.ProductCreatedEvent{
    Metadata: &events.EventMetadata{
        EventID:       uuid.New().String(),
        EventType:     "ProductCreated",
        Source:        "ecommerce-product-service",
        Timestamp:     time.Now().UnixMilli(),
        SchemaVersion: 1,
        TraceID:       extractTraceID(ctx),
        CorrelationID: extractCorrelationID(ctx),
    },
    Payload: &events.ProductCreatedPayload{
        ProductID:  product.ID,
        Name:       product.Name,
        Price:      product.Price,
        Quantity:   product.Quantity,
        Enabled:    product.Enabled,
        Version:    product.Version,
        CreatedAt:  product.CreatedAt.UnixMilli(),
        ModifiedAt: product.ModifiedAt.UnixMilli(),
    },
}

data, err := serializer.Serialize("ProductCreatedEvent", event)

// Send to Kafka with product_id as partition key for ordering
producer.Send(ctx, kafka.Message{
    Topic: "catalog.product.events",
    Key:   []byte(product.ID),  // Ensures ordering for same product
    Value: data,
})
```

## CI/CD

GitHub Actions automatically:

1. Validates OpenAPI and AsyncAPI specs
2. Generates Go and JS code
3. Publishes Avro schemas to Schema Registry
4. Creates releases with generated artifacts

Triggered on changes to:

- `openapi/openapi.yml`
- `asyncapi/**`
- `avro/**`

## Schema Evolution

Avro supports backward/forward compatible schema evolution:

- ✅ Add optional fields (with defaults)
- ✅ Remove optional fields
- ❌ Change field types
- ❌ Rename fields without aliases
- ❌ Remove required fields

**Workflow**:
1. Modify Avro schema files in `avro/` directory
2. AsyncAPI automatically references the updated schemas
3. CI/CD validates compatibility and publishes to Schema Registry
4. Consumers receive schema updates automatically

## Related Projects

- [ecommerce-product-service](https://github.com/Sokol111/ecommerce-product-service) - REST API server & event producer
- [ecommerce-product-query-service](https://github.com/Sokol111/ecommerce-product-query-service) - Event consumer
- [ecommerce-infrastructure](https://github.com/Sokol111/ecommerce-infrastructure) - CI/CD workflows

## Documentation

- [SCHEMA_APPROACH.md](SCHEMA_APPROACH.md) - Schema-first strategy and metadata+payload pattern
- [TOPIC_NAMING.md](TOPIC_NAMING.md) - Kafka topic naming strategy and best practices
- [USAGE_EXAMPLE.md](USAGE_EXAMPLE.md) - Code examples for producers and consumers

## Resources

- [OpenAPI Specification](https://spec.openapis.org/oas/v3.0.0)
- [AsyncAPI Specification](https://www.asyncapi.com/docs/reference/specification/v3.0.0)
- [Apache Avro](https://avro.apache.org/docs/current/)
- [Confluent Schema Registry](https://docs.confluent.io/platform/current/schema-registry/)
