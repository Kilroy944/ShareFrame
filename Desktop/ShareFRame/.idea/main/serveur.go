package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Picture struct {
	Name   string `json:"name,omitempty"`
	Code64 string `json:"code,omitempty"`
}

func getDirectory(dirname string) []Picture {

	var listPicture []Picture
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {

			data, _ := ioutil.ReadFile(dirname + "/" + file.Name())

			var pict = Picture{file.Name(), string(data)}

			listPicture = append(listPicture, pict)
		}
	}

	return listPicture
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var pictures = getDirectory(params["id"])

	for _, picture := range pictures {
		json.NewEncoder(w).Encode(picture)
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := os.Mkdir(params["id"], os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

}

func AddPicture(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var picture Picture = Picture{"", ""}
	_ = json.NewDecoder(r.Body).Decode(&picture)

	ioutil.WriteFile(params["id"]+"/"+picture.Name, []byte(picture.Code64), os.ModePerm)
}

func DeletePicture(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := os.Remove(params["id"] + "/" + params["name"])
	if err != nil {
		log.Fatal(err)
	}
}

func GetRandomPicture(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var pictures = getDirectory(params["id"])

	if len(pictures) > 0 {
		json.NewEncoder(w).Encode(pictures[rand.Intn(len(pictures))])
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/new_account/{id}", CreateAccount).Methods("GET")
	router.HandleFunc("/get_account/{id}", GetAccount).Methods("GET")
	router.HandleFunc("/delete_picture/{id}/{name}", DeletePicture).Methods("GET")
	router.HandleFunc("/add_picture/{id}", AddPicture).Methods("GET")
	router.HandleFunc("/get_random_picture/{id}", GetRandomPicture).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
