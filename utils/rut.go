package utils

import (
	"errors"
	"strconv"
)

func toInt(toConvert string) int {
	converted, _ := strconv.Atoi(toConvert)
	return converted
}

func generateVerifier(rut string) (string, error) {
	if _, err := strconv.Atoi(rut); err != nil {
		return "", errors.New("invalid RUT")
	}

	var multiplier = 2
	var sum = 0
	var remainder int
	var division int
	var rutLength = len(rut)

	for i := rutLength - 1; i >= 0; i-- {
		sum = sum + toInt(rut[i:i+1]) * multiplier
		multiplier++
		if (multiplier == 8) {
			multiplier = 2
		}
	}

	division = sum / 11
	division = division * 11.0
	remainder = sum - int(division)

	if (remainder != 0) {
		remainder = 11 - remainder
	}

	if (remainder == 10) {
		return "K", nil
	} else {
		return strconv.Itoa(remainder), nil
	}
}
