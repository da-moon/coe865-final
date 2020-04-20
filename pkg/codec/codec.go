package codec

// Codec interface must be implemented
// by any encoder/decoder
type Codec interface {
	Encode(v interface{}) error
	Decode() (interface{}, error)
	Format() string
}
