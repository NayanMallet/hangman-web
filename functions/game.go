package functions

import (
	game_function "hangman-classic/g-func"
)

func Game(Data Infos) Infos {
	if Data.Lives > 0 {
		if string(Data.WordRune) == Data.Word {
			Data.Win = true
			Data.Points = Points(Data)
			Save(Data)
			Data.Scores = ReadScoreBoard()
			return Data
		}
		if len(Data.Propositon) > 1 {
			// case Mot
			if Data.Propositon == Data.Word {
				Data.Win = true
				Data.Points = Points(Data)
				Save(Data)
				Data.Scores = ReadScoreBoard()
				return Data
			} else {
				Data.Lives -= 2
			}
		} else {
			// case Letter
			if !(game_function.ContainsTable(Data.LetterSuggested, Data.Propositon)) {
				Data.LetterSuggested = append(Data.LetterSuggested, Data.Propositon)
				Data.LetterSuggested = SortLetterSuggested(Data.LetterSuggested)
				if game_function.ContainsString(Data.Word, Data.Propositon) {
					indexes := game_function.LetterInWorld(Data.Word, Data.Propositon)
					for _, i := range indexes {
						Data.WordRune[i] = rune(Data.Word[i])
					}
					// TODO: Print Proposition + WordToPrint
					Data.WordToPrint = WordToPrint(Data.WordRune)
				} else {
					Data.Lives--
				}
			}
		}
	} else {
		if string(Data.WordRune) == Data.Word {
			Data.Win = true
			Data.Points = Points(Data)
			Save(Data)
			Data.Scores = ReadScoreBoard()
			return Data
		}
		if len(Data.Propositon) > 1 {
			// case Mot
			if Data.Propositon == Data.Word {
				Data.Win = true
				Data.Points = Points(Data)
				Save(Data)
				Data.Scores = ReadScoreBoard()
				return Data
			} else {
				Data.Lives -= 2
			}
		} else {
			// case Letter
			if !(game_function.ContainsTable(Data.LetterSuggested, Data.Propositon)) {
				Data.LetterSuggested = append(Data.LetterSuggested, Data.Propositon)
				Data.LetterSuggested = SortLetterSuggested(Data.LetterSuggested)
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
				}
			}
			if string(Data.WordRune) == Data.Word {
				Data.Win = true
				Data.Points = Points(Data)
				Save(Data)
				Data.Scores = ReadScoreBoard()
				return Data
			}
		}
	}
	return Data
}
