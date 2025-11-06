package events

import (
	"fmt"

	"github.com/hamba/avro/v2"
)

// AvroSerializer handles Avro serialization/deserialization
type AvroSerializer struct {
	schemas map[string]avro.Schema
}

// NewAvroSerializer creates a new Avro serializer
func NewAvroSerializer() (*AvroSerializer, error) {
	return &AvroSerializer{
		schemas: make(map[string]avro.Schema),
	}, nil
}

// RegisterSchema registers an Avro schema
func (s *AvroSerializer) RegisterSchema(name string, schemaJSON string) error {
	schema, err := avro.Parse(schemaJSON)
	if err != nil {
		return fmt.Errorf("failed to parse schema %s: %w", name, err)
	}
	s.schemas[name] = schema
	return nil
}

// Serialize encodes data using Avro
func (s *AvroSerializer) Serialize(schemaName string, data interface{}) ([]byte, error) {
	schema, ok := s.schemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema %s not registered", schemaName)
	}
	
	return avro.Marshal(schema, data)
}

// Deserialize decodes Avro data
func (s *AvroSerializer) Deserialize(schemaName string, data []byte, v interface{}) error {
	schema, ok := s.schemas[schemaName]
	if !ok {
		return fmt.Errorf("schema %s not registered", schemaName)
	}
	
	return avro.Unmarshal(schema, data, v)
}
