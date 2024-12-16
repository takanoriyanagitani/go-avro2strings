package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	"database/sql"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func BooleanToValue(v bool) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[bool] = pw.BooleanWriter
			var wtr func(bool) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}

func NullableBooleanToValue(v sql.Null[bool]) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var nw NullableWriter = w.NullableWriter
			var vw Writer[sql.Null[bool]] = nw.BooleanWriter
			var wtr func(sql.Null[bool]) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
