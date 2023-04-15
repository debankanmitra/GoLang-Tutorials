package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2019, time.February, 15, 9, 15, 59, 4, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
