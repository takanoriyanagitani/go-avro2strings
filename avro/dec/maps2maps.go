package dec

import (
	"fmt"
	"iter"
	"log"

	as "github.com/takanoriyanagitani/go-avro2string"

	sw "github.com/takanoriyanagitani/go-avro2string/writer"
)

type AnyToValue func(any) sw.Value

type AnyToValueResolver func(as.PrimitiveType) AnyToValue

func MapsToValueMaps(
	resolver AnyToValueResolver,
	name2type func(string) as.PrimitiveType,
	m iter.Seq2[map[string]any, error],
) iter.Seq2[map[string]sw.Value, error] {
	return func(yield func(map[string]sw.Value, error) bool) {
		var mapd map[string]sw.Value = map[string]sw.Value{}
		for original, e := range m {
			clear(mapd)

			if nil == e {
				for key, val := range original {
					var typ as.PrimitiveType = name2type(key)
					var a2v AnyToValue = resolver(typ)
					var converted sw.Value = a2v(val)
					mapd[key] = converted
				}
			}

			if !yield(mapd, e) {
				return
			}
		}
	}
}

func MapsToValueMapsDefault(
	name2type func(string) as.PrimitiveType,
	m iter.Seq2[map[string]any, error],
) iter.Seq2[map[string]sw.Value, error] {
	return MapsToValueMaps(AnyToValResolverDefault, name2type, m)
}

var AnyToInvalid AnyToValue = func(a any) sw.Value {
	return sw.InvalidValueFromErr(fmt.Errorf("%w: %v", sw.ErrInvalidValue, a))
}

var type2any2valMap map[as.PrimitiveType]AnyToValue = map[as.
	PrimitiveType]AnyToValue{

	as.PrimitiveString:  sw.AnyToValueString,
	as.PrimitiveBytes:   sw.AnyToValueBytes,
	as.PrimitiveInt:     sw.AnyToValueInt,
	as.PrimitiveLong:    sw.AnyToValueLong,
	as.PrimitiveFloat:   sw.AnyToValueFloat,
	as.PrimitiveDouble:  sw.AnyToValueDouble,
	as.PrimitiveBoolean: sw.AnyToValueBoolean,

	as.NullableString:  sw.AnyToValueString,
	as.NullableInt:     sw.AnyToValueInt,
	as.NullableLong:    sw.AnyToValueLong,
	as.NullableFloat:   sw.AnyToValueFloat,
	as.NullableDouble:  sw.AnyToValueDouble,
	as.NullableBoolean: sw.AnyToValueBoolean,
}

type AnyToAny func(any) any

var NopConverter AnyToAny = func(original any) any { return original }

type AnyToAnyResolver func(as.PrimitiveType) AnyToAny

func ConverterFromMap[K comparable, V any](
	alt V,
	m map[K]V,
) func(K) V {
	return func(key K) V {
		val, found := m[key]
		switch found {
		case true:
			return val
		default:
			log.Printf("not found. key=%v, m=%v\n", key, m)
			return alt
		}
	}
}

func DefaultToMapToConverter[V any, K comparable](
	alt V,
) func(map[K]V) func(K) V {
	return func(m map[K]V) func(K) V {
		return ConverterFromMap(alt, m)
	}
}

var AnyToValResolverFromMap func(map[as.PrimitiveType]AnyToValue) func(
	as.PrimitiveType,
) AnyToValue = DefaultToMapToConverter[AnyToValue, as.PrimitiveType](
	AnyToInvalid,
)

var AnyToValResolverDefault AnyToValueResolver = AnyToValResolverFromMap(
	type2any2valMap,
)

func MapsToMaps(
	resolver AnyToAnyResolver,
	name2type func(string) as.PrimitiveType,
	m iter.Seq2[map[string]any, error],
) iter.Seq2[map[string]any, error] {
	return func(yield func(map[string]any, error) bool) {
		var mapd map[string]any = map[string]any{}
		for original, e := range m {
			clear(mapd)

			if nil == e {
				for key, val := range original {
					var typ as.PrimitiveType = name2type(key)
					var a2a AnyToAny = resolver(typ)
					var converted any = a2a(val)
					mapd[key] = converted
				}
			}

			if !yield(mapd, e) {
				return
			}
		}
	}
}

var NameToTypeFromMap func(map[string]as.PrimitiveType) func(
	string,
) as.PrimitiveType = DefaultToMapToConverter[as.PrimitiveType, string](
	as.PrimitiveUnknown,
)
