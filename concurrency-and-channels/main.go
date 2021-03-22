package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	students := []string{"Student1", "Student2", "Student3", "Student4"}
	results := make(chan bool)
	for _, student := range students {
		go func(student string) {
			graduate(student)
			results <- true
		}(student)
	}

	for i := 0; i < len(students); i++ {
		<-results
	}
	fmt.Printf("school took %v seconds\n", time.Since(start))
}

func graduate(student string) {
	fmt.Printf("%s is studying...\n", student)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s is graduated\n", student)
	fmt.Printf("")
}
