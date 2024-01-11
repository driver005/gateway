package utils

import "reflect"

func GetSetDifference[T any](original []T, compare []T) []T {
	var difference []T
	for _, element := range original {
		for _, comp := range compare {
			if reflect.DeepEqual(element, comp) {
				difference = append(difference, element)
			}
		}
	}

	return difference
}
