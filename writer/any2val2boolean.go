package writer

// This file is generated using any2val.tmpl. NEVER EDIT.

import (
	"database/sql"
	"fmt"
)

func AnyToValueBoolean(a any) Value {
	switch t := a.(type) {
	case bool:
		return BooleanToValue(t)
	case *bool:
		return NullableBooleanToValue(sql.Null[bool]{})
	case nil:
		return NullableBooleanToValue(sql.Null[bool]{})
	case sql.Null[bool]:
		return NullableBooleanToValue(t)
	default:
		return InvalidValueFromErr(fmt.Errorf("%w: %v", ErrInvalidValue, t))
	}
}
