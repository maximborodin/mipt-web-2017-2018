package main

import "net/http"

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strconv"
)

var data = make(map[string]string)

var count = 0
var baseKey = "url"


func reduceUrl (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post")
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
	newKey := baseKey + strconv.Itoa(count)
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
	fmt.Println("Get")
	vars := mux.Vars(r)
	key := vars["key"]
	w.Header().Add("Location", data[key])
	w.WriteHeader(http.StatusMovedPermanently)
}

func main () {
	/*for i := 0; i < 10; i++ {
		fmt.Println(string(i))
	}*/
	router := mux.NewRouter()
	router.HandleFunc("/{key}", getUrl)
	router.HandleFunc("/", reduceUrl)
	http.ListenAndServe(":8082", router)
}