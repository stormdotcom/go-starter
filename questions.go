// 1. What is a goroutine in Go?
// Answer: A goroutine is a lightweight thread managed by the Go runtime.
// Example:
package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printMessage("Goroutine") // Runs in a separate goroutine
	printMessage("Main")        // Runs in the main goroutine
}

// Output:
// Alternating lines of "Goroutine" and "Main"

// 2. How does Go handle concurrency?
// Answer: Go uses goroutines and channels for concurrency.
// Example:
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	message := <-ch // Receive data from the channel
	fmt.Println(message)
}

// Output:
// Hello from goroutine

// 3. What are Go interfaces and how are they used?
// Answer: Interfaces define methods that a type must implement.
// Example:
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var s Shape

	s = Circle{Radius: 5}
	fmt.Printf("Circle Area: %.2f\n", s.Area())

	s = Rectangle{Width: 4, Height: 6}
	fmt.Printf("Rectangle Area: %.2f\n", s.Area())
}

// Output:
// Circle Area: 78.50
// Rectangle Area: 24.00


// 4. What is the difference between a slice and an array in Go?
// Answer:
// - Arrays have fixed size.
// - Slices are dynamic and can grow/shrink.

package main

import "fmt"

func main() {
	// Array
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice
	slice := []int{4, 5, 6}
	fmt.Println("Slice:", slice)

	// Append to slice
	slice = append(slice, 7)
	fmt.Println("Updated Slice:", slice)
}

// Output:
// Array: [1 2 3]
// Slice: [4 5 6]
// Updated Slice: [4 5 6 7]

// 5. Explain Go's defer statement.
// Answer: `defer` schedules a function to run after the surrounding function completes.
// Example:
package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("Deferred: This runs at the end")

	fmt.Println("End")
}

// Output:
// Start
// End
// Deferred: This runs at the end


// 6. How does Go handle errors?
// Answer: Go uses the `error` type for error handling.
// Example:
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Output:
// Result: 5
// Error: cannot divide by zero

// 7. How are maps used in Go?
// Answer: Maps are key-value pairs, similar to dictionaries in Python.
// Example:
package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	fmt.Println("Alice's Age:", ages["Alice"])

	// Add or update
	ages["Charlie"] = 35
	fmt.Println("Updated Map:", ages)

	// Delete
	delete(ages, "Bob")
	fmt.Println("After Deletion:", ages)
}

// Output:
// Alice's Age: 25
// Updated Map: map[Alice:25 Bob:30 Charlie:35]
// After Deletion: map[Alice:25 Charlie:35]

