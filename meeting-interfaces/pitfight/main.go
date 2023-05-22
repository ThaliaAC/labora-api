package main

import (
	"fmt"
	"math/rand"

	"time"

	"github.com/ThaliaAC/labora-api/meeting-interfaces/fighters"
)

func main() {
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
			Life: 20,
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
		intensity := contenders[randomOrderOfFighter].ThrowAttack()
		randomDirectionAttack := rand.Intn(2)
		var recieverFighter string
		if randomDirectionAttack == 0 { //Atack to the right
			if randomOrderOfFighter == len(contenders)-1 {
				contenders[0].RecieveAttack(intensity)
				recieverFighter = contenders[0].GetName()
			} else if randomOrderOfFighter == len(contenders)-2 {
				contenders[2].RecieveAttack(intensity)
				recieverFighter = contenders[2].GetName()
			} else {
				contenders[1].RecieveAttack(intensity)
				recieverFighter = contenders[1].GetName()
			}
		} else { //Atack to the left
			if randomOrderOfFighter == len(contenders)-1 {
				contenders[1].RecieveAttack(intensity)
				recieverFighter = contenders[1].GetName()
			} else if randomOrderOfFighter == len(contenders)-2 {
				contenders[0].RecieveAttack(intensity)
				recieverFighter = contenders[0].GetName()
			} else {
				contenders[2].RecieveAttack(intensity)
				recieverFighter = contenders[2].GetName()
			}
		}

		fmt.Println(contenders[randomOrderOfFighter].GetName(), "throw punch with intensity", intensity, "to", recieverFighter)
		fmt.Printf("PoliceLife = %d, CriminalLife = %d, PaladinLife = %d\n", police.Life, criminal.Life, paladin.Life)
		areAllAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
		time.Sleep(1 * time.Second)
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
