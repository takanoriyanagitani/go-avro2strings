package writer

import (
	"context"
	"io"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func BytesWriterNew(w io.Writer) Writer[[]byte] {
	return func(b []byte) IO[Void] {
		return func(_ context.Context) (Void, error) {
			_, e := w.Write(b)
			return Empty, e
		}
	}
}
