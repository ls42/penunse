package main

import (
	"fmt"
	"strconv"
)

// parseTimeFilterString parses a string and returns a string that is usable in
// GetTransactionsWithFilters
func parseTimeFilterString(filter string) string {
	// months := []string{
	// 	"jan", "january",
	// 	"feb", "february",
	// 	"mar", "march",
	// 	"apr", "april",
	// 	"may", "may",
	// 	"jun", "june",
	// 	"jul", "july",
	// 	"aug", "august",
	// 	"sep", "september",
	// 	"oct", "october",
	// 	"nov", "november",
	// 	"dec", "december",
	// }

	// If time is just a number its an ISO week -> return data from that week
	if parsed, err := strconv.Atoi(filter); err == nil {
		return fmt.Sprintf("strftime('%%W', date) = '%s'", strconv.Itoa(parsed))
	}

	// Default -> return date from current month
	return "date(date) >= date('now', 'start of month') AND " +
		"date(date) <= date('now', '+1 month', 'start of month', '-1 day')"
}
