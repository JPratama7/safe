package safetypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResultOk(t *testing.T) {
	res := result_test_ok()

	assert.Equal(t, res.IsOk(), true)
	assert.Equal(t, res.IsErr(), false)
	assert.NotEmpty(t, res.val)
	assert.NotNil(t, res.val)
}

func TestResultErr(t *testing.T) {
	res := result_test_none()

	assert.Equal(t, res.IsOk(), false)
	assert.Equal(t, res.IsErr(), true)
	assert.Empty(t, res.val)
	assert.Nil(t, res.val)
}

func result_test_ok() (res Result[int]) {
	return Ok(7)
}

func result_test_none() (res Result[int]) {
	return Err[int]("some fancy error message")
}
