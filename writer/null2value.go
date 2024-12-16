package writer

// This file is generated using toval.tmpl. NEVER EDIT.

import (
	"context"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

func NullToValue(v struct{}) Value {
	return func(w ValueWriter) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var pw PrimitiveWriter = w.PrimitiveWriter
			var vw Writer[struct{}] = pw.NullWriter
			var wtr func(struct{}) IO[Void] = vw
			return wtr(v)(ctx)
		}
	}
}
