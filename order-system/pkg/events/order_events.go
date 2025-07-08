package events

type OrderItem struct {
	ProductId string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderCreatedEvent struct {
	OrderId     string      `json:"order_id"`
	CustomerID  string      `json:"customer_id"`
	TotalAmount float64     `json:"total_amount"`
	Items       []OrderItem `json:"items"`
	CreatedAt   string      `json:"created_at"`
}
