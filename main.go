package main

import (
	"fmt"
	"encoding/json"

	app_store "./app/store"
	uber_store "./uber/store"
)

func main() {
	uberStoreData := uber_store.GetStoreDetails("69cea054-981c-47fe-a321-0465d3ba0cc5")
	appStore := app_store.UberToApp(uberStoreData)

	js, err := json.Marshal(appStore)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", string(js))
}