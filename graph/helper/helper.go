package helper

import "github.com/Kanishka5/goGraph/graph/model"

// MapItemsFromInput maps each item
func MapItemsFromInput(newItem []*model.NewItem) []*model.Item {
	var items []*model.Item
	for _, itemInput := range newItem {
		items = append(items, &model.Item{
			Name:   itemInput.Name,
			Cost:   itemInput.Cost,
			Seller: itemInput.Seller,
		})
	}
	return items
}
