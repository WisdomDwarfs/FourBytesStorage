package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"a", "b", "c", "d", "e", "f", "0", "1", "2", "00x", "01x", "10x", "11x"}
	counter := 3
	for i := 0; i < len(str); i++ {
		if i == counter && len(str) != counter {
			counter += 3
			if counter > len(str) {
				counter -= 3
			}

		}
		CodonRead(str[i:counter])
	}
}

func CodonRead(s []string) {
	if len(s) == 3 {
		fmt.Println(strings.Join(s, ""))
	}

}
