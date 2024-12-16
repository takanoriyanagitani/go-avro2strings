package writer

import (
	"bytes"
	"context"
	"io"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func StringWriterNew(w io.Writer) Writer[string] {
	var buf bytes.Buffer
	return func(s string) IO[Void] {
		return func(_ context.Context) (Void, error) {
			buf.Reset()
			_, _ = buf.WriteString(s) // error is always nil or panic
			_, e := w.Write(buf.Bytes())
			return Empty, e
		}
	}
}
