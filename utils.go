package main

import "fmt"

func GetName() string {
	name := "" // or var name string

	fmt.Println("Welcome to Tim's casino...")
	fmt.Println("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s let's play! \n ", name)
	return name

}

func GetBet(balance uint) uint {
	var bet uint
	for {
		fmt.Printf("Enter yout bet or 0 to quit (balance = $%d) :", balance)
		if _, err := fmt.Scan(&bet); err != nil {
			fmt.Println("Invalid input .Please enter a valid number.")

			var discard string
			fmt.Scanln(&discard)
			continue

		}

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance")
		} else {
			break
		}

	}
	return bet
}
