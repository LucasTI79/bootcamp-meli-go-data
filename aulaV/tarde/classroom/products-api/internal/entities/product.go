package entities

type Product struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

type ProductFullData struct {
	ID               uint64  `json:"id"`
	Name             string  `json:"name"`
	Category         string  `json:"category"`
	Count            int     `json:"count"`
	Price            float64 `json:"price"`
	Warehouse        string  `json:"warehouse"`
	WarehouseAddress string  `json:"warehouse_address"`
}
