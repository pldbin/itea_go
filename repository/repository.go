package repository

import (
	"HW_2/types"
	"errors"
)

type Repository struct {
	products  map[string]types.Product
	customers map[string]types.Customer
}

func New() *Repository {
	return &Repository{
		products:  make(map[string]types.Product),
		customers: make(map[string]types.Customer),
	}
}

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrCustomerNotFound  = errors.New("customer not found")
	ErrNotEnoughQuantity = errors.New("not enough quantity")
)

func (r *Repository) ViewProducts() []types.Product {
	products := make([]types.Product, 0, len(r.products))

	for _, product := range r.products {
		products = append(products, product)
	}

	return products
}

func (r *Repository) ViewProductDetails(productName string) (types.Product, error) {
	product, ok := r.products[productName]
	if !ok {
		return types.Product{}, ErrProductNotFound
	}

	return product, nil
}

func (r *Repository) AddToCart(customerEmail string, productName string, quantity int) error {
	customer, ok := r.customers[customerEmail]
	if !ok {
		return ErrCustomerNotFound
	}

	product, ok := r.products[productName]
	if !ok {
		return ErrProductNotFound
	}

	if quantity > product.Quantity {
		return ErrNotEnoughQuantity
	}

	customer.Cart[productName] = quantity
	product.Quantity -= quantity

	r.customers[customerEmail] = customer
	r.products[productName] = product

	return nil
}

func (r *Repository) ViewCart(customerEmail string) ([]types.Order, error) {
	customer, ok := r.customers[customerEmail]
	if !ok {
		return nil, ErrCustomerNotFound
	}

	orders := make([]types.Order, 0, len(customer.Cart))

	for productName, quantity := range customer.Cart {
		orders = append(orders, types.Order{
			Product:  productName,
			Quantity: quantity,
		})
	}

	return orders, nil
}
