package avro2str

type PrimitiveType string

const (
	PrimitiveUnknown PrimitiveType = "UNKNOWN"
	PrimitiveString  PrimitiveType = "string"
	PrimitiveBytes   PrimitiveType = "bytes"
	PrimitiveInt     PrimitiveType = "int"
	PrimitiveLong    PrimitiveType = "long"
	PrimitiveFloat   PrimitiveType = "float"
	PrimitiveDouble  PrimitiveType = "double"
	PrimitiveBoolean PrimitiveType = "boolean"
	PrimitiveNull    PrimitiveType = "null"
)

const (
	NullableString  PrimitiveType = "null-string"
	NullableInt     PrimitiveType = "null-int"
	NullableLong    PrimitiveType = "null-long"
	NullableFloat   PrimitiveType = "null-float"
	NullableDouble  PrimitiveType = "null-double"
	NullableBoolean PrimitiveType = "null-boolean"
)

const BlobSizeMaxDefault int = 1048576

type InputConfig struct {
	BlobSizeMax int
}

var InputConfigDefault InputConfig = InputConfig{
	BlobSizeMax: BlobSizeMaxDefault,
}
