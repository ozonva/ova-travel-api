package main

import (
	"fmt"

	"github.com/ozonva/ova-travel-api/internal/utils"
)

func main() {
	welcome := `Initial entry point for the ova travel api project`
	fmt.Printf(welcome)

	emptyBatch := make([]int, 0)
	fmt.Println(emptyBatch)
	fmt.Println(utils.SplitByBatch(emptyBatch, 3))
	smallBatch := []int{1, 2}
	fmt.Println(smallBatch)
	fmt.Println(utils.SplitByBatch(smallBatch, 3))
	tailBatch := []int{1, 2, 3, 4}
	fmt.Println(tailBatch)
	fmt.Println(utils.SplitByBatch(tailBatch, 3))
	multipleBatch := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(multipleBatch)
	fmt.Println(utils.SplitByBatch(multipleBatch, 3))

	emptyMap := make(map[int]string, 0)
	fmt.Println(emptyMap)
	fmt.Println(utils.InvertMap(emptyMap))
	usualMap := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(usualMap)
	fmt.Println(utils.InvertMap(usualMap))

	filterArray := []int{1, 3, 5, 8}
	emptyArr := make([]int, 0)
	fmt.Println(emptyArr)
	fmt.Println(utils.FilterByArray(emptyArr, filterArray))
	duplicatedArr := []int{1, 3, 5, 8}
	fmt.Println(duplicatedArr)
	fmt.Println(utils.FilterByArray(duplicatedArr, filterArray))
	usualArr := []int{1, 2, 4, 7, 8}
	fmt.Println(usualArr)
	fmt.Println(utils.FilterByArray(usualArr, filterArray))
}
