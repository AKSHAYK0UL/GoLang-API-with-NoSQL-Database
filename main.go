package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("Welcome To Golang Api")

	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/get", Get).Methods("GET")
	r.HandleFunc("/post", Post).Methods("POST")
	r.HandleFunc("/search{id}", Search).Methods("GET")
	r.HandleFunc("/update{id}", Update).Methods("PUT")
	r.HandleFunc("/del{id}", Delete).Methods("DELETE")
	r.HandleFunc("/del", DeleteAll).Methods("DELETE")
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	fmt.Println("Server is Running on 192.168.29.1:", port)
	go keepServerAlive()
	log.Fatal(http.ListenAndServe("192.168.29.1:"+port, r))

}

// Keep the Server alive
func keepServerAlive() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		sendRequest()
	}
}

func sendRequest() {
	fmt.Println("SERVER ALIVE")
	resp, err := http.Get("https://go-api-with-doc-database.onrender.com")
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("Request sent successfully at %s\n", time.Now().Format(time.RFC3339))
	} else {
		fmt.Printf("Failed to send request. Status code: %d\n", resp.StatusCode)
	}
}

// home Page
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("<h1>Welcome to Golang Api with Doc database</h1>"))

}

// post data
func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newDataFromReq Course //current reqdata
	var courseData []Course   //all the data in json file
	json.NewDecoder(r.Body).Decode(&newDataFromReq)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	newDataFromReq.CourseId = strconv.Itoa(rand.Intn(1000))
	file, _ := os.OpenFile("./courses.json", os.O_RDWR|os.O_CREATE, 0644)
	json.NewDecoder(file).Decode(&courseData)
	courseData = append(courseData, newDataFromReq)
	encoded, _ := json.MarshalIndent(courseData, "", "\t") //convert data in json formate
	file.Seek(0, 0)
	file.Write(encoded) //write the encoded json in file
	json.NewEncoder(w).Encode("Data Saved Sucessfully")
	defer file.Close()
}

// get data
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := os.Open("./courses.json")
	var courses []Course
	json.NewDecoder(file).Decode(&courses) //decode the file data and save in courses slice
	json.NewEncoder(w).Encode(courses)
	defer file.Close()

}

// Search data
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := os.Open("./courses.json")
	var courses []Course
	json.NewDecoder(file).Decode(&courses) //decode the file data in course slice
	params := mux.Vars(r)
	for _, cour := range courses {
		if cour.CourseId == params["id"] {
			json.NewEncoder(w).Encode(cour)
			return
		}
	}
	json.NewEncoder(w).Encode("No Data Found With This Id")
	defer file.Close()
}

// Update data
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := os.OpenFile("./courses.json", os.O_RDWR|os.O_CREATE, 0644)
	var courses []Course
	params := mux.Vars(r)
	json.NewDecoder(file).Decode(&courses)
	for index, cour := range courses {
		if cour.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var reqData Course
			json.NewDecoder(r.Body).Decode(&reqData)
			reqData.CourseId = params["id"]
			courses = append(courses, reqData)
			encoded, _ := json.MarshalIndent(courses, "", "\t")
			file.Seek(0, 0)  //set file curser to start
			file.Truncate(0) //clear the old data
			file.Write(encoded)
			json.NewEncoder(w).Encode(reqData)
			return
		}
	}
	json.NewEncoder(w).Encode("No Data Found With This Id")
	defer file.Close()
}

// delete one course
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _ := os.OpenFile("./courses.json", os.O_RDWR|os.O_CREATE, 0644)
	var courses []Course
	json.NewDecoder(file).Decode(&courses)
	params := mux.Vars(r)
	for index, cour := range courses {
		if cour.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			file.Seek(0, 0)
			file.Truncate(0)
			encoded, _ := json.MarshalIndent(courses, "", "\t")
			file.Write(encoded)
			json.NewEncoder(w).Encode(courses)
			return

		}
	}
	json.NewEncoder(w).Encode("No Data Found With This Id")

}

// delete all courses
func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	file, _ := os.OpenFile("./courses.json", os.O_RDWR, 0644)
	file.Truncate(0)
	json.NewEncoder(w).Encode("Data Deleted")
}
