package job_sequencing

import "testing"

func TestJobSequencing(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := job_sequencing([]Job{{job: "A", profit: 50, deadline: 2}, {job: "B", profit: 20, deadline: 1}, {job: "C", profit: 30, deadline: 2}, {job: "D", profit: 25, deadline: 1}, {job: "E", profit: 15, deadline: 3}})
		if result != "DAEBC" {
			t.Errorf("Expected DEA, got %s", result)
		}
	})
}
