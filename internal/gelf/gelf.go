package gelf

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/KlyuchnikovV/logist/internal"
// )

// var levelsToCode = map[internal.Level]int{
// 	internal.DebugLevel:   9,
// 	internal.TraceLevel:   8,
// 	internal.InfoLevel:    7,
// 	internal.WarningLevel: 5,
// 	internal.ErrorLevel:   3,
// 	internal.FatalLevel:   2,
// }

// type GELF struct {
// 	host     string
// 	file     string
// 	version  string
// 	facility string

// 	fieldNaming map[FieldName]string
// }

// func New(
// 	host,
// 	file,
// 	facility,
// 	version string,
// 	opts ...Option,
// ) *GELF {
// 	return &GELF{
// 		host:     host,
// 		file:     file,
// 		version:  version,
// 		facility: facility,
// 	}
// }

// func (gelf *GELF) Format(
// 	level internal.Level,
// 	message,
// 	line string, // TODO: in fields?
// 	timestamp string, // TODO: in fields?
// 	fields ...internal.Field,
// ) []byte {
// 	var result = map[FieldName]interface{}{
// 		HostFieldName:     gelf.host,
// 		FileFieldName:     gelf.file,
// 		VersionFieldName:  gelf.version,
// 		FacilityFieldName: gelf.facility,

// 		LineFieldName:         line,
// 		LevelFieldName:        levelsToCode[level],
// 		TimestampFieldName:    timestamp,
// 		FullMessageFieldName:  message, // TODO: how to form it?
// 		ShortMessageFieldName: message,
// 	}

// 	// var result = map[string]interface{}{
// 	// 	"host":      gelf.host,
// 	// 	"version":   gelf.version,
// 	// 	"timestamp": time.Now().UTC().Format(time.RFC3339),
// 	// 	"level":     levelsToCode[level],
// 	// }

// 	// var (
// 	// 	fieldArgs      = make([]internal.Field, 0, len(arguments)/2)
// 	// 	formattingArgs = make([]interface{}, 0, len(arguments)/2)
// 	// )
// 	// for i := range arguments {
// 	// 	switch typed := arguments[i].(type) {
// 	// 	case error:
// 	// 		fieldArgs = append(fieldArgs, internal.Error(typed))
// 	// 	case internal.Field:
// 	// 		fieldArgs = append(fieldArgs, typed)
// 	// 	default:
// 	// 		if len(fieldArgs) > 0 {
// 	// 			// TODO: error
// 	// 		} else {
// 	// 			formattingArgs = append(formattingArgs, typed)
// 	// 		}
// 	// 	}
// 	// }

// 	result["message"] = fmt.Sprintf(message, formattingArgs...)

// 	bytes, _ := json.Marshal(result)

// 	return bytes

// }
