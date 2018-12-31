package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// parseTimeFilterString parses a string and returns a string that is usable in
// GetTransactionsWithFilters
func parseTimeFilterString(filter string) string {
	filter = strings.TrimSpace(filter)
	now := time.Now()

	// If time is just a number its an ISO week
	// -> return data from that week
	if weekOfYear, err := strconv.Atoi(filter); err == nil {
		_, currentWeek := now.ISOWeek()
		if weekOfYear <= currentWeek {
			return fmt.Sprintf("strftime('%%W', date) = '%d'", weekOfYear)
		}
		// Display last years data
		return fmt.Sprintf("strftime('%%W', date, '-1 year') = '%d'", weekOfYear)

	}

	// Check if user input is a month
	// -> return data from that month (either this or last year's)
	if month, exists := months[filter]; exists {
		// 1. Check if month is in the past or future.
		// 2. If in future, then use data from last year
		//    Otherwise show data from this year
		fmt.Printf("would use this string for sql: 2018-%s-*", month)
	}

	// Default
	// -> return data from current month
	return "date(date) >= date('now', 'start of month') AND " +
		"date(date) <= date('now', '+1 month', 'start of month', '-1 day')"
}
