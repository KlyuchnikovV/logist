package logist

import (
	"fmt"
	"io"
	"sort"

	"github.com/KlyuchnikovV/logist/internal"
	"github.com/KlyuchnikovV/logist/internal/types"
)

// func WithBufferedLogger(
// 	writer io.StringWriter,
// 	bufferSize int,
// ) Option {
// 	return func(l *Logist) error {
// 		logger, err := internal.NewLogger(l.level, writer)
// 		if err != nil {
// 			return err
// 		}

// 		l.loggers = append(l.loggers, logger)
// 		return nil
// 	}
// }

func sortOptions(options []Option) []Option {
	sort.SliceStable(options, func(i, j int) bool {
		return options[i].priority < options[j].priority
	})

	return options
}

func WithLogger(
	writer io.StringWriter,
	options ...internal.Option,
) Option {
	return newOption(1, func(l *Logist) error {
		logger, err := internal.NewLogger(l.level, writer, l.formatter, options...)
		if err != nil {
			return err
		}

		l.loggers = append(l.loggers, logger)
		return nil
	})
}

func WithLevel(level types.Level) Option {
	var err error

	switch level {
	case types.TraceLevel, types.DebugLevel, types.InfoLevel,
		types.WarningLevel, types.ErrorLevel, types.FatalLevel:
	default:
		err = fmt.Errorf("level is not defined: '%s'", level)
	}

	return newOption(0, func(l *Logist) error {
		if err != nil {
			return err
		}

		l.level = level
		return nil
	})
}

func WithFormat(format internal.Formatter) Option {
	var err error
	if format == nil {
		err = fmt.Errorf("format is nil")
	}

	return newOption(0, func(l *Logist) error {
		if err != nil {
			return err
		}

		l.formatter = format
		return nil
	})
}

func TraceLevel() Option {
	return newOption(0, func(l *Logist) error {
		l.level = types.TraceLevel
		return nil
	})
}

func DebugLevel() Option {
	return newOption(0, func(l *Logist) error {
		l.level = types.DebugLevel
		return nil
	})
}

func InfoLevel() Option {
	return newOption(0, func(l *Logist) error {
		l.level = types.InfoLevel
		return nil
	})
}

func WarningLevel() Option {
	return newOption(0, func(l *Logist) error {
		l.level = types.WarningLevel
		return nil
	})
}

func ErrorLevel() Option {
	return newOption(0, func(l *Logist) error {
		l.level = types.ErrorLevel
		return nil
	})
}

func FatalLevel() Option {
	return newOption(0, func(l *Logist) error {
		l.level = types.FatalLevel
		return nil
	})
}

func PreventFatalFromPanic() Option {
	return newOption(0, func(l *Logist) error {
		l.preventFatalFromPanic = true
		return nil
	})
}
