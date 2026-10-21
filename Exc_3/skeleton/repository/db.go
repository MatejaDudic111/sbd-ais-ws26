package repository

import (
"ordersystem/model"
"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	drinks := []model.Drink{
	    {ID: 1, Name: "Water", Price:2.5, Description: "Still water"},
	    {ID: 2, Name: "Juice", Price:3.5, Description: "Orange juice"},
	    {ID: 3, Name: "Milk", Price:2, Description: "Cows Milk"},
	}

	orders := []model.Order{
        {DrinkID: 1, Amount: 2, CreatedAt: time.Now()},
        {DrinkID: 2, Amount: 1, CreatedAt: time.Now()},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
    totalledOrders := make(map[uint64]uint64)

    for _, order := range db.orders{
        totalledOrders[order.DrinkID] += order.Amount
    }
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	db.orders = append(db.orders, *order)
}
