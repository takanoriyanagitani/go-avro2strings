package writer

// This file is generated using any2val.tmpl. NEVER EDIT.

import (
	"database/sql"
	"fmt"
)

func AnyToValueBytes(a any) Value {
	switch t := a.(type) {
	case []byte:
		return BytesToValue(t)
	case *[]byte:
		return NullableBytesToValue(sql.Null[[]byte]{})
	case nil:
		return NullableBytesToValue(sql.Null[[]byte]{})
	case sql.Null[[]byte]:
		return NullableBytesToValue(t)
	default:
		return InvalidValueFromErr(fmt.Errorf("%w: %v", ErrInvalidValue, t))
	}
}
