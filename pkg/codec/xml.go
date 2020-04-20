package codec

import (
	"encoding/xml"
	"io"
)

type xmlCodec struct {
	enc *xml.Encoder
	dec *xml.Encoder
}

// NewXMLCodec ...
func NewXMLCodec(w io.Writer, r io.Reader) Codec {
	enc := xml.NewEncoder(w)
	dec := xml.NewDecoder(r)
	result := &xmlCodec{
		enc: enc,
		dec: dec,
	}
	return result
}

// Encode ...
func (j *xmlCodec) Encode(v interface{}) error {
	return j.enc.Encode(&v)
}

// Decode ...
func (j *xmlCodec) Decode() (interface{}, error) {
	var v interface{}
	err := j.dec.Decode(&v)
	return v, err
}

// Format ...
func (j *xmlCodec) Format() string {
	return XML.String()
}
