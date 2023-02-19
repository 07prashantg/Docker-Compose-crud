package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

type Detail struct {
	Student_Id   string `json:"Student_Id"`
	Student_Name string `json:"Student_Name"`
	Course       string `json:"Course"`
	Department   string `json:"Department"`
	Place        string `json:"Place"`
}

func connectDB() {
	db, err := sql.Open("mysql", "crud_user:password@tcp(db:3306)/crud_database")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("My Connection ......")
	dbm = db

	// defer db.Close()
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []*Detail

	results, err := dbm.Query("SELECT * FROM Student")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var u Detail
		err = results.Scan(&u.Student_Id, &u.Student_Name, &u.Course, &u.Department, &u.Place)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, &u)
	}

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}

func getDetailsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	param := keyVal["Student_Id"]

	var detail []*Detail

	result, err := dbm.Query("select * from Student where Student_Id = ?", param)
	if err != nil {
		log.Fatal(err.Error())
	}

	for result.Next() {
		var u Detail
		err = result.Scan(&u.Student_Id, &u.Student_Name, &u.Course, &u.Department, &u.Place)
		if err != nil {
			log.Fatal(err.Error())
		}
		detail = append(detail, &u)
	}

	fmt.Println("Endpoint Hit: getelementbyid")
	json.NewEncoder(w).Encode(detail)
}

func updateDetails(w http.ResponseWriter, r *http.Request) {
	result, err := dbm.Prepare("update Student set Student_NAME = ?, Course = ?, Department = ?, Place = ? where Student_Id = ?")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	param := keyVal["Student_Id"]
	newStudent_Name := keyVal["Student_Name"]
	newCourse := keyVal["Course"]
	newDepartment := keyVal["Department"]
	newPlace := keyVal["Place"]

	_, err = result.Exec(newStudent_Name, newCourse, newDepartment, newPlace, param)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(w, "Post Updated")
}

func insertDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := dbm.Prepare("Insert into Student (Student_NAME,Course,Department,Place) values (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	Student_Name := keyVal["Student_Name"]
	Course := keyVal["Course"]
	Department := keyVal["Department"]
	Place := keyVal["Place"]

	_, err = result.Exec(Student_Name, Course, Department, Place)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(w, "New Post Created")
}

func deleteDetails(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	param := keyVal["Student_Id"]

	result, err := dbm.Prepare("delete from Student where Student_Id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.Exec(param)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(w, "Post deleted")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page")
	fmt.Println("EndPoint hit: homePage")
}

func main() {
	connectDB()

	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/detailById", getDetailsById)
	http.HandleFunc("/insert", insertDetails)
	http.HandleFunc("/update", updateDetails)
	http.HandleFunc("/deleteById", deleteDetails)

	fmt.Printf("Starting server at port 8082\n")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
