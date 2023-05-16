package safe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// make a test for IsNotEmpty Function for every type available in golang

func TestNotEmpty(t *testing.T) {
	// string
	str := "Hello World"
	assert.Equal(t, IsNotEmpty(str), true)
	// int
	integer := 10
	assert.Equal(t, IsNotEmpty(integer), true)
	// float
	float := 10.0
	assert.Equal(t, IsNotEmpty(float), true)
	// bool
	boolean := true
	assert.Equal(t, IsNotEmpty(boolean), true)
	// slice
	slice := []int{1, 2, 3}
	assert.Equal(t, IsNotEmpty(slice), true)
	var sliceEmpty []int
	assert.Equal(t, IsNotEmpty(sliceEmpty), false)
	// map
	mapping := map[string]int{"Hello": 1, "World": 2}
	assert.Equal(t, IsNotEmpty(mapping), true)
	// struct
	structure := TestingWithStruct{
		OuterField:  "Hellow World",
		InnerStruct: InnerStruct{"Hellow World World"},
	}
	assert.Equal(t, IsNotEmpty(structure), true)
	// pointer
	pointer := &TestingWithStruct{
		OuterField:  "Hellow World",
		InnerStruct: InnerStruct{"Hellow World World"},
	}
	assert.Equal(t, IsNotEmpty(pointer), true)
}

// create benchmark IsNotEmpty function in with all possible type in golang
