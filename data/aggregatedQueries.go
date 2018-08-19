package data

import (
	"errors"
	"verspaetung/coordinate"
)

/* Retrieves the names of the lines located at the provided stop at the provided time
 */
func LinesAtStopAtTime(stop coordinate.Coordinate, timestamp string) (linesAtRequestedStopOnRequestedTime []string, error error) {
	var stopID string
	var linesAtStop stops
	var lines []string
	var found bool

	if stopID, found = Stops().StopID(stop); !found {
		error = errors.New("Coordinate not found in database.")
		return
	}

	if linesAtStop, found = Times().OccupiedStopsAtTime(timestamp); !found {
		error = errors.New("Timestamp not found in database.")
		return
	}

	if lines, found = linesAtStop[stopID]; !found {
		error = errors.New("No lines were found at the provided time in the provided stop.")
		return
	}

	for _, lineID := range lines {
		lineName, _ := Lines().LineName(lineID)
		linesAtRequestedStopOnRequestedTime = append(linesAtRequestedStopOnRequestedTime, lineName)
		error = nil
	}

	return linesAtRequestedStopOnRequestedTime, error
}
