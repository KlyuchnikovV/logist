package options

import (
	"fmt"

	"github.com/KlyuchnikovV/logist/internal"
	"github.com/KlyuchnikovV/logist/internal/types"
)

func WithBufferSize(capacity int) internal.Option {
	var err error
	if capacity == 0 {
		err = fmt.Errorf("now function can't be nil")
	}

	return func(logger *internal.Logger) error {
		if err != nil {
			return err
		}

		logger.Capacity = capacity
		return nil
	}
}

func WithCaller(skip int) internal.Option {
	return func(logger *internal.Logger) error {
		if skip < 0 {
			return fmt.Errorf("skip caller can't be negative")
		}

		logger.SkipCaller = skip
		return nil
	}
}

func WithTime(now func() int64) internal.Option {
	return func(logger *internal.Logger) error {
		if now == nil {
			return fmt.Errorf("now function can't be nil")
		}

		logger.Now = now
		return nil
	}
}

func WithFormatter(formatter internal.Formatter) internal.Option {
	return func(logger *internal.Logger) error {
		if formatter == nil {
			return fmt.Errorf("no formatter found")
		}

		logger.Formatter = formatter

		return nil
	}
}

func WithKey(previous types.Key, new string) internal.Option {
	var err error
	switch {
	case len(previous) == 0 && len(new) == 0:
		err = fmt.Errorf("no keys specified")
	case len(previous) == 0:
		err = fmt.Errorf("no previous key specified")
	case len(new) == 0:
		err = fmt.Errorf("no new key specified")
	}

	return func(logger *internal.Logger) error {
		if err != nil {
			return err
		}

		logger.Hooks = append(logger.Hooks, keyHook(previous, new))
		return nil
	}
}

func WithMessageKey(key string) internal.Option {
	return WithKey(types.MessageKey, key)
}

func WithTimeKey(key string) internal.Option {
	return WithKey(types.TimeKey, key)
}

func WithLevelKey(key string) internal.Option {
	return WithKey(types.LevelKey, key)
}

func WithFileKey(key string) internal.Option {
	return WithKey(types.FileKey, key)
}

func WithLineKey(key string) internal.Option {
	return WithKey(types.LineKey, key)
}

func keyHook(previous types.Key, new string) internal.Hook {
	return func(entry types.Entry) {
		entry[types.Key(new)] = entry[previous]
		delete(entry, previous)
	}
}
