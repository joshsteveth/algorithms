package divcon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKaratsubaMult(t *testing.T) {
	a, b := -123, 4571
	assert.Equal(t, a*b, KaratsubaMult(a, b))
}
