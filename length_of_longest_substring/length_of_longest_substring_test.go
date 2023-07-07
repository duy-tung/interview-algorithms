package length_of_longest_substring

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := lengthOfLongestSubstring("abcabcbb")
		if result != 3 {
			t.Errorf("Expected 3, got %d", result)
		}
	})
}
