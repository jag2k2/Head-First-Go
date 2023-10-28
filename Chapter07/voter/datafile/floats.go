package datafile

import "bufio"
import "os"
import "strconv"

func ReadFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var scanner *bufio.Scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}