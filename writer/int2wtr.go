package writer

import (
	"context"
	"io"
	"strconv"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

const IntWriterBaseDefault int = 10

func IntegerWriterNew[T any](
	toLong func(T) int64,
	base int,
	w io.Writer,
) Writer[T] {
	var buf []byte
	return func(i T) IO[Void] {
		return func(_ context.Context) (Void, error) {
			buf = buf[:0]
			var lng int64 = toLong(i)
			buf = strconv.AppendInt(buf, lng, base)
			_, e := w.Write(buf)
			return Empty, e
		}
	}
}

func IntWriterNew(base int, w io.Writer) Writer[int32] {
	return IntegerWriterNew(
		func(i int32) int64 { return int64(i) },
		base,
		w,
	)
}

var IntWriterNewDefault func(io.Writer) Writer[int32] = Curry(
	IntWriterNew,
)(IntWriterBaseDefault)

func LongWriterNew(base int, w io.Writer) Writer[int64] {
	return IntegerWriterNew(
		func(i int64) int64 { return i },
		base,
		w,
	)
}

var LongWriterNewDefault func(io.Writer) Writer[int64] = Curry(
	LongWriterNew,
)(IntWriterBaseDefault)
