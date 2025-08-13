package main

import (
	"fightrobots/clearterminal"
	"fightrobots/robot"
	"fmt"
	"os"
)

func main() {
	app()
}

func app() {
	var listRobots []robot.Robot

	for {
		var input string
		fmt.Print("1.Начать игру\n2.Закрыть приложение\n")
		fmt.Scanln(&input)

		switch input {
		case "1":
			clearterminal.Clear()
			listRobots = robot.GenerateRobots()
			clearterminal.Clear()
			robot.Arena(&listRobots)

		case "2":
			clearterminal.Clear()
			os.Exit(0)
		default:
			clearterminal.Clear()
			continue
		}

	}

}
