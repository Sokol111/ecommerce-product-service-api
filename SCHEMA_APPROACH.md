# Schema-First Approach

## Overview

This project follows a **Schema-First approach with Avro as the single source of truth** and uses a **metadata + payload pattern** for all events.

## Topic Strategy

**Single Topic for All Product Events**: `catalog.product.events`

All product lifecycle events (created, updated, deleted) are published to one topic:

```
Topic: catalog.product.events
├── ProductCreated events
├── ProductUpdated events
└── ProductDeleted events (future)
```

**Why Single Topic?**

1. **Event Ordering** - All events for a single product are guaranteed to be in order (using product_id as partition key)
2. **Simpler Consumption** - Consumers subscribe to one topic instead of multiple
3. **Event Sourcing** - Complete event stream in chronological order
4. **Schema Evolution** - Version managed by Schema Registry, not topic name
5. **Scalability** - Different products go to different partitions (parallel processing)

**Partition Strategy**:
```go
// product_id as partition key
producer.Send(kafka.Message{
    Topic: "catalog.product.events",
    Key:   []byte(productID),  // Same product → same partition
    Value: eventData,
})
```

This ensures:
- ✅ Events for product-123 are always in order
- ✅ Events for different products can be processed in parallel
- ✅ No need for external ordering coordination

## Event Structure Pattern

All events follow a consistent two-field structure:

```json
{
  "metadata": {
    "event_id": "unique-uuid",
    "event_type": "ProductCreated",
    "source": "ecommerce-product-service",
    "timestamp": 1699123456789,
    "schema_version": 1,
    "trace_id": "optional-opentelemetry-trace-id",
    "correlation_id": "optional-correlation-id"
  },
  "payload": {
    // Business-specific data
  }
}
```

### Why Metadata in Avro (not Kafka Headers)?

**Chosen Approach: Metadata in Avro Payload** ✅

**Advantages**:
- ✅ Everything versioned through Schema Registry
- ✅ Guaranteed delivery with payload (no header loss)
- ✅ Better for event sourcing and debugging
- ✅ Complete event in one place
- ✅ Easier to replay events from event store

**Alternative (Kafka Headers)** ❌ Rejected:
- ❌ Headers not versioned in Schema Registry
- ❌ Can be lost during system transfers
- ❌ Split between headers and payload complicates debugging
- ❌ Less reliable for event sourcing

### Benefits of Metadata + Payload Pattern

1. **Separation of Concerns**
   - Technical/observability fields in `metadata`
   - Business data in `payload`

2. **Reusability**
   - `EventMetadata` type shared across all events
   - No duplication of common fields

3. **Clean Domain Model**
   - Payload contains only business-relevant data
   - Metadata doesn't pollute domain objects

4. **Easy to Extend**
   - Add new metadata fields without touching business logic
   - Add observability fields globally

## Architecture Decision

### Before (Flat Structure with Duplication)
```
❌ ProductCreatedEvent
   ├── event_id (duplicated in every event)
   ├── event_type (duplicated in every event)
   ├── source (duplicated in every event)
   ├── created_at (duplicated in every event)
   ├── trace_id (duplicated in every event)
   ├── correlation_id (duplicated in every event)
   └── payload
       └── ProductCreatedPayload

❌ ProductUpdatedEvent
   ├── event_id (duplicated again)
   ├── event_type (duplicated again)
   └── ... (same duplication)
   
Problem: Metadata fields duplicated in every event schema
```

### After (Metadata + Payload with Shared Type)
```
✅ EventMetadata (event_metadata.avsc) - REUSABLE TYPE
   ├── event_id
   ├── event_type
   ├── source
   ├── timestamp
   ├── schema_version
   ├── trace_id
   └── correlation_id

✅ ProductCreatedEvent
   ├── metadata: EventMetadata (reference)
   └── payload: ProductCreatedPayload

✅ ProductUpdatedEvent
   ├── metadata: EventMetadata (reference)
   └── payload: ProductUpdatedPayload
   
Benefits: Metadata defined once, reused everywhere
```

## Why Avro as Source of Truth?

1. **Runtime Usage** - Avro schemas are used for actual serialization/deserialization
2. **Schema Registry** - Avro schemas are registered and validated
3. **Schema Evolution** - Avro has built-in compatibility checking
4. **Type Safety** - Code generation from Avro ensures type correctness
5. **Binary Format** - Compact and efficient for Kafka

## Why Keep AsyncAPI?

AsyncAPI serves a different purpose:

1. **Documentation** - Human-readable event flow documentation
2. **Channels & Operations** - Describes Kafka topics and message routing
3. **API Contract** - Defines the messaging API contract
4. **Tooling** - Enables AsyncAPI Studio visualization and code generation

## Workflow

### Modifying Event Schemas

1. **Edit Avro schema** in `avro/` directory
   ```bash
   vim avro/product_created.avsc
   # Only modify the 'payload' section for business logic changes
   # Metadata is shared and rarely needs changes
   ```

2. **AsyncAPI automatically references** the updated schema
   - No changes needed to `asyncapi.yaml`
   
3. **CI/CD validates** schema compatibility
   - Ensures backward/forward compatibility
   - Publishes to Schema Registry

4. **Code regeneration** happens automatically
   - Go event types updated
   - JS clients updated

### Adding New Events

1. **Create Avro schema** (e.g., `product_deleted.avsc`)
   ```json
   {
     "type": "record",
     "name": "ProductDeletedEvent",
     "namespace": "com.ecommerce.events.product",
     "fields": [
       {
         "name": "metadata",
         "type": "com.ecommerce.events.EventMetadata"
       },
       {
         "name": "payload",
         "type": {
           "type": "record",
           "name": "ProductDeletedPayload",
           "fields": [...]
         }
       }
     ]
   }
   ```

2. **Add message reference** in `asyncapi.yaml`:
   ```yaml
   ProductDeleted:
     contentType: application/avro
     payload:
       schemaFormat: application/vnd.apache.avro+json;version=1.9.0
       schema:
         $ref: '../avro/product_deleted.avsc'
   ```

### Modifying Metadata Fields (Rare)

If you need to add a new metadata field (e.g., `user_id`):

1. Edit `avro/event_metadata.avsc`
2. Add field with default value for backward compatibility
3. All events automatically inherit the new field

## Benefits

✅ **No Duplication** - Schema defined once in Avro  
✅ **Single Source of Truth** - Avro is the authority  
✅ **Consistency** - Impossible to have mismatched schemas  
✅ **Maintainability** - Changes in one place only  
✅ **Best of Both** - Avro for runtime, AsyncAPI for docs  
✅ **Reusable Metadata** - EventMetadata type shared across all events  
✅ **Clean Separation** - Technical vs business data clearly separated  
✅ **Reliable** - All data in Avro payload, not split between headers and body  

## Related Files

- `avro/event_metadata.avsc` - Shared metadata type (reusable)
- `avro/product_*.avsc` - Event schemas (source of truth)
- `asyncapi/asyncapi.yaml` - Event flow documentation
- `README.md` - Updated with metadata+payload approach
- CI/CD workflows remain unchanged
