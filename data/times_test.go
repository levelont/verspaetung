package data

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestLoadTimesData(t *testing.T) {
	loadedData := loadTimesData("times.csv")

	expected := make(stops, 0)
	expected["0"] = lines{"0"}
	require.Equal(t, expected, loadedData["10:00:00"])

	expected = make(stops, 0)
	expected["11"] = lines{"2"}
	require.Equal(t, expected, loadedData["10:15:00"])

	expected = make(stops, 0)
	expected["3"] = lines{"1", "2"}
	require.Equal(t, expected, loadedData["10:08:00"])

	expected = make(stops, 0)
	expected["4"] = lines{"0"}
	expected["9"] = lines{"2"}
	require.Equal(t, expected, loadedData["10:09:00"])

	expected = nil
	require.Equal(t, expected, loadedData["NON_EXISTENT"])
}

func TestQueryOccupiedStopsAtTime(t *testing.T) {
	loadedData := loadTimesData("times.csv")

	value, found := loadedData.OccupiedStopsAtTime("10:00:00")
	expected := make(stops, 0)
	expected["0"] = lines{"0"}
	require.True(t, found)
	require.Equal(t, expected, value)

	value, found = loadedData.OccupiedStopsAtTime("10:15:00")
	expected = make(stops, 0)
	expected["11"] = lines{"2"}
	require.True(t, found)
	require.Equal(t, expected, value)

	value, found = loadedData.OccupiedStopsAtTime("10:08:00")
	expected = make(stops, 0)
	expected["3"] = lines{"1", "2"}
	require.True(t, found)
	require.Equal(t, expected, value)

	value, found = loadedData.OccupiedStopsAtTime("10:09:00")
	expected = make(stops, 0)
	expected["4"] = lines{"0"}
	expected["9"] = lines{"2"}
	require.True(t, found)
	require.Equal(t, expected, value)

	value, found = loadedData.OccupiedStopsAtTime("NON_EXISTENT")
	expected = nil
	require.False(t, found)
	require.Equal(t, expected, value)
}