package product

type IMoney interface{}

type Product struct {
	Price     string `json:"price"`
	Available bool   `json:"avail"`
}
