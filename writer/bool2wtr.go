package writer

import (
	"bytes"
	"context"
	"io"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

type BoolString string

const (
	BoolTrue  BoolString = "true"
	BoolFalse BoolString = "false"
)

func BoolToString(b bool) string {
	switch b {
	case true:
		return string(BoolTrue)
	default:
		return string(BoolFalse)
	}
}

type BoolWriter func(bool) IO[Void]

func BooleanWriterNew(w io.Writer) func(bool) IO[Void] {
	var buf bytes.Buffer
	return func(b bool) IO[Void] {
		return func(_ context.Context) (Void, error) {
			buf.Reset()
			var s string = BoolToString(b)
			_, _ = buf.WriteString(s) // error is always nil or panic
			_, e := w.Write(buf.Bytes())
			return Empty, e
		}
	}
}
