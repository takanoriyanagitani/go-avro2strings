package dec

import (
	ha "github.com/hamba/avro/v2"

	as "github.com/takanoriyanagitani/go-avro2string"

	ad "github.com/takanoriyanagitani/go-avro2string/avro/dec"
)

var type2prim map[ha.Type]as.PrimitiveType = map[ha.Type]as.PrimitiveType{
	ha.String:  as.NullableString,
	ha.Int:     as.NullableInt,
	ha.Long:    as.NullableLong,
	ha.Float:   as.NullableFloat,
	ha.Double:  as.NullableDouble,
	ha.Boolean: as.NullableBoolean,
}

var type2primNN map[ha.Type]as.PrimitiveType = map[ha.Type]as.PrimitiveType{
	ha.String:  as.PrimitiveString,
	ha.Bytes:   as.PrimitiveBytes,
	ha.Int:     as.PrimitiveInt,
	ha.Long:    as.PrimitiveLong,
	ha.Float:   as.PrimitiveFloat,
	ha.Double:  as.PrimitiveDouble,
	ha.Boolean: as.PrimitiveBoolean,
	ha.Null:    as.PrimitiveNull,
}

var convertType func(ha.Type) as.PrimitiveType = ad.
	DefaultToMapToConverter[as.PrimitiveType, ha.Type](
	as.PrimitiveUnknown,
)(type2prim)

var convertTypeNN func(ha.Type) as.PrimitiveType = ad.
	DefaultToMapToConverter[as.PrimitiveType, ha.Type](
	as.PrimitiveUnknown,
)(type2primNN)

func ConvertType(t ha.Type) as.PrimitiveType { return convertType(t) }

func FieldsToNameToTypeMap(fields []*ha.Field) map[string]as.PrimitiveType {
	ret := map[string]as.PrimitiveType{}
	for _, field := range fields {
		var name string = field.Name()
		var s ha.Schema = field.Type()

		var u *ha.UnionSchema
		switch t := s.(type) {
		case *ha.PrimitiveSchema:
			var htyp ha.Type = t.Type()
			var mapdType as.PrimitiveType = convertTypeNN(htyp)
			ret[name] = mapdType
			continue
		case *ha.UnionSchema:
			u = t
		default:
			continue
		}

		var types []ha.Schema = u.Types()
		for _, typ := range types {
			var p *ha.PrimitiveSchema
			switch t := typ.(type) {
			case *ha.PrimitiveSchema:
				p = t
			default:
				continue
			}

			var ptyp ha.Type = p.Type()
			var converted as.PrimitiveType = ConvertType(ptyp)
			ret[name] = converted
		}
	}
	return ret
}

func FieldsToNameToType(fields []*ha.Field) func(string) as.PrimitiveType {
	var m map[string]as.PrimitiveType = FieldsToNameToTypeMap(fields)
	return ad.NameToTypeFromMap(m)
}

func RecordToNameToType(r *ha.RecordSchema) func(string) as.PrimitiveType {
	return FieldsToNameToType(r.Fields())
}

func SchemaToNameToTypeHamba(s ha.Schema) func(string) as.PrimitiveType {
	switch t := s.(type) {
	case *ha.RecordSchema:
		return RecordToNameToType(t)
	default:
		return func(_ string) as.PrimitiveType { return as.PrimitiveUnknown }
	}
}

func SchemaToNameToType(schema string) func(string) as.PrimitiveType {
	parsed, e := ha.Parse(schema)
	if nil != e {
		return func(_ string) as.PrimitiveType { return as.PrimitiveUnknown }
	}
	return SchemaToNameToTypeHamba(parsed)
}
