package logist

// func (log *Ladder) Trace(format string, args ...interface{}) {
// 	log.Writef(Trace, format, args...)
// }

// func (log *Ladder) Info(format string, args ...interface{}) {
// 	log.Writef(Info, format, args...)
// }

// func (log *Ladder) Warning(format string, args ...interface{}) {
// 	log.Writef(Warning, format, args...)
// }

// func (log *Ladder) Error(format string, args ...interface{}) {
// 	log.Writef(Error, format, args...)
// }

// func (logist *Logist) Info(message string, args ...interface{}) {
// 	logist.write(message, args...)
// }

// func (logist *Logist) write(level, message string, args ...interface{}) {
// 	var entry = entry{
// 		message: message,
// 		agrs:    args,
// 		time:    time.Now().UTC().Unix(), // TODO: parametrisied
// 	}

// 	if _, file, line, ok := runtime.Caller(1); ok {
// 		// TODO: parameter + i
// 		entry.file = file
// 		entry.line = line
// 	}

// 	// TODO: write to a buffer
// }

// type entry struct {
// 	file string
// 	line int

// 	message string
// 	level level
// 	agrs    []interface{}
// 	time    int64
// }
