/*
* In java when functions go into the classes we call them methods
* Then methods are called by object of the classes
* since in golang as we dont have classes we use structs
 */
package main

import "fmt"

func main() {
	debankan := User{"Dev", "debankanmitra@gmail.com", true, 21}
	debankan.isSingle()
	debankan.newmail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) isSingle() {
	fmt.Println("Is user active", u.Status)
}

func (u User) newmail() {
	fmt.Println("teh email is", u.Email)
}
