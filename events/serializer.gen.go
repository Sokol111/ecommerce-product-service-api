package events

import (
	"fmt"
	"sync"

	"github.com/hamba/avro/v2"
)

var (
	// Pre-parsed schemas (initialized once at startup)
	compiledSchemas     = make(map[string]avro.Schema)
	compiledSchemasOnce sync.Once
	schemaInitError     error
)

// initSchemas parses all embedded schemas once
func initSchemas() error {
	compiledSchemasOnce.Do(func() {
		schemas := map[string][]byte{
			"EventMetadata": EventMetadataSchema,
			"ProductCreatedEvent": ProductCreatedSchema,
			"ProductUpdatedEvent": ProductUpdatedSchema,
		}

		for name, schemaJSON := range schemas {
			schema, err := avro.Parse(string(schemaJSON))
			if err != nil {
				schemaInitError = fmt.Errorf("failed to parse schema %s: %w", name, err)
				return
			}
			compiledSchemas[name] = schema
		}
	})
	return schemaInitError
}

// Marshal serializes an event using its pre-compiled schema
func Marshal(schemaName string, v interface{}) ([]byte, error) {
	if err := initSchemas(); err != nil {
		return nil, err
	}
	
	schema, ok := compiledSchemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema %s not found", schemaName)
	}
	
	return avro.Marshal(schema, v)
}

// Unmarshal deserializes an event using its pre-compiled schema
func Unmarshal(schemaName string, data []byte, v interface{}) error {
	if err := initSchemas(); err != nil {
		return err
	}
	
	schema, ok := compiledSchemas[schemaName]
	if !ok {
		return fmt.Errorf("schema %s not found", schemaName)
	}
	
	return avro.Unmarshal(schema, data, v)
}

// Typed serialization helpers

// MarshalProductCreatedEvent serializes ProductCreatedEvent to Avro bytes
func MarshalProductCreatedEvent(event *ProductCreatedEvent) ([]byte, error) {
	return Marshal("ProductCreatedEvent", event)
}

// UnmarshalProductCreatedEvent deserializes ProductCreatedEvent from Avro bytes
func UnmarshalProductCreatedEvent(data []byte) (*ProductCreatedEvent, error) {
	var event ProductCreatedEvent
	if err := Unmarshal("ProductCreatedEvent", data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// MarshalProductUpdatedEvent serializes ProductUpdatedEvent to Avro bytes
func MarshalProductUpdatedEvent(event *ProductUpdatedEvent) ([]byte, error) {
	return Marshal("ProductUpdatedEvent", event)
}

// UnmarshalProductUpdatedEvent deserializes ProductUpdatedEvent from Avro bytes
func UnmarshalProductUpdatedEvent(data []byte) (*ProductUpdatedEvent, error) {
	var event ProductUpdatedEvent
	if err := Unmarshal("ProductUpdatedEvent", data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

