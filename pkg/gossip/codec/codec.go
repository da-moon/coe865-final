package codec
// Codec interface must be implemented
// by any encoder/decoder
// it Encodes data and then flushes it
// to a stream (io.Writer)
// or decodes data as it is reading
// it from a stream (io.Reader)
type Codec interface {
	Encode(v interface{}) error
	Decode() (interface{}, error)
}
