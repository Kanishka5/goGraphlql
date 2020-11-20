package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Kanishka5/goGraph/graph/generated"
	"github.com/Kanishka5/goGraph/graph/helper"
	"github.com/Kanishka5/goGraph/graph/model"
)

func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	item := model.Item{
		Name:   input.Name,
		Cost:   input.Cost,
		Seller: input.Seller,
	}

	return &item, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}
	model.DB.Create(&user)

	return &user, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	order := model.Order{
		// ID:        input.ID,
		Total:     input.Total,
		OrderDate: input.OrderDate,
		Customer:  input.Customer,
		Items:     helper.MapItemsFromInput(input.Items),
	}
	model.DB.Create(&order)

	return &order, nil
}

func (r *queryResolver) Order(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	model.DB.Preload("Items").Find(&orders)

	return orders, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
