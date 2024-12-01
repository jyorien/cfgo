package prediction

// Computes Item-Item predictions based on a given user, the entire ratings matrix, and the itemSimilarity matrix
func RecommendByItems(targetUser int, ratings, itemSimilarityMatrix map[int]map[int]float64) map[int]float64 {
	targetUserRatings := ratings[targetUser]
	predictedRatings := make(map[int]float64)
	normFactor := make(map[int]float64)
	// incrementally find the predicted scores of unrated items
	for item, rating := range targetUserRatings {
		for similarItem, similarity := range itemSimilarityMatrix[item] {
			// don't compute for items that user already rated
			if _, isExist := targetUserRatings[similarItem]; isExist {
				continue
			}
			predictedRatings[similarItem] += similarity * rating
			normFactor[similarItem] += similarity
		}
	}
	// normalize the predictions
	for item, rating := range predictedRatings {
		predictedRatings[item] = rating / normFactor[item]
	}
	return predictedRatings
}
