package clif

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type sliceprop struct {
	name string
	num  int
}

func createSliceToConvert() []*sliceprop {
	return []*sliceprop{
		&sliceprop{
			name: "prop1",
			num:  1,
		},
		&sliceprop{
			name: "prop2",
			num:  2,
		},
		&sliceprop{
			name: "prop3",
			num:  3,
		},
	}
}

func createMapToConvert() map[string]*sliceprop {
	return map[string]*sliceprop{
		"prop1": {
			name: "prop1",
			num:  1,
		},
		"prop2": {
			name: "prop2",
			num:  2,
		},
		"prop3": {
			name: "prop3",
			num:  3,
		},
	}
}

func TestSliceConversion(t *testing.T) {

	t.Run("slice passed as interface is properly converted to a slice of interfaces", func(t *testing.T) {

		// Arrange
		sliceToConvert := createSliceToConvert()

		// Act
		convertedSlice := interfaceToSlice(sliceToConvert)

		// Assert
		assert.Len(t, convertedSlice, 3)
	})

	t.Run("map passed as interface is properly converted to a slice of interfaces", func(t *testing.T) {

		// Arrange
		mapToConvert := createMapToConvert()

		// Act
		convertedSlice := interfaceToSlice(mapToConvert)

		// Assert
		assert.Len(t, convertedSlice, 3)
	})
}
