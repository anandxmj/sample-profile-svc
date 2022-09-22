package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var env1, env2, env3 string = "ENV1_NO_VALUE", "ENV2_NO_VALUE", "ENV3_NO_VALUE"

func main() {
	if v, ok := os.LookupEnv("ENV1"); ok {
		env1 = v
	}
	if v, ok := os.LookupEnv("ENV2"); ok {
		env2 = v
	}
	if v, ok := os.LookupEnv("ENV3"); ok {
		env3 = v
	}
	log.Println(env1, env2, env3)
	router := mux.NewRouter()
	router.HandleFunc("/profile", Create).Methods("POST")
	router.HandleFunc("/profile", Read).Methods("GET")
	router.HandleFunc("/profile", Update).Methods("PATCH")
	router.HandleFunc("/profile", Delete).Methods("DELETE")
	http.ListenAndServe(":5000", router)
}

func Create(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(fmt.Sprintf("Will Create Profile ENV1 %s, ENV2 %s, ENV3 %s", env1, env2, env3)))
}

func Read(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(fmt.Sprintf("Will Read Profile ENV1 %s, ENV2 %s, ENV3 %s", env1, env2, env3)))
}

func Update(wr http.ResponseWriter, r *http.Request) {
	log.Println("Will Update Profile")
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(fmt.Sprintf("Will Update Profile ENV1 %s, ENV2 %s, ENV3 %s", env1, env2, env3)))
}

func Delete(wr http.ResponseWriter, r *http.Request) {
	log.Println("Will Delete Profile")
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte(fmt.Sprintf("Will Delete Profile ENV1 %s, ENV2 %s, ENV3 %s", env1, env2, env3)))
}
