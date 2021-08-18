package main

import (
	"fmt"

	"github.com/ozonva/ova-travel-api/internal/travel"
	"github.com/ozonva/ova-travel-api/internal/utils"
)

func main() {
	welcome := `Initial entry point for the ova travel api project`
	fmt.Printf(welcome)

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

	a1 := travel.Trip{UserID: 0, FromLocation: "1", DestLocation: "2"}
	a2 := travel.Trip{UserID: 1, FromLocation: "3", DestLocation: "4"}
	a3 := travel.Trip{UserID: 2, FromLocation: "1", DestLocation: "5"}
	slice := []travel.Trip{a1, a2, a3}
	fmt.Println(utils.SplitByBatch(slice, 1))

	fmt.Println(utils.ConvertSpliceToMap(slice))

	for range slice {
		fmt.Println(utils.GetFileContent("Makefile"))
	}

}
