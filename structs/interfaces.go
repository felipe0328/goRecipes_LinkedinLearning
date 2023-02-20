package structs

import "fmt"

// this method will show how to work with interfaces in go

var cars map[int]carData

type carData struct {
	id   int
	name string
}

type ICar interface {
	getCarData(carId int) carData
	createNewCar(carName string) carData
	getCars() []carData
}

type carImplementation struct{}

func testInterfaces() {
	cars = make(map[int]carData)
	var carIntImpl ICar
	carIntImpl = &carImplementation{}

	firstCar := carIntImpl.createNewCar("testCar")
	carIntImpl.createNewCar("secondCar")

	getFirstCar := carIntImpl.getCarData(firstCar.id)

	allCars := carIntImpl.getCars()

	fmt.Printf("We created an interface, and implemented that into a struct, creating two cars, getting the first: %v, and all the cars: %v", getFirstCar, allCars)

}

func (c *carImplementation) getCarData(carId int) carData {
	return cars[carId]
}

func (c *carImplementation) createNewCar(carName string) carData {
	newCar := carData{
		id:   len(cars) + 1,
		name: carName,
	}

	cars[newCar.id] = newCar

	return newCar
}

func (c *carImplementation) getCars() []carData {
	result := make([]carData, 0)

	for _, value := range cars {
		result = append(result, value)
	}

	return result
}
