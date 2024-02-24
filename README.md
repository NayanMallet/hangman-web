# Hangman Game

## Introduction
Hangman is a word guessing game where the player must guess a secret word letter by letter within a certain number of attempts.

![image](https://github.com/NayanMallet/hangman-web/assets/81246812/7353283d-0934-49e4-a76b-8e67dbed145c)

This project implements a Hangman game using Golang with a RESTful API to facilitate communication between the client and server sides. It includes features such as updating high scores, responsive web design, and more.

![image](https://github.com/NayanMallet/hangman-web/assets/81246812/d23a6dbe-8504-49d0-909c-49b0e7ebe7ab)

## Game Rules 📝
- The player has 10 lives.
- The player must guess the word by either guessing a letter or the entire word.
- If the player guesses the entire word correctly, they win. Otherwise, they lose two lives.
- If the player guesses a letter incorrectly, they lose a life.
- If the player loses all their lives, they lose the game.
- If the player correctly guesses all the letters of the word, they win.

## How to Play 🎮
1. Clone the repository:
   ```bash
   git clone https://github.com/NayanMallet/hangman-web.git
   ```
2. Navigate to the project directory:
   ```bash
    cd hangman-web
   ```
3. Run the game using Go:
   ```bash
    go run .
   ```
4. Open a web browser and navigate to [http://localhost:8080](http://localhost:8080).
5. Enjoy playing Hangman!
