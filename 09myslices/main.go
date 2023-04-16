package main

import (
	"fmt"
	"sort"
)

func main() {
	var itemlist = []string{"apple"} // slices works like array under the hood and there is no fixed values

	fmt.Printf("Type is %T\n", itemlist)

	itemlist = append(itemlist, "pear", "banana")
	fmt.Println("after appending", itemlist[1:2])

	highscores := make([]int, 4)
	highscores[0] = 12
	highscores[1] = 19
	highscores[2] = 17
	highscores[3] = 11

	fmt.Println("highscores using make:", highscores)
	fmt.Printf("type of highscores %T\n", highscores) // this is acting as a array

	highscores = append(highscores, 66, 55, 44)
	fmt.Println("highscores after using append on make:", highscores) // after using append it will reallocate the memory

	// soring
	sort.Ints(highscores)
	fmt.Println("after sorting ", highscores)
	fmt.Println("are sorted", sort.IntsAreSorted(highscores))

	// Remove a value from slices based on index
	var fruits = []string{"apple", "guava", "papaya", "lichi"}
	var index int = 2
	fruits = append(fruits[:index], fruits[index+1:]...)
	fmt.Println("after removing index", fruits)

}
