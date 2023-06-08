package max_pairwise_product

import "testing"

func TestMaxPairwiseProduct(t *testing.T) {
	result := max_pairwise_product([]int{7, 5, 14, 2, 8, 8, 10, 1, 2, 3})
	if result != 140 {
		t.Errorf("Max pairwise product was incorrect, got: %d, want: %d.", result, 140)
	}
}
