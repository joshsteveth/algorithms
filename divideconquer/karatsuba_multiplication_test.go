package divcon

import (
	"fmt"
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

func TestPowerOf(t *testing.T) {
	t.Skip()
	expected := "52000"
	result := powerOf("52", 3)
	assert.Equal(t, expected, result)
}

func TestAddString(t *testing.T) {
	t.Skip()
	a, b := "49200", "68888"
	expected := "118088"
	assert.Equal(t, expected, addString(a, b))
}

func TestAtomicMult(t *testing.T) {
	t.Skip()
	a, b := "3", "4"
	expected := "12"
	assert.Equal(t, expected, multString(a, b))
}

func TestMultStringLongSingle(t *testing.T) {
	t.Skip()
	a, b := "123", "4"
	expected := "492"
	assert.Equal(t, expected, multStringLongSingle(a, b))
}

func TestMultStringLong(t *testing.T) {
	t.Skip()
	a, b := "123", "456"
	expected := fmt.Sprintf("%d", 123*456)
	assert.Equal(t, expected, multStringLong(a, b))
}

func TestProgrammingAssignment(t *testing.T) {
	t.Skip()
	a := "3141592653589793238462643383279502884197169399375105820974944592"
	b := "2718281828459045235360287471352662497757247093699959574966967627"

	t.Log(multStringLong(a, b))
}

func TestDeductString(t *testing.T) {
	a, b := "5301", "25"
	expected := "5276"
	assert.Equal(t, expected, deductString(a, b))

	expected = "-5276"
	assert.Equal(t, expected, deductString(b, a))
}

func TestLess(t *testing.T) {
	a, b := "25", "531123"
	assert.Equal(t, true, less(a, b))
	assert.Equal(t, false, less(b, a))

	a, b = "501", "503"
	assert.Equal(t, true, less(a, b))
	assert.Equal(t, false, less(b, a))
}

func TestKaratsuba(t *testing.T) {
	//t.Skip()
	a := "7315"
	b := "123"
	expected := "13259436476"
	assert.Equal(t, expected, karatsuba(a, b))
}

func TestSeparateIntsString(t *testing.T) {
	a := "5301"
	exp1, exp2 := separateIntsString(a, 2)
	assert.Equal(t, "53", exp1)
	assert.Equal(t, "01", exp2)

	b := "250"
	exp1, exp2 = separateIntsString(b, 2)
	assert.Equal(t, "2", exp1)
	assert.Equal(t, "50", exp2)
}
