package main

import "fmt"

func main() {
	data := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(checkDuplicate(data))
}

func checkDuplicate(data [5]string) bool {
	counter := 0
	for a := 0; a < len(data); a++ {
		for b := 0; b < len(data); b++ {
			if data[a] == data[b] {
				counter++
			}

			if counter == 2 {
				return true
			}
		}
		counter = 0
	}

	return false
}
