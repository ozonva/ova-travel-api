package utils

import (
	"fmt"

	"github.com/ozonva/ova-travel-api/internal/travel"
)

func SplitByBatch(arr []travel.Trip, batch int) [][]travel.Trip {
	batchSlice := make([][]travel.Trip, 0)
	for i := 0; i < len(arr); i += batch {
		batchSlice = append(batchSlice, arr[i:MinInt(i+batch, len(arr))])
	}

	return batchSlice
}

func ConvertSpliceToMap(arr []travel.Trip) map[int]travel.Trip {
	result := make(map[int]travel.Trip)
	for _, entity := range arr {
		if _, found := result[entity.UserID]; found {
			panic(fmt.Sprintf("UserID is duplicated %v", entity.UserID))
		}

		result[entity.UserID] = entity
	}

	return result
}
