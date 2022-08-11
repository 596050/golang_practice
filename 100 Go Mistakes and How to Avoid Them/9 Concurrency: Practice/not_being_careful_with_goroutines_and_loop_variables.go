s := []int{1, 2, 3}

for _, i := range s {
	go func() {
		fmt.Print(i)
	}()
}

// In this example, we create new goroutines from a closure. As a reminder, a closure is a function value that references variables from outside its body; here the i variable. We have to know that when a closure goroutine is executed, it doesn’t capture the values when the goroutine is created. Instead, all the goroutines refer to the same exact variable. When a goroutine runs, it prints the value of i at the time fmt.Println is executed. Hence, i may have been modified since the goroutine was launched.

for _, i := range s {
	val := i
	go func() {
		fmt.Print(val)
	}()
}
// Why is this code working? In each iteration, we create a new local val variable. This variable captures the current value of i before creating the goroutine. Hence, when each closure goroutine executes the prin

for _, i := range s {
	go func(val int) {
		fmt.Print(val)
	}(i)
}

// We still execute as a goroutine an anonymous function (we don’t run go f(i) for example), but this time it isn’t a closure. The function doesn’t reference val as a variable from outside its body; val is now part of the function input. By doing so, we fix i in each iteration and make our application work as expected.

