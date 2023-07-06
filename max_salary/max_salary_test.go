package max_salary

import "testing"

func TestLargestNumber(t *testing.T) {
	test := []struct {
		a        []string
		expected string
	}{
		{[]string{"9", "4", "6", "1", "9"}, "99641"},
		{[]string{"23", "39", "92"}, "923923"},
		{[]string{"21", "2"}, "221"},
		{[]string{"68", "6", "61", "73"}, "7368661"},
	}

	for _, test := range test {
		actual := largestNumber(test.a)
		if actual != test.expected {
			t.Errorf("largestNumber(%v) = %s; expected %s", test.a, actual, test.expected)
		}
	}
}
