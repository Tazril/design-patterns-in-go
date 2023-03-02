package strategy

import (
	"reflect"
	"testing"
)

func TestSearch_SearchQuery(t *testing.T) {
	// given
	ps := PrefixSearchStrategy{}
	ss := SuffixSearchStrategy{}
	us := UserService{users: []string{"jane", "jack", "mike", "raja"}, strategies: []SearchStrategy{&ps, &ss}}
	// when
	results := us.SearchQuery("ja")
	// then
	expected := []string{"jane", "jack", "raja"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Got results= %v Expected %v", results, expected)
	}
}
