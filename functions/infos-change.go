package functions

import "strconv"

func WordToPrint(runeArray []rune) string {
	return string(runeArray)
}

func HangmanStepLink(Lives int) string {
	return "/static/images/hangman-steps/" + strconv.Itoa(11-Lives) + ".png/"
}
