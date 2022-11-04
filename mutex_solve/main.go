package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {

	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("start factory\n")

	wg.Add(3)
	go makeBody(tireCh)
	go installTire(tireCh, paintCh)
	go paintCar(paintCh)

	wg.Wait()
}

func makeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func installTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "black tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func paintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "mint"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
