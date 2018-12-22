package pc

import (
	"errors"
	"io"

	"github.com/ugorji/go/codec"
)

// EncodeListToCbor encodes this List as CBOR to the given writer.
func (l List) EncodeListToCbor(output io.Writer) error {
	var enc *codec.Encoder = codec.NewEncoder(output, new(codec.CborHandle))
	return enc.Encode(l)
}

// DecodeListFromCbor tries to create a List based on a given reader. The
// filename-field will be still empty after this function.
func DecodeListFromCbor(input io.Reader) (l List, err error) {
	var dec *codec.Decoder = codec.NewDecoder(input, new(codec.CborHandle))
	err = dec.Decode(&l)

	if err != nil {
		for _, e := range l.Entries {
			if e.AbsValue() > valueAbsMax {
				err = errors.New("Entry's absolute value is greater than max")
				break
			}
		}
	}

	return
}
