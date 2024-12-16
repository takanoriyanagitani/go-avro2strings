package writer

// This file is generated using any2val.tmpl. NEVER EDIT.

import (
	"database/sql"
	"fmt"
)

func AnyToValueString(a any) Value {
	switch t := a.(type) {
	case string:
		return StringToValue(t)
	case *string:
		return NullableStringToValue(sql.Null[string]{})
	case nil:
		return NullableStringToValue(sql.Null[string]{})
	case sql.Null[string]:
		return NullableStringToValue(t)
	default:
		return InvalidValueFromErr(fmt.Errorf("%w: %v", ErrInvalidValue, t))
	}
}
