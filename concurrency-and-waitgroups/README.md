# Concurrency and Waitgroups

## Concurrency

>Concurrency is an ability of a program to do multiple things at the same time. This means a program that has two or more tasks that run individually of each other, at about the same time but remain part of the same program.

> In Golang, we can specifically manage our concurrent computations with the `go` keyword. In the `main.go` file, we put the `go` keyword in front of the function inside of for loop. In standard for loops, the computation starts from the initial first value to the last value sequentially and this synchronous behaviour might cause breaks. Because every step should be done in order and the program does not move to the next block before the calculation done.

> `go` keyword gives us a nice parallel computation feature. Thanks to the keyword, I've calculated every step of for loop concurrently.

## Waitgroups

> If you forget to use the `sync` package, the main function will be called before our goroutines and you will get an error. Because our function can not reach the inside of goroutines without `Waitgroup`. 

> To solve this problem, I've added the length of the students array time wait groups before for loop and put the wait method right after the for loop while calling done method inside of for loop.

> Thanks to this solution, I've got a concurrent for loop.

## Results

>  Call `go run main.go` in the command line 
