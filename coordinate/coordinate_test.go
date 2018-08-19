package coordinate

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	coordinate, err := New("1", "2")
	require.Equal(t, coordinate.x, 1)
	require.Equal(t, coordinate.y, 2)
	require.Equal(t, err, nil)
}

func TestConstructorError(t *testing.T) {
	{
		coordinate, err := New("a", "2")
		require.Equal(t, coordinate.x, 0)
		require.Equal(t, coordinate.y, 0)
		require.NotNil(t, err)
		require.EqualError(t, err, "strconv.Atoi: parsing \"a\": invalid syntax")
	}

	{
		coordinate, err := New("1", "b")
		require.Equal(t, coordinate.x, 0)
		require.Equal(t, coordinate.y, 0)
		require.NotNil(t, err)
		require.EqualError(t, err, "strconv.Atoi: parsing \"b\": invalid syntax")
	}
}