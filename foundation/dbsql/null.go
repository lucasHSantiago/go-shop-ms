package dbsql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func StringToText(s *string) sql.NullString {
	text := sql.NullString{
		Valid: s != nil,
	}

	if text.Valid {
		text.String = *s
	}

	return text
}

func TimeToTimestamptz(t *time.Time) sql.NullTime {
	time := sql.NullTime{
		Valid: t != nil,
	}

	if time.Valid {
		time.Time = *t
	}

	return time
}

func BoolToBool(b *bool) sql.NullBool {
	bool := sql.NullBool{
		Valid: b != nil,
	}

	if bool.Valid {
		bool.Bool = *b
	}

	return bool
}

func UUIDToUUID(u *uuid.UUID) NullUUID {
	uuid := NullUUID{
		Valid: u != nil,
	}

	if uuid.Valid {
		uuid.UUID = *u
	}

	return uuid
}

func Float64ToFloat8(f *float64) sql.NullFloat64 {
	float := sql.NullFloat64{
		Valid: f != nil,
	}

	if float.Valid {
		float.Float64 = *f
	}

	return float
}
