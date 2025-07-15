package car

// Car Interface defines the methods that a car should implement
type Car Interface {
	// Start the car
	Start() error
	// Stop the car
	Stop() error
	// Get the current speed of the car
	GetCurrentSpeed() (int, error)
	// Get the make of the car
	GetMake() string
	// Get the model of the car
	GetModel() string
}	


type BaseCar struct {
	Id string
	Make string 
	Model string 
	Year int 
}