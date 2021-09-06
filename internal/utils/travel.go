package utils

import (
	"github.com/ozonva/ova-travel-api/internal/travel"
)

func SplitByBatch(arr []travel.Trip, batch int) [][]travel.Trip {
	batchSlice := make([][]travel.Trip, 0)
	for i := 0; i < len(arr); i += batch {
		batchSlice = append(batchSlice, arr[i:MinInt(i+batch, len(arr))])
	}

	return batchSlice
}
