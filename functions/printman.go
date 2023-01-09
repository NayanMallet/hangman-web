package functions

func PrintHangMan(Lives int) string {
	var Hangman []string
	switch Lives {
	case 10:
		Hangman = []string{
			"",
			"           ",
			"           ",
			"           ",
			"           ",
			"           ",
			"           ",
			"           ",
			"_______    ",
		}
	case 9:
		Hangman = []string{
			"",
			"             ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 8:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/        ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 7:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 6:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |      O  ",
			"   |         ",
			"   |         ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 5:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |      O  ",
			"   |      |  ",
			"   |         ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 4:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |      O  ",
			"   |     /|  ",
			"   |         ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 3:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |      O  ",
			"   |     /|\\",
			"   |         ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 2:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |      O  ",
			"   |     /|\\",
			"   |     /   ",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 1:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |      O  ",
			"   |     /|\\",
			"   |     / \\",
			"   |         ",
			"   |         ",
			"___|___      ",
		}
	case 0:
		Hangman = []string{
			"",
			"   _______   ",
			"   |/     |  ",
			"   |      O  ",
			"   |         ",
			"   |         ",
			"   |     /|\\",
			"   |     / \\",
			"___|___      ",
		}
	}
	var HangmanString string
	for _, v := range Hangman {
		HangmanString += v + "\n"
	}
	return HangmanString
}
