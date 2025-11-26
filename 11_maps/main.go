package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["name"] = "Golang"
	m["area"] = "Backend"
	fmt.Println(m["name"])

	delete(m, "area")

	clear(m)
	fmt.Println(m)

	mp := map[string]int{
		"price": 100,
		"tax":   10,
	}
	_, ok := mp["price"]
	if ok {
		fmt.Println("Key 'price' exists in the map")
	} else {
		fmt.Println("Key 'price' does not exist in the map")
	}

	mmp := map[string]int{
		"apple":  150,
		"banana": 100,
		"cherry": 200,
	}

	fmt.Println(maps.Equal(mp, mmp))
}
