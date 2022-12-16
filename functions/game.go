package functions

import (
	game_function "hangman-classic/g-func"
)

func Game(Data Infos) Infos {
	if Data.Lives > 1 {
		if string(Data.WordRune) == Data.Word {
			Data.Win = true
			return Data
		}
		if len(Data.Propositon) > 1 {
			// case Mot
			if Data.Propositon == Data.Word {
				Data.Win = true
				return Data
			} else {
				// TODO: Print lives lefts & hangman
				Data.Lives -= 2
				Data.Url = PrintMan(Data.Lives)
				return Data
			}
		} else {
			// case Letter
			if game_function.ContainsTable(Data.LetterSuggested, Data.Propositon) {
				// TODO: Letter already proposed
			} else {
				if game_function.ContainsString(Data.Word, Data.Propositon) {
					Data.LetterSuggested = append(Data.LetterSuggested, Data.Propositon)
					indexes := game_function.LetterInWorld(Data.Word, Data.Propositon)
					for _, i := range indexes {
						Data.WordRune[i] = rune(Data.Word[i])
					}
					// TODO: Print Proposition + WordToPrint
					Data.WordToPrint = WordToPrint(Data.WordRune)
				} else {
					Data.Lives--
					Data.Url = PrintMan(Data.Lives)
					// TODO: Print lives lefts & hangman
					return Data
				}
			}
		}
		return Data
	} else {
		if string(Data.WordRune) == Data.Word {
			Data.Win = true
			return Data
		}
		if len(Data.Propositon) > 1 {
			// case Mot
			if Data.Propositon == Data.Word {
				Data.Win = true
				return Data
			} else {
				// TODO: Print Loose + The Word
				Data.Lives -= 2
				Data.Win = false
				return Data
			}
		} else {
			// case Letter
			if game_function.ContainsTable(Data.LetterSuggested, Data.Propositon) {
				// TODO: Letter already proposed
			} else {
				if game_function.ContainsString(Data.Word, Data.Propositon) {
					Data.LetterSuggested = append(Data.LetterSuggested, Data.Propositon)
					indexes := game_function.LetterInWorld(Data.Word, Data.Propositon)
					for _, i := range indexes {
						Data.WordRune[i] = rune(Data.Word[i])
					}
					// TODO: Print Proposition + WordToPrint
					Data.WordToPrint = WordToPrint(Data.WordRune)
				} else {
					// TODO: Print Loose + The Word
					Data.Lives--
					Data.Win = false
					return Data
				}
			}
		}
		return Data
	}
}
