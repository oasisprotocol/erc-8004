package logging

import (
	"fmt"
	"io"
	"os"

	"github.com/oasisprotocol/oasis-core/go/common/logging"

	"github.com/oasisprotocol/erc-8004/validator-agent/config"
)

// InitLogging initializes logging based on the provided configuration.
func InitLogging(cfg *config.LogConfig) error {
	var w io.Writer = os.Stdout
	format := logging.FmtLogfmt
	level := logging.LevelDebug
	if cfg != nil {
		var err error
		if w, err = getLoggingStream(cfg); err != nil {
			return fmt.Errorf("opening log file: %w", err)
		}
		if err := format.Set(cfg.Format); err != nil {
			return err
		}
		if err := level.Set(cfg.Level); err != nil {
			return err
		}
	}

	return logging.Initialize(w, format, level, nil)
}

func getLoggingStream(cfg *config.LogConfig) (io.Writer, error) {
	if cfg == nil || cfg.File == "" {
		return os.Stdout, nil
	}
	w, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return nil, err
	}
	return w, nil
}
