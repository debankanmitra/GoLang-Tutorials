package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course file
type Course struct {
	CourseId    string  `json:"CourseId"`
	CourseName  string  `json:"CourseName"`
	CoursePrice int     `json:"CoursePrice"`
	Author      *Author `json:"Author"`
}

type Author struct {
	Fullname string `json:"Fullname"`
	Website  string `json:"Website"`
}

// Fake DB
var Courses []Course

// Middleware , helper - file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("API WORING")

	r := mux.NewRouter()

	// seeding
	Courses = append(Courses, Course{CourseId: "2", CourseName: "React", CoursePrice: 70, Author: &Author{Fullname: "Dev", Website: "lo.io"}})
	Courses = append(Courses, Course{CourseId: "5", CourseName: "js", CoursePrice: 90, Author: &Author{Fullname: "pk", Website: "plo.io"}})
	Courses = append(Courses, Course{CourseId: "7", CourseName: "java", CoursePrice: 170, Author: &Author{Fullname: "ouy", Website: "ldgyddfo.io"}})

	// routing , api end points
	r.HandleFunc("/", Servehome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getSpecificCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCouerse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
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

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in JSON")
		return
	}

	// generate unique id , string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("content-Type", "application/json")

	// grab id from request,remove, add with my id
	params := mux.Vars(r)
	for index, val := range Courses {
		if val.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...) // remove
			var course Course
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
