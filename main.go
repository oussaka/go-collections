package gocollections

import (
	"fmt"
)

func Map[T, R any](slice []T, predicate func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = predicate(v)
	}
	return result
}

func Find[T any](slice []T, predicate func(T) bool) (T, error) {
	var notFound T
	for i := range slice {
		if predicate(slice[i]) {
			return slice[i], nil
		}
	}
	return notFound, fmt.Errorf("Element not found")
}

func Filter[T any](slice []T, predicate func(T) bool) ([]T, error) {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no elements found")
	}
	return result, nil
}
