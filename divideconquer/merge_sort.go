//if we have set of numbers e.g. [5,4,1,8,7,2,6,3]
//we separate them into 2 sets [5,4,1,8] and [7,2,6,3]
//then separate again recursively so that it can be sorted into [1,4,5,8] and [2,3,6,7]
//finally join them together so we get [1,2,3,4,5,6,7,8]
package divcon

func MergeSort(inp []int) []int {
	divided := divide(inp)

	for len(divided) != 1 {
		container := [][]int{}
		numLoop := len(divided) / 2
		if len(divided)&1 > 0 {
			numLoop++
		}

		for i := 0; i < numLoop; i++ {
			var C []int
			//merge a pair of []int if available
			//check the length first to prevent panic
			if (i*2)+1 < len(divided) {
				C = merge(divided[i*2], divided[i*2+1])
			} else {
				C = divided[i*2]
			}
			container = append(container, C)
		}

		divided = container
	}

	return divided[0]
}

//divide the input sets into the littlest possible
//that means it has not more than 2 elements each
//sort all of them before appending them to the result
func divide(inp []int) [][]int {
	var result [][]int

	numLoop := len(inp) / 2
	if len(inp)&1 > 0 {
		numLoop++
	}

	for i := 0; i < numLoop; i++ {
		tuple := []int{inp[i*2]}
		if (i*2)+1 < len(inp) {
			tuple = append(tuple, inp[i*2+1])
		}
		sortTuple(tuple)
		result = append(result, tuple)
	}
	return result
}

//sort input integers
//return immediately if input has less than 2 elements
//otherwise it will be effective if it has exactly 2 elements
func sortTuple(inp []int) {
	if len(inp) < 2 {
		return
	}
	if inp[1] < inp[0] {
		inp[0], inp[1] = inp[1], inp[0]
	}
}

//merge algorithm is as follows:
//say we have 2 sets A (index i with length n) and B (index j with length m)
//we want to merge them as C (index k with lenght n+m)
//so for k := 0 to (n + m), we compare A[i] and B[j]
//if A[i] is actually less than B[j], we use that value and then increment i by 1
//the same applies for B as well
func merge(A, B []int) []int {
	C := make([]int, len(A)+len(B))
	i, j := 0, 0

	for k := 0; k < len(C); k++ {
		//make sure that i and j doesn't violate length of both slices
		if i >= len(A) {
			C[k] = B[j]
			j++
			continue
		}

		if j >= len(B) {
			C[k] = A[i]
			i++
			continue
		}

		if A[i] < B[j] {
			C[k] = A[i]
			i++
		} else {
			C[k] = B[j]
			j++
		}
	}

	return C
}
