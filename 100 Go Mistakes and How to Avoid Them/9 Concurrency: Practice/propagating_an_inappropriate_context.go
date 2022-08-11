
// Problem:
// We expose an HTTP handler that performs some tasks and returns a response. However, just before returning the response, we want also to send it to a Kafka topic. As we don’t want to penalize the HTTP consumer latency-wise, we want the publish action to be handled asynchronously within a new goroutine. We assume that we have at our disposal a publish function, accepting a context so that the action of publishing a message can be interrupted if the context is canceled, for example.
func handler(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		// this is problematic because the context could be closed not because the above request failed but because it was successful, in which case we still need this function to execute
		err := publish(r.Context(), response)
		// Do something with err
	}()

	writeResponse(response)
}

// We have to know that the context attached to an HTTP request can cancel in different conditions:

// - When the client’s connection closes
// - In the case of an HTTP/2 request, when the request is canceled
// - Or when the response has been written back to the client

// In the first two cases, we probably handle things correctly. For example, if we just got a response from doSomeTask, but the client has closed the connection, it’s probably OK to call publish with a context already canceled so that the message isn’t published. But what about the last case?

// When the response has been written to the client, the context associated with the request will be canceled. Therefore, we are facing a race condition:

// If writing the response is done after the Kafka publication, we both return a response and publish a message successfully
// However, if writing the response is done before or during the Kafka publication, the message may not be published


// Solutions:
err := publish(context.Background(), response)

// So how can we fix this issue? One idea could be not to propagate the parent context. Instead, we would call publish with an empty context:
// Yet, what if the context contained some useful values? For example, if the context contained a correlation ID used for distributed tracing, we can correlate the HTTP request and the Kafka publication. Ideally, we would like to have a new context, detached from the potential parent cancellation, but that still conveys the values.


// The standard package doesn’t provide an immediate solution to this problem. Hence, a possible solution would be to implement our own Go context similar to the context provided, except that it wouldn’t carry the cancellation signal.

type Context interface {
  Deadline() (deadline time.Time, ok bool)
  Done() <-chan struct{}
  Err() error
  Value(key any) any
}

type detach struct {
        ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) {
        return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
        return nil
}

func (d detach) Err() error {
        return nil
}

func (d detach) Value(key any) any {
        return d.ctx.Value(key)
}

err := publish(detach{ctx: r.Context()}, response)


// Now, the context passed to publish is a context that would never expire nor get canceled but that would carry the parent context’s values.

// In summary, propagating a context should be done cautiously. We illustrated this section with an example of handling an asynchronous action based on a context associated with an HTTP request. As the context is canceled once we return the response, the asynchronous action can also get stopped unexpectedly. Let’s bear in mind the impacts of propagating a given context and if necessary, let’s also keep in mind that it would always be possible to create our custom context for a specific action.

