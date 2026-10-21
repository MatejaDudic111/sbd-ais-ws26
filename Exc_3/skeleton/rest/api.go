package rest

import (
    "encoding/json"
	"net/http"
	"ordersystem/repository"
    "ordersystem/model"

	"github.com/go-chi/render"
)

// GetMenu 			godoc
// @tags 			Menu
// @Description 	Returns the menu of all drinks
// @Produce  		json
// @Success 		200 {array} model.Drink
// @Router 			/api/menu [get]
func GetMenu(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		drinks := db.GetDrinks()
		render.Status(r, http.StatusOK)
		render.JSON(w,r,drinks)
	}
}

func GetOrders(db *repository.DatabaseHandler) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
        orders := db.GetOrders()
        render.Status(r, http.StatusOK)
        render.JSON(w,r,orders)

    }
}

func GetOrdersTotal(db *repository.DatabaseHandler) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
        totalledOrders := db.GetTotalledOrders()
        render.Status(r, http.StatusOK)
        render.JSON(w,r,totalledOrders)

    }
}

// PostOrder 		godoc
// @tags 			Order
// @Description 	Adds an order to the db
// @Accept 			json
// @Param 			b body model.Order true "Order"
// @Produce  		json
// @Success 		200
// @Failure     	400
// @Router 			/api/order [post]
func PostOrder(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order model.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
		    render.Status(r, http.StatusBadRequest)
		    render.JSON(w,r,map[string]string{
		        "error": "Invalid order JSON",
		    })
            return
		}
        db.AddOrder(&order)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, "ok")
	}
}
