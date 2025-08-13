package robot

import (
	"fightrobots/clearterminal"
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit/v7"
)

type Actions interface {
	Shot()
	CookBorscht()
	ShowIndicators()
}

type Robot struct {
	Name          string
	AmountBorscht int
	HP            int
	Force         int
	Alive         bool
}

func (r *Robot) Shot(opponent *Robot) {
	// attackeName := r.Name
	if r.AmountBorscht < 5 {
		fmt.Printf("У %s не хватает борща. Сейчас будет попытка сварить борщ \n", r.Name)
		Pause()
		newBorsch, explosion := AttamtCoocBorsch(r)
		if explosion {
			fmt.Printf("К сожалению для %s, кастрюля взорвалась при попытке сварить борщ. Он погибает\n", r.Name)
			Pause()
			return
		}
		if !newBorsch {
			fmt.Printf("У %s не получилось сварить борщ. Он пропускает ход \n", r.Name)
			Pause()
			return
		} else if newBorsch {
			fmt.Printf("У %s получилось сварить борщ. Он пропускает ход \n", r.Name)
			Pause()
			return
		}

	}

	fmt.Printf("%s кидет тарелку борща в  %s\n", r.Name, opponent.Name)
	Pause()
	evasion := Random()
	if evasion {
		fmt.Printf("%s уклоняется от тарелки борща\n", opponent.Name)
		Pause()
		r.AmountBorscht -= 5

		return
	} else if !evasion {
		fmt.Printf("Тарелка борща попадает в %s\n", opponent.Name)
		r.AmountBorscht -= 5
		opponent.HP -= r.Force
		Pause()
		return
	}

}

func AttamtCoocBorsch(r *Robot) (bool, bool) {
	explosion := gofakeit.IntN(100)
	if explosion < 30 {

		r.Alive = false

		return false, true
	}

	if Random() {
		r.AmountBorscht = 20
		return true, false
	}
	return false, false

}

func Random() bool {
	random := gofakeit.IntN(100)
	if random > 50 {
		return true
	} else if random < 50 {
		return false
	} else {
		return true
	}

}

func GenerateRobots() []Robot {
	var value int
	robots := []Robot{}
	for {
		var input string
		fmt.Println("Выберете четное кол-во роботов которые будут участвовать от 2 до 6")
		fmt.Scanln(&input)
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введите число")
			clearterminal.Clear()
			continue
		}
		if val < 2 || val > 6 || val%2 != 0 {
			clearterminal.Clear()
			continue
		}
		value = val
		break

	}

	for i := 0; i < value; i++ {
		robot := NewRobot()

		robots = append(robots, robot)
	}
	return robots
}

func NewRobot() Robot {
	r := Robot{
		Name:          gofakeit.SongArtist(),
		AmountBorscht: 20,
		HP:            gofakeit.Number(80, 100),
		Force:         gofakeit.Number(10, 30),
		Alive:         true,
	}
	return r

}
