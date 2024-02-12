package utils

import (
	"github.com/google/uuid"
)

func GenerateComponentID() string {
	return uuid.New().String()
}

// Go utils

func SliceContainsItem[K comparable](slice []K, item K) (result bool) {
	for _, v := range slice {
		if item == v {
			return true
		}
	}
	return false
}

// Shorthand to _, ok := someMap[someKey]
func MapContainsKey[K comparable, V any](map_ map[K]V, item K) (result bool) {
	_, result = map_[item]
	return
}
