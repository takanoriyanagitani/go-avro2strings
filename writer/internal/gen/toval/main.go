package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"text/template"

	. "github.com/takanoriyanagitani/go-avro2string/util"
)

var (
	ErrInvalidArg error = errors.New("invalid arguments")
)

var argLen int = len(os.Args)

var GetArgByIndex func(int) IO[string] = Lift(
	func(ix int) (string, error) {
		if ix < argLen {
			return os.Args[ix], nil
		}
		return "", ErrInvalidArg
	},
)

// e.g, String, Int, Long, ...
var typeHint IO[string] = GetArgByIndex(1)

// e.g, string, int32, int64, ...
var primitive IO[string] = GetArgByIndex(2)

// e.g, string2value.go
var basename IO[string] = Bind(
	typeHint,
	Lift(func(s string) (string, error) {
		var low string = strings.ToLower(s)
		return low + "2" + "value.go", nil
	}),
)

type Config struct {
	TypeHint  string
	Primitive string
	Filename  string
}

var config IO[Config] = Bind(
	All(typeHint, primitive, basename),
	Lift(func(s []string) (Config, error) {
		return Config{
			TypeHint:  s[0],
			Primitive: s[1],
			Filename:  s[2],
		}, nil
	}),
)

var tmpl *template.Template = template.Must(template.ParseFiles(
	"./internal/gen/toval/toval.tmpl",
))

func ExecuteTemplate(
	t *template.Template,
	w io.Writer,
	cfg Config,
) error {
	var bw *bufio.Writer = bufio.NewWriter(w)
	defer bw.Flush()
	e := t.Execute(bw, cfg)
	return errors.Join(e, bw.Flush())
}

func ExecuteToFile(
	t *template.Template,
	file *os.File,
	cfg Config,
) error {
	defer file.Close()
	return ExecuteTemplate(t, file, cfg)
}

func ExecuteToFilename(
	t *template.Template,
	cfg Config,
) error {
	file, e := os.Create(cfg.Filename)
	if nil != e {
		return e
	}
	return ExecuteToFile(t, file, cfg)
}

var config2file IO[Void] = Bind(
	config,
	Lift(func(cfg Config) (Void, error) {
		return Empty, ExecuteToFilename(tmpl, cfg)
	}),
)

var sub IO[Void] = func(ctx context.Context) (Void, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return config2file(ctx)
}

func main() {
	_, e := sub(context.Background())
	if nil != e {
		panic(e)
	}
}
