// Channels are a mechanism for communicating across goroutines via signaling. A signal can be either with or without data. However, it’s not always straightforward for Go programmers how the latter case should be tackled.

// Let’s delve into it with a concrete example. We will create a channel that would notify whenever a certain disconnection occurs. One idea could be to handle it as a chan bool:

disconnectCh := make(chan bool)

// In Go, an empty struct is a struct without any fields. Regardless of the architecture, it occupies zero bytes of storage as we can verify using unsafe.Sizeof:

var s struct{}
fmt.Println(unsafe.Sizeof(s))

// An empty struct is a de-facto standard to convey an absence of meaning. For example, if we need a hash set structure (a collection of unique elements), we should use an empty struct as a value (e.g., map[string]struct{}).

