package main

import "fmt"

func main() {
	var itemlist [4]string

	itemlist[0] = "soap"
	itemlist[1] = "oil"
	itemlist[2] = "vegies"
	itemlist[3] = "papers"

	fmt.Println(itemlist)

	anotherList := [2]string{"dev", "sam"}
	fmt.Println("anotherlsit ", anotherList)
}
