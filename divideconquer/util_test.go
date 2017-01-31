package divcon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsInt(t *testing.T) {
	a := -10
	assert.Equal(t, 10, absInt(a))
}
