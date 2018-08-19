package data

import (
	"testing"
	"github.com/stretchr/testify/require"
	"verspaetung/coordinate"
)

func TestLoadStopsData(t *testing.T) {
	loadedData := loadStopsData("stops.csv")

	location, _ := coordinate.New("1", "1")
	require.Equal(t, "0", loadedData[location])

	location, _ = coordinate.New("5", "7")
	require.Equal(t, "11", loadedData[location])

	location, _ = coordinate.New("1", "NON_EXISTENT")
	require.Equal(t, "", loadedData[location])

	location, _ = coordinate.New("NON_EXISTENT", "1")
	require.Equal(t, "", loadedData[location])
}

func TestQueryStopID(t *testing.T) {
	loadedData := loadStopsData("stops.csv")

	location, _ := coordinate.New("1", "1")
	value, found := loadedData.StopID(location)
	require.True(t, found)
	require.Equal(t, "0", value)

	location, _ = coordinate.New("5", "7")
	value, found = loadedData.StopID(location)
	require.True(t, found)
	require.Equal(t, "11", value)

	location, _ = coordinate.New("1", "NON_EXISTENT")
	value, found = loadedData.StopID(location)
	require.False(t, found)
	require.Equal(t, "", value)

	location, _ = coordinate.New("NON_EXISTENT", "1")
	value, found = loadedData.StopID(location)
	require.False(t, found)
	require.Equal(t, "", value)
}