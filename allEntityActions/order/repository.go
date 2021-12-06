package order

import "github.com/ermiasashebr/housing/entity"

type OrderRepository interface {
	Orders() ([]entity.Order, []error)
	Order(id uint) (*entity.Order, []error)
	CustomerOrders(customer *entity.User) (entity.Order, []error)
	UpdateOrder(order *entity.Order) (*entity.Order, []error)
	DeleteOrder(id uint) (*entity.Order, []error)
	StoreOrder(order *entity.Order) (*entity.Order, []error)
}


