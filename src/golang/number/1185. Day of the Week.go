package number

import (
	"fmt"
	"time"
)

func DayOfTheWeek(day int, month int, year int) string {
	layout := "2006-01-02"
	dt := fmt.Sprintf("%d-%.2d-%.2d",year,month,day)
	t1, _ := time.Parse(layout, dt)
	return t1.Weekday().String()
}
