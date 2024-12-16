package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	"database/sql"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func StringToValue(v string) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[string] = pw.StringWriter
			var wtr func(string) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}

func NullableStringToValue(v sql.Null[string]) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var nw NullableWriter = w.NullableWriter
			var vw Writer[sql.Null[string]] = nw.StringWriter
			var wtr func(sql.Null[string]) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
