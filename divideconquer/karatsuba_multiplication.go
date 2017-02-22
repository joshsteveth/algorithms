//Separate x into x1. B^m + x0
//and y into y1.B^m + y0
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
//handle negative with absolute value and then multiply it with -1 if needed
func KaratsubaMult(a, b int) int {
	var useNeg bool
	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		useNeg = true
	}
	a, b = absInt(a), absInt(b)

	m := getMValue(a, b)

	x1, x0 := separateInts(a, m)
	y1, y0 := separateInts(b, m)

	z2, z0 := x1*y1, x0*y0
	z1 := (x1+x0)*(y1+y0) - z0 - z2

	result := (z2 * int(math.Pow10(m*2))) + (z1 * int(math.Pow10(m))) + z0
	if useNeg {
		result = result * -1
	}

	return result
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

//recursive mode
//make sure it only ends in 1 digt * 1 digt multiplication
func karatsuba(a, b string) string {
	if len(a) <= 4 || len(b) <= 4 {
		return multString(a, b)
	}

	m := getMValueString(a, b)

	//separate x into x1 and x0
	x1, x0 := separateIntsString(a, m)
	y1, y0 := separateIntsString(b, m)

	z2, z0 := karatsuba(x1, y1), karatsuba(x0, y0)
	z1 := karatsuba(addString(x1, x0), addString(y1, y0))
	z1 = deductString(deductString(z1, z0),z2)

	result1 := multStringLong(z2, powerOf("1", m*2))
	result2 := multStringLong(z1, powerOf("1", m))

	return addString(addString(result1, result2), z0)
}

func getMValueString(a, b string) int {
	l := len(a)
	if len(b) < len(a) {
		l = len(b)
	}
	return l / 2
}

func separateIntsString(a string, m int) (string, string) {
	return a[:len(a)-m], a[len(a)-m:]
}

func powerOf(a string, pow int) string {
	for i := 0; i < pow; i++ {
		a += "0"
	}
	return a
}

func addString(a, b string) string {
	ma := map[int]int{}
	mb := map[int]int{}

	for i := 0; i < len(a); i++ {
		val, _ := strconv.Atoi(a[len(a)-1-i : len(a)-i])
		ma[i] = val
	}

	for i := 0; i < len(b); i++ {
		val, _ := strconv.Atoi(b[len(b)-1-i : len(b)-i])
		mb[i] = val
	}

	l := len(ma)
	if len(ma) < len(mb) {
		l = len(mb)
	}

	mresult := map[int]int{}
	var overflow int
	var isOverflow bool

	for i := 0; i < l; i++ {
		var result int
		if isOverflow {
			result = overflow
		}
		isOverflow = false

		if val, ok := ma[i]; ok {
			result += val
		}
		if val, ok := mb[i]; ok {
			result += val
		}

		if result > 9 {
			isOverflow = true
			resultStr := fmt.Sprintf("%d", result)

			overflow, _ = strconv.Atoi(string(resultStr[0]))
			result, _ = strconv.Atoi(string(resultStr[1]))
		}

		mresult[i] = result
	}

	if isOverflow {
		mresult[l] = overflow
		l++
	}

	var str string
	for i := 0; i < l; i++ {
		str += fmt.Sprintf("%d", mresult[l-1-i])
	}
	return str
}

func deductString(a, b string) string {
	ma, mb := map[int]int{}, map[int]int{}

	var isNeg bool

	if less(a, b) {
		a, b = b, a
		isNeg = true
	}

	for i := 0; i < len(a); i++ {
		val, _ := strconv.Atoi(a[len(a)-1-i : len(a)-i])
		ma[i] = val
	}

	for i := 0; i < len(b); i++ {
		val, _ := strconv.Atoi(b[len(b)-1-i : len(b)-i])
		mb[i] = val
	}

	l := len(ma)
	if len(ma) < len(mb) {
		l = len(mb)
	}

	mresult := map[int]int{}
	var overflow int
	var isOverflow bool

	for i := 0; i < l; i++ {
		var result int
		if isOverflow {
			result -= overflow
		}
		isOverflow = false

		if val, ok := ma[i]; ok {
			result += val
		}
		if val, ok := mb[i]; ok {
			result -= val
		}

		if result < 0 {
			isOverflow = true

			overflow = 1
			result += 10
		}

		mresult[i] = result
	}

	if isOverflow {
		mresult[l] = overflow
		l++
	}

	var str string
	for i := 0; i < l; i++ {
		str += fmt.Sprintf("%d", mresult[l-1-i])
	}

	if isNeg {
		str = fmt.Sprintf("-%s", str)
	}

	return str
}

func less(a, b string) bool {
	if len(a) < len(b) {
		return true
	} else if len(a) > len(b) {
		return false
	}

	ma, mb := map[int]int{}, map[int]int{}

	for i := 0; i < len(a); i++ {
		val, _ := strconv.Atoi(a[len(a)-1-i : len(a)-i])
		ma[i] = val
	}

	for i := 0; i < len(b); i++ {
		val, _ := strconv.Atoi(b[len(b)-1-i : len(b)-i])
		mb[i] = val
	}

	for i := 0; i < len(ma); i++ {
		if ma[i] < mb[i] {
			return true
		} else if ma[i] > mb[i] {
			return false
		}
	}

	return false
}

//function to multiply only limited number of string
func multString(a, b string) string {
	a2, _ := strconv.Atoi(a)
	b2, _ := strconv.Atoi(b)
	var c int
	for i := 0; i < b2; i++ {
		c += a2
	}
	return fmt.Sprintf("%d", c)
}

//support long multiplication, however b only consists of 1 element
//create series of integer using map for a and then multiply it one by one
//the order between a and b matters
func multStringLongSingle(a, b string) string {
	var result string

	ma := map[int]int{}
	for i := 0; i < len(a); i++ {
		val, _ := strconv.Atoi(a[len(a)-1-i : len(a)-i])
		ma[i] = val
	}

	for i := 0; i < len(a); i++ {
		result = addString(result, powerOf(multString(b, fmt.Sprintf("%d", ma[i])), i))
	}

	return result
}

func multStringLong(a, b string) string {
	var result string

	mb := map[int]int{}
	for i := 0; i < len(b); i++ {
		val, _ := strconv.Atoi(b[len(b)-1-i : len(b)-i])
		mb[i] = val
	}

	for i := 0; i < len(b); i++ {
		singleMultResult := multStringLongSingle(a, fmt.Sprintf("%d", mb[i]))
		powered := powerOf(singleMultResult, i)
		result = addString(result, powered)
	}

	return result
}
