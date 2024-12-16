package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	"database/sql"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func BytesToValue(v []byte) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[[]byte] = pw.BytesWriter
			var wtr func([]byte) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}

func NullableBytesToValue(v sql.Null[[]byte]) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var nw NullableWriter = w.NullableWriter
			var vw Writer[sql.Null[[]byte]] = nw.BytesWriter
			var wtr func(sql.Null[[]byte]) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
