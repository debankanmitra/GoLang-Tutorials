package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["jv"] = "java"
	languages["rb"] = "ruby"
	languages["ts"] = "typescript"

	fmt.Println("map is ", languages)
	fmt.Println(languages["js"])

	// deletion in map
	delete(languages, "js")
	fmt.Println("map is ", languages)

	// looping through map
	for _, val := range languages {
		fmt.Printf("for key , value is %v \n", val)
	}
}
