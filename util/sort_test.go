package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortString(t *testing.T) {
	ao := assert.New(t)
	myString := "dskro wecd"
	expectedString := " cddekorsw"

	wantString := SortString(myString)
	ao.Equal(expectedString, wantString)
}
