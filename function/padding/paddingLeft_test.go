package padding

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFnPaddingLeft(t *testing.T) {
	f := &fnPaddingLeft{}

	v, err := function.Eval(f, "test", 10, "x")
	assert.Nil(t, err)
	assert.Equal(t, "xxxxxxtest", v)
}
