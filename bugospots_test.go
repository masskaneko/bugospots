package main

import "testing"

func TestIsFixMessage(t *testing.T) {
	fmj := NewFixMessageJudger("(?i)(^| )(fi(x|xed|xes)|clos(e|es|ed))")

	t.Run("TrueCases", func(t *testing.T) {
		trueStrings := []string{
			"fix", "fixed", "fixes",
			"close", "closed", "closes",
			"fix #123", "fixed #123", "fixes #123",
			"close #123", "closed #123", "closes #123",
			"bug fix", "bug close"}
		expected := true
		for _, s := range trueStrings {
			actual := fmj.IsFixMessage(s)
			if expected != actual {
				t.Log("[Expected]", expected, "[Actual]", actual, "[Input]", s)
			}
		}
	})

	t.Run("FalseCases", func(t *testing.T) {
		trueStrings := []string{"f i x", "ｆ ｉ ｘ", "suffix", "disclose"}
		expected := false
		for _, s := range trueStrings {
			actual := fmj.IsFixMessage(s)
			if expected != actual {
				t.Log("[Expected]", expected, "[Actual]", actual, "[Input]", s)
			}
		}
	})
}
