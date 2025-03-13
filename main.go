package main

import (
	"fmt"

	games "github.com/AlexSilverson/MRPO_1lab/Games"
)

func main() {
	fmt.Println("Welcome to the Brain Games!\nMay I have your name? ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hellow " + name)
	for {
		fmt.Println("Which game would you like to play?")
		fmt.Println("1 - lcm")
		fmt.Println("2 - progression")
		fmt.Println("-1 - leave")
		var choice string
		var game games.Game
		fmt.Scan(&choice)
		switch choice {
		case "1":
			game = games.LcmGame()
		case "2":
			game = games.ProgressionGame()
		case "-1":
			{
				fmt.Println("Goodbye " + name + "!")
				return
			}
		default:
			{
				fmt.Println("Invalid input.")
				continue
			}
		}
		fmt.Println(game.GetRules())
		game.Generate()
		fmt.Print("Question: ")
		for _, v := range game.GetQuestion() {
			fmt.Print(v, " ")
		}
		fmt.Println()
		fmt.Println("Your answer:")
		var userAns int
		fmt.Scan(&userAns)
		if userAns == game.GetAnswer() {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\n", userAns, game.GetAnswer())
			fmt.Printf("Let's try again, %s!\n", name)
		}
	}
}
