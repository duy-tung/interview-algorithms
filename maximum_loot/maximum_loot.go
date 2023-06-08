package maximum_loot

import (
	"math"
	"sort"
)

type Item struct {
	Weight int
	Cost   int
}

func maximum_loot(capacity int, weights []int, costs []int) float64 {
	if capacity == 0 {
		return 0
	}

	items := make([]Item, len(weights))
	for i := range weights {
		items[i] = Item{Weight: weights[i], Cost: costs[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return float64(items[i].Cost)/float64(items[i].Weight) > float64(items[j].Cost)/float64(items[j].Weight)
	})

	value := 0.0

	for _, item := range items {
		if capacity == 0 {
			break
		}
		amount := math.Min(float64(capacity), float64(item.Weight))
		capacity -= int(amount)
		value += amount * float64(item.Cost) / float64(item.Weight)
	}

	return math.Round(value*1000) / 1000
}
