# Event Usage Examples

## Event Structure

All events follow the **metadata + payload** pattern:

```json
{
  "metadata": {
    "event_id": "550e8400-e29b-41d4-a716-446655440000",
    "event_type": "ProductCreated",
    "source": "ecommerce-product-service",
    "timestamp": 1699123456789,
    "schema_version": 1,
    "trace_id": "80f198ee56343ba864fe8b2a57d3eff7",
    "correlation_id": "req-12345"
  },
  "payload": {
    "product_id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Laptop",
    "price": 999.99,
    "quantity": 10,
    "image_id": "img-001",
    "enabled": true,
    "version": 1,
    "created_at": 1699123456789,
    "modified_at": 1699123456789
  }
}
```

## Producer Example (Go)

### Publishing ProductCreated Event

```go
package main

import (
    "context"
    "time"
    
    "github.com/google/uuid"
    "github.com/Sokol111/ecommerce-product-service-api/events"
)

func PublishProductCreatedEvent(ctx context.Context, product *Product) error {
    // Extract tracing info from context
    traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
    correlationID := ctx.Value("correlation_id").(string)
    
    // Build event with metadata + payload
    event := &events.ProductCreatedEvent{
        Metadata: &events.EventMetadata{
            EventID:       uuid.New().String(),
            EventType:     "ProductCreated",
            Source:        "ecommerce-product-service",
            Timestamp:     time.Now().UnixMilli(),
            SchemaVersion: 1,
            TraceID:       &traceID,
            CorrelationID: &correlationID,
        },
        Payload: &events.ProductCreatedPayload{
            ProductID:  product.ID.String(),
            Name:       product.Name,
            Price:      product.Price,
            Quantity:   product.Quantity,
            ImageID:    product.ImageID,
            Enabled:    product.Enabled,
            Version:    product.Version,
            CreatedAt:  product.CreatedAt.UnixMilli(),
            ModifiedAt: product.ModifiedAt.UnixMilli(),
        },
    }
    
    // Serialize
    data, err := serializer.Serialize("ProductCreatedEvent", event)
    if err != nil {
        return fmt.Errorf("failed to serialize event: %w", err)
    }
    
    // Publish to Kafka
    return kafkaProducer.Send(ctx, "catalog.product.events", data)
}
```

### Publishing ProductUpdated Event

```go
func PublishProductUpdatedEvent(ctx context.Context, product *Product) error {
    traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
    correlationID := ctx.Value("correlation_id").(string)
    
    event := &events.ProductUpdatedEvent{
        Metadata: &events.EventMetadata{
            EventID:       uuid.New().String(),
            EventType:     "ProductUpdated",
            Source:        "ecommerce-product-service",
            Timestamp:     time.Now().UnixMilli(),
            SchemaVersion: 1,
            TraceID:       &traceID,
            CorrelationID: &correlationID,
        },
        Payload: &events.ProductUpdatedPayload{
            ProductID:  product.ID.String(),
            Name:       product.Name,
            Price:      product.Price,
            Quantity:   product.Quantity,
            ImageID:    product.ImageID,
            Enabled:    product.Enabled,
            Version:    product.Version,
            CreatedAt:  product.CreatedAt.UnixMilli(),
            ModifiedAt: product.ModifiedAt.UnixMilli(),
        },
    }
    
    data, err := serializer.Serialize("ProductUpdatedEvent", event)
    if err != nil {
        return fmt.Errorf("failed to serialize event: %w", err)
    }
    
    return kafkaProducer.Send(ctx, "catalog.product.events", data)
}
```

### Using Partition Key for Ordering

```go
// Ensure all events for the same product go to the same partition
func PublishEvent(ctx context.Context, productID string, eventData []byte) error {
    return kafkaProducer.Send(ctx, kafka.Message{
        Topic: "catalog.product.events",
        Key:   []byte(productID),  // Partition key = product_id
        Value: eventData,
    })
}

// This guarantees:
// - All events for product-123 go to the same partition
// - Events for product-123 are ordered
// - Events for different products can be in different partitions (parallelism)
```

## Consumer Example (Go)

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/Sokol111/ecommerce-product-service-api/events"
)

func ConsumeProductEvents(ctx context.Context) error {
    consumer := kafkaConsumer.Subscribe("catalog.product.events")
    
    for {
        msg, err := consumer.ReadMessage(ctx)
        if err != nil {
            return err
        }
        
        // Deserialize
        var event interface{}
        eventType, err := deserializer.Deserialize(msg.Value, &event)
        if err != nil {
            log.Printf("Failed to deserialize: %v", err)
            continue
        }
        
        // Handle based on event type from metadata
        switch eventType {
        case "ProductCreated":
            handleProductCreated(ctx, event.(*events.ProductCreatedEvent))
        case "ProductUpdated":
            handleProductUpdated(ctx, event.(*events.ProductUpdatedEvent))
        default:
            log.Printf("Unknown event type: %s", eventType)
        }
    }
}

func handleProductCreated(ctx context.Context, event *events.ProductCreatedEvent) {
    // Access metadata
    log.Printf("Processing event %s from %s", 
        event.Metadata.EventID, 
        event.Metadata.Source)
    
    // Use trace_id for distributed tracing
    if event.Metadata.TraceID != nil {
        ctx = trace.ContextWithTraceID(ctx, *event.Metadata.TraceID)
    }
    
    // Access business data from payload
    product := &Product{
        ID:         event.Payload.ProductID,
        Name:       event.Payload.Name,
        Price:      event.Payload.Price,
        Quantity:   event.Payload.Quantity,
        Enabled:    event.Payload.Enabled,
        Version:    event.Payload.Version,
        // ...
    }
    
    // Process the product
    if err := productRepo.Save(ctx, product); err != nil {
        log.Printf("Failed to save product: %v", err)
    }
}

func handleProductUpdated(ctx context.Context, event *events.ProductUpdatedEvent) {
    log.Printf("Processing update event %s", event.Metadata.EventID)
    
    // Similar processing as ProductCreated
    // ...
}
```

## Benefits of This Structure

### 1. Clear Separation

```go
// Metadata - technical/observability
event.Metadata.EventID        // For deduplication
event.Metadata.TraceID        // For distributed tracing
event.Metadata.CorrelationID  // For request correlation

// Payload - business data only
event.Payload.ProductID
event.Payload.Name
event.Payload.Price
```

### 2. Reusable Metadata

Every event has the same metadata structure:
- No need to remember different field names
- Consistent logging and monitoring
- Easy to add new metadata fields globally

### 3. Type Safety

```go
// Compiler catches errors
event.Metadata.EventType = "ProductCreated"  // ✅ String
event.Payload.Price = 99.99                   // ✅ float64

// This won't compile:
event.Metadata.EventType = 123               // ❌ Type error
```

### 4. Easy Testing

```go
func TestEventSerialization(t *testing.T) {
    event := &events.ProductCreatedEvent{
        Metadata: &events.EventMetadata{
            EventID:   "test-id",
            EventType: "ProductCreated",
            Source:    "test-service",
            Timestamp: 1699123456789,
            SchemaVersion: 1,
        },
        Payload: &events.ProductCreatedPayload{
            ProductID: "product-123",
            Name:      "Test Product",
            Price:     99.99,
            Quantity:  10,
            Enabled:   true,
            Version:   1,
        },
    }
    
    // Serialize and deserialize
    data, err := serializer.Serialize("ProductCreatedEvent", event)
    assert.NoError(t, err)
    
    var decoded events.ProductCreatedEvent
    err = deserializer.Deserialize(data, &decoded)
    assert.NoError(t, err)
    assert.Equal(t, event.Metadata.EventID, decoded.Metadata.EventID)
    assert.Equal(t, event.Payload.Name, decoded.Payload.Name)
}
```

## Common Patterns

### Adding Correlation ID from HTTP Request

```go
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
    // Extract or generate correlation ID
    correlationID := r.Header.Get("X-Correlation-ID")
    if correlationID == "" {
        correlationID = uuid.New().String()
    }
    
    ctx := context.WithValue(r.Context(), "correlation_id", correlationID)
    
    // ... create product ...
    
    // Publish event with correlation ID
    PublishProductCreatedEvent(ctx, product)
}
```

### Event Filtering by Type

```go
// Consumer can filter by event type from metadata
if event.Metadata.EventType == "ProductCreated" {
    // Only process creation events
}
```

### Debugging with Metadata

```go
log.Printf("Event: %s | Type: %s | Source: %s | Trace: %s",
    event.Metadata.EventID,
    event.Metadata.EventType,
    event.Metadata.Source,
    *event.Metadata.TraceID)
```
