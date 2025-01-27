package main

import (
	"fmt"
)

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}
	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {

	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}
	symbolArr := GenerateSymbolArray(symbols)

	balance := uint(200) // or var balance uint=20
	GetName()

	for balance > 0 {
		bet := GetBet((balance))
		if bet == 0 {
			break
		}
		balance -= bet //balance = balance-bet
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		//check win and update the balance
		winningLines := checkWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("won $%d , (%dx) on line #%d\n ", win, multi, i+1)
			}

		}

	}
	fmt.Printf("you left with,$%d \n", balance)

}
