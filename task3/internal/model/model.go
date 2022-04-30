package model

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Discount bool    `json:"discount,omitempty"`
}

type Basket struct {
	Products  []*Product `json:"products"`
	LinkAB    int
	LinkDE    int
	LinkEFG   int
	LinkAklm  int
	Link3     int
	Link4     int
	Link5     int
	ResultSum float64
}

func NewBasket() *Basket {
	return &Basket{}
}
