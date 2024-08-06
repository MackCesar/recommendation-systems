package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Review struct {
	UserID string  `json:"user_id"`
	ASIN   string  `json:"asin"`
	Rating float64 `json:"rating"`
}

var reviews []Review

// LoadReviews loads reviews from a JSON Lines file
func loadReviews(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var review Review
		if err := json.Unmarshal(scanner.Bytes(), &review); err != nil {
			return fmt.Errorf("error unmarshalling review: %v", err)
		}
		reviews = append(reviews, review)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	return nil
}

func GetRecommendations(asin string) []string {
	similarityScores := make(map[string]float64)

	for _, review := range reviews {
		if review.ASIN == asin {
			for _, otherReview := range reviews {
				if review.UserID == otherReview.UserID && otherReview.ASIN != asin {
					similarityScores[otherReview.ASIN] += otherReview.Rating
				}
			}
		}
	}
	type itemScore struct {
		ASIN  string
		Score float64
	}
	var sortedItems []itemScore
	for asin, score := range similarityScores {
		sortedItems = append(sortedItems, itemScore{asin, score})
	}
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Score > sortedItems[j].Score
	})
	var recommendations []string
	for i := 0; i < 5 && i < len(sortedItems); i++ {
		recommendations = append(recommendations, sortedItems[i].ASIN)
	}
	return recommendations
}
func recommendHandler(w http.ResponseWriter, r *http.Request) {
	asin := strings.TrimPrefix(r.URL.Path, "/recommend/")

	recommendations := GetRecommendations(asin)
	response, _ := json.Marshal(map[string][]string{"recommendations": recommendations})
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	filePath := flag.String("file", "data/Video_Games.jsonl", "Path to the reviews JSONL file")
	flag.Parse()

	if err := loadReviews(*filePath); err != nil {
		log.Fatalf("Error loading reviews.json: %v", err)
	}
	http.HandleFunc("/recommend/", recommendHandler)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
