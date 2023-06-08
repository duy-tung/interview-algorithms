package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
import (
	"fmt"
	"strconv"
)

func Solution(A string, B string) int {
	// Implement your solution here
	hoursA, minutesA := parseTime(A)
	hoursB, minutesB := parseTime(B)

	if hoursA == hoursB && minutesA == minutesB {
		return 0
	}

	startMinutes := hoursA*60 + minutesA

	if startMinutes%15 != 0 {
		startMinutes = startMinutes + (15 - startMinutes%15)
	}

	endMinutes := hoursB*60 + minutesB

	if endMinutes%15 != 0 {
		endMinutes = endMinutes - endMinutes%15
	}

	num := (endMinutes - startMinutes) / 15

	if num < 0 {
		num += 24 * 60 / 15
	}

	fmt.Println(num)

	return num
}

func parseTime(time string) (int, int) {
	hours, _ := strconv.Atoi(time[:2])
	minutes, _ := strconv.Atoi(time[3:])

	return hours, minutes
}
