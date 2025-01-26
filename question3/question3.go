package question3

import "strings"

func PieFireDire(text string) map[string]interface{} {
	words := strings.Fields(text)

	meatCount := make(map[string]int)

	for _, word := range words {
		word = strings.Trim(word, ",. ")
		word = strings.ToLower(word)
		if word != "" {
			meatCount[word]++
		}
	}

	result := make(map[string]interface{})
	for key, value := range meatCount {
		result[key] = value
	}

	return result
}
