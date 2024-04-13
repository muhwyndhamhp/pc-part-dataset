package utils

import "strconv"

func ToFloat64(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return res
}

func ToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return res
}

func ToBool(s string) bool {
	return s == "true"
}
