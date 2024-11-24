package similarity

import (
	"errors"

	"github.com/jyorien/cfgo/vectors"
)

func CosineSimilarity(vecA, vecB []float64) (float64, error) {
	var (
		similarity float64
		dot        float64
		normA      float64
		normB      float64
	)
	dot, err := vectors.DotProduct(vecA, vecB)
	if err != nil {
		return 0, err
	}
	normA, _ = vectors.Norm(vecA)
	normB, _ = vectors.Norm(vecB)

	if normA == 0 || normB == 0 {
		return 0, errors.New("norm cannot be 0")
	}
	similarity = dot / (normA * normB)
	return similarity, nil
}
