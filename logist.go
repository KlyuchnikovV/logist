package logist

import (
	"fmt"
	"os"

	"github.com/KlyuchnikovV/logist/internal"
	"github.com/KlyuchnikovV/logist/internal/types"
	"github.com/KlyuchnikovV/logist/options"
)

// TODO: e2e logging (context logging)
// TODO: save all logs using cycle buffering
// TODO: deferred/recovered write (with option to set defer message)
// TODO: option for unsupported loggers
// TODO: different log formats
// TODO: different formats for different loggers
// TODO: replace standart loggers
// TODO: compile fields on add

// CONCEPT: easy, fast and customizable
// CONCEPT: try to reduce repeatable logs

type Logist struct {
	level types.Level

	loggers []*internal.Logger // TODO: with config

	preventFatalFromPanic bool
	formatter             internal.Formatter
}

func Simple() *Logist {
	logger, err := internal.NewLogger(
		types.TraceLevel,
		os.Stdout,
		options.FormatJSON,
	)
	if err != nil {
		panic(fmt.Errorf("can't create simple logger: %w", err))
	}

	return &Logist{
		level:     types.TraceLevel,
		loggers:   []*internal.Logger{logger},
		formatter: options.FormatJSON,
	}
}

func New(options ...Option) (*Logist, error) {
	var logist = &Logist{
		level:   types.TraceLevel,
		loggers: make([]*internal.Logger, 0),
	}

	for _, option := range sortOptions(options) {
		if err := option.process(logist); err != nil {
			return nil, err
		}
	}

	if len(logist.loggers) == 0 {
		logger, err := internal.NewLogger(logist.level, os.Stdout, logist.formatter)
		if err != nil {
			return nil, err
		}

		logist.loggers = []*internal.Logger{logger}
	}

	return logist, nil
}

func (logist *Logist) Trace(message string, fields ...types.Field) {
	for i := range logist.loggers {
		logist.loggers[i].Trace(message, fields...)
	}
}

func (logist *Logist) Debug(message string, fields ...types.Field) {
	for i := range logist.loggers {
		logist.loggers[i].Debug(message, fields...)
	}
}

func (logist *Logist) Info(message string, fields ...types.Field) {
	for i := range logist.loggers {
		logist.loggers[i].Info(message, fields...)
	}
}

func (logist *Logist) Warning(message string, fields ...types.Field) {
	for i := range logist.loggers {
		logist.loggers[i].Warning(message, fields...)
	}
}

func (logist *Logist) Error(message string, fields ...types.Field) {
	for i := range logist.loggers {
		logist.loggers[i].Error(message, fields...)
	}
}

func (logist *Logist) Fatal(message string, fields ...types.Field) {
	for i := range logist.loggers {
		logist.loggers[i].Fatal(message, fields...)
	}

	if !logist.preventFatalFromPanic {
		panic(message) // TODO: need default message format for this case
	}
}

func (logist *Logist) Stop() {
	for i := range logist.loggers {
		logist.loggers[i].Stop()
	}
}
