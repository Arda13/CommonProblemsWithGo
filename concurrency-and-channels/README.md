# Concurrency and Channels

## Channels

> `Waitgroup` is a way to coordinate and synchronize go routines `channels` are a bit more, it is used for spesifically passing information across go routines. It can also be used as a synchronization method.

> Channels are like a slice of data, it can be consumed by different go routines as a pipe. You can basically send an information to end and get an information from another end. Also there are possible restrictions to turn channels into only receivers and only senders. Generally it is used as a communication method between go routines.

> `chan` keyword is essential to create channels. You can create many types of channels like `chan bool` or `chan int`.

## Results

> Call `go run main.go` in the command line
