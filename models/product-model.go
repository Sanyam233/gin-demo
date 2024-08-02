package models

type Product struct {
	ID       string `json:"id"`
	Title    string `json:"title" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Seller   *User  `json:"seller" binding:"required"`
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var Products = []Product{
	{
		ID:       "123",
		Title:    "Product 1",
		Quantity: 10,
		Seller: &User{
			FirstName: "John",
			LastName:  "Doe",
		},
	},
	{
		ID:       "234",
		Title:    "Product 2",
		Quantity: 3,
		Seller: &User{
			FirstName: "John",
			LastName:  "Doe",
		},
	},
	{
		ID:       "345",
		Title:    "Product 3",
		Quantity: 8,
		Seller: &User{
			FirstName: "Marry",
			LastName:  "Popins",
		},
	},
	{
		ID:       "456",
		Title:    "Product 4",
		Quantity: 16,
		Seller: &User{
			FirstName: "John",
			LastName:  "Doe",
		},
	},
}
