package writer

import (
	"database/sql"
	"errors"
	"io"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

var (
	ErrInvalidWriter error = errors.New("unimplemented")

	ErrInvalidValue error = errors.New("invalid value")
)

type WriteAny func(any) IO[Void]

type Writer[T any] func(T) IO[Void]

func InvalidWriterNew[T any]() Writer[T] {
	return func(_ T) IO[Void] {
		return Err[Void](ErrInvalidWriter)
	}
}

type PrimitiveWriter struct {
	StringWriter  Writer[string]
	BytesWriter   Writer[[]byte]
	IntWriter     Writer[int32]
	LongWriter    Writer[int64]
	FloatWriter   Writer[float32]
	DoubleWriter  Writer[float64]
	BooleanWriter Writer[bool]
	NullWriter    Writer[struct{}]
}

func (p PrimitiveWriter) ToNullableWriter(w io.Writer) NullableWriter {
	return NullableWriter{
		StringWriter:  NullableToWriterNew(w, p.StringWriter),
		BytesWriter:   NullableToWriterNew(w, p.BytesWriter),
		IntWriter:     NullableToWriterNew(w, p.IntWriter),
		LongWriter:    NullableToWriterNew(w, p.LongWriter),
		FloatWriter:   NullableToWriterNew(w, p.FloatWriter),
		DoubleWriter:  NullableToWriterNew(w, p.DoubleWriter),
		BooleanWriter: NullableToWriterNew(w, p.BooleanWriter),
	}
}

func WriterToPrimitiveWriterDefault(wtr io.Writer) PrimitiveWriter {
	return PrimitiveWriter{
		StringWriter:  StringWriterNew(wtr),
		BytesWriter:   BytesWriterNew(wtr),
		IntWriter:     IntWriterNewDefault(wtr),
		LongWriter:    LongWriterNewDefault(wtr),
		FloatWriter:   FloatWriterNewDefault(wtr),
		DoubleWriter:  DoubleWriterNewDefault(wtr),
		BooleanWriter: BooleanWriterNew(wtr),
		NullWriter:    NullWriterNew(wtr),
	}
}

var PrimitiveWriterDefault PrimitiveWriter = PrimitiveWriter{
	StringWriter:  InvalidWriterNew[string](),
	BytesWriter:   InvalidWriterNew[[]byte](),
	IntWriter:     InvalidWriterNew[int32](),
	LongWriter:    InvalidWriterNew[int64](),
	FloatWriter:   InvalidWriterNew[float32](),
	DoubleWriter:  InvalidWriterNew[float64](),
	BooleanWriter: InvalidWriterNew[bool](),
	NullWriter:    InvalidWriterNew[struct{}](),
}

type NullableWriter struct {
	StringWriter  Writer[sql.Null[string]]
	BytesWriter   Writer[sql.Null[[]byte]]
	IntWriter     Writer[sql.Null[int32]]
	LongWriter    Writer[sql.Null[int64]]
	FloatWriter   Writer[sql.Null[float32]]
	DoubleWriter  Writer[sql.Null[float64]]
	BooleanWriter Writer[sql.Null[bool]]
	NullWriter    Writer[sql.Null[struct{}]]
}

var NullableWriterDefault NullableWriter = NullableWriter{
	StringWriter:  InvalidWriterNew[sql.Null[string]](),
	BytesWriter:   InvalidWriterNew[sql.Null[[]byte]](),
	IntWriter:     InvalidWriterNew[sql.Null[int32]](),
	LongWriter:    InvalidWriterNew[sql.Null[int64]](),
	FloatWriter:   InvalidWriterNew[sql.Null[float32]](),
	DoubleWriter:  InvalidWriterNew[sql.Null[float64]](),
	BooleanWriter: InvalidWriterNew[sql.Null[bool]](),
	NullWriter:    InvalidWriterNew[sql.Null[struct{}]](),
}

//go:generate go run ./internal/gen/toval/main.go String string
//go:generate go run ./internal/gen/toval/main.go Bytes []byte
//go:generate go run ./internal/gen/toval/main.go Int int32
//go:generate go run ./internal/gen/toval/main.go Long int64
//go:generate go run ./internal/gen/toval/main.go Float float32
//go:generate go run ./internal/gen/toval/main.go Double float64
//go:generate go run ./internal/gen/toval/main.go Boolean bool
//go:generate go run ./internal/gen/toval/main.go Null struct{}
//go:generate gofmt -s -w .
type ValueWriter struct {
	PrimitiveWriter
	NullableWriter
}

func WriterToValueWriterDefault(w io.Writer) ValueWriter {
	var pw PrimitiveWriter = WriterToPrimitiveWriterDefault(w)
	var nw NullableWriter = pw.ToNullableWriter(w)
	return ValueWriter{
		PrimitiveWriter: pw,
		NullableWriter:  nw,
	}
}

var ValueWriterDefault ValueWriter = ValueWriter{
	PrimitiveWriter: PrimitiveWriterDefault,
	NullableWriter:  NullableWriterDefault,
}

type WriterTo func(ValueWriter) IO[Void]

//go:generate go run ./internal/gen/any2val/main.go String string
//go:generate go run ./internal/gen/any2val/main.go Bytes []byte
//go:generate go run ./internal/gen/any2val/main.go Int int32
//go:generate go run ./internal/gen/any2val/main.go Long int64
//go:generate go run ./internal/gen/any2val/main.go Float float32
//go:generate go run ./internal/gen/any2val/main.go Double float64
//go:generate go run ./internal/gen/any2val/main.go Boolean bool
//go:generate gofmt -s -w .
type Value = WriterTo

func InvalidWriterTo(_ ValueWriter) IO[Void] {
	return Err[Void](ErrInvalidValue)
}

func InvalidValueFromErr(err error) Value {
	return func(_ ValueWriter) IO[Void] {
		return Err[Void](err)
	}
}

var InvalidValue Value = InvalidWriterTo
