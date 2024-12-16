package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	"database/sql"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func FloatToValue(v float32) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[float32] = pw.FloatWriter
			var wtr func(float32) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}

func NullableFloatToValue(v sql.Null[float32]) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var nw NullableWriter = w.NullableWriter
			var vw Writer[sql.Null[float32]] = nw.FloatWriter
			var wtr func(sql.Null[float32]) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
