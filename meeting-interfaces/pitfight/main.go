package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"time"

	"github.com/ThaliaAC/labora-api/meeting-interfaces/fighters"
)

const col = 20

func drawBarUntil(cond func(t int) bool, valueEnteredSignal chan int) int {
	bar := fmt.Sprintf("[%%-%vs]", col)
	var t int = 0
	for cond(t) {
		fmt.Print("\033[H\033[2J")
		fmt.Printf(bar, strings.Repeat("=", t%col)+"🤜🏼")

		time.Sleep(20 * time.Millisecond)

		t++
	}
	valueEnteredSignal <- t % col
	return t % col
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var valueEnteredSignal chan int = make(chan int)

	var police fighters.Police = fighters.Police{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
		Armour: 5,
	}
	var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}
	var paladin fighters.Paladin = fighters.Paladin{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}

	/*var contenders []fighters.Contender = make([]fighters.Contender, 2)

	randomValueBetweenOneAndZero := rand.Intn(2)
	contenders[randomValueBetweenOneAndZero] = &police
	contenders[(randomValueBetweenOneAndZero+1)%2] = &criminal*/

	contenders := []fighters.Contender{&police, &criminal, &paladin}
	rand.Shuffle(len(contenders), func(i, j int) {
		contenders[i], contenders[j] = contenders[j], contenders[i]
	})

	var areAllAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
	for areAllAlive {
		randomOrderOfFighter := rand.Intn(3)
		var intensity float64 = float64(contenders[randomOrderOfFighter].ThrowAttack())

		wasEnterPressed := false
		go drawBarUntil(func(t int) bool {
			return !wasEnterPressed
		}, valueEnteredSignal)

		scanner.Scan()
		wasEnterPressed = true
		var value float64 = float64(<-valueEnteredSignal)

		intensity = intensity * (value / col)

		randomDirectionAttack := rand.Intn(2)
		var recieverFighter string
		if randomDirectionAttack == 0 { //Atack to the right
			if randomOrderOfFighter == len(contenders)-1 {
				contenders[0].RecieveAttack(int(intensity))
				recieverFighter = contenders[0].GetName()
			} else if randomOrderOfFighter == len(contenders)-2 {
				contenders[2].RecieveAttack(int(intensity))
				recieverFighter = contenders[2].GetName()
			} else {
				contenders[1].RecieveAttack(int(intensity))
				recieverFighter = contenders[1].GetName()
			}
		} else { //Atack to the left
			if randomOrderOfFighter == len(contenders)-1 {
				contenders[1].RecieveAttack(int(intensity))
				recieverFighter = contenders[1].GetName()
			} else if randomOrderOfFighter == len(contenders)-2 {
				contenders[0].RecieveAttack(int(intensity))
				recieverFighter = contenders[0].GetName()
			} else {
				contenders[2].RecieveAttack(int(intensity))
				recieverFighter = contenders[2].GetName()
			}
		}

		fmt.Printf("%s throw a punch with intensity %.2f to %s\n", contenders[randomOrderOfFighter].GetName(), intensity, recieverFighter)
		fmt.Printf("PoliceLife = %d, CriminalLife = %d, PaladinLife = %d\n", police.Life, criminal.Life, paladin.Life)
		areAllAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
		time.Sleep(2 * time.Second)
	}
}

/*func main_legacy() {

	var police fighters.Police = fighters.Police{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
		Armour: 5,
	}
	var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 10,
		},
	}

	randomValueBetweenOneAndZero := rand.Intn(2)
	policeHitFirst := randomValueBetweenOneAndZero == 1

	var areBothAlive = police.IsAlive() && criminal.IsAlive()
	for areBothAlive {
		if policeHitFirst {
			intesity := police.ThrowAttack()
			fmt.Println("Policia tira golpe con intensidad =", intesity)
			criminal.RecieveAttack(intesity)

			if criminal.IsAlive() {
				intesity := criminal.ThrowAttack()
				fmt.Println("Criminal tira golpe con intensidad =", intesity)
				police.RecieveAttack(intesity)
			}
		} else {
			intesity := criminal.ThrowAttack()
			fmt.Println("Criminal tira golpe con intensidad =", intesity)
			police.RecieveAttack(intesity)

			if police.IsAlive() {
				intesity := police.ThrowAttack()
				fmt.Println("Policia tira golpe con intensidad =", intesity)
				criminal.RecieveAttack(intesity)
			}
		}
		fmt.Printf("PoliceLife=%d, CriminalLife=%d\n", police.Life, criminal.Life)
		areBothAlive = police.IsAlive() && criminal.IsAlive()
		time.Sleep(3 * time.Second)
	}
}*/
