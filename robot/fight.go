package robot

import (
	"fightrobots/clearterminal"
	"fmt"
	"os"
	"time"
)

func Arena(listRobots *[]Robot) {

	for {
		var input string
		fmt.Print("1.Показать список участников\n2.Начать бои между роботами\n3.Выйти из игры\n4.Завершить приложение\n")
		fmt.Scanln(&input)

		switch input {
		case "1":
			input = ""
			clearterminal.Clear()
			printNameRobots(listRobots)
		case "2":
			input = ""
			clearterminal.Clear()
			champ, ok := ChampionExists(listRobots)
			if ok {

				fmt.Printf("Чемпион определен и это - %s\n", champ.Name)
				time.Sleep(2 * time.Second)
				clearterminal.Clear()
				continue
			}
			StartFight(listRobots)
		case "3":
			input = ""
			clearterminal.Clear()
			return
		case "4":
			clearterminal.Clear()
			os.Exit(0)

		default:
			clearterminal.Clear()
			continue
		}
	}

}

func StartFight(listRobots *[]Robot) {

	for {
		robotOne, robotTwo, champion := ChosePair(listRobots)

		if champion != nil {
			fmt.Printf("Определен победитель - %s\n", champion.Name)
			Pause()
			clearterminal.Clear()
			break
		}
		if robotOne == nil && robotTwo == nil {
			fmt.Println("Нет живых роботов")
			Pause()
			clearterminal.Clear()
			break
		} else {
			winner := Fight(robotOne, robotTwo)

			fmt.Printf("Определен победитель - %s\n", winner.Name)
			Pause()
			clearterminal.Clear()
			break
		}

	}

}

func ChosePair(listRobots *[]Robot) (robotOne, robotTwo *Robot, champion *Robot) {
	var first, second *Robot

	for i := range *listRobots {
		robot := &(*listRobots)[i]
		if robot.Alive {
			if first == nil {
				first = robot
			} else {
				second = robot
				break
			}
		}

	}

	if first != nil && second != nil {
		return first, second, nil
	}

	if first != nil && second == nil {
		return nil, nil, first
	}

	return nil, nil, nil

}

func Fight(one, two *Robot) *Robot {
	fmt.Printf("Начинается бой %s против %s", one.Name, two.Name)
	Pause()

	count := 1
	for {
		clearterminal.Clear()
		if !one.Alive {
			return two
		} else if !two.Alive {
			return one
		}

		switch count % 2 {
		case 1:
			one.Shot(two)
			count++
		case 0:
			two.Shot(one)
			count++
		}

	}

}

func printNameRobots(listRobots *[]Robot) {
	fmt.Println("Перед вами список участников:")
	for index, robot := range *listRobots {

		fmt.Printf("\nРобот под номером %d. \nИмя робота %s. Статус функциональности %v\n\n", index+1, robot.Name, robot.Alive)

	}
	fmt.Println("Нажмите Enter, чтобы продолжить.")
	fmt.Scanln()
	clearterminal.Clear()
}

func ChampionExists(listRobots *[]Robot) (Robot, bool) {
	var champion []Robot
	for _, robot := range *listRobots {
		if robot.Alive {
			champion = append(champion, robot)
		}
	}

	if len(champion) == 1 {
		return champion[0], true
	}
	return Robot{}, false
}

func Pause() {

	time.Sleep(2 * time.Second)
}
