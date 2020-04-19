package node

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	// setup listener
	l, err := net.Listen("tcp", ":3000")
	assert.NoError(t, err)
	assert.NotNil(t, l)
	defer l.Close()
	t.Run("json marshaller", func(t *testing.T) {
	})
	t.Run("json", func(t *testing.T) {
	})
	t.Run("encode message", func(t *testing.T) {

	})
	t.Run("decode message", func(t *testing.T) {

	})

}

// func TestJSONUtil_EncodeJSON(t *testing.T) {
// 	input := map[string]interface{}{
// 		"validation": "process",
// 		"test":       "data",
// 	}

// 	actualBytes, err := jsonutil.EncodeJSON(input)
// 	if err != nil {
// 		t.Fatalf("failed to encode JSON: %v", err)
// 	}

// 	actual := strings.TrimSpace(string(actualBytes))
// 	expected := `{"validation":"process","test":"data"}`

// 	if actual != expected {
// 		t.Fatalf("bad: encoded JSON: expected:%s\nactual:%s\n", expected, string(actualBytes))
// 	}
// }

// func TestJSONUtil_DecodeJSON(t *testing.T) {
// 	input := `{"test":"data","validation":"process"}`

// 	var actual map[string]interface{}

// 	err := jsonutil.DecodeJSON([]byte(input), &actual)
// 	if err != nil {
// 	}

// 	expected := map[string]interface{}{
// 		"test":       "data",
// 		"validation": "process",
// 	}
// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Fatalf("bad: expected:%#v\nactual:%#v", expected, actual)
// 	}
// }

func TestConn(t *testing.T) {
	message := "Hi there!\n"

	go func() {
		conn, err := net.Dial("tcp", ":3000")
		assert.NoError(t, err)
		assert.NotNil(t, conn)

		defer conn.Close()
		_, err = fmt.Fprintf(conn, message)
		assert.NoError(t, err)
	}()

	l, err := net.Listen("tcp", ":3000")
	assert.NoError(t, err)
	assert.NotNil(t, l)

	defer l.Close()
	for {
		conn, err := l.Accept()
		assert.NoError(t, err)
		assert.NotNil(t, conn)

		defer conn.Close()
		buf, err := ioutil.ReadAll(conn)
		assert.NoError(t, err)
		assert.NotEmpty(t, buf)
		msg := string(buf[:])
		assert.Equal(t, message, msg)
		return
	}
}
