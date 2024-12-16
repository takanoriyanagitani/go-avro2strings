package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	"database/sql"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func IntToValue(v int32) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[int32] = pw.IntWriter
			var wtr func(int32) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}

func NullableIntToValue(v sql.Null[int32]) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var nw NullableWriter = w.NullableWriter
			var vw Writer[sql.Null[int32]] = nw.IntWriter
			var wtr func(sql.Null[int32]) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
