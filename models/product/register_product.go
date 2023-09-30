package product

type ProductData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Stock     int    `json:"stock"`
	Price     int    `json:"price"`
	Publisher string `json:"publisher"`
}
