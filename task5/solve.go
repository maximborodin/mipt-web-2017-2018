package main

import "net/http"

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)

var data = make(map[string]string)

var count = 0


func reduceUrl (w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		panic(erro)
	}
	var request map[string]string
	erro = json.Unmarshal(body, &request)
	if erro != nil {
		panic(erro)
	}
	url := request["url"]
	newKey := string (count)
	count += 1
	data[newKey] = url
	responseJson := make(map[string]string)
	responseJson["key"] = newKey
	var response []byte
	response, erro = json.Marshal(responseJson)
	if erro != nil {
		panic(erro)
	}
	w.Write(response)

}

func getUrl (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	w.Header().Add("Location", data[key])
	w.WriteHeader(http.StatusMovedPermanently)
}

func main () {
	router := mux.NewRouter()
	router.HandleFunc("/", reduceUrl)
	router.HandleFunc("/{key}", getUrl)
	http.ListenAndServe(":8082", nil)
}