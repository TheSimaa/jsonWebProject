package main

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var users []User

type User struct {
	ID        int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       string `json:"age"`
	Location  string `json:"location"`
	Email     string `json:"email"`
}

func readJsonFile(jsonPath string) {

	jsonFile, err := os.Open(jsonPath)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	jsonByteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(jsonByteValue, &users)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func getUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idUser, _ := strconv.Atoi(params["id"])

	for _, item := range users {
		if item.ID == idUser {
			json.NewEncoder(w).Encode(item)
		}
	}a

}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = users[len(users)-1].ID + 1

	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}

func main() {

	r := mux.NewRouter()
	readJsonFile("./json/users.json")

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/user", createUser).Methods("POST")
	//r.HandleFunc("/api/user/{id}", updateUser).Methods("PUT")
	//r.HandleFunc("/api/user/{id}", delпeteUser).Methods("DELETE")

	http.HandleFunc("/", nil)
	http.ListenAndServe(":8000", r)

}
