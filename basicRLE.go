package main

import "fmt"

func main() {
	fmt.Println(rle("ABBBCCCCJDD"))
}

func rle(line string) string {
	var answer string
	var counter int = 1

	lastsym := string(line[0])

	for j := 1; j < len(line); j++ {
		if string(line[j]) != lastsym {
			if counter != 1 {
				answer += lastsym + fmt.Sprintf("%d", counter)
			} else {
				answer += lastsym
			}
			lastsym = string(line[j])
			counter = 0
		}
		counter++
	}
	answer += lastsym + fmt.Sprintf("%d", counter)
	return answer
}
