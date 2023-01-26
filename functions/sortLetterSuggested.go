package functions

func SortLetterSuggested(LetterSuggested []string) []string {
	// Sort the letter suggested
	for i := 0; i < len(LetterSuggested); i++ {
		for j := i + 1; j < len(LetterSuggested); j++ {
			if LetterSuggested[i] > LetterSuggested[j] {
				LetterSuggested[i], LetterSuggested[j] = LetterSuggested[j], LetterSuggested[i]
			}
		}
	}
	return LetterSuggested
}
