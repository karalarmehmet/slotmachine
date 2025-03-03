package main

import (
	"fmt"
)

func CheckWin(spin [][]string, multipliers map[string]uint) []uint {
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
		"🍒": 4,
		"🍋": 7,
		"🔔": 12,
		"🍎": 20,
	}

	multipliers := map[string]uint{
		"🍒": 20,
		"🍋": 10,
		"🔔": 5,
		"🍎": 2,
	}
	symbolArr := GenerateSymbolArray(symbols)

	balance := uint(200) // or => var balance uint = 200
	GetName()
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		winningLines := CheckWin(spin, multipliers)
		fmt.Println(winningLines)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("Won $%d ,(%dx) on Line #%d\n ", win, multi, i+1)
			}

		}
	}

	fmt.Printf("You left with $%d \n.", balance)

}
