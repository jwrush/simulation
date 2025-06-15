package main

import (
	"fmt"
	//"simulation"
	"simulation/lib/fight"
	"simulation"
)

func main() {
	fmt.Println("THE MONSTER FIGHT")

	// Create some attacks
	bite := fight.NewAttack("Bite", "The monster bites ferociously!", 8, 2)
	claw := fight.NewAttack("Claw", "Sharp claws rake across the target!", 6, 1)
	
	// Create some monsters
	monster1 := fight.NewMonster("Goblin", 30, 3, []fight.Attack{bite})
	monster2 := fight.NewMonster("Orc", 50, 2, []fight.Attack{claw})

	//sim := simulation.NewSimulation(0)

	fmt.Printf("%s (HP: %d) vs %s (HP: %d)\n", monster1.Name, monster1.HitPoints, monster2.Name, monster2.HitPoints)

	// Example of how to use the Attack method with error handling
	if monster1.IsAlive() && monster2.IsAlive() {
		recovery, err := monster1.Attack(monster2, 0)
		if err != nil {
			fmt.Printf("Attack failed: %v\n", err)
		} else {
			fmt.Printf("Attack successful! Recovery time: %d\n", recovery)
			fmt.Printf("%s now has %d HP\n", monster2.Name, monster2.HitPoints)
		}
	}

	fmt.Println("The fight is over!")
	fmt.Println("THE END")
}