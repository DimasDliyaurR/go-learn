package main

import "fmt"

const (
	umur   int32   = 12
	hobi   string  = "Makan"
	income float32 = 20000
)

func main() {
	type health int32
	type damage int32
	type lari bool
	type defense bool

	// Warrior
	var (
		WarriorHealth  health  = -12000
		WarriorDamage  damage  = 1300
		WarriorLari    lari    = false
		WarriorDefense defense = false
	)

	// Enemy
	var (
		EnemyHealth  health  = 300
		EnemyDamage  damage  = 100
		EnemyLari    lari    = false
		EnemyDefense defense = false
	)

	/*
		Warrior State
	*/
	fmt.Printf("Health Warrior = %d\n", WarriorHealth)
	fmt.Printf("Attach Warrior = %d\n", WarriorDamage)
	fmt.Printf("Warrior escape = %t\n", WarriorLari)
	fmt.Printf("Warrior Defense = %t\n", WarriorDefense)

	/*
		Enemy State
	*/
	fmt.Printf("Health Enemy = %d\n", EnemyHealth)
	fmt.Printf("Attach Enemy = %d\n", EnemyDamage)
	fmt.Printf("Enemy escape = %t\n", EnemyLari)
	fmt.Printf("Enemy Defense = %t\n", EnemyDefense)

}
