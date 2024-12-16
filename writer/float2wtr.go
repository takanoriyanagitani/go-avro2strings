package writer

import (
	"context"
	"io"
	"strconv"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

const FloatWriterFormatDefault byte = 'g'
const FloatWriterPrecDefault int = -1

func FloatNumberWriterNew[T any](
	prec int,
	format byte,
	toDouble func(T) float64,
	bitSize int,
	writer io.Writer,
) Writer[T] {
	var buf []byte
	return func(i T) IO[Void] {
		return func(_ context.Context) (Void, error) {
			buf = buf[:0]
			var dbl float64 = toDouble(i)
			buf = strconv.AppendFloat(
				buf,
				dbl,
				format,
				prec,
				bitSize,
			)
			_, e := writer.Write(buf)
			return Empty, e
		}
	}
}

func FloatNumberWriterNewDefault[T any](
	toDouble func(T) float64,
	bitSize int,
	writer io.Writer,
) Writer[T] {
	return FloatNumberWriterNew(
		FloatWriterPrecDefault,
		FloatWriterFormatDefault,
		toDouble,
		bitSize,
		writer,
	)
}

type FloatToWriterConfig[T any] struct {
	ToDouble func(T) float64
	BitSize  int
}

func FloatNumberWriterNewDefaultFromConfig[T any](
	cfg FloatToWriterConfig[T],
	writer io.Writer,
) Writer[T] {
	return FloatNumberWriterNewDefault[T](
		cfg.ToDouble,
		cfg.BitSize,
		writer,
	)
}

var FloatWriterNewDefault func(
	io.Writer,
) Writer[float32] = Curry[FloatToWriterConfig[float32]](
	FloatNumberWriterNewDefaultFromConfig[float32],
)(
	FloatToWriterConfig[float32]{
		ToDouble: func(i float32) float64 { return float64(i) },
		BitSize:  32,
	},
)

var DoubleWriterNewDefault func(
	io.Writer,
) Writer[float64] = Curry[FloatToWriterConfig[float64]](
	FloatNumberWriterNewDefaultFromConfig[float64],
)(
	FloatToWriterConfig[float64]{
		ToDouble: func(i float64) float64 { return i },
		BitSize:  64,
	},
)
