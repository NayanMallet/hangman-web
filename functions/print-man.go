package functions

import (
	"strconv"
)

func PrintMan(lives int) (url string) {
	url = "http://localhost:8080/static/hangman-steps/" + strconv.Itoa(lives) + ".png"
	return url
}
