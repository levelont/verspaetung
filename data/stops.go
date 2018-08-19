package data

import (
	"verspaetung/utilities"
	"verspaetung/coordinate"
)

const STOPS_DATA_PATH = "data/stops.csv"

/* Caching

The _stopsCache variable contains the loaded data from the stops.csv file.
It has @c coordinates as its type, which is a map with the following structure:

   coordinate -> stop_id

In words: each pair of cartesian coordinates x,y get a stop_id assigned as their identifier.
In real world terms: stop identifiers are used to uniquely identify geographic locations.

The custom type @c Coordinate from the coordinates package is used to represent
the cartesian coordinates of a stop. The stop_id field itself is handled as a string.
 */
type coordinates map[coordinate.Coordinate]string
var _stopsCache coordinates

func loadStopsData(path string) (loadedData coordinates){
	records := utilities.LoadCSVData(path)
	loadedData = make(coordinates, len(records))
	for _, record := range records {
		stopID := record[0]
		stopLocation, _ := coordinate.New(record[1], record[2])
		loadedData[stopLocation] = stopID
	}
	return
}

func Stops() (stopsCache coordinates) {
	if len(_stopsCache) != 0 {
		stopsCache = _stopsCache
	} else {
		stopsCache = loadStopsData(STOPS_DATA_PATH)
	}
	return
}

/* Retrieves the stop_id of the provided coordinate
 */
func (coordinatesData coordinates) StopID(c coordinate.Coordinate) (stopID string, coordinateFoundInCache bool) {
	stopID, coordinateFoundInCache = coordinatesData[c]
	return
}
