package similarity

import (
	"fmt"
)

func GenerateItemItemSimilarityMatrix(ratings map[int]map[int]float64, numItems int) map[int]map[int]float64 {
	similarityMatrix := make(map[int]map[int]float64)
	for i := 0; i < numItems; i++ {
		if similarityMatrix[i] == nil {
			similarityMatrix[i] = make(map[int]float64)
		}

		for j := i + 1; j < numItems; j++ {
			if i == j {
				continue
			}

			vec1 := getItemRatingVector(ratings, i)
			vec2 := getItemRatingVector(ratings, j)
			sim, err := CosineSimilarity(vec1, vec2)
			if err != nil {
				fmt.Println(err)
				break
			}
			similarityMatrix[i][j] = sim
			if similarityMatrix[j] == nil {
				similarityMatrix[j] = make(map[int]float64)
			}
			similarityMatrix[j][i] = sim
		}
	}

	return similarityMatrix
}

func getItemRatingVector(ratings map[int]map[int]float64, itemIndex int) []float64 {
	// num users -> number of key-value pairs in map
	numUsers := len(ratings)
	itemRatings := make([]float64, numUsers)
	i := 0
	for _, userRatings := range ratings {
		// if user has rated the given item
		if rating, ok := userRatings[itemIndex]; ok {
			itemRatings[i] = rating
		} else {
			// 0 assigned as missing value
			itemRatings[i] = 0
		}
		i += 1
	}
	return itemRatings
}
