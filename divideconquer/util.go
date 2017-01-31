package divcon

import (
	"math"
)

//since standard library only supports float64
//create our own func that supports int as well
func absInt(i int) int {
	return int(math.Abs(float64(i)))
}
