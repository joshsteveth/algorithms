//Separate x into x1. B^m + x0
//and y into y1.B*m + y0
//calculate z2 = x1.y1
//calculate z0 = x0.y0
//calculate z1 = (x1 + x0).(y1 + y0) - z0 - z2
//the result is: z2. B^2m + z1. B^m + z0
//Example for multiplication x=5678 and y=1234:
//separate x and y into a=56; b=78; c=12; and d=34 (use B = 10 and m = 2)
//STEP 1: compute a.c = 672
//STEP 2: compute b.d = 2652
//STEP 3: compute (a + b) (c + d) = 6164
//STEP 4: compute S3 - S2 - S1 = 2840
//Result: (S1 * 10^4) + (S4 * 10^2) +  S2

package divcon

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

//basic multiplication func which use B=10 and m=half power of smallest number
//sort the input ints first and determine value of m
func KaratsubaMult(a, b int) int {
	m := getMValue(a, b)
	fmt.Printf("m is: %d\n", m)

	x1, x0 := separateInts(a, m)
	y1, y0 := separateInts(b, m)
	fmt.Printf("x1, x0 : %d, %d\n", x1, x0)
	fmt.Printf("y1, y0 : %d, %d\n", y1, y0)

	z2, z0 := x1*y1, x0*y0
	z1 := (x1+x0)*(y1+y0) - z0 - z2

	return (z2 * int(math.Pow10(m*2))) + (z1 * int(math.Pow10(m))) + z0
}

func getMValue(a, b int) int {
	i := []int{a, b}
	sort.Ints(i)
	return len(fmt.Sprintf("%d", i[0])) / 2
}

func separateInts(a, m int) (int, int) {
	intStr := fmt.Sprintf("%d", a)
	x1, _ := strconv.Atoi(intStr[:len(intStr)-m])
	x0, _ := strconv.Atoi(intStr[len(intStr)-m:])
	return x1, x0
}
