package utils

import "fmt"

func minInt(values ...int) int {
	minValue := values[0]
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func SplitByBatch(arr []int, batch int) [][]int {
	batchSlice := make([][]int, 0)
	for i := 0; i < len(arr); i += batch {
		batchSlice = append(batchSlice, arr[i:minInt(i+batch, len(arr))])
	}

	return batchSlice
}

func InvertMap(m map[int]string) map[string]int {
	result := make(map[string]int)
	for key, value := range m {
		if _, found := result[value]; found {
			panic(fmt.Sprintf("Key is duplicated %v", value))
		}
		result[value] = key
	}

	return result
}

func contains(arr []int, value int) bool {
	for _, elem := range arr {
		if elem == value {
			return true
		}
	}

	return false
}

func FilterByArray(input []int, filterArray []int) []int {
	// Is it make sense to use map for filterArray for the fast find operation?
	output := make([]int, 0)
	for _, elem := range input {
		if !contains(filterArray, elem) {
			output = append(output, elem)
		}
	}

	return output
}
