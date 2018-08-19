package data

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestLoadDelaysData(t *testing.T) {
	loadedData := loadDelaysData("delays.csv")
	require.Equal(t, "1", loadedData["M4"])
	require.Equal(t, "10", loadedData["S75"])
	require.Equal(t, "", loadedData["NON_EXISTENT"])
}

func TestQueryDelayForLine(t *testing.T) {
	loadedData := loadDelaysData("delays.csv")
	value, found := loadedData.DelayForLine("M4")
	require.True(t, found)
	require.Equal(t, "1", value)

	value, found = loadedData.DelayForLine("S75")
	require.True(t, found)
	require.Equal(t, "10", value)

	value, found = loadedData.DelayForLine("NON_EXISTENT")
	require.False(t, found)
	require.Equal(t, "", value)
}