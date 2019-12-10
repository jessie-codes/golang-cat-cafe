# Concurrency & Mutex

<img src="concurrency.jpg"
     style="height:500px;" />

To run our cat cafe, we'll need the ability to let customers reserve a cat. Fortunately for us, someone already built an API endpoint
to handle this. Unfortunately, it runs `time.Sleep` in the main thread, which prevents any other processing from happening. Let's see what
we can do to fix it.

## Task Five: Go Routines

Go supports multi-threading with go routines. Each routine runs in it's own thread. For routines that can run on their own, we can create one just by using the `go` keyword before a function call.

```golang
go myFunc()
```

Let's use this to update the `Reserve` function in `cat.go`. Instead of handling adjusting the availability inline, we should create a new function on the `Cat` object that does it for us instead.

```golang
func (c *Cat) startAppointment() {
	c.Available = false
	time.Sleep(5 * time.Second)
	c.Available = true
}
```

We can then update `Reserve` to call this function in a go routine, unblocking our main thread.

For cases where you need data from your go routine, you can use `channels`. We won't be covering them in today's workshop, but you can read more about them [here](https://gobyexample.com/channels).

## Task Six: Mutex

Great! We've unblocked our main thread, but we've introduced a new problem, race conditions. We now have a scenario where two people could reserve the same cat. This can be solved by adding a read/write mutex.

With a mutex we can ensure that only one process can read and update our `List` of cats.

Golang has a built in package for working with a mutex, `sync`. To get started, let's import the `sync` in `Cat.go`.

You can add a mutex to a struct by adding an additional property. We should make this one private so that only our `cat` package has access to it.

```go
mux  sync.Mutex
```

We can then control read/write access to an object by locking/unlocking it.

```go
mux.Lock()
mux.Unlock()
```

See if you can update the `List` struct as well as the `Reserve` function to ensure only one go routine can access the value at a time.

If you get stuck, you can look at `Cat.go` in the `/complete` folder for a reference for how to implement a `mutex`.