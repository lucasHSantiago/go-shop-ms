package logger

import (
	"log"

	zlog "github.com/rs/zerolog/log"
)

type ZeroLogAdapter struct{}

func (z ZeroLogAdapter) Write(p []byte) (n int, err error) {
	zlog.Error().Msg(string(p))
	return len(p), nil
}

func NewStdLogger() *log.Logger {
	return log.New(ZeroLogAdapter{}, "", 0)
}
