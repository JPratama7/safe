package safetypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSome(t *testing.T) {
	res := option_test_some()

	assert.Equal(t, res.IsSome(), true)
	assert.Equal(t, res.IsNone(), false)
	assert.NotEmpty(t, res.Val)
	assert.NotNil(t, res.Val)
}

func TestNone(t *testing.T) {
	res := option_test_none()

	assert.Equal(t, res.IsSome(), false)
	assert.Equal(t, res.IsNone(), true)
	assert.Empty(t, res.Val)
	assert.Nil(t, res.Val)
}

func option_test_some() (opt Option[int]) {
	return Some[int](7)
}

func option_test_none() (opt Option[int]) {
	return None[int]()
}
