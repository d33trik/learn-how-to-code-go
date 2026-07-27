package main

import (
	"errors"
	"fmt"
	"log"
)

type coordinatesError struct {
	lat  float64
	long float64
	err  error
}

func (ce coordinatesError) Error() string {
	return fmt.Sprintf("coordinates error: %v %v %v", ce.lat, ce.long, ce.err)
}

func main() {
	_, err := parseCoordinates(180.283, 152.829)
	if err != nil {
		log.Println(err)
	}
}

func parseCoordinates(lat, long float64) (float64, error) {
	if lat < -90 || lat > 90 {
		return 0, coordinatesError{lat, long, errors.New("lat must be -90 < lat < 90")}
	}

	if long < -180 || long > 180 {
		return 0, coordinatesError{lat, long, errors.New("long must be -180 < long < 180")}
	}
	return 42, nil
}
