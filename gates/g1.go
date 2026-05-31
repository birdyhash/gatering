package gates

import "github.com/birdyhash/gatering"

/*
g1 is a mirror gate.

the gate reverses the byte order of the payload.

example:

hello

becomes

olleh
*/
var G1 = gatering.Gate{
	Name: "mirror",

	Run: func(data []byte) []byte {

		out := make([]byte, len(data))
		copy(out, data)

		for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}

		return out
	},
}
