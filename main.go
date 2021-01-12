package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

// Position does the thing
type Position struct {
	Latitude  float64
	Longitude float64
}

const (
	meanRadius = 6371000.7900
	piOver180  = math.Pi / 180
)

func radians(degrees float64) float64 { return degrees * piOver180 }

// Distance returns the distance (haversine) between two positions
func (me *Position) Distance(p *Position) float64 {
	const diameter = 2 * meanRadius
	lat1 := radians(me.Latitude)
	lat2 := radians(p.Latitude)
	latH := math.Sin((lat1 - lat2) / 2)
	latH *= latH
	lonH := math.Sin(radians(me.Longitude-p.Longitude) / 2)
	lonH *= lonH
	tmp := latH + math.Cos(lat1)*math.Cos(lat2)*lonH
	return diameter * math.Asin(math.Sqrt(tmp))
}

// Stop holds info about a... well, stop...
type Stop struct {
	ID        int
	Position  *Position
	Sequence  int
	StopTime  time.Time
	Timepoint bool
}

// String returns a string representation of Stop.
func (me *Stop) String() string {
	return fmt.Sprintf("%d", me.ID)
}

// InterpolateDeltas gets proportional time deltas (seconds) between stops
func InterpolateDeltas(stoplist []*Stop) ([]int, error) {
	numStops := len(stoplist)

	if numStops == 0 {
		return nil, errors.New("no stops")
	}

	if numStops == 1 {
		return []int{0}, nil
	}

	// Identify indexes for timepoints
	timepoints := make([]int, 0)
	for idx, stop := range stoplist {
		if stop.Timepoint == true && !stop.StopTime.IsZero() {
			timepoints = append(timepoints, idx)
		}
	}

	// without 2 timepoints we can't interpolate times
	if len(timepoints) < 2 {
		return nil, errors.New("invalid timepoints")
	}

	deltas := make([]int, numStops)

	// Prefill zeros for all elements before the first timepoint
	for i := 0; i < timepoints[0]; i++ {
		deltas[i] = 0
	}

	// Prefill zeros for all elements after the first timepoint
	for i := timepoints[len(timepoints)-1]; i < numStops; i++ {
		deltas[i] = 0
	}

	// With two consecutive timepoint indexes, interpolate times for the stops between them. The
	// values at the low and high timpoint indexes represent lower and upper indexes of the stops
	// slice.
	lowTimepointIdx := 0
	highTimepointIdx := 1

	// Keep track of the earliest timepoint. This will be used to calculate as an offset
	initialTimepointIdx := timepoints[lowTimepointIdx]
	initialTimepointTime := stoplist[initialTimepointIdx].StopTime
	for {
		if highTimepointIdx >= len(timepoints) {
			break
		}

		// Grab the subset of stops between the low and high timepoint indexes
		lowStopIdx := timepoints[lowTimepointIdx]
		highStopIdx := timepoints[highTimepointIdx]
		subset := stoplist[lowStopIdx : highStopIdx+1]

		// Grab the seconds difference between timepoints
		lowTime := stoplist[lowStopIdx].StopTime
		highTime := stoplist[highStopIdx].StopTime
		cumulativeDelta := highTime.Sub(lowTime).Seconds()

		// Iterate over elements in the subset, calculating Haversine distance between consecutive
		// stops. Keep track of the individual distances as well as a cumulative distance. distance
		// and cumulative distance are relative to the lower timepoint. Relative distances
		// effectively syncs the distances to the specific range.
		lowSubsetIdx := 0
		highSubsetIdx := 1
		distances := make([]float64, len(subset))
		cumulativeDistance := float64(0)
		distances[0] = 0
		for {
			if highSubsetIdx >= len(subset) {
				break
			}

			lowStop := subset[lowSubsetIdx]
			highStop := subset[highSubsetIdx]
			distance := lowStop.Position.Distance(highStop.Position)

			distances[highSubsetIdx] = distance
			cumulativeDistance += distance

			lowSubsetIdx++
			highSubsetIdx++
		}

		// With a slice of pairwise distances, calculate the proportional distance of two stops to
		// the entire slice. Using the proportional distance, estimate proportional time. Update the
		// deltas slice accordingly.

		// Grab a time offset between the overall earliest timpoint and the lowest timepoint of this
		// subset of stops. Since distances are relative to other stops within the same timepoint
		// bounds, we use the offset to make sure deltas increase monotonically.
		timeOffset := int(lowTime.Sub(initialTimepointTime).Seconds())
		cumulativeTime := 0
		for i, distance := range distances {
			deltaIdx := i + lowStopIdx
			distanceProportion := distance / cumulativeDistance
			proportionalTime := int(float64(cumulativeDelta) * distanceProportion)
			deltas[deltaIdx] = cumulativeTime + proportionalTime + timeOffset
			cumulativeTime += proportionalTime
		}

		lowTimepointIdx++
		highTimepointIdx++
	}

	return deltas, nil
}

// GetETA returns the eta between two stops
func GetETA(originIdx, destinationIdx int, stoplist []*Stop) (int, error) {
	if destinationIdx < originIdx {
		return 0, errors.New("invalid ETA order")
	}

	deltas, err := InterpolateDeltas(stoplist)
	if err != nil {
		return 0, err
	}

	return deltas[destinationIdx] - deltas[originIdx], nil
}

func main() {
	fmt.Println("Hello, World!")
}
