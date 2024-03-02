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

// Tries to get an index of the item in slice.
func SliceItemIndex[K comparable](slice []K, item K) (index int) {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// Removes the item at the given index.
func SliceRemoveAtIndex[K comparable](s []K, index int) []K {
	return append(s[:index], s[index+1:]...)
}

// Shorthand to _, ok := someMap[someKey]
func MapContainsKey[K comparable, V any](map_ map[K]V, item K) (result bool) {
	_, result = map_[item]
	return
}
