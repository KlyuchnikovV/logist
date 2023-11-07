package logist_test

import (
	"testing"

	"github.com/KlyuchnikovV/logist"
	"github.com/KlyuchnikovV/logist/internal/types"
	"github.com/KlyuchnikovV/logist/options"
	"github.com/stretchr/testify/assert"
)

type stringWriterMock []string

func (mock *stringWriterMock) WriteString(s string) (int, error) {
	*mock = append(*mock, s)
	return len(s), nil
}

func (mock *stringWriterMock) Clear() {
	*mock = []string{}
}

func TestTrace(t *testing.T) {
	var writer = make(stringWriterMock, 0)

	logger, err := logist.New(
		logist.TraceLevel(),
		logist.PreventFatalFromPanic(),
		logist.WithLogger(&writer),
		logist.WithFormat(options.FormatString(' ', types.MessageKey)),
	)
	assert.NoError(t, err)
	defer logger.Stop()

	var cases = []struct {
		name   string
		action func(*logist.Logist, string, ...types.Field)
		params string
		result []string
	}{
		{
			name:   "Try to trace",
			action: (*logist.Logist).Trace,
			params: "trace",
			result: []string{"trace\n"},
		},
		{
			name:   "Try to debug",
			action: (*logist.Logist).Debug,
			params: "debug",
			result: []string{"debug\n"},
		},
		{
			name:   "Try to info",
			action: (*logist.Logist).Info,
			params: "info",
			result: []string{"info\n"},
		},
		{
			name:   "Try to warning",
			action: (*logist.Logist).Warning,
			params: "warning",
			result: []string{"warning\n"},
		},
		{
			name:   "Try to error",
			action: (*logist.Logist).Error,
			params: "error",
			result: []string{"error\n"},
		},
		{
			name:   "Try to fatal",
			action: (*logist.Logist).Fatal,
			params: "fatal",
			result: []string{"fatal\n"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.action(logger, c.params)

			assert.Len(t, writer, len(c.result))
			for i, item := range c.result {
				assert.Equal(t, item, writer[i])
			}

			writer.Clear()
		})
	}
}

func TestWarning(t *testing.T) {
	var writer = make(stringWriterMock, 0)

	logger, err := logist.New(
		logist.WarningLevel(),
		logist.PreventFatalFromPanic(),
		logist.WithLogger(&writer),
		logist.WithFormat(options.FormatString(' ', types.MessageKey)),
	)
	assert.NoError(t, err)
	defer logger.Stop()

	var cases = []struct {
		name   string
		action func(*logist.Logist, string, ...types.Field)
		params string
		result []string
	}{
		{
			name:   "Try to trace",
			action: (*logist.Logist).Trace,
			params: "trace",
			result: []string{},
		},
		{
			name:   "Try to debug",
			action: (*logist.Logist).Debug,
			params: "debug",
			result: []string{},
		},
		{
			name:   "Try to info",
			action: (*logist.Logist).Info,
			params: "info",
			result: []string{},
		},
		{
			name:   "Try to warning",
			action: (*logist.Logist).Warning,
			params: "warning",
			result: []string{"warning"},
		},
		{
			name:   "Try to error",
			action: (*logist.Logist).Error,
			params: "error",
			result: []string{"error"},
		},
		{
			name:   "Try to fatal",
			action: (*logist.Logist).Fatal,
			params: "fatal",
			result: []string{"fatal"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.action(logger, c.params)

			assert.Len(t, writer, len(c.result))
			for i, item := range c.result {
				assert.Equal(t, item, writer[i])
			}

			writer.Clear()
		})
	}
}

func TestFatal(t *testing.T) {
	var (
		writer  = make(stringWriterMock, 0)
		options = []logist.Option{
			logist.FatalLevel(),
			logist.WithLogger(&writer),
			logist.WithFormat(options.FormatString(' ', types.MessageKey)),
		}
	)

	var cases = []struct {
		name      string
		params    []logist.Option
		panicking bool
	}{
		{
			name:      "Try to fatal",
			params:    []logist.Option{logist.PreventFatalFromPanic()},
			panicking: false,
		},
		{
			name:      "Try to fatal, not prevented",
			params:    nil,
			panicking: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			logger, err := logist.New(append(options, c.params...)...)
			assert.NoError(t, err)
			defer logger.Stop()

			if c.panicking {
				assert.PanicsWithValue(t, "fatal", func() {
					logger.Fatal("fatal")
				})
			} else {
				assert.NotPanics(t, func() {
					logger.Fatal("fatal")
				})
			}
		})
	}
}
