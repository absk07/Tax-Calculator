package conversion

import (
	"strconv"
)

func StringToFloat(strs []string) ([]float64, error) {
	var floats []float64
	for _, str := range strs {
		floatPrice, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
