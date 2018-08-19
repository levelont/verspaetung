package data

import "verspaetung/utilities"

const LINES_DATA_PATH = "data/lines.csv"

/* Caching

The _lineNamesCache variable contains the loaded data from the lines.csv file.
It has @c lineNames as its type, which is a map with the following structure:

   line_id -> line_name

In words: each line_id gets a name assigned.
Both value types (line_id, line_name) are handled as strings.
 */
type lineNames map[string]string
var _lineNamesCache lineNames

func loadLinesData(path string) (loadedData lineNames) {
	records := utilities.LoadCSVData(path)
	loadedData = make(lineNames, len(records))
	for _, record := range records {
		lineID := record[0]
		lineName := record[1]
		loadedData[lineID] = lineName
	}
	return loadedData
}

func Lines() (linesCache lineNames) {
	if len(_lineNamesCache) != 0 {
		linesCache = _lineNamesCache
	} else {
		linesCache = loadLinesData(LINES_DATA_PATH)
	}
	return linesCache
}

/* Retrieves the line name of the provided line_id
 */
func (linesData lineNames) LineName(lineID string) (lineName string, foundLineInCache bool){
	lineName, foundLineInCache = linesData[lineID]
	return
}
