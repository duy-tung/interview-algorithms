package job_sequencing

import (
	"sort"
	"strings"
)

type Job struct {
	job      string
	profit   int
	deadline int
}

func job_sequencing(jobs []Job) string {
	// create a array to store the result init with empty string
	result := make([]string, 5)

	// sort by profit desc
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].profit > jobs[j].profit
	})

	// find max deadline
	maxDeadline := 0

	for i := 0; i < len(jobs); i++ {
		if jobs[i].deadline > maxDeadline {
			maxDeadline = jobs[i].deadline
		}
	}

	for i := 0; i < len(jobs); i++ {
		if result[jobs[i].deadline-1] == "" {
			result[jobs[i].deadline-1] = jobs[i].job
		} else {
			result[maxDeadline+jobs[i].deadline-1] = jobs[i].job
		}
	}

	return strings.Join(result, "")
}
