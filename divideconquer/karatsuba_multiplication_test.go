package divcon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKaratsubaMult(t *testing.T) {
	//positive positive
	a, b := 123, 4571
	assert.Equal(t, a*b, KaratsubaMult(a, b))

	//positive negative
	a, b = -123, 4571
	assert.Equal(t, a*b, KaratsubaMult(a, b))

	//negative negative
	a, b = -123, -4571
	assert.Equal(t, a*b, KaratsubaMult(a, b))
}
