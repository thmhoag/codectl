package clif_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/thmhoag/codectl/pkg/clif"
	"testing"
)

func TestFormatWrite(t *testing.T) {

	t.Run("correctly prints table from slice", func(t *testing.T) {

		// Arrange
		testSlice := []*testItem{{Name: "item1", Num:  1,}, {Name: "item2", Num:  2,}, {Name: "item3", Num:  3,},}
		var b bytes.Buffer

		// Act
		err := clif.New("table {{ .Name }} {{ .Num }}").Output(&b).Write(testSlice)

		actual := b.String()
		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, b)
		assert.Equal(t, "NAME                NUM\nitem1               1\nitem2               2\nitem3               3\n", actual)
	})

	t.Run("correctly prints table from map", func(t *testing.T) {

		// Arrange
		testMap := map[string]*testItem{"item1": {Name: "item1", Num:  1,}, "item2": {Name: "item2", Num:  2,}, "item3": {Name: "item3", Num:  3,},}
		var b bytes.Buffer

		// Act
		err := clif.New("table {{ .Name }} {{ .Num }}").Output(&b).Write(testMap)

		actual := b.String()
		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, b)
		assert.Equal(t, "NAME                NUM\nitem1               1\nitem2               2\nitem3               3\n", actual)
	})
}

type testItem struct {
	Name string
	Num  int
}