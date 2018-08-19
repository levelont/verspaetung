package data

import "verspaetung/utilities"

const TIMES_DATA_PATH = "data/times.csv"

/* Caching

The _timesCache variable contains the loaded data from the times.csv file.
It has @c times as its type, which is a map with the structure:

   time -> [ stop_id -> [ line_id ] ]

In words: a time key can hold many stop values. Each stop on its turn is also
a key and can hold many line values.
In real world terms: in a transport system with many stops, several lines
can be at the same time on a single stop.

All three value types stored in the times.csv file (time, stop_id, line_id) are
handled as strings.
 */
type lines []string
type stops map[string]lines
type times map[string]stops
var _timesCache times

func loadTimesData(path string) (loadedData times){
	records := utilities.LoadCSVData(path)
	loadedData = make(times, len(records))
	for _, record := range records {
		lineId := record[0]
		stopId := record[1]
		timestamp := record[2]
		if stopsForCurrentTime, found := loadedData[timestamp]; found {
			if linesForCurrentStop, found := stopsForCurrentTime[stopId]; found {
				stopsForCurrentTime[stopId] = append(linesForCurrentStop, lineId)
			} else {
				tmp := append(make(lines, 0), lineId) //initialize new list of lines for this stopId
				stopsForCurrentTime[stopId] = tmp
			}
		} else {
			loadedData[timestamp] = make(stops, 0) // initialize new list of stops for this timestamp
			tmp := append(make(lines, 0), lineId) //initialize new list of lines for this stopId
			loadedData[timestamp][stopId] = tmp
		}
	}
	return loadedData
}

func Times() (timesCache times) {
	if len(_timesCache) != 0 {
		timesCache = _timesCache
	} else {
		timesCache = loadTimesData(TIMES_DATA_PATH)
	}
	return
}

/* Retrieves the list of stops occupied by lines at a given time
 */
func (timesData times) OccupiedStopsAtTime(timestamp string) (occupiedStops stops, timestampFoundInCache bool) {
	occupiedStops, timestampFoundInCache = timesData[timestamp]
	return
}

