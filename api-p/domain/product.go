package domain

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImgUrl      string `json:"imageUrl"`
}
