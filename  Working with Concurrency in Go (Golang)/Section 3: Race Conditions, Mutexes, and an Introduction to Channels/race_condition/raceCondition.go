package race_condition

import (
	"fmt"
	"sync"

	"section_3/types"
)

// var msg string
// var wg sync.WaitGroup

// func updateMessage(s string, m *sync.Mutex) {
//   defer wg.Done()
//   m.Lock()
//   msg = s
//   m.Unlock()
// }

// func raceCondition() {
//   msg = "Hello World"
//   wg.Add(2)
//   go updateMessage("Hello, universe", nil)
//   go updateMessage("Hello, cosmos", nil)
//   wg.Wait()
//   fmt.Println(msg)
// }

// func mutextRaceCondition() {
//   msg = "Hello World"
//   var mutex sync.Mutex

//   wg.Add(2)
//   go updateMessage("Hello, universe", &mutex)
//   go updateMessage("Hello, cosmos", &mutex)
//   wg.Wait()
//   fmt.Println(msg)
// }

func raceConditionBankExample(wg *sync.WaitGroup) {
	// variable for bank balance
	var bankBalance int

	// print out starting value
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []types.Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
	}
	wg.Add(len(incomes))
	// loop through 52 weeks and print out how much is made keep a running total
	for i, income := range incomes {
		go func(i int, income types.Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()
	// print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
}
