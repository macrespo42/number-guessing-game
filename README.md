# [number-guessing-game](https://roadmap.sh/projects/number-guessing-game)

A simple number guessing game to test your luck.

```bash
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50.

Enter your guess: 25
Incorrect! The number is greater than 25.

Enter your guess: 35
Incorrect! The number is less than 35.

Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
```

## Requirements

go installed

## installation

To install or run locally, clone the repo and make the script executable:

```bash
git clone https://github.com/macrespo42/number-guessing-game
cd number-guessing-game
go run .

```

or to installed it globally

```bash
go install https://github.com/macrespo42/number-guessing-game@latest
# if $GOPATH isn't in your .zshrc/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin > ~/.zshrc
source ~/.zshrc
# Then run the program from anywhere
number-guessing-game
```
