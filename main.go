package main

import (
	"encoding/json"
	"flag"
	"log"

	app_item "github.com/Zoraone/order-web-scraper/app/item"
	app_store "github.com/Zoraone/order-web-scraper/app/store"
	uber_item "github.com/Zoraone/order-web-scraper/uber/item"
	uber_store "github.com/Zoraone/order-web-scraper/uber/store"

	"github.com/Zoraone/order-web-scraper/util"
)

func main() {
	store := flag.String("store", "", "The uuid of the store to retrieve data from.")
	live := flag.Bool("live", false, "If included, will post to the live API.")

	flag.Parse()

	if *store == "" {
		log.Fatalln("Please specify a store.")
	}

	uberStoreData := uber_store.GetStoreDetails(*store)
	appStore := app_store.UberToApp(uberStoreData)

	uberItems := getItemDetails(appStore)
	var items []app_item.Item
	for _, uberItem := range uberItems {
		items = append(items, app_item.UberToApp(uberItem))
	}
	js, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	util.WriteToFile(js)

	
	if *live {
		app_store.AddNewStore(appStore)
		for _, item := range items {
			app_item.AddNewItem(item)
		}
	}
}

func getItemDetails(store app_store.Store) []uber_item.Item {
	var items []uber_item.Item
	for _, section := range store.Sections {
		for _, subsection := range section.Subsections {
			for _, item := range subsection.Items {
				reqBody := uber_item.RequestBody{
					MenuItemUuid: item.Uuid,
					SectionUuid: section.Uuid,
					SubsectionUuid: subsection.Uuid,
					StoreUuid: store.Uuid,
				}
				items = append(items, uber_item.GetItemDetails(reqBody))
			}
		}
	}
	return items
}