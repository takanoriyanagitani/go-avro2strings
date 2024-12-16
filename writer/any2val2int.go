package writer

// This file is generated using any2val.tmpl. NEVER EDIT.

import (
	"database/sql"
	"fmt"
)

func AnyToValueInt(a any) Value {
	switch t := a.(type) {
	case int32:
		return IntToValue(t)
	case *int32:
		return NullableIntToValue(sql.Null[int32]{})
	case nil:
		return NullableIntToValue(sql.Null[int32]{})
	case sql.Null[int32]:
		return NullableIntToValue(t)
	default:
		return InvalidValueFromErr(fmt.Errorf("%w: %v", ErrInvalidValue, t))
	}
}
