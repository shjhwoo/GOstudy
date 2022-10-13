package main

import "time"

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt         *Product
	ShippedTime time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(product *Product) *Parcel{
	var resultParcel = Parcel{}
	resultParcel.ShippedTime = time.Now()
	resultParcel.Pdt = product
	return &resultParcel
}

func (p *Parcel) Delivered() *Product{
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {

}