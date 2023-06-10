package models

import "time"

type User struct {
	ID         int
	Username   string
	Password   string
	Email      string
	Phone      string
	CreatedAt  time.Time
	RoleID     int
	Passengers []*Passenger
}

type Role struct {
	ID          int
	Name        string
	Description string
	Permission  []Permission
}

type Permission struct {
	ID          int
	Name        string
	Description string
}

type Gender int

const (
	Male Gender = iota
	Female
)

type Passenger struct {
	ID           int
	FirstName    string
	LastName     string
	NationalCode string
	Email        string
	Gender       Gender
	Phone        string
	Age          uint
	Address      string
	Tickets      []*Ticket
}

type FlightClass int

const (
	FirstClass FlightClass = iota
	BusinessClass
	EconomyClass
)

type Flight struct {
	ID            int
	FlightNo      string
	Departure     Airport
	Destination   Airport
	DepartureTime time.Time
	ArrivalTime   time.Time
	Airlines      Airline
	FlightClass   FlightClass
	Price         float64
	Capacity      int
}

type Ticket struct {
	ID        int
	Flight    *Flight
	Passenger *Passenger
	User      User
	StatusPay string
	Refund    bool
}

type Airline struct {
	ID   int
	Name string
	Logo string
}

type Airport struct {
	ID       int
	Name     string
	City     City
	Terminal string
}

type City struct {
	ID   int
	Name string
}
