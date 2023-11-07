package logist

type (
	Logger interface {
		Write(string, string, ...interface{})
		Trace(string, ...interface{})
		Debug(string, ...interface{})
		Info(string, ...interface{})
		Warn(string, ...interface{})
		Error(string, ...interface{})
	}
)

type Option struct {
	priority int
	process  func(*Logist) error
}

func newOption(priority int, process func(*Logist) error) Option {
	return Option{
		priority: priority,
		process:  process,
	}
}
