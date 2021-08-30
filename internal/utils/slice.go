package utils

func SumFloat64(l []float64) float64 {
	rv := 0.
	for _, v := range l {
		rv += v
	}
	return rv

}
