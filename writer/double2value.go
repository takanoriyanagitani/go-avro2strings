package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	"database/sql"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func DoubleToValue(v float64) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[float64] = pw.DoubleWriter
			var wtr func(float64) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}

func NullableDoubleToValue(v sql.Null[float64]) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var nw NullableWriter = w.NullableWriter
			var vw Writer[sql.Null[float64]] = nw.DoubleWriter
			var wtr func(sql.Null[float64]) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
