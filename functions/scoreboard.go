package functions

func Points(Data Infos) int {
	if Data.Difficulty == "easy" {
		Data.Points += Data.Lives
	} else if Data.Difficulty == "medium" {
		Data.Points += Data.Lives * 2
	}
	if Data.Difficulty == "hard" {
		Data.Points += Data.Lives * 3
	}

	return Data.Points
}
