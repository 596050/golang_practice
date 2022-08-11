package race_condition

import (
	"fmt"
	"section_3/types"
	"sync"
)

func solution_raceCondition(wg *sync.WaitGroup) {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

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
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
				balance.Unlock()
			}
		}(i, income)
	}
	wg.Wait()
	// print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()

}
