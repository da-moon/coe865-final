package codec

import (
	"encoding/json"
	"io"
)

type jsonCodec struct {
	enc *json.Encoder
	dec *json.Decoder
}

// NewJSONCodec ...
func NewJSONCodec(w io.Writer, r io.Reader) Codec {

	enc := json.NewEncoder(w)
	dec := json.NewDecoder(r)
	result := &jsonCodec{
		enc: enc,
		dec: dec,
	}
	return result
}

// Encode ...
func (j *jsonCodec) Encode(v interface{}) error {

	return j.enc.Encode(&v)
}

// Decode ...
func (j *jsonCodec) Decode() (interface{}, error) {

	var v interface{}
	err := j.dec.Decode(&v)
	return v, err
}
