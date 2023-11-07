package writer

// import (
// 	"encoding/json"
// 	"fmt"
// 	"time"

// 	"github.com/KlyuchnikovV/logist/internal/types"
// )

// // TODO: move to ladder package

// // TODO time format option
// func GELF(host, version string, opts ...interface{}) func(types.Level, string, ...interface{}) []byte {
// 	return func(level types.Level, message string, arguments ...interface{}) []byte {
// 		var result = map[string]interface{}{
// 			"host":      host,
// 			"version":   version,
// 			"timestamp": time.Now().UTC().Format(time.RFC3339),
// 			"level":     level,
// 		}

// 		var (
// 			fieldArgs      = make([]Field, 0, len(arguments)/2)
// 			formattingArgs = make([]interface{}, 0, len(arguments)/2)
// 		)
// 		for i := range arguments {
// 			switch typed := arguments[i].(type) {
// 			case error:
// 				fieldArgs = append(fieldArgs, Error(typed))
// 			case Field:
// 				fieldArgs = append(fieldArgs, typed)
// 			default:
// 				if len(fieldArgs) > 0 {
// 					// TODO: error
// 				} else {
// 					formattingArgs = append(formattingArgs, typed)
// 				}
// 			}
// 		}

// 		result["message"] = fmt.Sprintf(message, formattingArgs...)

// 		bytes, _ := json.Marshal(result)

// 		return bytes
// 	}
// }

// // host (the creator of the message)
// // version
// // timestamp
// // long and short version of the message

// type Field func() (string, interface{})

// func Error(err error) func() (string, interface{}) {
// 	return func() (string, interface{}) {
// 		return "error", err.Error()
// 	}
// }
