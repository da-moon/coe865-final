package codec_test

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/da-moon/coe865-final/pkg/codec"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

type sample struct {
	XMLName    xml.Name `xml:"sample"`
	Validation string   `json:"validation" xml:"validation"`
	Test       string   `json:"test" xml:"test"`
}

func TestCodec(t *testing.T) {

	input := sample{
		Validation: "process",
		Test:       "data",
	}
	t.Run("json", func(t *testing.T) {
		jsonEncoded := `{"validation":"process","test":"data"}`
		writer := new(bytes.Buffer)
		reader := bytes.NewBufferString(jsonEncoded)
		jsonCodec := codec.NewJSONCodec(writer, reader)
		t.Run("encode", func(t *testing.T) {
			err := jsonCodec.Encode(input)
			assert.NoError(t, err)
			actual := strings.TrimSpace(writer.String())
			assert.Equal(t, jsonEncoded, actual)
		})
		t.Run("decode", func(t *testing.T) {
			decodedMap, err := jsonCodec.Decode()
			assert.NoError(t, err)
			assert.NotNil(t, decodedMap)
			var res sample
			err = mapstructure.Decode(decodedMap, &res)
			assert.NoError(t, err)
			assert.Equal(t, input, res)
		})
	})

}
