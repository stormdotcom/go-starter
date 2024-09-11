package main

import "fmt"

func main() {
	m := make(map[string]string)
	m1 := map[string]int{
		"apple":  3,
		"orange": 4,
	}
	m2 := make(map[string]interface{})
	m["name"] = "ajmal"
	m["age"] = "3"
	fmt.Println(m)
	fmt.Println((m1))

	value, exists := m1["apple"]
	if exists {
		fmt.Println("apple map value :", value)
	}

	delete(m, "age")

	m2["zip"] = 1111
	m2["addressLine"] = "sdfsd"
	m2["phone"] = 11111
	m2["hasWifi"] = true
	fmt.Println(m2)

	fmt.Println(m)
}
