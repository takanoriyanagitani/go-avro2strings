package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	"database/sql"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func LongToValue(v int64) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[int64] = pw.LongWriter
			var wtr func(int64) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}

func NullableLongToValue(v sql.Null[int64]) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var nw NullableWriter = w.NullableWriter
			var vw Writer[sql.Null[int64]] = nw.LongWriter
			var wtr func(sql.Null[int64]) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
