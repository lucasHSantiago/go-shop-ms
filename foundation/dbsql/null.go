package dbsql

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToText(s *string) pgtype.Text {
	text := pgtype.Text{
		Valid: s != nil,
	}

	if text.Valid {
		text.String = *s
	}

	return text
}

func TimeToTimestamptz(t *time.Time) pgtype.Timestamptz {
	time := pgtype.Timestamptz{
		Valid: t != nil,
	}

	if time.Valid {
		time.Time = *t
	}

	return time
}

func BoolToBool(b *bool) pgtype.Bool {
	nb := pgtype.Bool{
		Valid: b != nil,
	}

	if nb.Valid {
		nb.Bool = *b
	}

	return nb
}

func UUIDToUUID(u *uuid.UUID) pgtype.UUID {
	uuid := pgtype.UUID{
		Valid: u != nil,
	}

	if uuid.Valid {
		uuid.Bytes = *u
	}

	return uuid
}

func Float64ToFloat8(f *float64) pgtype.Float8 {
	float := pgtype.Float8{
		Valid: f != nil,
	}

	if float.Valid {
		float.Float64 = *f
	}

	return float
}
