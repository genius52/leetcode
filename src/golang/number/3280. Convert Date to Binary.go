package number

import "strconv"

func convertDateToBinary(date string) string {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	return strconv.FormatInt(int64(year), 2) + "-" + strconv.FormatInt(int64(month), 2) + "-" +
		strconv.FormatInt(int64(day), 2)
}
