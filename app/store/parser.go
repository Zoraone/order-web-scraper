package store

import (
	uber "github.com/Zoraone/order-web-scraper/uber/store"
)

func UberToApp(uberData uber.Data) Store {
	return parseStore(uberData)
}

func parseStore(data uber.Data) Store {
	sectionEntitiesMap := buildSectionEntitiesMap(data)

	var store Store
	store.Uuid = data.Uuid
	store.Title = data.Title

	var sections []Section
	for _, dataSection := range data.Sections {
		section := parseSection(dataSection, data, sectionEntitiesMap)
		sections = append(sections, section)
	}
	store.Sections = sections

	return store
}

func parseSection(dataSection uber.Section, data uber.Data, sectionEntitiesMap map[string]uber.SectionEntities) Section {
	var section Section
	section.Uuid = dataSection.Uuid
	section.Title = dataSection.Title
	section.Subtitle = dataSection.Subtitle
	section.IsOnSale = dataSection.IsOnSale

	var subsections []Subsection
	for _, subsectionUuid := range dataSection.SubsectionUuids {
		dataSubsection := data.SubsectionsMap[subsectionUuid]
		subsection := parseSubsection(dataSubsection, sectionEntitiesMap[dataSection.Uuid])
		subsections = append(subsections, subsection)
	}
	section.Subsections = subsections
	return section
}

func parseSubsection(dataSubsection uber.Subsection, dataSectionItemMap uber.SectionEntities) Subsection {
	var subsection Subsection
	subsection.Uuid = dataSubsection.Uuid
	subsection.Title = dataSubsection.Title
	subsection.Subtitle = dataSubsection.Subtitle

	var items []Item
	for _, itemUuid := range dataSubsection.ItemUuids {
		dataItem := dataSectionItemMap.ItemsMap[itemUuid]
		item := parseItem(dataItem)
		items = append(items, item)
	}
	subsection.Items = items
	return subsection
}

func parseItem(uberItem uber.Item) Item {
	item := Item(uberItem)
	return item
}

func buildSectionEntitiesMap(data uber.Data) map[string]uber.SectionEntities {
	var sectionEntitiesMap = make(map[string]uber.SectionEntities)
	for sectionUuid, itemMap := range data.SectionEntitiesMap {
		var sectionEntities uber.SectionEntities
		var itemsMap = make(map[string]uber.Item)
		for itemUuid, item := range itemMap.(map[string]interface{}) {
			var uberItem uber.Item
			for k, v := range item.(map[string]interface{}) {
				if v != nil {
					switch k {
					case "uuid":
						uberItem.Uuid = v.(string)
					case "title":
						uberItem.Title = v.(string)
					case "description":
						uberItem.Description = v.(string)
					case "imageUrl":
						uberItem.ImageUrl = v.(string)
					case "price":
						uberItem.Price = int(v.(float64))
					}
				}
			}
			itemsMap[itemUuid] = uberItem
		}
		sectionEntities.ItemsMap = itemsMap
		sectionEntitiesMap[sectionUuid] = sectionEntities
	}
	return sectionEntitiesMap
}