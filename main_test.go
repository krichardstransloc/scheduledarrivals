package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getTestStops() []*Stop {
	t := time.Now()
	return []*Stop{
		{
			ID: 780017,
			Position: &Position{
				Latitude:  35.995952,
				Longitude: -78.906005,
			},
			Sequence:  0,
			Timepoint: true,
			StopTime:  time.Date(t.Year(), t.Month(), t.Day(), 7, 30, 0, 0, t.Location()),
		},
		{
			ID: 2583089,
			Position: &Position{
				Latitude:  35.995579,
				Longitude: -78.908444,
			},
			Sequence: 1,
		},
		{
			ID: 779732,
			Position: &Position{
				Latitude:  35.999935,
				Longitude: -78.909003,
			},
			Sequence: 2,
		},
		{
			ID: 779290,
			Position: &Position{
				Latitude:  36.0016,
				Longitude: -78.9117,
			},
			Sequence: 3,
		},
		{
			ID: 780096,
			Position: &Position{
				Latitude:  36.00379,
				Longitude: -78.91528,
			},
			Sequence: 4,
		},
		{
			ID: 2562471,
			Position: &Position{
				Latitude:  36.005842,
				Longitude: -78.919316,
			},
			Sequence: 5,
		},
		{
			ID: 779303,
			Position: &Position{
				Latitude:  36.01058,
				Longitude: -78.91947,
			},
			Sequence:  6,
			Timepoint: true,
			StopTime:  time.Date(t.Year(), t.Month(), t.Day(), 7, 38, 0, 0, t.Location()),
		},
		{
			ID: 2562472,
			Position: &Position{
				Latitude:  36.010586,
				Longitude: -78.922118,
			},
			Sequence: 7,
		},
		{
			ID: 779425,
			Position: &Position{
				Latitude:  36.010538,
				Longitude: -78.925131,
			},
			Sequence: 8,
		},
		{
			ID: 779436,
			Position: &Position{
				Latitude:  36.010649,
				Longitude: -78.927815,
			},
			Sequence: 9,
		},
		{
			ID: 779363,
			Position: &Position{
				Latitude:  36.012039,
				Longitude: -78.930439,
			},
			Sequence: 10,
		},
		{
			ID: 779150,
			Position: &Position{
				Latitude:  36.01302,
				Longitude: -78.9316812,
			},
			Sequence: 11,
		},
		{
			ID: 779443,
			Position: &Position{
				Latitude:  36.014786,
				Longitude: -78.933818,
			},
			Sequence: 12,
		},
		{
			ID: 779416,
			Position: &Position{
				Latitude:  36.0162,
				Longitude: -78.93737,
			},
			Sequence: 13,
		},
		{
			ID: 779565,
			Position: &Position{
				Latitude:  36.016499389233,
				Longitude: -78.936650263442,
			},
			Sequence: 14,
		},
		{
			ID: 779190,
			Position: &Position{
				Latitude:  36.01117,
				Longitude: -78.93717,
			},
			Sequence: 15,
		},
		{
			ID: 779198,
			Position: &Position{
				Latitude:  36.008813,
				Longitude: -78.93743,
			},
			Sequence:  16,
			Timepoint: true,
			StopTime:  time.Date(t.Year(), t.Month(), t.Day(), 7, 44, 0, 0, t.Location()),
		},
		{
			ID: 780069,
			Position: &Position{
				Latitude:  36.008863,
				Longitude: -78.943709,
			},
			Sequence: 17,
		},
		{
			ID: 780071,
			Position: &Position{
				Latitude:  36.006834,
				Longitude: -78.947226,
			},
			Sequence: 18,
		},
		{
			ID: 780072,
			Position: &Position{
				Latitude:  36.0048335155034,
				Longitude: -78.9493562745149,
			},
			Sequence: 19,
		},
		{
			ID: 780073,
			Position: &Position{
				Latitude:  36.003907218545,
				Longitude: -78.9507185925409,
			},
			Sequence: 20,
		},
		{
			ID: 780074,
			Position: &Position{
				Latitude:  36.00382,
				Longitude: -78.95275,
			},
			Sequence: 21,
		},
		{
			ID: 780118,
			Position: &Position{
				Latitude:  36.005668,
				Longitude: -78.954668,
			},
			Sequence: 22,
		},
		{
			ID: 780077,
			Position: &Position{
				Latitude:  36.007155,
				Longitude: -78.954687,
			},
			Sequence: 23,
		},
		{
			ID: 780078,
			Position: &Position{
				Latitude:  36.008762,
				Longitude: -78.954635,
			},
			Sequence: 24,
		},
		{
			ID: 780079,
			Position: &Position{
				Latitude:  36.013029,
				Longitude: -78.95629,
			},
			Sequence: 25,
		},
		{
			ID: 780080,
			Position: &Position{
				Latitude:  36.015459,
				Longitude: -78.962606,
			},
			Sequence: 26,
		},
		{
			ID: 780081,
			Position: &Position{
				Latitude:  36.016633,
				Longitude: -78.966937,
			},
			Sequence: 27,
		},
		{
			ID: 2562298,
			Position: &Position{
				Latitude:  36.020903,
				Longitude: -78.967099,
			},
			Sequence: 28,
		},
		{
			ID: 2562474,
			Position: &Position{
				Latitude:  36.026323,
				Longitude: -78.967301,
			},
			Sequence:  29,
			Timepoint: true,
			StopTime:  time.Date(t.Year(), t.Month(), t.Day(), 7, 56, 0, 0, t.Location()),
		},
	}
}

func TestInterpolateDeltas(t *testing.T) {
	stoplist := getTestStops()[0:7]
	deltas, err := InterpolateDeltas(stoplist)
	expected := []int{0, 45, 143, 204, 285, 371, 477}
	assert.Equal(t, expected, deltas)
	assert.Nil(t, err)
}

func TestInterpolateDeltas_MultipleTimepoints(t *testing.T) {
	stoplist := getTestStops()
	deltas, err := InterpolateDeltas(stoplist)
	expected := []int{
		0, 45, 143, 204, 285, 371, 480, 511, 546, 577, 613, 633, 668, 714, 723, 800, 840, 924, 982,
		1025, 1048, 1075, 1115, 1139, 1165, 1239, 1333, 1394, 1464, 1554,
	}
	assert.Equal(t, expected, deltas)
	assert.Nil(t, err)
}

func TestInterpolateDeltas__NoStops(t *testing.T) {
	deltas, err := InterpolateDeltas([]*Stop{})
	assert.Nil(t, deltas)
	assert.Error(t, err)
}

func TestInterpolateDeltas__OneStop(t *testing.T) {
	stoplist := getTestStops()
	deltas, err := InterpolateDeltas([]*Stop{stoplist[0]})
	expected := []int{0}
	assert.Equal(t, expected, deltas)
	assert.Nil(t, err)
}

func TestInterpolateDeltas_NoTimepoints(t *testing.T) {
	stoplist := getTestStops()[0:3]
	stoplist[0].Timepoint = false
	stoplist[0].StopTime = time.Time{}

	deltas, err := InterpolateDeltas(stoplist)
	assert.Nil(t, deltas)
	assert.Error(t, err)
}

func TestInterpolateDeltas_ConsecutiveTimepoints(t *testing.T) {
	stoplist := getTestStops()[0:8]
	stoplist[7].Timepoint = stoplist[6].Timepoint
	stoplist[7].StopTime = stoplist[6].StopTime

	deltas, err := InterpolateDeltas(stoplist)
	expected := []int{0, 45, 143, 204, 285, 371, 480, 480}
	assert.Equal(t, expected, deltas)
	assert.Nil(t, err)
}

func TestInterpolateDeltas_StopsBeforeTimepoint(t *testing.T) {
	stoplist := getTestStops()[0:7]

	stoplist[2].Timepoint = stoplist[0].Timepoint
	stoplist[2].StopTime = stoplist[0].StopTime
	stoplist[0].Timepoint = false
	stoplist[0].StopTime = time.Time{}

	deltas, err := InterpolateDeltas(stoplist)
	// the first 2 elements are before the first timepoint and should be 0
	expected := []int{0, 0, 0, 87, 203, 326, 477}
	assert.Equal(t, expected, deltas)
	assert.Nil(t, err)
}

func TestInterpolateDeltas_StopsAfterTimepoint(t *testing.T) {
	stoplist := getTestStops()[0:8]
	deltas, err := InterpolateDeltas(stoplist)
	// the 8th element is after a timepoint, so it should be 0
	expected := []int{0, 45, 143, 204, 285, 371, 477, 0}
	assert.Equal(t, expected, deltas)
	assert.Nil(t, err)
}

func TestGetETA(t *testing.T) {
	stoplist := getTestStops()[0:7]
	eta, err := GetETA(1, 2, stoplist)
	assert.Equal(t, 98, eta)
	assert.Nil(t, err)
}

func TestGetETA_InvalidOrder(t *testing.T) {
	stoplist := getTestStops()[0:7]
	eta, err := GetETA(2, 1, stoplist)
	assert.Equal(t, 0, eta)
	assert.Error(t, err)
}

func TestGetETA_ErrorPropagation(t *testing.T) {
	stoplist := getTestStops()[0:3]
	stoplist[0].Timepoint = false
	stoplist[0].StopTime = time.Time{}
	eta, err := GetETA(1, 2, stoplist)
	assert.Equal(t, 0, eta)
	assert.Error(t, err)
}
