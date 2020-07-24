package number

import (
	"math"
	"time"
)

func DaysBetweenDates(date1 string, date2 string) int {
	layout := "2006-01-02"

	t1, _ := time.Parse(layout, date1)
	t2, _ := time.Parse(layout, date2)
	return int(math.Abs(float64(t1.Sub(t2).Hours() / 24)))
}
