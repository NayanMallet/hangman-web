package functions

import game_function "hangman-classic/g-func"

type Infos struct {
	Word            string
	WordRune        []rune
	WordToPrint     string
	Propositon      string
	LetterSuggested []string
	Lives           int
	Url             string
}

func NewGamePrep() (StartData Infos) {
	Word, WordRune := game_function.NewGamePrep([]string{"words.txt"})
	StartData = Infos{
		Word:            Word,
		WordRune:        WordRune,
		WordToPrint:     WordToPrint(WordRune),
		Propositon:      "",
		LetterSuggested: nil,
		Lives:           10,
		Url:             PrintMan(10),
	}
	return StartData
}
