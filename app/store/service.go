package store

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const localhost = "http://localhost:8080/"
const prodhost = ""

func AddNewStore(store Store) {
	path := "api/store/add"

	requestBody, err := json.Marshal(store)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(localhost + path, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}