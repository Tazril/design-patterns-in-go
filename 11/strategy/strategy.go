package strategy

import (
	"strings"
)

type SearchStrategy interface {
	Search(users []string, query string) []string
}

type UserService struct {
	users      []string
	strategies []SearchStrategy
}

type PrefixSearchStrategy struct {
}

func (s PrefixSearchStrategy) Search(users []string, query string) []string {
	var searched []string
	for _, user := range users {
		if strings.HasPrefix(user, query) {
			searched = append(searched, user)
		}
	}
	return searched
}

type SuffixSearchStrategy struct {
}

func (s SuffixSearchStrategy) Search(users []string, query string) []string {
	var searched []string
	for _, user := range users {
		if strings.HasSuffix(user, query) {
			searched = append(searched, user)
		}
	}
	return searched
}

func (s *UserService) SearchQuery(query string) []string {
	var searched []string
	present := map[string]bool{}
	for _, strategy := range s.strategies {
		results := strategy.Search(s.users, query)
		for _, r := range results {
			if !present[r] { // avoid duplicates
				searched = append(searched, r)
				present[r] = true
			}
		}
	}
	return searched
}
