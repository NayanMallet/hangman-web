package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Points(Data Infos) int {
	// calculate the points of the user by the difficulty
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

type ScoreInfos struct {
	Name   string
	Points int
}

func Save(Data Infos) {
	// verify if the user already play, if yes and if the score is better, the func replace the old score
	var Score = ScoreInfos{
		Name:   Data.Name,
		Points: Data.Points,
	}
	var Scores []ScoreInfos
	var NewScores []ScoreInfos
	var Exist bool
	Exist = false
	// Read file
	file, err := ioutil.ReadFile("score-board.txt")
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal
	err = json.Unmarshal(file, &Scores)
	if err != nil {
		fmt.Println(err)
	}
	// Check if Name already exist
	for _, s := range Scores {
		if s.Name == Score.Name {
			Exist = true
			if s.Points < Score.Points {
				NewScores = append(NewScores, Score)
			} else {
				NewScores = append(NewScores, s)
			}
		} else {
			NewScores = append(NewScores, s)
		}
	}
	if Exist == false {
		NewScores = append(NewScores, Score)
	}
	// Marshal
	file, err = json.Marshal(NewScores)
	if err != nil {
		fmt.Println(err)
	}
	// Write file
	err = ioutil.WriteFile("score-board.txt", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadScoreBoard() []ScoreInfos {
	// read the file and return the scoreboard in descendent order
	var Scores []ScoreInfos
	// Read file
	file, err := ioutil.ReadFile("score-board.txt")
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal
	err = json.Unmarshal(file, &Scores)
	if err != nil {
		fmt.Println(err)
	}
	// Sort
	for i := 0; i < len(Scores); i++ {
		for j := i + 1; j < len(Scores); j++ {
			if Scores[i].Points < Scores[j].Points {
				Scores[i], Scores[j] = Scores[j], Scores[i]
			}
		}
	}

	// return only first five scores
	if len(Scores) > 5 {
		return Scores[:5]
	} else {
		return Scores
	}
}
