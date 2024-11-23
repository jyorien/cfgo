package vectors

import (
	"errors"
	"math"
)

func DotProduct(vecA, vecB []float64) (float64, error) {
	if len(vecA) != len(vecB) {
		return 0, errors.New("vector lengths do not match")
	}
	var (
		length = len(vecA)
		dot    float64
	)
	for i := 0; i < length; i++ {
		dot += vecA[i] * vecB[i]
	}

	return dot, nil
}

func Norm(vec []float64) (float64, error) {
	var norm float64
	length := len(vec)
	for i := 0; i < length; i++ {
		norm += vec[i] * vec[i]
	}
	norm = math.Sqrt(norm)
	return norm, nil
}
