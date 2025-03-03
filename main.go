package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
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

func getBet(balance uint) uint {
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

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func getRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf(symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}

		}
		fmt.Println() // yeni satıra geçmek için
	}
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}
	//Kod, her kolon için selected kullanarak aynı elemanı tekrar seçmeyi engelliyor.
	//Ama satırlar için değil, kolon için benzersiz eleman seçiyo
	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for {
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

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
	symbolArr := generateSymbolArray(symbols)

	balance := uint(200) // or => var balance uint = 200
	getName()
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := getSpin(symbolArr, 3, 3)
		printSpin(spin)
		winningLines := checkWin(spin, multipliers)
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
