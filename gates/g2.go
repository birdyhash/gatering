package gates

import "github.com/birdyhash/gatering"

/*
g2 is a prefix gate.

the gate injects a static marker at the
beginning of the payload.

this gate exists mainly to demonstrate
how multiple transformations can be chained.
*/
var G2 = gatering.Gate{
	Name: "threshold",

	Run: func(data []byte) []byte {

		prefix := []byte("gate:")

		out := make([]byte, 0, len(prefix)+len(data))

		out = append(out, prefix...)
		out = append(out, data...)

		return out
	},
}
