package writer

// This file is generated using any2val.tmpl. NEVER EDIT.

import (
	"database/sql"
	"fmt"
)

func AnyToValueLong(a any) Value {
	switch t := a.(type) {
	case int64:
		return LongToValue(t)
	case *int64:
		return NullableLongToValue(sql.Null[int64]{})
	case nil:
		return NullableLongToValue(sql.Null[int64]{})
	case sql.Null[int64]:
		return NullableLongToValue(t)
	default:
		return InvalidValueFromErr(fmt.Errorf("%w: %v", ErrInvalidValue, t))
	}
}
