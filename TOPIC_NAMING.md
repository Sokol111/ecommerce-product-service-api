# Kafka Topic Naming Strategy

## Current Approach

**Topic**: `catalog.product.events`

All product events (created, updated, deleted) are published to a single topic.

## Topic Naming Breakdown

```
catalog . product . events
   │         │        │
   │         │        └─── Event stream indicator
   │         └──────────── Entity/aggregate type
   └────────────────────── Domain/namespace
```

## Why This Approach?

### ✅ Single Topic for Entity Events

```
catalog.product.events
├── ProductCreated
├── ProductUpdated
└── ProductDeleted
```

**Advantages**:
1. **Event Ordering** - All events for a product are in order
2. **Simpler Consumer** - Subscribe to one topic
3. **Event Sourcing** - Complete event stream
4. **Schema Evolution** - Via Schema Registry, not topic versioning

**Use partition key** (`product_id`) to ensure ordering:
```go
producer.Send(kafka.Message{
    Topic: "catalog.product.events",
    Key:   []byte(productID),  // Partition key
    Value: eventData,
})
```

## Alternative Approaches (NOT Used)

### ❌ Separate Topics per Event Type

```
catalog.product.created.v1
catalog.product.updated.v1
catalog.product.deleted.v1
```

**Problems**:
- Hard to maintain event ordering
- Consumers need multiple subscriptions
- More topics to manage
- Difficult event sourcing

### ❌ Version in Topic Name

```
catalog.product.events.v1
catalog.product.events.v2
```

**Problems**:
- Need to create new topic for schema changes
- Complex migration between versions
- Schema Registry handles versioning better

### ❌ Domain-Wide Topic

```
catalog.events  (all catalog events)
```

**When to use**:
- Very small domains
- Need global ordering across entities
- Legacy event sourcing systems

**Our choice**: Entity-level topics for better scaling and isolation

## Topic Configuration

### Recommended Settings

```yaml
# Kafka topic config for catalog.product.events
topic: catalog.product.events
partitions: 10              # Adjust based on product volume
replication-factor: 3       # High availability
retention.ms: 604800000     # 7 days (or longer for event sourcing)
cleanup.policy: delete      # or 'compact' for event sourcing
min.insync.replicas: 2      # Data durability
```

### Partition Strategy

```go
// Hash of product_id determines partition
partitionKey := productID

// All events for same product go to same partition
// → Guaranteed ordering per product
// → Different products can be processed in parallel
```

### Scaling Considerations

**Current**: 10 partitions
- Can handle ~10 concurrent consumers (one per partition)
- Each partition maintains order
- Add more partitions if needed (note: can't reduce partitions)

## Event Type Identification

Event type is in **metadata**, not topic:

```json
{
  "metadata": {
    "event_type": "ProductCreated"  // ← Identifies event type
  },
  "payload": { ... }
}
```

**Consumer filtering**:
```go
switch event.Metadata.EventType {
case "ProductCreated":
    // Handle creation
case "ProductUpdated":
    // Handle update
}
```

## Schema Versioning

**NOT in topic name**, managed by Schema Registry:

```
Topic: catalog.product.events (never changes)
Schema: ProductCreatedEvent v1, v2, v3... (in Schema Registry)
```

**Schema evolution**:
1. Update Avro schema in `avro/product_created.avsc`
2. Schema Registry validates compatibility
3. Publish new schema version
4. Consumers auto-detect version

**Compatibility modes**:
- `BACKWARD` - New consumers can read old data
- `FORWARD` - Old consumers can read new data
- `FULL` - Both directions (recommended)

## Multi-Domain Topics

For other domains, follow same pattern:

```
catalog.product.events      # Product events
catalog.category.events     # Category events
inventory.stock.events      # Stock events
order.order.events          # Order events
```

**Benefits**:
- Clear domain boundaries
- Independent scaling per domain
- Easy to manage permissions/ACLs

## When to Create New Topic

Create a new topic when:
- ✅ New entity/aggregate (e.g., `catalog.category.events`)
- ✅ Different domain (e.g., `inventory.stock.events`)
- ✅ Different access patterns (e.g., public vs internal events)
- ✅ Different retention requirements

Do NOT create new topic for:
- ❌ New event type for same entity
- ❌ Schema version changes
- ❌ Different event sources (use metadata.source)

## Migration from Old Naming

If migrating from:
```
catalog.product.created.v1  → catalog.product.events
catalog.product.updated.v1  → catalog.product.events
```

**Migration steps**:
1. Create new topic `catalog.product.events`
2. Dual-publish to both old and new topics (transition period)
3. Update consumers to read from new topic
4. Monitor until all consumers migrated
5. Stop publishing to old topics
6. Eventually delete old topics

## Best Practices

1. **Use namespace** - Prefix with domain (`catalog`, `inventory`, etc.)
2. **Entity-level topics** - One topic per entity type
3. **Partition by ID** - Use entity ID as partition key
4. **Event type in metadata** - Not in topic name
5. **Version via Schema Registry** - Not in topic name
6. **Descriptive names** - Clear what events are on the topic
7. **Plural events** - Use `.events` suffix to indicate stream

## Examples from Other Domains

```
# E-commerce platform
catalog.product.events
catalog.category.events
inventory.stock.events
pricing.price.events
order.order.events
payment.transaction.events
shipping.shipment.events
user.account.events

# Each topic:
# - Contains all event types for that entity
# - Uses entity ID as partition key
# - Versioned via Schema Registry
```
