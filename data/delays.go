package data

import "verspaetung/utilities"

const DELAYS_DATA_PATH = "data/delays.csv"

/* Caching

The _delaysCache variable contains the loaded data from the delays.csv file.
It has @c delays as its type, which is a map with the following structure:

   line_name -> delay

In words: if a line is delayed, its line_name is assigned the amount of the
delay (in minutes).

Both value types (line_name, delay) are handled as strings.
 */
type delays map[string]string
var _delaysCache delays

func loadDelaysData(path string) (loadedData delays) {
	records := utilities.LoadCSVData(path)
	loadedData = make(delays, len(records))
	for _, record := range records {
		lineName := record[0]
		delay := record[1]
		loadedData[lineName] = delay
	}
	return loadedData
}

func Delays() (delaysCache delays) {
	if len(_delaysCache) != 0 {
		delaysCache = _delaysCache
	} else {
		delaysCache = loadDelaysData(DELAYS_DATA_PATH)
	}
	return
}

/* Retrieves the delay of the provided line name
 */
func (delaysData delays) DelayForLine(lineName string) (delay string, foundLineInCache bool){
	delay, foundLineInCache = delaysData[lineName]
	return
}
