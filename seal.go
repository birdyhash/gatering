package gatering

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
a seal represents the final result produced by
the chamber.

the hash field contains the generated digest.

the route field contains the sequence of gates
that participated in the transformation process.

the size field contains the payload size at the
moment the seal was generated.
*/
type Seal struct {
	Hash  string
	Size  int
	Route []string
}

/*
createseal generates a sha256 digest from the
provided payload and associates metadata about
the transformation route.
*/
func CreateSeal(data []byte, route []string) Seal {

	sum := sha256.Sum256(data)

	return Seal{
		Hash:  hex.EncodeToString(sum[:]),
		Size:  len(data),
		Route: route,
	}
}
