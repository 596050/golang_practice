// Let’s imagine we want to implement a goroutine that needs to receive from two channels:

// messageCh for new messages to be processed.
// disconnectedCh to receive notifications conveying some disconnections. In that case, we want to return from the parent function.
// Between those two channels, we want to prioritize messageCh. For example, if a disconnection occurs, we want to ensure that we have received all the messages before returning.
for {
	select {
	case v := <-messageCh:
		fmt.Println(v)
	case <-disconnectCh:
		fmt.Println("disconnection, return")
		return
	}
}
// We use select to receive from multiple channels. As we want to prioritize messageCh, an assumption could be to write the messageCh case first and the disconnectedCh one next.
for i := 0; i < 10; i++ {
	messageCh <- i
}
disconnectCh <- struct{}{}

// 0
// 1
// 2
// 3
// 4
// disconnection, return

// Unlike a switch statement where the first case with a match wins, the select statement will select one randomly if multiple options are possible.

// This behavior might look odd at first, but there’s a good reason for that: to prevent possible starvation. Indeed, suppose the first possible communication chosen is based on the source order. In that case, we may fall into the situation where we would solely receive from one single channel because of a fast sender, for example. To prevent this, the language designers have decided to use a random selection.

// If there’s a single producer goroutine, we have two options:

// Either to make messageCh an unbuffered channel instead of a buffered one. As the sender goroutine blocks until the receiver goroutine is ready, it would guarantee that all the messages from messageCh are received before receiving the disconnection from disconnectedCh.

// Or to use a single channel instead of two. For example, we could define a struct that would either convey a new message or a disconnection. As channels guarantee that the order for the messages sent is the same as the messages received, we could ensure that the disconnection would be received in the end.

for {
	select {
	case v := <-messageCh:
		fmt.Println(v)
	case <-disconnectCh:
		for {
			select {
			case v := <-messageCh:
				fmt.Println(v)
			default:
				fmt.Println("disconnection, return")
				return
			}
		}
	}
}

// When using select with multiple channels, we must remember that if multiple options are possible, it’s not the first case in the source order that will win. Instead, Go will select it randomly, so there’s no guarantee about which one will be chosen. To overcome this behavior, in the case of a single producer goroutine, we can either use unbuffered channels or use a single channel instead. In the case of multiple producer goroutines, we can use inner selects and default to handle prioritizations.


