package item

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const localhost = "http://localhost:8080/"
const prodhost = ""

func AddNewItem(item Item) {
	path := "api/item/add/"

	requestBody, err := json.Marshal(item)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(localhost + path, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(body)
}