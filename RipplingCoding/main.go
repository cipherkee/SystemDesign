package main

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"fmt"
)

func addNumbers(a uint32, b uint32) uint32 {
	return (a + b)
}

// billingSystem.AddDriver(1, 35.10)
// billingSystem.AddDriver(2, 15.15)
// billingSystem.AddDriver(3, 8.55)
// billingSystem.AddDriver(4, 11.28)
// billingSystem.AddDriver(5, 0.1)

// billingSystem.RecordDelivery(1, 0, 3600)       // Jan 1 1970 00:00:00 - Jan 1 1970 01:00:00
// billingSystem.RecordDelivery(5, 0, 3600)      // Jan 1 1970 00:00:00 - Jan 1 1970 01:00:00
// billingSystem.RecordDelivery(2, 0, 5400)      // Jan 1 1970 00:00:00 - Jan 1 1970 01:30:00
// billingSystem.RecordDelivery(2, 5400, 7200)   // Jan 1 1970 01:30:00 - Jan 1 1970 02:00:00
// billingSystem.RecordDelivery(5, 3600, 7200)   // Jan 1 1970 01:00:00 - Jan 1 1970 02:00:00
// billingSystem.RecordDelivery(5, 7200, 10800)  // Jan 1 1970 02:00:00 - Jan 1 1970 03:00:00
// billingSystem.RecordDelivery(5, 10800, 14400) // Jan 1 1970 03:00:00 - Jan 1 1970 04:00:00
// billingSystem.RecordDelivery(5, 14400, 18000) // Jan 1 1970 04:00:00 - Jan 1 1970 05:00:00

func main() {
	s := NewDriverDelivery()

	d1 := NewDriver(1, 35.10)
	d2 := NewDriver(2, 15.15)
	d3 := NewDriver(3, 8.55)
	d4 := NewDriver(4, 11.28)
	d5 := NewDriver(5, 0.1)

	s.AddDriver(d1)
	s.AddDriver(d2)
	s.AddDriver(d3)
	s.AddDriver(d4)
	s.AddDriver(d5)

	s.RecordDelivery(d1, NewDelivery(Epoch(0), Epoch(3600), 1))
	s.RecordDelivery(d5, NewDelivery(Epoch(0), Epoch(3600), 2))
	s.RecordDelivery(d2, NewDelivery(Epoch(0), Epoch(5400), 3))      // Jan 1 1970 00:00:00 - Jan 1 1970 01:30:00
	s.RecordDelivery(d2, NewDelivery(Epoch(5400), Epoch(7200), 4))   // Jan 1 1970 01:30:00 - Jan 1 1970 02:00:00
	s.RecordDelivery(d5, NewDelivery(Epoch(3600), Epoch(7200), 5))   // Jan 1 1970 01:00:00 - Jan 1 1970 02:00:00
	s.RecordDelivery(d5, NewDelivery(Epoch(7200), Epoch(10800), 6))  // Jan 1 1970 02:00:00 - Jan 1 1970 03:00:00
	s.RecordDelivery(d5, NewDelivery(Epoch(10800), Epoch(14400), 7)) // Jan 1 1970 03:00:00 - Jan 1 1970 04:00:00
	s.RecordDelivery(d5, NewDelivery(Epoch(14400), Epoch(18000), 8)) // Jan 1 1970 04:00:00 - Jan 1 1970 05:00:00

	cost := s.GetTotalCost()

	fmt.Println(cost)
}

func Epoch(s int) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: int64(s),
		Nanos:   0,
	}
}

type Driver struct {
	Id   int
	Rate float32
}

func NewDriver(Id int, rate float32) *Driver {
	return &Driver{
		Id:   Id,
		Rate: rate,
	}
}

type Delivery struct {
	Id        int
	StartTime *timestamppb.Timestamp
	EndTime   *timestamppb.Timestamp
	Price     float64
	Dr        *Driver
}

func NewDelivery(stTime, endTime *timestamppb.Timestamp, Id int) *Delivery {
	return &Delivery{
		Id:        Id,
		StartTime: stTime,
		EndTime:   endTime,
		Price:     0.0,
		Dr:        nil,
	}
}

type DriverDelivery struct {
	Deliveries map[int]*Delivery
	Drivers    map[int]*Driver
}

func NewDriverDelivery() *DriverDelivery {
	return &DriverDelivery{
		Deliveries: map[int]*Delivery{},
		Drivers:    map[int]*Driver{},
	}
}

func (d *DriverDelivery) AddDriver(driver *Driver) {
	d.Drivers[driver.Id] = driver
}

func (d *DriverDelivery) RecordDelivery(driver *Driver, delivery *Delivery) {
	d.Deliveries[delivery.Id] = delivery
	// price calculation
	price := (delivery.EndTime.AsTime().Sub(delivery.StartTime.AsTime()).Seconds() * float64(d.Drivers[driver.Id].Rate)) / float64(60*60)
	// TODO: update delivery price to round of to 2 decimal places
	price = price * 100
	price = float64(int(price)) / float64(100.0)
	d.Deliveries[delivery.Id].Price = price
}

func (d *DriverDelivery) GetTotalCost() float64 {
	cost := float64(0)
	for _, v := range d.Deliveries {
		cost += v.Price
	}
	return cost
}
