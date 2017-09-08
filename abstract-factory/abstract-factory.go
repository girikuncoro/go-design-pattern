package main

import "fmt"

type PickupTruck interface {
	Load() string
}
type SportCar interface {
	Race() string
}

type VehicleFactory interface {
	CreatePickupTruck() PickupTruck
	CreateSportCar() SportCar
}

type MazdaPickup struct{}

func (mp *MazdaPickup) Load() string {
	return "mazda loading 200 pounds"
}

type MazdaSport struct{}

func (ms *MazdaSport) Race() string {
	return "mazda racing 300 km/h"
}

type ToyotaPickup struct{}

func (tp *ToyotaPickup) Load() string {
	return "toyota loading 150 pounds"
}

type ToyotaSport struct{}

func (ts *ToyotaSport) Race() string {
	return "toyota racing 260 km/h"
}

type MazdaFactory struct{}

func (f *MazdaFactory) CreatePickupTruck() PickupTruck {
	return &MazdaPickup{}
}

func (f *MazdaFactory) CreateSportCar() SportCar {
	return &MazdaSport{}
}

type ToyotaFactory struct{}

func (f *ToyotaFactory) CreatePickupTruck() PickupTruck {
	return &ToyotaPickup{}
}

func (f *ToyotaFactory) CreateSportCar() SportCar {
	return &ToyotaSport{}
}

func main() {
	mazdaFactory := &MazdaFactory{}
	toyotaFactory := &ToyotaFactory{}

	mazdaPickup := mazdaFactory.CreatePickupTruck()
	mazdaSport := mazdaFactory.CreateSportCar()
	toyotaPickup := toyotaFactory.CreatePickupTruck()
	toyotaSport := toyotaFactory.CreateSportCar()

	fmt.Println(mazdaPickup.Load())
	fmt.Println(mazdaSport.Race())
	fmt.Println(toyotaPickup.Load())
	fmt.Println(toyotaSport.Race())
}
