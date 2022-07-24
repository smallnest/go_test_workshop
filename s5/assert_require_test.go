package s5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssert(t *testing.T) {
	var err error
	assert.NoError(t, err)

	assert.Equal(t, 1, 1)

	assert.Empty(t, len([]string{"abc"}))

	assert.Greater(t, 7, 6)

	t.Log("exit assert")
}

func TestRequire(t *testing.T) {
	var err error
	require.NoError(t, err)

	require.Equal(t, 1, 1)

	require.Empty(t, len([]string{"abc"}))

	require.Greater(t, 7, 6)

	t.Log("exit require")
}
