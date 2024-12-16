package writer

import (
	"context"
	"database/sql"
	"io"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func NullWriterNew(w io.Writer) func(struct{}) IO[Void] {
	return func(_ struct{}) IO[Void] {
		return func(_ context.Context) (Void, error) {
			_, e := w.Write([]byte("null"))
			return Empty, e
		}
	}
}

func NullableToWriterNew[T any](
	w io.Writer,
	wtr Writer[T],
) Writer[sql.Null[T]] {
	return func(n sql.Null[T]) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			if !n.Valid {
				_, e := w.Write([]byte("None"))
				return Empty, e
			}
			_, e := w.Write([]byte("Some("))
			if nil != e {
				return Empty, e
			}

			var val T = n.V
			_, e = wtr(val)(ctx)
			if nil != e {
				return Empty, e
			}

			_, e = w.Write([]byte(")"))

			return Empty, e
		}
	}
}
