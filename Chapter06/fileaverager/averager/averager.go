package averager

import ()

func FloatAverage(numbers []float64) float64 {
	var sum float64 = 0.0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	return sum/sampleCount
}