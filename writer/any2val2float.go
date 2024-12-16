package writer

// This file is generated using any2val.tmpl. NEVER EDIT.

import (
	"database/sql"
	"fmt"
)

func AnyToValueFloat(a any) Value {
	switch t := a.(type) {
	case float32:
		return FloatToValue(t)
	case *float32:
		return NullableFloatToValue(sql.Null[float32]{})
	case nil:
		return NullableFloatToValue(sql.Null[float32]{})
	case sql.Null[float32]:
		return NullableFloatToValue(t)
	default:
		return InvalidValueFromErr(fmt.Errorf("%w: %v", ErrInvalidValue, t))
	}
}
