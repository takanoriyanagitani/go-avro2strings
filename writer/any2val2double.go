package writer

// This file is generated using any2val.tmpl. NEVER EDIT.

import (
	"database/sql"
	"fmt"
)

func AnyToValueDouble(a any) Value {
	switch t := a.(type) {
	case float64:
		return DoubleToValue(t)
	case *float64:
		return NullableDoubleToValue(sql.Null[float64]{})
	case nil:
		return NullableDoubleToValue(sql.Null[float64]{})
	case sql.Null[float64]:
		return NullableDoubleToValue(t)
	default:
		return InvalidValueFromErr(fmt.Errorf("%w: %v", ErrInvalidValue, t))
	}
}
