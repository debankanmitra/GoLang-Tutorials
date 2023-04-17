package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	detail := []Userdetail{
		{"Dev", "dhfjkdh@gmail", 808909856},
		{"Ddgv", "dhfjkdh@gmail", 80309856},
		{"dgdg", "dhfjkdh@gmail", 80358956},
		{"tregh", "dhfjkdh@gmail", 80358909},
	}

	fmt.Println(detail)

	final, _ := json.Marshal(detail)

	fmt.Printf("%s\t", final)

	final2, _ := json.MarshalIndent(detail, "", "\t")
	fmt.Printf("%s\n", final2)

}

type Userdetail struct {
	Name  string
	Email string
	Phone int
}
