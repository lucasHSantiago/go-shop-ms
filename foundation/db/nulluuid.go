package db

import (
	"database/sql/driver"
	"errors"

	"github.com/google/uuid"
)

type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

// Scan implementa a interface sql.Scanner
func (n *NullUUID) Scan(value interface{}) error {
	if value == nil {
		n.Valid = false
		return nil
	}

	switch v := value.(type) {
	case []byte:
		u, err := uuid.ParseBytes(v)
		if err != nil {
			return err
		}
		n.UUID = u
	case string:
		u, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		n.UUID = u
	default:
		return errors.New("unsupported type for UUID")
	}

	n.Valid = true
	return nil
}

// Value implementa a interface driver.Valuer
func (n NullUUID) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UUID.String(), nil
}
