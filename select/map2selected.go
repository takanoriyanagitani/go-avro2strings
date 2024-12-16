package sel

import (
	"iter"

	wtr "github.com/takanoriyanagitani/go-avro2string/writer"
)

func MapsToSelected(
	m iter.Seq2[map[string]wtr.Value, error],
	name string,
) iter.Seq2[wtr.Value, error] {
	return func(yield func(wtr.Value, error) bool) {
		for original, e := range m {
			var val wtr.Value = wtr.InvalidValue
			if nil == e {
				src, found := original[name]
				if found {
					val = src
				}
			}

			if !yield(val, e) {
				return
			}
		}
	}
}
