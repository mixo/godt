package godt

import (
    "time"
)

const (
    timeLayout = "2006-01-02"
)

// Returns the time.Time object by parsing a timeString
func Get(timeString string) time.Time {
    time, err := time.Parse(timeLayout, timeString)
    if err != nil {
        panic(err)
    }

    return time
}

// Returns a string representation of a time.Time object, formatted according to the layout '2006-01-02'
func ToString(date time.Time) string {
    return date.Format(timeLayout)
}
