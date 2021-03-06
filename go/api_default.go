/*
 * Verspaetung API
 *
 * Verspaetung public transport information
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
	"fmt"
	"verspaetung/data"
	"verspaetung/coordinate"
	"strings"
	"regexp"
)

func parseLocation(w http.ResponseWriter, r *http.Request) (location coordinate.Coordinate, success bool){
	if len(r.FormValue("x")) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "X-coordinate missing.")
		return
	}

	if len(r.FormValue("y")) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Y-coordinate missing.")
		return
	}

	parsedCoordinate, error := coordinate.New(r.FormValue("x"), r.FormValue("y"))
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing input parameters: %v\n", error)
		return
	}

	location = parsedCoordinate
	success = true
	return
}

func parseTimestamp(w http.ResponseWriter, r *http.Request) (timestamp string, success bool){
	if len(r.FormValue("timestamp")) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Timestamp parameter missing.")
		return
	}

	parsedTimestamp := r.FormValue("timestamp")
	re := regexp.MustCompile("((\\d){2}:){2}(\\d){2}")
	if ! re.MatchString(parsedTimestamp) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error parsing timestamp: it doesn't match the required format 'HH:mm:ss'.")
		return
	}
	timestamp = parsedTimestamp
	success = true
	return
}

func LinesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	location, locationOK := parseLocation(w, r)
	timestamp, timestampOK := parseTimestamp(w, r)
	if ! locationOK || ! timestampOK{
		return
	}

	w.WriteHeader(http.StatusOK)

	lines, error := data.LinesAtStopAtTime(location, timestamp)
	if error != nil {
		fmt.Fprintln(w, error)
		return
	}

	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}

func LinesNameGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	lineName := strings.Split(r.URL.Path, "/")[2]
	_, foundLineNameInDelayCache := data.Delays().DelayForLine(lineName)
	fmt.Fprintf(w, "%t\n", foundLineNameInDelayCache)
}
