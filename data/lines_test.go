package data

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestLoadLinesData(t *testing.T) {
	loadedData := loadLinesData("lines.csv")
	require.Equal(t, "M4", loadedData["0"])
	require.Equal(t, "S75", loadedData["2"])
	require.Equal(t, "", loadedData["NON_EXISTENT"])
}

func TestQueryLineName(t *testing.T) {
	loadedData := loadLinesData("lines.csv")
	value, found := loadedData.LineName("0")
	require.True(t, found)
	require.Equal(t, "M4", value)

	value, found = loadedData.LineName("2")
	require.True(t, found)
	require.Equal(t, "S75", value)

	value, found = loadedData.LineName("NON_EXISTENT")
	require.False(t, found)
	require.Equal(t, "", value)
}