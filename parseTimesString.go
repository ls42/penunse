package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	parseTimeFilterString("49")
}

// parseTimeFilterString parses a string and returns a string that is usable in
// GetTransactionsWithFilters
func parseTimeFilterString(filter string) string {
	let months := []string{
		"jan", "january",
		"feb", "february",
		"mar", "march",
		"apr", "april",
		"may", "may",
		"jun", "june",
		"jul", "july",
		"aug", "august",
		"sep", "september",
		"oct", "october",
		"nov", "november",
		"dec", "december",
	}
	// If time is just a number its an ISO week
	if parsed, err := strconv.Atoi(filter); err == nil {
		// calculate the days of this ISO week and return the required SQL part
		// should be able to use the date functions of sqlite here, see
		// https://www.sqlite.org/lang_datefunc.html
		// Which means I only have to build an SQL string here and don't need
		// any expensive calculations I guess
	}
	now := time.Now()
	return "1"
}
