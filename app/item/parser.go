package item

import (
	uber "github.com/Zoraone/order-web-scraper/uber/item"
)

func UberToApp(uberItem uber.Item) Item {
	return parseItem(uberItem)
}

func parseItem(uberItem uber.Item) Item {
	var item Item
	item.Uuid = uberItem.Uuid
	item.Title = uberItem.Title
	item.Price = uberItem.Price
	item.ItemDescription = uberItem.ItemDescription
	item.ImageUrl = uberItem.ImageUrl

	var customizations []Customization
	for _, uberCustomization := range uberItem.CustomizationsList {
		customization := parseCustomization(uberCustomization)
		customizations = append(customizations, customization)
	}
	item.CustomizationsList = customizations
	return item
}

func parseCustomization(uberCustomization uber.Customization) Customization {
	var cust Customization
	cust.Uuid = uberCustomization.Uuid
	cust.Title = uberCustomization.Title
	cust.MinPermitted = uberCustomization.MinPermitted
	cust.MaxPermitted = uberCustomization.MaxPermitted
	cust.DisplayState = uberCustomization.DisplayState

	var options []Option
	for _, uberOption := range uberCustomization.Options {
		option := parseOption(uberOption)
		options = append(options, option)
	}
	cust.Options = options
	return cust
}

func parseOption(uberOption uber.Option) Option {
	option := Option(uberOption)
	return option
}