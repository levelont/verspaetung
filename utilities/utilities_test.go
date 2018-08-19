package utilities

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestLoadCSVData(t *testing.T) {
	data := LoadCSVData("test_data.csv")
	expectedData := [][]string {{"id1","val11","va12"}, {"id2","val21","va22"}, {"id3","val31","va32"}}
	for index, line := range data {
		require.Equal(t, expectedData[index], line)
	}
}

func TestRemoveHeaderFromRecords(t *testing.T) {
	data := [][]string{{"Header"}, {"Data"}}
	require.Equal(t, [][]string{{"Data"}}, removeHeaderFromRecords(data))
}
