package main

import (
	"context"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"strings"

	as "github.com/takanoriyanagitani/go-avro2string"
	. "github.com/takanoriyanagitani/go-avro2string/util"

	ss "github.com/takanoriyanagitani/go-avro2string/select"
	sw "github.com/takanoriyanagitani/go-avro2string/writer"

	ad "github.com/takanoriyanagitani/go-avro2string/avro/dec"
	dh "github.com/takanoriyanagitani/go-avro2string/avro/dec/hamba"
)

var EnvValByKey func(string) IO[string] = Lift(
	func(key string) (string, error) {
		val, found := os.LookupEnv(key)
		switch found {
		case true:
			return val, nil
		default:
			return "", fmt.Errorf("env key %s missing", key)
		}
	},
)

var selectedColumn IO[string] = EnvValByKey("ENV_COLUMN_NAME")

var schemaFilename IO[string] = EnvValByKey("ENV_SCHEMA_FILENAME")

func FilenameToStringLimited(limit int64) func(string) IO[string] {
	return Lift(func(filename string) (string, error) {
		f, e := os.Open(filename)
		if nil != e {
			return "", e
		}
		defer f.Close()

		limited := &io.LimitedReader{
			R: f,
			N: limit,
		}

		var buf strings.Builder

		_, e = io.Copy(&buf, limited)
		return buf.String(), e
	})
}

const SchemaSizeMaxDefault int64 = 1048576

var schemaContent IO[string] = Bind(
	schemaFilename,
	FilenameToStringLimited(SchemaSizeMaxDefault),
)

var name2type IO[func(string) as.PrimitiveType] = Bind(
	schemaContent,
	Lift(func(s string) (func(string) as.PrimitiveType, error) {
		return dh.SchemaToNameToType(s), nil
	}),
)

var stdin2maps IO[iter.Seq2[map[string]any, error]] = dh.StdinToMapsDefault

var valueMaps IO[iter.Seq2[map[string]sw.Value, error]] = Bind(
	name2type,
	func(
		n2t func(string) as.PrimitiveType,
	) IO[iter.Seq2[map[string]sw.Value, error]] {
		return Bind(
			stdin2maps,
			Lift(func(
				m iter.Seq2[map[string]any, error],
			) (iter.Seq2[map[string]sw.Value, error], error) {
				return ad.MapsToValueMapsDefault(
					n2t,
					m,
				), nil
			}),
		)
	},
)

var selected IO[iter.Seq2[sw.Value, error]] = Bind(
	selectedColumn,
	func(colname string) IO[iter.Seq2[sw.Value, error]] {
		return Bind(
			valueMaps,
			Lift(func(
				m iter.Seq2[map[string]sw.Value, error],
			) (iter.Seq2[sw.Value, error], error) {
				return ss.MapsToSelected(
					m,
					colname,
				), nil
			}),
		)
	},
)

func selected2stdout(m iter.Seq2[sw.Value, error]) IO[Void] {
	return func(ctx context.Context) (Void, error) {
		var buf strings.Builder
		var vw sw.ValueWriter = sw.WriterToValueWriterDefault(&buf)
		for selected, e := range m {
			buf.Reset()
			select {
			case <-ctx.Done():
				return Empty, ctx.Err()
			default:
			}

			if nil != e {
				return Empty, e
			}

			_, e = selected(vw)(ctx)
			if nil != e {
				return Empty, e
			}

			fmt.Printf("%s\n", buf.String())
		}
		return Empty, nil
	}
}

var stdin2maps2selected2stdout IO[Void] = Bind(
	selected,
	selected2stdout,
)

var sub IO[Void] = func(ctx context.Context) (Void, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return stdin2maps2selected2stdout(ctx)
}

func main() {
	_, e := sub(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
