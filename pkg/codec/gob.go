package codec

import (
	"encoding/gob"
	"io"
)

type gobCodec struct {
	enc *gob.Encoder
	dec *gob.Encoder
}

// NewGOBCodec ...
// make sure to register structs if
// using this codec, eg:
// func init(){
// 		gob.Register(PayloadStruct{})
// }
func NewGOBCodec(w io.Writer, r io.Reader) Codec {
	enc := gob.NewEncoder(w)
	dec := gob.NewDecoder(r)
	result := &gobCodec{
		enc: enc,
		dec: dec,
	}
	return result
}

// Encode ...
func (j *gobCodec) Encode(v interface{}) error {
	return j.enc.Encode(&v)
}

// Decode ...
func (j *gobCodec) Decode() (interface{}, error) {
	var v interface{}
	err := j.dec.Decode(&v)
	return v, err
}

// Format ...
func (j *gobCodec) Format() string {
	return GOB.String()
}
