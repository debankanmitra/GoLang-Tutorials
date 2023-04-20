package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course file
type course struct {
	CourseId    string  `json:"Id"`
	CourseName  string  `json:"Name"`
	CoursePrice int     `json:"Price"`
	Author      *Author `json:"Author"`
}

type Author struct {
	Fullname string `json:"Fullname"`
	Website  string `json:"Website"`
}

// Fake DB
var Courses []course

// Middleware , helper - file
func (c *course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

// controller - file
// serve home route

func Servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Api</h1>"))
}

// another route where we will show all the values frm our database there
func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-Type", "application/json")
	// throwing json values
	json.NewEncoder(w).Encode(Courses)
}

func getSpecificCourse(w http.ResponseWriter, r *http.Request) { // this way it is not used in production
	fmt.Println("Get specific course")
	w.Header().Set("content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses find matching id and return the response
	for _, val := range Courses {
		if val.CourseId == params["id"] {
			json.NewEncoder(w).Encode(val)
		}
	}
	json.NewEncoder(w).Encode("No value found")
	return
}

// Create , Update , Delete operation on course - file
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("content-Type", "application/json")

	// body should not be empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Send some data")
	}

	var course course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in input")
		return
	}

	// generate unique id , string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode("COurse added")

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("content-Type", "application/json")

	// grab id from request,remove, add with my id
	params := mux.Vars(r)
	for index, val := range Courses {
		if val.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...) // remove
			var course course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}
}

func deleteCouerse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting Course")
	w.Header().Set("content-Type", "application/json")

	// grab id from request,remove
	params := mux.Vars(r)
	for index, val := range Courses {
		if val.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...) // remove
			json.NewEncoder(w).Encode(" deleted")
			break
		}
	}
}
