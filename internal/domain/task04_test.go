package domain

import "testing"

func TestRuleCategoryMatchingNormalizesInput(t *testing.T) {
	rule := &Rule{MatchCategory: "WAGE"}
	item := &RightsCase{Category: " wage "}
	if !rule.Matches(item) {
		t.Fatal("category matching ignored case or surrounding whitespace")
	}
}
