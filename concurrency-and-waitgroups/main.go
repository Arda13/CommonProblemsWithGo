package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	students := []string{"Student1", "Student2", "Student3", "Student4"}
	var wg sync.WaitGroup
	wg.Add(len(students))
	for _, student := range students {
		go func(student string) {
			graduate(student)
			wg.Done()
		}(student)
	}
	wg.Wait()
	fmt.Printf("school took %v seconds\n", time.Since(start))
}

func graduate(student string) {
	fmt.Printf("%s is studying...\n", student)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s is graduated\n", student)
	fmt.Printf("")
}
