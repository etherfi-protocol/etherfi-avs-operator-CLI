package logging

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger = zerolog.New(os.Stdout)
}

func Info(module string, msgs ...string) {
	msg := strings.Join(msgs, " ")
	logger.Info().Str("module", module).Msg(msg)
}

func Error(module string, err error) {
	logger.Error().Str("module", module).Err(err)
}
