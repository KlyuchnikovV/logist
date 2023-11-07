package writer

// import (
// 	"io"

// 	"github.com/KlyuchnikovV/logist/internal/types"
// )

// type Buffer struct {
// 	writer *io.Writer

// 	hooks     []Hook
// 	formatter Formatter // TODO: Default formatter
// }

// func New(writer io.Writer, options ...Option) *Buffer {
// 	return &Buffer{
// 		writer:    &writer,
// 		formatter: FormatJSON,
// 	}
// }

// func (buffer *Buffer) Write(entry types.Entry) error {
// 	// TODO: make write
// 	return nil
// }

// func (buffer *Buffer) WriteError(err error) {
// 	// TODO: make write
// }

// func (buffer *Buffer) Sync() error {
// 	// TODO: make sync
// 	return nil
// }
