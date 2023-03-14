package main

import "fmt"

// import "time"

// type Courier struct {
// 	Name string
// }

// type Product struct {
// 	Name  string
// 	Price int
// 	ID    int
// }

// type Parcel struct {
// 	Pdt         *Product
// 	ShippedTime time.Time
// 	DeliveredTime time.Time
// }

// func (c *Courier) SendProduct(product *Product) *Parcel{
// 	var resultParcel = Parcel{}
// 	resultParcel.ShippedTime = time.Now()
// 	resultParcel.Pdt = product
// 	return &resultParcel
// }

// func (p *Parcel) Delivered() *Product{
// 	p.DeliveredTime = time.Now()
// 	return p.Pdt
// }

func main() {
	b := make([]int, 0, 5) //len은 0이고, 즉 요소 없고 최대 5개까지 넣기가능
	c := b[:2] //3번쨰 인수를 넣지 않아서 cap도 2고 len도 2다.
	d := c[1:5] // this is equivalent to d := b[1:5]
	d[0] = 1
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
}
