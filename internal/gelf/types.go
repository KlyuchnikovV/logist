package gelf

// type FieldName string

// const (
// 	VersionFieldName      FieldName = "version"
// 	HostFieldName         FieldName = "host"
// 	ShortMessageFieldName FieldName = "short_message"
// 	FullMessageFieldName  FieldName = "full_message"
// 	TimestampFieldName    FieldName = "timestamp"
// 	LevelFieldName        FieldName = "level"
// 	FacilityFieldName     FieldName = "facility"
// 	LineFieldName         FieldName = "line"
// 	FileFieldName         FieldName = "file"
// )

// version string (UTF-8)
// GELF spec version – “1.1”; MUST be set by client library.

// host string (UTF-8)
// the name of the host, source or application that sent this message; MUST be set by client library.

// short_message string (UTF-8)
// a short descriptive message; MUST be set by client library.

// full_message string (UTF-8)
// a long message that can i.e. contain a backtrace; optional.

// timestamp number
// Seconds since UNIX epoch with optional decimal places for milliseconds; SHOULD be set by client library. Will be set to the current timestamp (now) by the server if absent.

// level number
// the level equal to the standard syslog levels; optional, default is 1 (ALERT).

// facility string (UTF-8)
// optional, deprecated. Send as additional field instead.

// line number
// the line in a file that caused the error (decimal); optional, deprecated. Send as additional field instead.

// file string (UTF-8)
// the file (with path if you want) that caused the error (string); optional, deprecated. Send as additional field instead.

// _[additional field] string (UTF-8) or number
// every field you send and prefix with an underscore (_) will be treated as an additional field. Allowed characters in field names are any word character (letter, number, underscore), dashes and dots. The verifying regular expression is: ^[\w\.\-]*$. Libraries SHOULD not allow to send id as additional field (_id). Graylog server nodes omit this field automatically.

// func (gelf *GELF) getTemplateMessage() map[FieldName]interface{} {
// 	return map[FieldName]interface{}{
// 		HostFieldName:     gelf.host,
// 		FileFieldName:     gelf.file,
// 		VersionFieldName:  gelf.version,
// 		FacilityFieldName: gelf.facility,

// 		LineFieldName:         nil,
// 		LevelFieldName:        nil,
// 		TimestampFieldName:    nil,
// 		FullMessageFieldName:  nil,
// 		ShortMessageFieldName: nil,
// 	}
// }
