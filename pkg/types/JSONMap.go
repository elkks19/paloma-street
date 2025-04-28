package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONMap map[string]any

// Para base de datos (guardar)
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Para base de datos (leer)
func (j *JSONMap) Scan(value any) error {
	if value == nil {
		*j = JSONMap{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to convert database value to []byte")
	}
	return json.Unmarshal(bytes, j)
}

// Para API (serializar JSON)
func (j JSONMap) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(map[string]any(j))
}

// Para API (deserializar JSON)
func (j *JSONMap) UnmarshalJSON(data []byte) error {
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*j = m
	return nil
}
