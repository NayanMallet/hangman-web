package functions

import (
	gamefunction "github.com/NayanMallet/hangman-classic/g-func"
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
				Data.Hangman = HangmanStepLink(Data.Lives)
			}
		} else {
			// case Letter
			if !(gamefunction.ContainsTable(Data.LetterSuggested, Data.Propositon)) {
				Data.LetterSuggested = append(Data.LetterSuggested, Data.Propositon)
				Data.LetterSuggested = SortLetterSuggested(Data.LetterSuggested)
				if gamefunction.ContainsString(Data.Word, Data.Propositon) {
					indexes := gamefunction.LetterInWorld(Data.Word, Data.Propositon)
					for _, i := range indexes {
						Data.WordRune[i] = rune(Data.Word[i])
					}
					Data.WordToPrint = WordToPrint(Data.WordRune)
				} else {
					Data.Lives--
					Data.Hangman = HangmanStepLink(Data.Lives)
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
				Data.Hangman = HangmanStepLink(Data.Lives)
			}
		} else {
			// case Letter
			if !(gamefunction.ContainsTable(Data.LetterSuggested, Data.Propositon)) {
				Data.LetterSuggested = append(Data.LetterSuggested, Data.Propositon)
				Data.LetterSuggested = SortLetterSuggested(Data.LetterSuggested)
				if gamefunction.ContainsString(Data.Word, Data.Propositon) {
					Data.LetterSuggested = append(Data.LetterSuggested, Data.Propositon)
					indexes := gamefunction.LetterInWorld(Data.Word, Data.Propositon)
					for _, i := range indexes {
						Data.WordRune[i] = rune(Data.Word[i])
					}
					Data.WordToPrint = WordToPrint(Data.WordRune)
				} else {
					Data.Lives--
					Data.Hangman = HangmanStepLink(Data.Lives)
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
