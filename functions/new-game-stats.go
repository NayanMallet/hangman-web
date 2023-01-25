package functions

import gamefunction "github.com/NayanMallet/hangman-classic/g-func"

type Infos struct {
	Word            string
	WordRune        []rune
	WordToPrint     string
	Propositon      string
	LetterSuggested []string
	Lives           int
	Name            string
	Difficulty      string
	Points          int
	Hangman         string
	Scores          []ScoreInfos
	Win             bool
}

func NewGamePrep(Difficulty string) (StartData Infos) {
	Word := ""
	var WordRune []rune
	switch Difficulty {
	case "easy":
		Word, WordRune = gamefunction.NewGamePrep([]string{"words.txt"})
	case "medium":
		Word, WordRune = gamefunction.NewGamePrep([]string{"words2.txt"})
	case "hard":
		Word, WordRune = gamefunction.NewGamePrep([]string{"words3.txt"})
	}

	StartData = Infos{
		Word:            Word,
		WordRune:        WordRune,
		WordToPrint:     WordToPrint(WordRune),
		Propositon:      "",
		LetterSuggested: []string{},
		Lives:           10,
		Name:            "",
		Difficulty:      Difficulty,
		Points:          0,
		Hangman:         HangmanStepLink(10),
		Win:             false,
	}
	return StartData
}
