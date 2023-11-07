package types

type Level string

const (
	TraceLevel   Level = "TRACE"
	DebugLevel   Level = "DEBUG"
	InfoLevel    Level = "INFO"
	WarningLevel Level = "WARN"
	ErrorLevel   Level = "ERROR"
	FatalLevel   Level = "FATAL"
	OffLevel     Level = "OFF"
)

// type Entry struct {
// 	Message string
// 	Level   Level
// 	Agrs    []interface{}
// 	Time    int64

// 	File string
// 	Line int
// }

type Field struct {
	Name  string
	Value interface{}
}

// func NewField(name string, value interface{}) Field {
// 	return Field{
// 		Name:  name,
// 		Value: value,
// 	}
// }

type Key string

const (
	MessageKey Key = "message"
	TimeKey    Key = "timestamp"
	LevelKey   Key = "level"
	FileKey    Key = "file"
	LineKey    Key = "line"
)

type Entry map[Key]interface{}
