package internal

import (
	"context"
	"io"
	"runtime"
	"time"

	"github.com/KlyuchnikovV/logist/internal/buffer"
	"github.com/KlyuchnikovV/logist/internal/types"
)

type LogFunc func(string, ...types.Field)

func NoLogFunc(string, ...types.Field) {}

type (
	Option    func(*Logger) error
	Hook      func(types.Entry)
	Formatter func(types.Entry) (string, error)
)

type Logger struct {
	level types.Level

	ctx context.Context

	buffer   *buffer.CycleBuffer
	Capacity int

	Trace   LogFunc
	Debug   LogFunc
	Info    LogFunc
	Warning LogFunc
	Error   LogFunc
	Fatal   LogFunc

	// Options
	SkipCaller int
	Now        func() int64

	Hooks     []Hook
	Formatter Formatter
}

func NewLogger(
	level types.Level,
	writer io.StringWriter,
	format Formatter,
	options ...Option,
) (*Logger, error) {
	var logger = &Logger{
		level:     level,
		Capacity:  1,
		Trace:     NoLogFunc,
		Debug:     NoLogFunc,
		Info:      NoLogFunc,
		Warning:   NoLogFunc,
		Error:     NoLogFunc,
		Fatal:     NoLogFunc,
		Now:       time.Now().UTC().Unix,
		Hooks:     make([]Hook, 0),
		Formatter: format,
		ctx:       context.Background(),
	}

	switch level {
	case types.TraceLevel:
		logger.Trace = logger.wrapWriter(types.TraceLevel)
		fallthrough
	case types.DebugLevel:
		logger.Debug = logger.wrapWriter(types.DebugLevel)
		fallthrough
	case types.InfoLevel:
		logger.Info = logger.wrapWriter(types.InfoLevel)
		fallthrough
	case types.WarningLevel:
		logger.Warning = logger.wrapWriter(types.WarningLevel)
		fallthrough
	case types.ErrorLevel:
		logger.Error = logger.wrapWriter(types.ErrorLevel)
		fallthrough
	case types.FatalLevel:
		logger.Fatal = logger.wrapWriter(types.FatalLevel)
	}

	for _, option := range options {
		if err := option(logger); err != nil {
			return nil, err
		}
	}

	var err error
	if logger.buffer, err = buffer.New(writer, logger.Capacity); err != nil {
		return nil, err
	}

	logger.buffer.Start(logger.ctx)

	return logger, nil
}

func (logger *Logger) wrapWriter(level types.Level) func(string, ...types.Field) {
	return func(message string, arguments ...types.Field) {
		logger.write(level, message, arguments...)
	}
}

func (logger *Logger) write(level types.Level, message string, args ...types.Field) {
	var entry = types.Entry{
		types.MessageKey: message,
		types.TimeKey:    logger.Now(),
		types.LevelKey:   level,
	}

	if _, file, line, ok := runtime.Caller(1 + logger.SkipCaller); ok {
		entry[types.FileKey] = file
		entry[types.LineKey] = line
	}

	for _, arg := range args {
		entry[types.Key(arg.Name)] = arg.Value
	}

	for _, hook := range logger.Hooks {
		hook(entry)
	}

	str, err := logger.Formatter(entry)
	if err != nil {
		panic(err)
	}

	logger.buffer.Add(str)
}

func (logger *Logger) Stop() {
	logger.buffer.Stop()
}
