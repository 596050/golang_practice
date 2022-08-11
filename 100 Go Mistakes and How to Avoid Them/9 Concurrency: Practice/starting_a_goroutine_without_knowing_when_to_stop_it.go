ch := foo()
go func() {
	for v := range ch {
		// ...
	}
}()


func main() {
        newWatcher()

        // Run the application
}

type watcher struct { /* Some resources */ }

func newWatcher() {
        w := watcher{}
        go w.watch()
}

// We call newWatcher, which creates a watcher struct and spins up a goroutine in charge of watching the configuration. The problem with this code is that when the main goroutine exits (perhaps because of an OS signal or because it has a finite workload), the application will be stopped. Hence, the resources created by watcher won’t be gracefully closed. How can we prevent this from happening?


// One option could be to pass to newWatcher a context that will be cancelled when main returns:
func main() {
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()

        newWatcher(ctx)

        // Run the application
}

func newWatcher(ctx context.Context) {
        w := watcher{}
        go w.watch(ctx)
}

// ------------------------------------------------------------------------------------

// The problem here is that we used signaling to convey that a goroutine had to be stopped. We didn’t block the parent goroutine until the resources had been closed. Let’s make sure now that we do:
func main() {
        w := newWatcher()
        defer w.close()

        // Run the application
}

func newWatcher() watcher {
        w := watcher{}
        go w.watch()
        return w
}

func (w watcher) close() {
        // Close the resources
}

// watcher has a new method: close. Instead of signaling watcher that it’s time to close its resources, we now call this close method using defer to guarantee the resources are closed before the application exits.


// In summary, let’s be mindful that a goroutine is a resource like any other which have to be eventually closed, be it to free memory or other resources. Starting a goroutine without knowing when to stop it is a design issue. Whenever a goroutine is started, we should have a clear plan about when it will stop. Last but not least, if a goroutine creates resources and its lifetime is bound to the lifetime of the application, it’s probably safer to wait for it to be closed instead of notifying it. This way, we can ensure the resources are freed before exiting the application.