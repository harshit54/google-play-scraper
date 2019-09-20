package search

import (
	"github.com/n0madic/google-play-scraper/pkg/app"
)

// Results of search
type Results []*app.App

// Append result
func (results *Results) Append(res ...app.App) {
	for _, result := range res {
		if !results.searchDuplicate(result.ID) {
			results.append(result)
		}
	}
}

func (results *Results) append(result app.App) {
	*results = append(*results, &result)
}

func (results *Results) searchDuplicate(id string) bool {
	for _, result := range *results {
		if id == result.ID {
			return true
		}
	}
	return false
}
