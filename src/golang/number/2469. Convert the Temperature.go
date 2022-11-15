package number

//开氏度 = 摄氏度 + 273.15
//华氏度 = 摄氏度 * 1.80 + 32.00
func convertTemperature(celsius float64) []float64 {
	var res []float64 = make([]float64,2)
	res[0] = celsius + 273.15
	res[1] = celsius * 1.8 + 32
	return res
}