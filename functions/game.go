package functions

import (
	game_function "hangman-classic/g-func"
)

func WordToPrint(runeArray []rune) string {
	return string(runeArray)
}

func NewGamePrep() (Word string, WordRune []rune) {
	Word, WordRune = game_function.NewGamePrep([]string{"words.txt"})
	return Word, WordRune
}

func Game(word string, wordrune []rune, proposition string, letterssuggested []string, lives int) {
	if lives > 0 {
		if string(wordrune) == word {
			// TODO: Print Congrats
			return
		}
		if len(proposition) > 1 {
			// case Mot
			if proposition == word {
				// TODO: Print Congrats
				return
			} else {
				lives -= 2
				// TODO: Print lives lefts & hangman
			}
		} else {
			// case Letter
			if game_function.ContainsTable(letterssuggested, proposition) {
				// TODO: Letter already proposed
			} else {
				if game_function.ContainsString(word, proposition) {
					letterssuggested = append(letterssuggested, proposition)
					indexes := game_function.LetterInWorld(word, proposition)
					for _, i := range indexes {
						wordrune[i] = rune(word[i])
					}
					// TODO: Print Proposition + WordToPrint
				} else {
					lives--
					// TODO: Print lives lefts & hangman
					//PrintMan(attemps)
				}
			}
		}
	}
	// TODO: Print Loose + The Word
}
