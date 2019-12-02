package godt

import (
    "time"
)

// Returns the number of days between startTime and endTime
func CountDaysInPeriod(startTime, endTime time.Time) int64 {
    return int64(endTime.Sub(startTime).Hours() / 24)
}

func GetPeriod(start, end time.Time) (period []time.Time) {
    for current := start; current.Before(end) || current.Equal(end); current = current.AddDate(0, 0, 1) {
        period = append(period, current)
    }

    return
}
