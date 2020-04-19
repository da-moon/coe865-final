package identity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultIdentity(t *testing.T) {

	id, err := Default()
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.Equal(t, defaultSize/8, id.private.Size())
	assert.Equal(t, defaultSize/8, id.private.PublicKey.Size())
	shabytes, err := id.Sha256()
	assert.NoError(t, err)
	assert.NotEmpty(t, shabytes)
	str, err := id.Sha256String()
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
}
func TestNewIdentity(t *testing.T) {
	size := 16
	id, err := New(size)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.Equal(t, size/8, id.private.Size())
	assert.Equal(t, size/8, id.private.PublicKey.Size())
	shabytes, err := id.Sha256()
	assert.NoError(t, err)
	assert.NotEmpty(t, shabytes)
	str, err := id.Sha256String()
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
}
