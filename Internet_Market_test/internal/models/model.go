package models

type Seller struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
type SellerResponse struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	SellerID int32   `json:"seller_id"`
}
type ProductResponse struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	SellerId int32   `json:"seller_id"`
}

type Customer struct {
	Name  string
	Phone string
}

type Order struct {
	CustomerID uint
	Products   []Product
}
type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
