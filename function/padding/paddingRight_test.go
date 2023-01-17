package padding

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnPaddingRight(t *testing.T) {
	f := &fnPaddingRight{}

	v, err := function.Eval(f, "test", 10, "x")
	assert.Nil(t, err)
	assert.Equal(t, "testxxxxxx", v)
}
